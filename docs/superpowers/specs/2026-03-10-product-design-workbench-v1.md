# Product Design Workbench V1

**Date:** 2026-03-10

## Goal

Define a sharper product structure for StoreOperation so the system behaves like a daily operating product instead of a broad management platform. The design centers on three workbenches linked by one closed loop:

- store manager execution
- supervisor orchestration
- HQ strategy scaling

## Product Definition

StoreOperation should not be designed as three disconnected dashboards. It should be one operating system that converts on-site signals into business actions, then carries those actions through execution, replication, and ROI validation.

Core closed loop:

`现场信号 -> AI 判断 -> 动作执行 -> 结果归因 -> 策略沉淀`

## Core Design Principles

1. The product’s main object is not a chart or an alert. It is an action card.
2. Every role sees the same operating loop, but from a different decision layer.
3. Hardware is a first-class input layer, not hidden infrastructure.
4. AI should appear as structured judgment, not as a chat-first interface.
5. Static reporting should be weakened unless it directly supports action or strategy decisions.

## Unified Information Architecture

The system should be built around five layers:

### 1. On-Site Event Layer

Inputs from cameras, robots, IoT, and edge computing:

- queue congestion
- process violations
- equipment anomalies
- cold-chain deviations
- delivery / meal-routing bottlenecks

### 2. Business Judgment Layer

AI combines site events and business data to output:

- risk identification
- profit opportunity detection
- priority ranking
- recommended action
- expected gain / loss
- replication suitability

### 3. Action Execution Layer

The software platform turns judgment into:

- owner assignment
- SLA / deadline
- escalation
- multi-store rollout
- execution status
- result writeback

### 4. Attribution Layer

The system links actions to measurable outcomes:

- stockout reduction
- waste reduction
- response speed improvement
- labor efficiency improvement
- activity ROI improvement

### 5. Strategy Layer

Validated actions are elevated into reusable templates and strategies for regional and HQ use.

## Shared Operating Object: Action Card

Every action card must answer six questions:

1. What happened
2. Why it happened
3. What should be done now
4. Who owns it
5. What gain or loss is expected
6. How effectiveness will be judged after completion

This object should stay consistent across all workbenches:

- store manager sees action cards to execute
- supervisor sees action cards to push or replicate
- HQ sees action cards aggregated into strategy performance

## Workbench 1: Store Manager

### Positioning

`利润机会 + 风险告警双引擎首页`

### Product Goal

Within 10 seconds, the store manager should know:

- what profit can still be recovered today
- what risk must be handled immediately

### First Screen Structure

- top operating status bar
- profit opportunity area
- high-risk alert area
- execution status area
- weekly improvement trend

### Required Modules

#### Top Status Bar

- store name / shift
- today recoverable profit
- number of high-risk items
- actions in progress

#### Profit Opportunity Area

Only 3 top opportunity cards, ordered by expected benefit.

Each card includes:

- opportunity title
- why now
- recommended action
- expected gain
- latest action time
- accept / defer / reject actions

#### High-Risk Alert Area

Alerts must not be buried in normal tasks.

Each alert card includes:

- event type
- evidence image / video / device state
- estimated loss
- SLA
- escalation path
- immediate handling action

#### Execution Area

Only show:

- pending
- in progress
- pending verification

#### Trend Area

Keep lightweight:

- stockout trend
- waste trend
- response time trend
- labor trend

### Must Be Weakened or Removed

- full BI dashboards
- dense metric walls
- HQ configuration entry points
- low-priority reminder streams

## Workbench 2: Supervisor

### Positioning

`区域经营动作编排台`

### Product Goal

The supervisor should know each day:

- which actions are worth replicating
- which stores require intervention
- which unresolved items need active push

### First Screen Structure

- regional status bar
- replicable actions area
- high-risk stores interruption area
- unresolved execution area
- regional progress and template performance area

### Required Modules

#### Regional Status Bar

- replicable actions today
- high-risk stores
- overdue items

#### Replicable Actions Area

This is the strategic main entry.

Each card includes:

- source store
- scenario
- verified benefit
- applicable store count
- replication success rate
- replicate / adjust then replicate action

#### High-Risk Stores Area

This is an interruption layer, not the primary layout.

Each card includes:

- store name
- main issue
- risk level
- recent trend
- owner
- intervene / assign action options

#### Unresolved Execution Area

Show only matters requiring supervisor involvement:

- overdue
- repeated issues
- escalated items

#### Regional Effectiveness Area

- regional improvement progress
- action template effectiveness ranking

### Must Be Weakened or Removed

- all-store detail tables as default
- pure monitoring views with no action path
- store-level execution noise already handled by store managers

## Workbench 3: HQ

### Positioning

`策略实验与放大平台`

### Product Goal

HQ should be able to decide weekly:

- which strategies truly work
- where they should be expanded
- which ones should be adjusted or stopped

### First Screen Structure

- strategy status bar
- strategy ROI overview
- weekly decision recommendation
- replicable region / store list
- strategy health and abnormal strategy warning

### Required Modules

#### Strategy Status Bar

- weekly strategy ROI
- number of scalable strategies
- number of abnormal strategies

#### Strategy ROI Overview

Show strategy and template performance, including:

- coverage
- return
- payback period
- stability

#### Weekly Decision Recommendation

System should explicitly classify:

- expand
- hold
- adjust
- stop

#### Replicable Regions / Stores

Each card includes:

- benchmark store or region
- similar candidate stores
- expected value
- replication risk
- rollout planning action

#### Strategy Health Area

- adoption rate
- completion rate
- false-positive rate
- cross-region stability

#### Abnormal Strategy Warning

Surface:

- high false positives
- low adoption
- regional divergence
- performance decay

### Must Be Weakened or Removed

- frontline task streams
- line-by-line alert handling
- large executive dashboards without decision implications

## Low-Fidelity Wireframes

### Store Manager

```text
┌──────────────────────────────────────────────┐
│ 门店A 午市     今日可回收利润 ¥2,860   风险2 │
├───────────────────────┬──────────────────────┤
│ 利润机会 1            │ 高危风险 1           │
│ 缺货风险：招牌套餐    │ 冷链温度异常         │
│ 预计少卖 ¥680         │ 预计损失 ¥1,200      │
│ 建议：15分钟内补货    │ SLA：30分钟          │
│ [立即执行] [不采纳]   │ [立即处理] [升级]    │
├───────────────────────┼──────────────────────┤
│ 利润机会 2            │ 高危风险 2           │
├───────────────────────┴──────────────────────┤
│ 待执行 3 | 进行中 2 | 待确认 1               │
├──────────────────────────────────────────────┤
│ 本周趋势：缺货↓ 报损↓ 响应速度↑ 人效↑         │
└──────────────────────────────────────────────┘
```

### Supervisor

```text
┌──────────────────────────────────────────────┐
│ 华东一区   可复制动作6   高风险门店4   逾期9 │
├───────────────────────┬──────────────────────┤
│ 可复制动作            │ 高风险门店           │
│ 门店B：临期套餐组合   │ 门店F：出餐拥堵      │
│ 已验证收益 ¥1,320/周  │ 风险等级：高         │
│ 适配门店：12家        │ 连续3天异常          │
│ [复制到多店]          │ [介入] [下发动作]    │
├───────────────────────┼──────────────────────┤
│ 可复制动作 2          │ 高风险门店 2         │
├───────────────────────┴──────────────────────┤
│ 未闭环任务推进                                    │
│ 门店C 排班调整 逾期18h [催办] [升级]            │
│ 门店D 冷链整改 升级中 [评论] [继续跟踪]         │
├──────────────────────────────────────────────┤
│ 区域改善进度 | 模板效果榜                         │
└──────────────────────────────────────────────┘
```

### HQ

```text
┌──────────────────────────────────────────────┐
│ 总部策略台  ROI 18%   可放大策略3   异常2     │
├───────────────────────┬──────────────────────┤
│ 策略 ROI 总览         │ 本周决策建议         │
│ 补货策略 ROI 22%      │ 扩：补货模板-茶饮店  │
│ 临期策略 ROI 15%      │ 稳：冷链告警模板     │
│ 排班策略 ROI 9%       │ 调：活动控盘模板     │
│ [查看详情]            │ 停：某低采纳模板     │
├───────────────────────┴──────────────────────┤
│ 可复制区域/门店清单                                 │
│ 华东区 12家适配  预估增益 ¥8.6万 [生成计划]       │
│ 华南区 8家适配   风险中等       [查看差异]         │
├──────────────────────────────────────────────┤
│ 策略健康度 | 异常策略预警                           │
│ 采纳率 完成率 误报率 稳定性                          │
└──────────────────────────────────────────────┘
```

## Shared Navigation

Recommend only five primary navigation items:

- 首页
- 动作
- 风险
- 复盘
- 模板 / 策略

Role differences should mainly come from default landing page and permissions, not from completely different product structures.

## AI Presentation Rules

AI should output structured business judgment, not free-form chat by default.

Standard AI output format:

- judgment
- reason
- recommended action
- expected impact
- evidence
- whether to escalate or replicate

Chat can exist as an assistive layer, but not as the primary interaction model.

## Hardware Presentation Rules

Hardware value must be visible in product surfaces:

- camera events should show evidence screenshots / timestamps
- robot events should show route nodes and delay bottlenecks
- IoT events should show live state and threshold deviation
- edge events should indicate local recognition source

Users should clearly understand that StoreOperation is not pure software.

## Priority Recommendation

### P0

- store manager homepage redesign
- standardized action card
- evidence-based high-risk alerts
- accept / reject / writeback loop
- lightweight weekly review

### P1

- supervisor replication workflow
- high-risk store interruption mechanism
- batch push / escalation
- regional template performance comparison

### P2

- HQ strategy ROI view
- layered strategy rollout
- expand / hold / adjust / stop recommendations
- cross-region replication recommendation

## Immediate Next Design Step

The most important next design artifact is:

`Action Card Component Spec`

Because once the action card is stable, all three workbenches become much easier to align in UI, workflow, and data design.
