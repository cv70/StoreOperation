# Homepage Flow Spec V1

**Date:** 2026-03-10
**Scope:** Store manager, supervisor, HQ homepages
**Dependency:** `2026-03-10-action-card-component-spec.md`

## Goal

Define role-specific homepage flows that share one execution model:

- each role sees clear first-screen priorities
- first click leads to action, not browsing
- every path can end in validated outcome or explicit escalation

## Shared Flow Principles

1. Homepages are operational entry points, not BI overviews.
2. First click must be a high-value action decision.
3. Risk interruption always preempts normal opportunity browsing.
4. All flows operate on action cards and the same status machine.
5. Completion is not "marked done"; completion is "validated effective/ineffective."

## Store Manager Homepage Flow

### First-Screen Priority

1. Today recoverable profit
2. High-risk interruptions
3. Top opportunity cards
4. Execution lane (`pending`, `in_progress`, `pending_validation`)

### First Click Path (Primary Happy Path)

1. User opens homepage.
2. User clicks top opportunity card `recommended_action`.
3. Card detail panel opens with reason, evidence, impact, deadline.
4. User clicks `accept` and moves to `accepted`.
5. User starts execution and moves to `in_progress`.
6. User submits result writeback and moves to `done_pending_validation`.
7. System validates within configured window.
8. Card becomes `validated_effective` or `validated_ineffective`.

### Interruption Path (Risk Override)

1. A high-severity risk card appears in interruption zone.
2. User clicks risk card before any opportunity task.
3. User sees evidence-first layout and SLA countdown.
4. User chooses `handle_now` or `request_help`.
5. If SLA breaches, system auto-escalates and card enters `escalated`.

### Rejection Path

1. User clicks `reject`.
2. Reason taxonomy is mandatory.
3. Card enters `rejected_with_reason`.
4. Reason is logged for model and template feedback.

### Success Metrics (Store Manager Flow)

- time to first action
- opportunity acceptance rate
- on-time completion rate
- risk response within SLA
- validated-effective rate

## Supervisor Homepage Flow

### First-Screen Priority

1. Replicable action queue
2. High-risk store interruption list
3. Unresolved escalation queue
4. Regional improvement summary

### First Click Path (Replication Happy Path)

1. User opens homepage.
2. User clicks top replicable action card.
3. System shows source store proof and suitability list.
4. User selects target stores and clicks `replicate`.
5. Cards are generated for target stores in `new` state.
6. User assigns owners and deadlines in batch.
7. Supervisor tracks rollout performance and completion outcomes.

### Intervention Path (High-Risk Store)

1. User clicks high-risk store card in interruption list.
2. System opens unresolved risk cards and overdue actions.
3. User performs `push`, `reassign`, or `escalate`.
4. System logs intervention and updates SLA ownership.

### Escalation Follow-Up Path

1. User opens unresolved queue.
2. User filters by repeated issue and overdue duration.
3. User performs batch push/escalate.
4. Queue refreshes with new owners and deadlines.

### Success Metrics (Supervisor Flow)

- replication adoption rate
- replication success rate
- overdue reduction rate
- time from overdue to intervention
- repeated issue recurrence rate

## HQ Homepage Flow

### First-Screen Priority

1. Strategy ROI overview
2. Weekly decision panel (`expand`, `hold`, `adjust`, `stop`)
3. Replication candidate regions/stores
4. Strategy health and anomaly warnings

### First Click Path (Strategy Decision Happy Path)

1. User opens homepage.
2. User clicks a top strategy ROI card.
3. System opens strategy detail with:
   - ROI trend
   - adoption
   - false positive rate
   - stability by region
4. User chooses one control action:
   - `expand_scope`
   - `hold`
   - `adjust_rules`
   - `pause`
5. System publishes updated scope/rules and logs decision.

### Replication Planning Path

1. User clicks replication candidate card.
2. System shows benchmark store evidence and similarity reasoning.
3. User confirms target set and rollout wave.
4. System generates replication plan for supervisor execution.

### Anomaly Handling Path

1. User clicks abnormal strategy warning.
2. System shows root contributors:
   - low adoption
   - high false positives
   - regional divergence
3. User chooses `adjust` or `stop`.
4. Downstream rollout is updated and flagged for next weekly review.

### Success Metrics (HQ Flow)

- strategy decision cycle time
- strategy ROI improvement
- scale success rate after expansion
- stop-loss response time for failing strategies
- cross-region stability score

## Cross-Role Handoff Flows

### Handoff 1: Store Manager -> Supervisor

Trigger:

- card overdue
- repeated rejection
- high-severity risk with no closure

Behavior:

1. Card escalates to supervisor queue.
2. Supervisor receives intervention context and evidence.
3. Ownership updates are visible to store manager immediately.

### Handoff 2: Supervisor -> HQ

Trigger:

- replication pattern shows strong positive/negative trend
- repeated instability in same scenario template

Behavior:

1. Supervisor marks template performance summary.
2. HQ strategy panel receives template signal.
3. HQ decides scale/adjust/stop.

### Handoff 3: HQ -> Supervisor/Stores

Trigger:

- strategy scope or rule update

Behavior:

1. HQ publishes decision package.
2. Supervisor gets rollout tasks.
3. Store cards update with new template version tags.

## Required Homepage Components

All three homepages must include:

- interruption zone for high-risk items
- action queue with explicit ordering logic
- state visibility (`new` to validation states)
- quick access to evidence
- outcome validation reminders

## Mobile and Desktop Behavior

### Desktop

- two-column or three-zone layout allowed
- interruption zone pinned above fold
- one-click access to action detail drawer

### Mobile

- single-column priority stack
- interruption cards fixed at top of feed
- quick action buttons remain sticky in card detail

## Guardrails

1. Do not place full historical analytics above action entry areas.
2. Do not mix high-risk interruption cards into generic task lists.
3. Do not allow completion without required result writeback fields.
4. Do not allow HQ homepage to become frontline task board.
5. Do not allow supervisor homepage to default to full store data tables.

## Telemetry Requirements

Track these homepage events:

- `homepage_opened`
- `first_card_clicked`
- `card_accepted`
- `card_rejected`
- `card_started`
- `card_done_pending_validation`
- `card_validated_effective`
- `card_validated_ineffective`
- `card_escalated`
- `replication_created`
- `strategy_decision_published`

Each event must include:

- `role`
- `store_or_region_scope`
- `scenario_code`
- `card_type`
- `timestamp`

## Rollout Recommendation

1. Launch store manager homepage flow first with strict interruption logic.
2. Add supervisor replication and intervention flows.
3. Launch HQ decision panel after enough validated card outcomes are available.

This order minimizes strategy noise and keeps flows grounded in real execution data.
