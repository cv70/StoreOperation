# Pilot Risk Register Template V1

**Date:** 2026-03-10
**Scope:** 8-week pilot risk tracking and governance
**Dependencies:**
- `2026-03-10-mvp-delivery-checklist.md`
- `2026-03-10-pilot-kpi-baseline-template.md`
- `2026-03-10-pilot-weekly-review-deck-outline.md`

## Goal

Provide one standard risk register to track, prioritize, and resolve pilot risks with clear ownership and deadlines.

## Usage Rules

1. Review weekly in the pilot governance meeting.
2. Update status and next action for every open risk.
3. No risk can stay ownerless.
4. Escalate any `critical` risk within 24 hours.

## Risk Levels

- `critical`: immediate impact to pilot outcome or data trust
- `high`: likely to affect KPI trend or delivery milestone
- `medium`: manageable with planned mitigation
- `low`: monitor only

## Risk Register Table

| risk_id | category | title | level | trigger_signal | impact_area | owner | mitigation_plan | due_date | status | escalation_path |
|---|---|---|---|---|---|---|---|---|---|---|
| R-001 | data |  |  |  |  |  |  |  | open |  |
| R-002 | product |  |  |  |  |  |  |  | open |  |
| R-003 | operations |  |  |  |  |  |  |  | open |  |

## Category Definitions

- `data`: source quality, missing fields, delay, mismatch
- `product`: flow blockers, permission issues, state machine errors
- `model`: false positives, low confidence quality, drift
- `operations`: low adoption, incomplete writeback, owner inactivity
- `hardware`: evidence capture failure, device downtime, edge instability
- `commercial`: KPI miss risk, conversion risk, stakeholder alignment risk

## Mandatory Fields Guide

- `trigger_signal`: measurable symptom, not opinion
- `impact_area`: KPI, flow, scope, or customer trust
- `mitigation_plan`: concrete action with owner and ETA
- `escalation_path`: named role sequence

## Weekly Risk Review Template

### 1. New Risks (This Week)

| risk_id | title | level | owner | first_detected |
|---|---|---|---|---|

### 2. Escalated Risks

| risk_id | reason_for_escalation | current_owner | next_update_due |
|---|---|---|---|

### 3. Closed Risks

| risk_id | closure_reason | validated_by | closed_date |
|---|---|---|---|

## Redline Triggers (Auto Escalate)

1. KPI worsens 2 weeks in a row for 2+ core metrics.
2. High-risk SLA compliance drops below target.
3. Evidence availability falls below threshold for high-risk cards.
4. Writeback completeness drops below agreed floor.
5. Data pipeline failure lasts over 1 business day.

## Status Values

- `open`
- `mitigating`
- `blocked`
- `escalated`
- `closed`

## Closure Criteria

A risk can be marked `closed` only if:

1. Trigger signal is no longer present.
2. Impact metric is back within acceptable range.
3. Owner and reviewer confirm closure evidence.

## Minimum Governance Cadence

- Daily: critical/high risk check (15 min)
- Weekly: full risk register review
- Bi-weekly: executive escalation review

## Example Entries

| risk_id | category | title | level | trigger_signal | impact_area | owner | mitigation_plan | due_date | status | escalation_path |
|---|---|---|---|---|---|---|---|---|---|---|
| R-101 | model | High false positives in cold-chain alerts | high | false_positive_feedback_rate > target for 5 days | adoption + trust | PM + ML lead | tune threshold, add evidence filter, revalidate on 2 stores | 2026-03-17 | mitigating | Supervisor Lead -> HQ Ops -> Pilot Sponsor |
| R-102 | operations | Store writeback completeness low | medium | writeback_completeness_rate < 80% | attribution quality | CS lead | retrain store operators + enforce mandatory checklist | 2026-03-14 | open | CS Lead -> Customer Ops Owner |

## Definition of Good Risk Register

A good register is one where:

1. All critical/high risks have named owners and due dates.
2. Escalation path is explicit before escalation is needed.
3. Mitigation actions are measurable and time-bound.
4. Closed risks include evidence, not only statements.
