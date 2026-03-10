# Pilot UAT Cases V1

**Date:** 2026-03-10
**Scope:** 8-week pilot acceptance testing
**Dependencies:**
- `2026-03-10-action-card-component-spec.md`
- `2026-03-10-homepage-flow-spec.md`
- `2026-03-10-interaction-copy-spec.md`
- `2026-03-10-mvp-delivery-checklist.md`

## Goal

Provide executable UAT cases for pilot validation across all three roles:

- store manager
- supervisor
- HQ

Each case includes:

- preconditions
- test steps
- expected result
- pass criteria

## UAT Rules

1. Test data must include both opportunity and risk cards.
2. UAT must be run on desktop and mobile form factors.
3. All role actions must be tested with role-correct accounts.
4. Each critical flow must have at least one negative test.
5. UAT pass requires evidence capture (screens, logs, event IDs).

## Test Data Baseline

Prepare at minimum:

- 3 pilot stores
- 1 supervisor covering all pilot stores
- 1 HQ operator
- 10 opportunity cards across top scenarios
- 6 risk cards with mixed severities
- at least 2 cards expected to escalate

## Section A: Store Manager UAT

### UAT-SM-01 Homepage Priority Rendering

Preconditions:

- store manager account active
- homepage has at least 2 opportunities and 1 high-risk card

Steps:

1. Open store manager homepage.
2. Observe top section and interruption zone.
3. Check ordering of opportunity cards and risk cards.

Expected:

- today recoverable profit visible above fold
- high-risk interruption zone visible above fold
- opportunities sorted by expected impact

Pass criteria:

- all three checks true on desktop and mobile

### UAT-SM-02 Opportunity Execution Happy Path

Preconditions:

- one `new` opportunity card with confidence >=70

Steps:

1. Open top opportunity card.
2. Click `立即执行` or equivalent primary CTA.
3. Start task and move to in-progress.
4. Submit required writeback fields.

Expected:

- state transitions: `new -> accepted -> in_progress -> done_pending_validation`
- validation result produced within configured window

Pass criteria:

- no blocked transition
- all transitions logged with actor/timestamp

### UAT-SM-03 Opportunity Rejection Path

Preconditions:

- one `new` opportunity card

Steps:

1. Open card.
2. Click reject action.
3. Attempt submit without reason.
4. Submit with valid reason code.

Expected:

- submit blocked without reason
- card reaches `rejected_with_reason` with reason persisted

Pass criteria:

- reason is mandatory and queryable in card detail

### UAT-SM-04 High-Risk Handling Path

Preconditions:

- one high-severity risk card with SLA countdown

Steps:

1. Open risk card from interruption zone.
2. Verify evidence and SLA visibility.
3. Click primary handling action.
4. Submit handling writeback.

Expected:

- evidence links visible
- SLA visible
- state progresses without skipping required fields

Pass criteria:

- risk card can reach done_pending_validation with complete writeback

### UAT-SM-05 Risk Escalation Negative Path

Preconditions:

- one high-risk card near SLA breach

Steps:

1. Do not take action until SLA expires.
2. Refresh card status.

Expected:

- card auto-transitions to `escalated`
- escalation owner and level visible

Pass criteria:

- escalation event generated and visible in logs

## Section B: Supervisor UAT

### UAT-SV-01 Replication Queue Rendering

Preconditions:

- at least 2 validated-effective cards available as replication candidates

Steps:

1. Open supervisor homepage.
2. Inspect replicable action queue.
3. Open top candidate details.

Expected:

- queue visible above fold
- each card shows source evidence, expected value, target suitability

Pass criteria:

- all required fields present for top 2 candidates

### UAT-SV-02 Replicate to Multi-Store

Preconditions:

- one replicable candidate
- at least 2 target stores available

Steps:

1. Click `复制到多店`.
2. Select two stores.
3. Assign owner and deadline.
4. Confirm replication.

Expected:

- target cards created in `new` state
- ownership and SLA populated on created cards

Pass criteria:

- replication finished in <=4 steps
- `replication_created` event logged

### UAT-SV-03 Overdue Intervention

Preconditions:

- one overdue card in unresolved queue

Steps:

1. Open unresolved queue.
2. Click overdue card.
3. Perform `催办` then `升级处理`.

Expected:

- intervention action logged
- owner/escalation level updated

Pass criteria:

- queue reflects update within refresh cycle

### UAT-SV-04 Permission Negative Test

Preconditions:

- supervisor account active

Steps:

1. Attempt HQ-only action (for example `暂停策略`) from any available surface.

Expected:

- action blocked by permission policy

Pass criteria:

- clear permission error shown
- no strategy state change

## Section C: HQ UAT

### UAT-HQ-01 Strategy ROI Overview

Preconditions:

- strategy data available for at least 3 scenarios

Steps:

1. Open HQ homepage.
2. Open one strategy ROI card.
3. Inspect metrics and trend panel.

Expected:

- ROI, adoption, false-positive, stability shown
- no frontline task stream shown

Pass criteria:

- all required metrics visible

### UAT-HQ-02 Publish Strategy Decision

Preconditions:

- one active strategy card

Steps:

1. Open strategy card.
2. Choose one decision: `扩大范围` / `维持观察` / `调整策略` / `暂停策略`.
3. Confirm publish.

Expected:

- decision persisted
- audit record created
- updated scope visible on next fetch

Pass criteria:

- strategy decision end-to-end completed <=5 steps
- `strategy_decision_published` event logged

### UAT-HQ-03 Replication Planning

Preconditions:

- one replication candidate region

Steps:

1. Open replication candidate details.
2. Review similarity and risk rationale.
3. Confirm rollout wave.

Expected:

- supervisor-side rollout tasks generated
- plan linked to strategy version

Pass criteria:

- generated tasks visible to supervisor account

### UAT-HQ-04 Permission Negative Test

Preconditions:

- HQ account active

Steps:

1. Attempt frontline-only action such as direct mark-done on store card.

Expected:

- action blocked by policy

Pass criteria:

- no store card state changed by HQ restricted action

## Section D: Cross-Role Handoff UAT

### UAT-XR-01 Store Manager -> Supervisor Escalation

Preconditions:

- one card expected to breach SLA

Steps:

1. Let card breach SLA at store level.
2. Check supervisor unresolved queue.

Expected:

- escalated card appears with full context/evidence

Pass criteria:

- handoff latency within configured refresh interval

### UAT-XR-02 Supervisor -> HQ Strategy Signal

Preconditions:

- repeated pattern of high-performing replication

Steps:

1. Supervisor marks pattern summary.
2. Open HQ strategy panel.

Expected:

- strategy signal appears in HQ decision inputs

Pass criteria:

- signal references underlying cards and scenario

### UAT-XR-03 HQ -> Supervisor Rollout Update

Preconditions:

- HQ publishes strategy scope change

Steps:

1. Publish HQ decision.
2. Check supervisor replication queue and affected store cards.

Expected:

- updated template/scope visible in downstream queues

Pass criteria:

- version tag consistent across HQ, supervisor, store surfaces

## Section E: Telemetry and Audit UAT

### UAT-TM-01 Event Completeness

Preconditions:

- all role flows executed at least once

Steps:

1. Query telemetry sink for required event names.
2. Validate required dimensions.

Expected:

- all required events present
- each includes role, scope, scenario_code, card_type, timestamp

Pass criteria:

- event completeness >=99% for executed flows

### UAT-TM-02 State Transition Auditability

Preconditions:

- at least 5 cards with full lifecycle history

Steps:

1. Open audit logs for selected cards.
2. Verify transition chain and actor attribution.

Expected:

- no missing transition records
- invalid transitions absent

Pass criteria:

- 100% selected cards have continuous transition history

## Section F: Copy and UX Guardrail UAT

### UAT-CX-01 Card Copy Completeness

Preconditions:

- sample of 10 cards across types

Steps:

1. Review each card in homepage context.
2. Check for reason, action, impact, confidence, time.

Expected:

- all cards include required copy blocks

Pass criteria:

- completeness rate 100% in sample set

### UAT-CX-02 Low Confidence CTA Guardrail

Preconditions:

- one card with confidence <55

Steps:

1. Open low-confidence card.
2. Inspect primary CTA text.

Expected:

- primary CTA uses review-first wording, not immediate execute wording

Pass criteria:

- guardrail passes for all low-confidence samples

## UAT Exit Criteria

Pilot UAT is considered passed when:

1. All P0 cases pass.
2. No Sev-1 defects remain open.
3. No blocker in role permissions and escalation.
4. Telemetry completeness target reached.
5. At least one end-to-end closed loop validated per core scenario.

## Defect Severity Guideline

- `Sev-1`: blocks core action flow, wrong role authorization, missing escalation
- `Sev-2`: major friction with workaround available
- `Sev-3`: minor UI/copy issue with no flow break

## UAT Evidence Package

For each case retain:

- case ID
- tester name
- environment
- timestamp
- screenshots/video
- event IDs/log references
- pass/fail result
- defect ID if failed
