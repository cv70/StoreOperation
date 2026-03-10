# Pilot Go-Live Gate Checklist V1

**Date:** 2026-03-10
**Scope:** Go/No-Go decision before pilot launch
**Dependencies:**
- `2026-03-10-mvp-delivery-checklist.md`
- `2026-03-10-pilot-uat-cases.md`
- `2026-03-10-pilot-kpi-baseline-template.md`
- `2026-03-10-pilot-risk-register-template.md`

## Goal

Define strict launch gates for pilot go-live to avoid starting with unstable flows, unclear KPI baselines, or unresolved critical risks.

## Decision Rule

- Go-live allowed only if all `P0` gates pass.
- Any failed `P0` gate means `NO-GO`.
- `P1` gate failures require explicit risk acceptance by pilot sponsor.

## Gate Summary Table

| Gate ID | Gate Name | Priority | Owner | Status (`pass`/`fail`) | Notes |
|---|---|---|---|---|---|
| G-01 | Core Flow Readiness | P0 | Product Lead |  |  |
| G-02 | Role Permission and Escalation | P0 | Eng Lead |  |  |
| G-03 | UAT Critical Cases | P0 | QA Lead |  |  |
| G-04 | Data and Baseline Lock | P0 | Data Lead |  |  |
| G-05 | KPI and Attribution Readiness | P0 | Analytics Lead |  |  |
| G-06 | Telemetry and Audit | P0 | Platform Lead |  |  |
| G-07 | Pilot Ops Readiness | P1 | CS Lead |  |  |
| G-08 | Risk Register Health | P0 | PM + Sponsor |  |  |

## Detailed Gates

### G-01 Core Flow Readiness (P0)

Pass criteria:

1. Store manager can complete one opportunity card from `new` to validation state.
2. Store manager can process one high-risk card with evidence and SLA.
3. Supervisor can replicate one validated action to multiple stores.
4. HQ can publish one strategy decision and see updated scope.

Evidence required:

- recorded walkthrough
- state transition logs

### G-02 Role Permission and Escalation (P0)

Pass criteria:

1. Unauthorized actions blocked for each role.
2. SLA breach triggers automatic escalation.
3. Escalated cards show new owner and escalation level.

Evidence required:

- permission test results
- escalation event logs

### G-03 UAT Critical Cases (P0)

Pass criteria:

1. All critical UAT cases pass.
2. No open Sev-1 defects.
3. Sev-2 defects have workaround and owner/date.

Evidence required:

- UAT result sheet
- defect tracker snapshot

### G-04 Data and Baseline Lock (P0)

Pass criteria:

1. Baseline window locked with minimum required days.
2. Mandatory data sources connected and healthy.
3. Timezone/store calendar alignment verified.

Evidence required:

- baseline lock report
- data quality checklist

### G-05 KPI and Attribution Readiness (P0)

Pass criteria:

1. Core KPI definitions signed off.
2. Attribution method frozen for pilot period.
3. KPI owners assigned and review cadence scheduled.

Evidence required:

- KPI template
- sign-off record

### G-06 Telemetry and Audit (P0)

Pass criteria:

1. Required events are emitted end-to-end in test run.
2. Event payload includes required dimensions.
3. State transitions are audit-complete for sampled cards.

Evidence required:

- telemetry validation report
- audit sample report

### G-07 Pilot Ops Readiness (P1)

Pass criteria:

1. Weekly review deck owner assigned.
2. On-call contact list and escalation tree confirmed.
3. Training completed for store/supervisor/HQ key users.

Evidence required:

- runbook ownership list
- training attendance record

### G-08 Risk Register Health (P0)

Pass criteria:

1. No open `critical` risks.
2. All open `high` risks have mitigation owner and due date.
3. Redline trigger monitors active.

Evidence required:

- latest risk register
- redline monitor status

## Go/No-Go Meeting Template

### Meeting Inputs

1. Gate summary table with statuses.
2. UAT and defect report.
3. Baseline lock and KPI sign-off.
4. Risk register snapshot.

### Decision Output

- decision: `go` / `no-go`
- decision_date:
- decision_owner:
- blockers_if_no_go:
- must_fix_before_recheck:
- recheck_date:

## Exception Handling

Any P1 failure can proceed only if:

1. pilot sponsor signs risk acceptance
2. mitigation owner and due date are set
3. impact on KPI confidence is explicitly noted

P0 failures cannot be waived.

## Post Go-Live Day-1 Checklist

1. Confirm all role logins and homepage render.
2. Verify first action card processed successfully.
3. Verify one risk card escalation path works.
4. Verify telemetry events received.
5. Confirm support channel and on-call owner active.

## Definition of Launch Ready

Pilot is launch-ready only when:

1. End-to-end role flows work with auditability.
2. KPI baseline and attribution logic are locked.
3. Operational governance is staffed and scheduled.
4. No unresolved critical risk remains.
