# MVP Delivery Checklist V1

**Date:** 2026-03-10
**Scope:** 8-week pilot MVP for StoreOperation
**Dependencies:**
- `2026-03-10-product-design-workbench-v1.md`
- `2026-03-10-action-card-component-spec.md`
- `2026-03-10-homepage-flow-spec.md`
- `2026-03-10-interaction-copy-spec.md`

## Goal

Provide one execution checklist that aligns product, engineering, data, and delivery.

This checklist maps each MVP capability to:

- page/module
- API/data contract
- telemetry events
- acceptance criteria

## MVP Boundary

In scope:

- store manager homepage core flow
- supervisor replication and intervention core flow
- HQ strategy decision core flow
- action card state machine
- evidence-based risk handling
- result writeback and validation loop

Out of scope:

- full workflow builder
- custom per-customer state machine
- broad BI expansion
- fully autonomous execution without human confirmation

## Workstream A: Frontend Pages and Modules

### A1. Store Manager Homepage

Status: `todo`

Must deliver:

- today recoverable profit header
- top opportunity cards (max 3)
- high-risk interruption zone
- execution lane (`pending`, `in_progress`, `pending_validation`)
- weekly trend summary

Acceptance:

- first action click in <=10s in usability test
- risk cards always visible above fold on mobile and desktop

### A2. Supervisor Homepage

Status: `todo`

Must deliver:

- replicable action queue
- high-risk store interruption zone
- unresolved escalation queue
- regional effect summary

Acceptance:

- replicate-to-multi-store flow completed in <=4 steps
- overdue intervention action reachable in <=2 clicks

### A3. HQ Homepage

Status: `todo`

Must deliver:

- strategy ROI overview
- weekly decision panel (`expand/hold/adjust/stop`)
- replication candidate list
- strategy health and anomaly panel

Acceptance:

- one strategy decision can be published end-to-end in <=5 steps
- no frontline task list appears on HQ homepage

## Workstream B: Action Card and Workflow Engine

### B1. Action Card Schema

Status: `todo`

Must deliver:

- required identity fields
- reason/evidence/action/impact fields
- ownership and SLA fields
- result writeback contract fields

Acceptance:

- cards cannot be created without required fields
- card payload validates against schema in API tests

### B2. State Machine

Status: `todo`

Must deliver:

- states from `new` to validation terminal states
- transition guards
- escalation handling
- expiry handling

Acceptance:

- invalid transitions rejected with explicit reason
- risk card SLA breach triggers escalation event

### B3. Role Actions and Permissions

Status: `todo`

Must deliver:

- store manager actions: accept/reject/start/done/help
- supervisor actions: push/escalate/reassign/replicate
- HQ actions: expand/hold/adjust/stop

Acceptance:

- unauthorized actions blocked by role policy
- role-specific button sets match copy spec

## Workstream C: API and Data Contracts

### C1. Action Card APIs

Status: `todo`

Endpoints:

- create/list/get action cards
- state transition endpoint
- result writeback endpoint
- escalation endpoint

Acceptance:

- all transitions logged with actor and timestamp
- list API supports role-based homepage sorting

### C2. Evidence APIs

Status: `todo`

Endpoints:

- fetch evidence summary for card
- fetch raw evidence references

Acceptance:

- risk cards render with at least one evidence reference
- missing evidence path returns fallback status and message

### C3. Strategy APIs

Status: `todo`

Endpoints:

- strategy ROI summary
- decision publish (`expand/hold/adjust/stop`)
- replication candidate query

Acceptance:

- decision publish creates audit trail
- updated scope visible on next fetch

## Workstream D: Telemetry and Analytics

### D1. Event Instrumentation

Status: `todo`

Required events:

- homepage_opened
- first_card_clicked
- card_accepted
- card_rejected
- card_started
- card_done_pending_validation
- card_validated_effective
- card_validated_ineffective
- card_escalated
- replication_created
- strategy_decision_published

Acceptance:

- each event includes role, scope, scenario_code, card_type, timestamp
- event delivery success rate >=99%

### D2. KPI Dashboards (Internal)

Status: `todo`

Must expose:

- acceptance rate
- on-time completion rate
- validated-effective rate
- risk SLA compliance
- replication success rate
- strategy decision cycle time

Acceptance:

- KPI refresh available daily
- metrics definitions documented and versioned

## Workstream E: Copy and UX Quality

### E1. Copy Integration

Status: `todo`

Must deliver:

- card templates wired to interaction copy spec
- confidence wording threshold logic
- rejection reason taxonomy

Acceptance:

- no card shown without action + reason + impact + deadline
- confidence<55 cards use review-first primary CTA

### E2. UX Guardrails

Status: `todo`

Must deliver:

- interruption zone design lock
- mobile card truncation/expand rules
- explicit writeback checklist in execution flow

Acceptance:

- high-risk cards cannot be hidden below generic task lists
- done state always requires required writeback fields

## Workstream F: Pilot Readiness

### F1. Pilot Config Pack

Status: `todo`

Must deliver:

- scenario template configuration for top 5 P0 scenarios
- SLA defaults by scenario
- escalation rules by role

Acceptance:

- pilot tenant can enable templates without engineering changes

### F2. Operational Runbook

Status: `todo`

Must deliver:

- weekly review checklist
- incident handling checklist
- false-positive feedback loop

Acceptance:

- customer success can run weekly cycle with no product team intervention

## Verification Checklist Before Pilot Start

1. Store manager happy path runs from homepage to validated outcome.
2. High-risk interruption can escalate automatically after SLA breach.
3. Supervisor can replicate one validated action to multiple stores.
4. HQ can publish one strategy decision and see scope update.
5. Telemetry captures all required events.
6. KPI dashboard shows baseline and week-over-week movement.
7. Copy checks pass for all top 5 scenario templates.

## Suggested Delivery Sequence (8 Weeks)

Week 1-2:

- action card schema/state machine
- store manager homepage core
- core telemetry events

Week 3-4:

- risk evidence rendering
- result writeback and validation
- supervisor intervention queue

Week 5-6:

- supervisor replication flow
- HQ ROI summary and decision panel
- KPI internal dashboard v1

Week 7-8:

- pilot runbook and config pack
- UAT and bug fixes
- acceptance review and go-live gate

## Definition of Done (MVP)

MVP is done only when all are true:

1. Core role flows are executable in production-like environment.
2. Action cards can complete full closed loop to validated outcomes.
3. Escalation and audit logs are reliable.
4. Pilot team can operate the workflow weekly without ad-hoc engineering support.
5. KPI and telemetry support ROI review conversations with customers.
