package dbdao

import (
	"time"

	"gorm.io/gorm/clause"
)

type WorkbenchActionCard struct {
	CardID    string    `gorm:"column:card_id;primaryKey"`
	Status    string    `gorm:"column:status"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (WorkbenchActionCard) TableName() string {
	return "workbench_action_cards"
}

type WorkbenchActionEvent struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement"`
	CardID    string    `gorm:"column:card_id;index"`
	FromState string    `gorm:"column:from_state"`
	ToState   string    `gorm:"column:to_state"`
	Reason    string    `gorm:"column:reason"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (WorkbenchActionEvent) TableName() string {
	return "workbench_action_events"
}

func (d *DB) UpsertWorkbenchActionStatus(cardID string, status string) error {
	card := WorkbenchActionCard{
		CardID: cardID,
		Status: status,
	}
	return d.DB().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "card_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"status", "updated_at"}),
	}).Create(&card).Error
}

func (d *DB) ListWorkbenchActionStatuses() (map[string]string, error) {
	var cards []WorkbenchActionCard
	if err := d.DB().Find(&cards).Error; err != nil {
		return nil, err
	}
	out := make(map[string]string, len(cards))
	for _, c := range cards {
		out[c.CardID] = c.Status
	}
	return out, nil
}

func (d *DB) CreateWorkbenchActionEvent(cardID string, fromState string, toState string, reason string) error {
	event := WorkbenchActionEvent{
		CardID:    cardID,
		FromState: fromState,
		ToState:   toState,
		Reason:    reason,
	}
	return d.DB().Create(&event).Error
}

func (d *DB) ListWorkbenchActionEvents(cardID string) ([]WorkbenchActionEvent, error) {
	var events []WorkbenchActionEvent
	if err := d.DB().Where("card_id = ?", cardID).Order("id desc").Limit(50).Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}
