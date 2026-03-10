# Pilot KPI Baseline Template V1

**Date:** 2026-03-10
**Scope:** 8-week paid pilot KPI baseline and weekly review
**Dependencies:**
- `2026-03-10-mvp-delivery-checklist.md`
- `2026-03-10-pilot-uat-cases.md`

## Goal

Provide one standard KPI template for:

- baseline setup before pilot starts
- weekly tracking during pilot
- attribution logic for business outcomes
- go/no-go decision for annual conversion

## KPI Framework

Use three metric layers:

1. Business Value Metrics
2. Product Execution Metrics
3. Delivery Health Metrics

## Section A: Pilot Baseline Setup (Week 0)

### A1. Pilot Scope

- customer_name:
- pilot_period:
- stores_in_scope:
- scenarios_in_scope:
- excluded_scenarios:
- business_owner:
- product_owner:

### A2. Baseline Window

- baseline_start_date:
- baseline_end_date:
- baseline_days_count:
- baseline_data_quality_status (`pass`/`fail`):

Rule:

- baseline window should be at least 14 continuous operating days.

### A3. Baseline Data Checklist

- POS data connected (`yes`/`no`)
- inventory data connected (`yes`/`no`)
- procurement data connected (`yes`/`no`)
- camera/IoT evidence available (`yes`/`no`)
- timezone and store calendar aligned (`yes`/`no`)

Exit rule:

- baseline cannot be locked if any mandatory source is `no`.

## Section B: Business Value KPI Baseline

Record per store and aggregated pilot scope.

### B1. Core KPI Table

| KPI | Definition | Baseline Value | Week Target | Owner |
|---|---|---:|---:|---|
| stockout_rate | 重点SKU缺货率 |  |  |  |
| waste_rate | 临期/报损率 |  |  |  |
| critical_response_time | 严重告警平均响应时长 |  |  |  |
| labor_efficiency | 人效指标（客单或订单/工时） |  |  |  |
| campaign_roi | 活动ROI |  |  |  |

### B2. Optional KPI Table

| KPI | Definition | Baseline Value | Week Target | Owner |
|---|---|---:|---:|---|
| queue_delay_time | 高峰排队时长 |  |  |  |
| coldchain_exception_count | 冷链异常次数 |  |  |  |
| equipment_downtime | 关键设备故障时长 |  |  |  |

## Section C: Product Execution KPI Baseline

### C1. Action Loop Metrics

| KPI | Definition | Baseline/Start Value | Target by Week 8 |
|---|---|---:|---:|
| action_acceptance_rate | 建议采纳率 |  |  |
| ontime_completion_rate | 任务按时完成率 |  |  |
| validated_effective_rate | 完成后验证有效率 |  |  |
| avg_time_to_first_action | 首次动作平均耗时 |  |  |
| escalation_rate | 升级率 |  |  |

### C2. Quality Metrics

| KPI | Definition | Baseline/Start Value | Target by Week 8 |
|---|---|---:|---:|
| false_positive_feedback_rate | 误报反馈率 |  |  |
| writeback_completeness_rate | 回写完整率 |  |  |
| evidence_availability_rate | 证据可用率 |  |  |

## Section D: Weekly Review Template (Week 1-8)

### D1. Weekly Snapshot

- week_number:
- week_start:
- week_end:
- stores_active_count:
- key_incidents:
- major_changes:

### D2. KPI Movement

| KPI | Baseline | Last Week | This Week | Delta vs Baseline | Delta vs Last Week |
|---|---:|---:|---:|---:|---:|
| stockout_rate |  |  |  |  |  |
| waste_rate |  |  |  |  |  |
| critical_response_time |  |  |  |  |  |
| labor_efficiency |  |  |  |  |  |
| campaign_roi |  |  |  |  |  |

### D3. Action Funnel

| Stage | Count | Conversion |
|---|---:|---:|
| cards_created |  |  |
| cards_accepted |  |  |
| cards_in_progress |  |  |
| cards_done_pending_validation |  |  |
| cards_validated_effective |  |  |

### D4. Top 5 Effective Actions

| Rank | Scenario | Action | Stores | Estimated Impact | Observed Impact |
|---|---|---|---:|---:|---:|
| 1 |  |  |  |  |  |
| 2 |  |  |  |  |  |
| 3 |  |  |  |  |  |
| 4 |  |  |  |  |  |
| 5 |  |  |  |  |  |

### D5. Top 5 Blockers

| Rank | Blocker | Impacted KPI | Root Cause | Mitigation Owner | ETA |
|---|---|---|---|---|---|
| 1 |  |  |  |  |  |
| 2 |  |  |  |  |  |
| 3 |  |  |  |  |  |
| 4 |  |  |  |  |  |
| 5 |  |  |  |  |  |

## Section E: Attribution Method (Fixed Rules)

Use fixed attribution logic during pilot.

### E1. Attribution Unit

- primary unit: action card
- secondary unit: scenario template

### E2. Attribution Formula (Template)

`Attributed Impact = Observed KPI Change * Attribution Weight`

Weight reference factors:

- execution completion quality
- timing proximity to KPI change
- control group or historical baseline contrast
- confounder adjustment tag

### E3. Attribution Confidence

- High: strong temporal and control consistency
- Medium: partial confounder uncertainty
- Low: weak isolation, directional only

Rule:

- low-confidence attribution cannot be used as sole conversion evidence.

## Section F: Conversion Gate (Week 8)

### F1. Conversion Scorecard

| Dimension | Indicator | Target | Actual | Pass/Fail |
|---|---|---:|---:|---|
| business_outcome | 至少3个核心KPI改善达标 |  |  |  |
| execution_adoption | 采纳率与按时率达标 |  |  |  |
| stability | 误报率与证据可用率达标 |  |  |  |
| scalability | 至少1个动作模板具备跨店复制效果 |  |  |  |

### F2. Annual Conversion Decision

- decision: `convert` / `extend_pilot` / `stop`
- decision_date:
- decision_owner:
- rationale_summary:
- required_next_actions:

## Section G: Redline Rules

Any one of these should trigger escalation review:

1. KPI trend worsens for 2 consecutive weeks in 2+ core scenarios.
2. writeback completeness <80%.
3. evidence availability <90% for high-risk cards.
4. escalation rate remains high with no downward trend by Week 4.
5. data quality fails for 3 consecutive reporting days.

## Section H: Reporting Cadence

- Daily: operational monitor (store/supervisor)
- Weekly: KPI review (all roles)
- Bi-weekly: strategy review (HQ + customer leadership)
- Week 8: conversion committee review

## Section I: Template Usage Checklist

Before Week 1 review:

1. Baseline period locked.
2. KPI definitions signed off by both sides.
3. Attribution rules frozen.
4. Owners assigned for each KPI.
5. Reporting calendar agreed.

Before conversion decision:

1. Week 1-8 data complete.
2. KPI deltas validated.
3. Attribution confidence labeled.
4. Blockers and mitigations documented.
5. Conversion scorecard completed.
