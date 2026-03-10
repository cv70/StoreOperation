# Action Card Component Spec V1

**Date:** 2026-03-10
**Scope:** Store manager, supervisor, HQ workbenches

## Goal

Define one shared action object for the whole product so all workbenches operate on the same execution logic:

- store manager executes
- supervisor pushes and replicates
- HQ evaluates and scales

This spec focuses on product behavior and component contract, not UI style code.

## Core Definition

Action Card is the smallest closed-loop unit in StoreOperation.

Each card must answer:

1. What happened
2. Why it happened
3. What should be done now
4. Who owns it
5. What gain/loss is expected
6. How effectiveness will be validated

## Card Types

### 1. Opportunity Card

Purpose: recover profit through proactive actions.

Typical examples:

- replenishment adjustment
- near-expiry handling
- shift tuning
- campaign correction

Priority basis:

- expected gain
- urgency window
- confidence score

### 2. Risk Card

Purpose: reduce loss by handling high-impact incidents quickly.

Typical examples:

- cold-chain exception
- process violation
- queue congestion
- equipment anomaly

Priority basis:

- estimated loss
- SLA deadline
- severity level

## Required Fields

All action cards require these fields.

### A. Identity

- `action_id` (global unique)
- `card_type` (`opportunity` | `risk`)
- `scenario_code` (for template mapping)
- `created_at`
- `source_system` (`camera` | `iot` | `robot` | `pos` | `inventory` | `manual`)

### B. Problem Summary

- `title` (one-line statement)
- `event_summary` (what happened)
- `reason_summary` (why it happened)
- `evidence_refs` (image/video/device snapshot/log IDs)

### C. Decision Layer

- `recommended_action` (clear operator verb)
- `alternative_action` (optional)
- `expected_impact_type` (`profit_gain` | `loss_reduction` | `risk_reduction`)
- `expected_impact_value` (currency or metric delta)
- `confidence_score` (0-100)
- `latest_action_deadline`

### D. Ownership and Routing

- `owner_role` (`store_manager` | `supervisor` | `hq_ops`)
- `owner_id`
- `watchers` (optional)
- `sla_minutes`
- `escalation_level` (0..N)

### E. Execution and Validation

- `execution_status` (state machine value)
- `result_writeback_required` (`true` | `false`)
- `result_fields` (required outcome inputs)
- `validation_metric` (how to judge effectiveness)
- `validation_window` (time window for result validation)

## State Machine

Single state machine shared across all workbenches.

### States

1. `new`
2. `accepted`
3. `in_progress`
4. `done_pending_validation`
5. `validated_effective`
6. `validated_ineffective`
7. `rejected_with_reason`
8. `expired`
9. `escalated`

### Transition Rules

- `new -> accepted` when user commits to action
- `new -> rejected_with_reason` only with mandatory reason
- `accepted -> in_progress` on first execution signal
- `in_progress -> done_pending_validation` when operator marks done and submits required result fields
- `done_pending_validation -> validated_effective` when metric meets threshold
- `done_pending_validation -> validated_ineffective` when metric fails threshold
- any non-terminal state -> `escalated` when SLA breached or risk severity increases
- `new` or `accepted` -> `expired` when deadline passes without execution

### Guardrails

- No card can reach validated states without `result_writeback_required` fields complete.
- Risk cards cannot be rejected without second confirmation if severity is high.
- Escalated cards must show escalation owner and next SLA.

## User Actions by Role

### Store Manager

Allowed primary actions:

- accept
- reject_with_reason
- start
- mark_done_with_result
- request_help

Not primary:

- cross-store replicate
- template publish

### Supervisor

Allowed primary actions:

- push (nudge owner)
- escalate
- reassign
- replicate_to_stores
- comment_review

Not primary:

- direct strategy publish

### HQ

Allowed primary actions:

- approve_template
- adjust_strategy_scope
- pause_or_stop_template
- trigger_replication_plan

Not primary:

- line-by-line frontline handling

## Opportunity vs Risk Behavior

### Opportunity Card Behavior

- default sort: expected impact value descending
- default SLA: shorter for high confidence + high impact
- can be batched for supervisor replication
- rejection allowed with reason taxonomy

Reason taxonomy:

- not feasible now
- local condition mismatch
- expected value not credible
- dependency missing

### Risk Card Behavior

- default sort: severity + loss + SLA urgency
- always evidence-first rendering
- high severity requires immediate action path on top area
- escalation is mandatory when SLA breached

## AI Output Contract

AI payload should be structured, not free-form narrative.

Required sections:

1. `judgment`
2. `reason`
3. `recommended_action`
4. `expected_impact`
5. `evidence_summary`
6. `escalate_or_replicate_advice`

Quality constraints:

- no recommendation without reason and evidence summary
- no impact value without confidence score
- confidence below threshold must surface uncertainty hint

## Ranking Logic (V1)

### Store Manager Ranking

`priority_score = impact_weight * expected_impact + urgency_weight * deadline_factor + confidence_weight * confidence_score`

Hard override:

- high-severity risk cards always appear in risk interruption zone

### Supervisor Ranking

Two separate queues:

- replication queue by `validated_effective + similarity_score + estimated_rollout_value`
- intervention queue by `overdue + repeated_issue + severity`

### HQ Ranking

Strategy view groups cards by template and scenario:

- ROI
- stability
- adoption rate
- false positive rate

## Result Writeback Contract

Result writeback is mandatory before validation.

Minimum writeback fields:

- action_completed_at
- action_done_by
- local_result_observation
- blocker_or_exception (optional but recommended)

Scenario-specific fields (examples):

- replenishment: replenished_qty, stockout_avoided
- near-expiry: disposed_qty, waste_reduction
- risk incident: response_time, issue_resolved_flag

## Component Layout Contract (Low Fidelity)

```text
[Card Header]
- Type Badge | Severity/Impact Badge | Deadline

[Problem]
- Title
- What happened
- Why happened
- Evidence quick links

[Decision]
- Recommended action
- Expected impact + confidence

[Ownership]
- Owner | SLA | Escalation level

[Execution]
- Status
- Required writeback checklist
- Primary action buttons
```

## Event Logging Requirements

Every state change should log:

- actor
- timestamp
- previous_state
- new_state
- reason_code (if reject/escalate/expire)

This is required for:

- auditability
- attribution analysis
- model quality feedback

## KPIs for Action Card Quality

Primary:

- acceptance rate
- on-time completion rate
- validation-effective rate

Secondary:

- false positive feedback rate
- average time to first action
- escalation rate

## Non-Goals for V1

- full workflow builder
- fully automatic execution without operator confirmation
- open-ended natural language card schema
- per-customer custom state machine

## Release Recommendation

1. Ship store manager card first (opportunity + risk)
2. Add supervisor replication and push interactions
3. Add HQ template-level aggregation and strategy controls

This sequence keeps the product aligned with execution reality before scaling strategy controls.
