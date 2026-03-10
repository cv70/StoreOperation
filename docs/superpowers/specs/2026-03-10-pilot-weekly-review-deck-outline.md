# Pilot Weekly Review Deck Outline V1

**Date:** 2026-03-10
**Scope:** Weekly pilot review meeting deck (Week 1-8)
**Dependencies:**
- `2026-03-10-pilot-kpi-baseline-template.md`
- `2026-03-10-mvp-delivery-checklist.md`
- `2026-03-10-pilot-uat-cases.md`

## Goal

Provide a fixed slide-by-slide structure for weekly pilot review so meetings stay outcome-driven and comparable week over week.

## Audience

- customer business owner
- HQ operations leader
- supervisor representative
- product/CS lead
- data analyst

## Deck Rules

1. Keep total slides between 10 and 14.
2. One core message per slide.
3. Every KPI slide must show baseline, last week, this week, and delta.
4. Every conclusion slide must include owner and next action.
5. Avoid feature updates unless linked to KPI movement.

## Slide Outline

### Slide 1: Executive Snapshot

Purpose:

- summarize weekly outcome in 30 seconds

Content:

- pilot week number
- overall status (`on_track`/`at_risk`/`off_track`)
- top 3 wins
- top 3 risks
- conversion confidence trend

Template:

`本周结论：<一句话>`

### Slide 2: Scope and Data Health

Purpose:

- confirm this week data is reliable enough for decisions

Content:

- active stores count
- scenario coverage
- data source health (POS/inventory/evidence)
- missing data notes

Decision Gate:

- if data quality fail, flag KPI confidence downgrade

### Slide 3: Core KPI Movement

Purpose:

- show business outcome trend for top KPIs

Content table:

| KPI | Baseline | Last Week | This Week | Delta vs Baseline | Delta vs Last Week |
|---|---:|---:|---:|---:|---:|
| stockout_rate |  |  |  |  |  |
| waste_rate |  |  |  |  |  |
| critical_response_time |  |  |  |  |  |
| labor_efficiency |  |  |  |  |  |
| campaign_roi |  |  |  |  |  |

### Slide 4: Value Attribution Summary

Purpose:

- separate observed change from attributed impact

Content:

- attributed value this week
- attributed value cumulative
- confidence split (`high`/`medium`/`low`)
- top confounders

Template:

`本周归因收益：¥X（高置信度占比Y%）`

### Slide 5: Action Funnel Performance

Purpose:

- evaluate execution engine effectiveness

Content:

- cards created
- accepted
- in progress
- done pending validation
- validated effective

Key ratios:

- acceptance rate
- on-time completion rate
- validated-effective rate

### Slide 6: Top Effective Actions

Purpose:

- show what worked and why

Content table:

| Rank | Scenario | Action | Stores | Estimated Impact | Observed Impact | Confidence |
|---|---|---|---:|---:|---:|---|
| 1 |  |  |  |  |  |  |
| 2 |  |  |  |  |  |  |
| 3 |  |  |  |  |  |  |

### Slide 7: Risk and Escalation Review

Purpose:

- ensure high-risk handling is under control

Content:

- high-severity incidents count
- SLA compliance rate
- escalated card count and trend
- unresolved high-risk items

Decision:

- confirm immediate containment actions

### Slide 8: Supervisor Replication Progress

Purpose:

- measure cross-store scaling quality

Content:

- replicated actions count
- replication adoption rate
- replication success rate
- top replication blockers

### Slide 9: Product/Model Quality Signals

Purpose:

- identify system quality issues affecting trust and adoption

Content:

- false-positive feedback rate
- evidence availability rate
- writeback completeness rate
- low-confidence card handling quality

### Slide 10: Customer Feedback and Adoption

Purpose:

- include qualitative signals from frontline and operations

Content:

- store manager feedback themes
- supervisor feedback themes
- friction points
- quick wins implemented

### Slide 11: Decisions Required This Week

Purpose:

- force explicit management decisions

Content:

- decision item
- options
- recommendation
- decision owner
- deadline

Format:

| Decision | Options | Recommendation | Owner | Due Date |
|---|---|---|---|---|

### Slide 12: Next Week Execution Plan

Purpose:

- align actions and accountability

Content:

- top 5 next-week actions
- owner and ETA
- expected KPI impact
- dependency/risk

## Optional Slides (Use Only If Needed)

### Optional A: Scenario Deep Dive

Use when one scenario dominates impact.

### Optional B: Store Cluster Comparison

Use when performance diverges by region/store type.

### Optional C: Conversion Readiness Tracker

Use from Week 6 onward.

## Meeting Runbook

### Recommended Time Allocation (60 minutes)

1. Executive snapshot and data health: 10 min
2. KPI and attribution: 15 min
3. execution/risk/replication: 20 min
4. decisions and next week plan: 15 min

### Role Responsibilities in Meeting

- business owner: approve weekly priorities
- operations leader: commit execution changes
- supervisor: confirm field feasibility
- product/CS: track actions and unblock
- data analyst: defend metric and attribution integrity

## Slide Quality Checklist

Before presenting, verify:

1. Numbers match KPI baseline template.
2. All deltas have clear sign (+/-).
3. Attribution confidence is explicitly labeled.
4. Every risk has an owner and ETA.
5. Every decision item has due date.
6. Deck ends with concrete next-week plan.

## Week-by-Week Emphasis

### Week 1-2

- baseline stability
- adoption and execution hygiene

### Week 3-4

- risk control stability
- first attributable value evidence

### Week 5-6

- replication quality
- strategy adjustment signals

### Week 7-8

- conversion scorecard readiness
- annual conversion recommendation

## Definition of Good Weekly Deck

A weekly deck is good when:

1. Decision makers can state if pilot is on track in under 2 minutes.
2. KPI movement and attribution confidence are both visible.
3. Field execution issues are translated into owned actions.
4. Next week plan has measurable expected impact.
