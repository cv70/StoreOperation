package workbench

import (
	"errors"
	"sync"
)

var ErrCardNotFound = errors.New("action card not found")
var ErrInvalidTransition = errors.New("invalid transition")

var allowedTransitions = map[string]map[string]bool{
	"new":                     {"accepted": true, "rejected_with_reason": true, "escalated": true},
	"accepted":                {"in_progress": true, "rejected_with_reason": true, "escalated": true},
	"in_progress":             {"done_pending_validation": true, "escalated": true},
	"done_pending_validation": {"validated_effective": true, "validated_ineffective": true},
	"escalated":               {"in_progress": true, "done_pending_validation": true},
}

type Store struct {
	mu    sync.RWMutex
	repo  Repository
	roles map[string]OverviewResp
	index map[string]string
}

func NewStore(repo Repository) *Store {
	if repo == nil {
		repo = noopRepo{}
	}
	roles := map[string]OverviewResp{
		"store_manager": storeManagerOverview(),
		"supervisor":    supervisorOverview(),
		"hq":            hqOverview(),
	}
	s := &Store{
		repo:  repo,
		roles: roles,
		index: make(map[string]string),
	}
	s.reindex()
	return s
}

func (s *Store) Overview(role string) OverviewResp {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if overview, ok := s.roles[role]; ok {
		return s.applyPersistedStatuses(overview)
	}
	return s.applyPersistedStatuses(s.roles["store_manager"])
}

func (s *Store) Transition(cardID string, toState string, reason string) (Card, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	role, ok := s.index[cardID]
	if !ok {
		return Card{}, ErrCardNotFound
	}

	overview := s.roles[role]
	for i, card := range overview.Opportunities {
		if card.ID != cardID {
			continue
		}
		fromState := card.Status
		updated, err := transitionCard(card, toState, reason)
		if err != nil {
			return Card{}, err
		}
		overview.Opportunities[i] = updated
		s.roles[role] = overview
		if err := s.repo.UpsertStatus(cardID, updated.Status); err != nil {
			return Card{}, err
		}
		if err := s.repo.AppendEvent(cardID, fromState, updated.Status, reason); err != nil {
			return Card{}, err
		}
		return updated, nil
	}

	for i, card := range overview.Risks {
		if card.ID != cardID {
			continue
		}
		fromState := card.Status
		updated, err := transitionCard(card, toState, reason)
		if err != nil {
			return Card{}, err
		}
		overview.Risks[i] = updated
		s.roles[role] = overview
		if err := s.repo.UpsertStatus(cardID, updated.Status); err != nil {
			return Card{}, err
		}
		if err := s.repo.AppendEvent(cardID, fromState, updated.Status, reason); err != nil {
			return Card{}, err
		}
		return updated, nil
	}

	return Card{}, ErrCardNotFound
}

func (s *Store) Events(cardID string) ([]ActionEvent, error) {
	return s.repo.ListEvents(cardID)
}

func (s *Store) reindex() {
	for role, overview := range s.roles {
		for _, card := range overview.Opportunities {
			s.index[card.ID] = role
		}
		for _, card := range overview.Risks {
			s.index[card.ID] = role
		}
	}
}

func (s *Store) applyPersistedStatuses(overview OverviewResp) OverviewResp {
	statuses, err := s.repo.ListStatuses()
	if err != nil {
		return overview
	}
	for i, c := range overview.Opportunities {
		if status, ok := statuses[c.ID]; ok {
			overview.Opportunities[i].Status = status
		}
	}
	for i, c := range overview.Risks {
		if status, ok := statuses[c.ID]; ok {
			overview.Risks[i].Status = status
		}
	}
	return overview
}

func transitionCard(card Card, toState string, reason string) (Card, error) {
	if !allowedTransitions[card.Status][toState] {
		return Card{}, ErrInvalidTransition
	}
	if toState == "rejected_with_reason" && reason == "" {
		return Card{}, ErrInvalidTransition
	}
	card.Status = toState
	return card, nil
}

func storeManagerOverview() OverviewResp {
	return OverviewResp{
		Role:     "store_manager",
		Headline: "今日先抢利润机会，再处理高危风险。",
		Metrics: []Metric{
			{Key: "recoverable_profit", Label: "今日可回收利润", Value: "¥2,860", Trend: "+12%"},
			{Key: "high_risk_count", Label: "高危风险", Value: "2", Trend: "-1"},
			{Key: "in_progress", Label: "进行中动作", Value: "5", Trend: "+2"},
		},
		Opportunities: []Card{
			{ID: "opp-1", Title: "招牌套餐缺货风险，建议15分钟内补货", Reason: "午市客流上升，安全库存低于阈值。", Action: "补货20份并回写实际补货量", Impact: "预计减少损失¥680", Deadline: "14:30", Priority: "high", CardType: "opportunity", OwnerRole: "store_manager", Status: "new"},
			{ID: "opp-2", Title: "临期库存积压，建议组合促销处理", Reason: "2小时后进入临期窗口。", Action: "上架组合售卖并标注处理方式", Impact: "预计减少报损¥430", Deadline: "16:00", Priority: "medium", CardType: "opportunity", OwnerRole: "store_manager", Status: "new"},
		},
		Risks: []Card{
			{ID: "risk-1", Title: "冷链温度异常，需立即排查", Reason: "冷柜温度连续30分钟高于阈值。", Action: "立即排查并回写复位时间", Impact: "预计避免损失¥1,200", Deadline: "SLA 12分钟", Priority: "critical", CardType: "risk", OwnerRole: "store_manager", Status: "new"},
		},
		Focus: []string{
			"机会卡优先按收益排序",
			"高危风险必须插队处理",
			"完成后必须回写结果用于验证",
		},
	}
}

func supervisorOverview() OverviewResp {
	return OverviewResp{
		Role:     "supervisor",
		Headline: "先复制已验证动作，再介入高风险门店。",
		Metrics: []Metric{
			{Key: "replicable_actions", Label: "今日可复制动作", Value: "6", Trend: "+3"},
			{Key: "high_risk_stores", Label: "高风险门店", Value: "4", Trend: "-1"},
			{Key: "overdue_items", Label: "逾期事项", Value: "9", Trend: "-2"},
		},
		Opportunities: []Card{
			{ID: "opp-s-1", Title: "门店B临期套餐组合可复制到12店", Reason: "同业态门店验证有效率高。", Action: "批量下发模板并设置截止时间", Impact: "预计周增益¥8,600", Deadline: "今日18:00", Priority: "high", CardType: "opportunity", OwnerRole: "supervisor", Status: "new"},
		},
		Risks: []Card{
			{ID: "risk-s-1", Title: "门店F出餐拥堵连续3天", Reason: "高峰时段任务闭环率低。", Action: "介入并升级到区域负责人", Impact: "预计降低超时率22%", Deadline: "SLA 30分钟", Priority: "high", CardType: "risk", OwnerRole: "supervisor", Status: "new"},
		},
		Focus: []string{
			"复制动作优先于全量巡检",
			"逾期事项批量催办和升级",
			"跨店复用前先看适配条件",
		},
	}
}

func hqOverview() OverviewResp {
	return OverviewResp{
		Role:     "hq",
		Headline: "按策略ROI做扩、稳、调、停决策。",
		Metrics: []Metric{
			{Key: "strategy_roi", Label: "本周策略ROI", Value: "18%", Trend: "+2.3%"},
			{Key: "scalable_strategies", Label: "可放大策略", Value: "3", Trend: "+1"},
			{Key: "abnormal_strategies", Label: "异常策略", Value: "2", Trend: "0"},
		},
		Opportunities: []Card{
			{ID: "opp-h-1", Title: "补货模板在华东区可扩大到12店", Reason: "采纳率与有效率连续两周达标。", Action: "发布扩店策略并生成复制计划", Impact: "预计月增益¥86,000", Deadline: "本周五", Priority: "high", CardType: "opportunity", OwnerRole: "hq", Status: "new"},
		},
		Risks: []Card{
			{ID: "risk-h-1", Title: "活动控盘模板在华南误报偏高", Reason: "误报率高于阈值且采纳率下降。", Action: "下调权重并分层灰度发布", Impact: "预计减少无效任务18%", Deadline: "48小时", Priority: "medium", CardType: "risk", OwnerRole: "hq", Status: "new"},
		},
		Focus: []string{
			"总部只做策略决策，不做一线处理",
			"所有放大动作必须绑定归因口径",
			"异常策略优先止损再优化",
		},
	}
}
