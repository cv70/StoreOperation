# Interaction Copy Spec V1

**Date:** 2026-03-10
**Scope:** Action card copy for store manager, supervisor, HQ
**Dependencies:**
- `2026-03-10-action-card-component-spec.md`
- `2026-03-10-homepage-flow-spec.md`

## Goal

Standardize action-card interaction copy so users can make fast decisions with minimal ambiguity.

This spec defines:

- copy structure
- sentence templates
- role-specific wording
- confidence and uncertainty language
- warning and escalation wording

## Copy Principles

1. One card communicates one clear decision.
2. Use action-first language, not analysis-first language.
3. Every recommendation must include reason and expected impact.
4. Avoid model-centric wording such as "AI thinks".
5. Risk copy must be evidence-first and time-sensitive.
6. Keep copy concise and scannable on mobile.

## Global Copy Structure

Each card uses this top-down order:

1. `Title`
2. `What happened`
3. `Why now`
4. `What to do`
5. `Expected impact`
6. `Deadline/SLA`
7. `Primary actions`

## Field-Level Copy Templates

### 1. Title

Template:

`[对象/场景] + [问题/机会] + [动作方向]`

Examples:

- `招牌套餐缺货风险，建议15分钟内补货`
- `冷链温度异常，需立即排查并复位`
- `临期库存积压，建议组合促销处理`

Rules:

- Max 24 Chinese chars on homepage card
- Must include verb or actionable direction
- Avoid abstract labels like "经营优化建议"

### 2. What Happened

Template:

`过去[时间窗口]，[对象]出现[可观测现象]。`

Examples:

- `过去2小时，A类套餐销量高于基线42%。`
- `过去30分钟，冷柜温度持续高于阈值2.8°C。`

Rules:

- Include time window and observable evidence
- No speculative interpretation in this line

### 3. Why Now

Template:

`当前触发原因是[关键因子1] + [关键因子2]。`

Examples:

- `当前触发原因是午市客流上升 + 安全库存偏低。`
- `当前触发原因是设备波动 + 门店客流高峰临近。`

Rules:

- Max 2 factors in first-screen summary
- Detailed analysis moves to detail panel

### 4. What To Do

Template:

`请在[时限]内完成[动作]，并[结果回写动作]。`

Examples:

- `请在15分钟内补货20份，并回写实际补货量。`
- `请在30分钟内完成设备排查，并回写复位时间。`

Rules:

- Must specify concrete operator verb
- Must include time bound
- Must include writeback requirement if needed

### 5. Expected Impact

Template:

`预计可[收益/避免损失][数值]，置信度[等级或分值]。`

Examples:

- `预计可减少损失¥680，置信度78。`
- `预计可提升毛利¥1,200，置信度高。`

Rules:

- Show currency or metric delta, not vague benefits
- Confidence must always appear with impact

### 6. Deadline/SLA

Template:

`最晚处理时间：[时间点]` or `SLA剩余：[倒计时]`

Examples:

- `最晚处理时间：14:30`
- `SLA剩余：12分钟`

Rules:

- Risk cards default to countdown format
- Opportunity cards default to latest-time format

## Confidence Language Standard

### Numeric to Label Mapping

- `85-100`: 高
- `70-84`: 中高
- `55-69`: 中
- `<55`: 低

### Required Wording by Confidence

High:

- `建议优先执行`

Medium:

- `建议执行，并留意现场差异`

Low:

- `建议人工复核后执行`

Guardrail:

- If confidence <55, primary button text cannot be "立即执行"; use "复核后执行".

## Role-Specific Copy Tone

### Store Manager Tone

Objective:

- immediate execution clarity

Style:

- short imperative sentences
- explicit time and benefit

Preferred verbs:

- 补货
- 处理
- 调整
- 确认
- 回写

Avoid:

- strategy-level wording
- long explanation paragraphs

### Supervisor Tone

Objective:

- intervention and replication decisions

Style:

- comparison and prioritization language

Preferred verbs:

- 复制
- 催办
- 升级
- 重分配
- 跟踪

Avoid:

- frontline micro-step wording

### HQ Tone

Objective:

- scale/adjust/stop strategy decisions

Style:

- decision-grade summary language

Preferred verbs:

- 扩大
- 维持
- 调整
- 暂停
- 下线

Avoid:

- incident-level handling details

## Button Copy Standard

### Store Manager Buttons

- `立即执行`
- `稍后处理`
- `不采纳`
- `标记完成`
- `请求协助`

### Supervisor Buttons

- `复制到多店`
- `催办`
- `升级处理`
- `重分配`
- `加入重点跟踪`

### HQ Buttons

- `扩大范围`
- `维持观察`
- `调整策略`
- `暂停策略`

Rules:

- button labels must start with verb
- max 6 Chinese chars preferred
- avoid ambiguous text like `确认` as primary action

## Rejection Reason Copy Taxonomy

For opportunity cards:

- `当前不具备执行条件`
- `门店实际情况不匹配`
- `预估收益可信度不足`
- `依赖项未准备完成`

For risk cards:

- `已线下处理，待补证据`
- `事件识别疑似误报`
- `已转交处理责任人`

Rules:

- reason selection required
- optional free text for detail
- reason must feed model feedback loop

## Escalation Copy Standard

### Escalation Trigger Message

Template:

`该事项已超过处理时限，系统已升级至[角色/人员]。`

### Escalation Status Tag

- `已升级-L1`
- `已升级-L2`
- `已升级-L3`

### Escalation Guidance

Template:

`请在[时限]内完成处理或说明阻塞原因。`

## Error and Empty-State Copy

### No Action Cards

Store manager:

- `当前无高优先动作，系统将持续监测新的利润机会和风险。`

Supervisor:

- `当前无待介入高优事项，可优先查看复制动作机会。`

HQ:

- `当前无异常策略，建议查看可扩复制区域清单。`

### Missing Evidence

Template:

`证据数据暂不可用，请先按标准流程处理并补充现场记录。`

### Validation Pending

Template:

`动作已完成，正在等待结果验证（预计[时间]完成）。`

## Mobile Copy Constraints

1. Title <= 18 chars preferred
2. Each summary line <= 26 chars preferred
3. Show max 3 summary lines before expand
4. Primary button text <= 4 chars preferred

## Localization and Terminology

Use these fixed terms consistently:

- `利润机会` (not "收益点")
- `高危风险` (not "高优警报")
- `执行回写` (not "反馈登记")
- `策略下线` (not "策略关闭")
- `复制动作` (not "同步动作")

## QA Checklist for Copy Review

1. Can a user decide the next step in under 5 seconds?
2. Does each card include reason + action + impact + time?
3. Are impact and confidence always paired?
4. Are risk cards evidence-first and SLA-visible?
5. Are buttons verb-led and unambiguous?
6. Is role wording consistent with role responsibility?

## Release Recommendation

1. Apply this copy spec to top 5 opportunity templates.
2. Apply to high-severity risk templates.
3. Validate with store manager and supervisor usability sessions.
4. Tune confidence wording thresholds based on rejection reasons.
