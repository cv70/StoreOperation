# 数据字典 V1

## 1. 文档目标

这份数据字典不追求覆盖未来所有能力，只服务当前 V1 MVP：

- 智能补货防缺货
- 临期库存处置
- 高危告警闭环
- 高峰排班优化
- 活动 ROI 控盘

设计原则：

1. 先保证关键口径可信
2. 先支撑试点和 ROI 归因
3. 先满足店长、督导、总部三类角色的最小使用需求
4. 先保留智能建议和预测的可解释输入

---

## 2. V1 数据范围

V1 只要求八类核心数据：

1. 门店主数据
2. SKU 主数据
3. 员工主数据
4. 销售与库存事实
5. 建议、告警、任务事实
6. ROI 与日级经营聚合指标
7. 设备主数据与设备事件
8. 机器人链路事件

暂不作为 V1 必需：

- 会员全量标签
- 视频原始流长期存储
- 复杂设备遥测全量历史
- 复杂供应商结算数据

---

## 3. 主数据

## 3.1 `dim_store`

粒度：门店

| 字段名 | 类型 | 必填 | 说明 |
|---|---|---|---|
| store_id | string | 是 | 门店唯一 ID |
| store_code | string | 是 | 门店编码 |
| store_name | string | 是 | 门店名称 |
| brand_id | string | 是 | 品牌 ID |
| region_id | string | 是 | 区域 ID |
| store_type | string | 否 | 门店类型 |
| open_date | date | 否 | 开业日期 |
| status | string | 是 | `open/closed/suspended` |
| created_time | datetime | 是 | 创建时间 |
| updated_time | datetime | 是 | 更新时间 |

## 3.2 `dim_sku`

粒度：SKU

| 字段名 | 类型 | 必填 | 说明 |
|---|---|---|---|
| sku_id | string | 是 | SKU 唯一 ID |
| sku_code | string | 是 | SKU 编码 |
| sku_name | string | 是 | SKU 名称 |
| category_lv1 | string | 否 | 一级品类 |
| category_lv2 | string | 否 | 二级品类 |
| unit | string | 是 | 单位 |
| shelf_life_days | int | 否 | 保质期天数 |
| min_order_qty | decimal | 否 | 最小起订量 |
| enabled_flag | tinyint | 是 | 是否启用 |
| created_time | datetime | 是 | 创建时间 |
| updated_time | datetime | 是 | 更新时间 |

## 3.3 `dim_employee`

粒度：员工

| 字段名 | 类型 | 必填 | 说明 |
|---|---|---|---|
| employee_id | string | 是 | 员工 ID |
| employee_no | string | 否 | 工号 |
| employee_name | string | 是 | 姓名 |
| store_id | string | 是 | 所属门店 |
| role_code | string | 是 | `manager/cashier/cook/supervisor` |
| status | string | 是 | `active/inactive` |
| hire_date | date | 否 | 入职日期 |
| created_time | datetime | 是 | 创建时间 |
| updated_time | datetime | 是 | 更新时间 |

## 3.4 `dim_device`

粒度：设备

| 字段名 | 类型 | 必填 | 说明 |
|---|---|---|---|
| device_id | string | 是 | 设备 ID |
| store_id | string | 是 | 门店 ID |
| device_type | string | 是 | `camera/iot/robot/edge_box` |
| device_code | string | 是 | 设备编码 |
| location_code | string | 否 | 区域位置 |
| status | string | 是 | `online/offline/fault` |
| vendor_name | string | 否 | 厂商 |
| installed_time | datetime | 否 | 安装时间 |
| created_time | datetime | 是 | 创建时间 |
| updated_time | datetime | 是 | 更新时间 |

---

## 4. 经营事实层

## 4.1 `fact_sales_order_item`

粒度：订单行

| 字段名 | 类型 | 必填 | 说明 |
|---|---|---|---|
| order_id | string | 是 | 订单 ID |
| order_item_id | string | 是 | 订单行 ID |
| store_id | string | 是 | 门店 ID |
| sku_id | string | 是 | SKU ID |
| channel | string | 是 | `dine_in/takeaway/delivery` |
| quantity | decimal | 是 | 销售数量 |
| unit_price | decimal | 是 | 标准单价 |
| sale_amount | decimal | 是 | 实收金额 |
| discount_amount | decimal | 否 | 折扣金额 |
| refund_amount | decimal | 否 | 退款金额 |
| order_status | string | 是 | `paid/refunded/canceled` |
| cashier_employee_id | string | 否 | 收银员工 ID |
| order_time | datetime | 是 | 下单时间 |
| paid_time | datetime | 否 | 支付时间 |
| business_date | date | 是 | 营业日期 |
| created_time | datetime | 是 | 入库时间 |

## 4.2 `fact_inventory_snapshot`

粒度：门店 + SKU + 快照时间

| 字段名 | 类型 | 必填 | 说明 |
|---|---|---|---|
| snapshot_id | string | 是 | 快照 ID |
| store_id | string | 是 | 门店 ID |
| sku_id | string | 是 | SKU ID |
| on_hand_qty | decimal | 是 | 在手库存 |
| available_qty | decimal | 是 | 可用库存 |
| in_transit_qty | decimal | 否 | 在途库存 |
| expiry_date | date | 否 | 到期日 |
| snapshot_time | datetime | 是 | 快照时间 |
| business_date | date | 是 | 营业日期 |
| created_time | datetime | 是 | 入库时间 |

## 4.3 `fact_purchase_order_item`

粒度：采购单行

| 字段名 | 类型 | 必填 | 说明 |
|---|---|---|---|
| po_id | string | 是 | 采购单 ID |
| po_item_id | string | 是 | 采购单行 ID |
| store_id | string | 是 | 门店 ID |
| sku_id | string | 是 | SKU ID |
| supplier_id | string | 否 | 供应商 ID |
| suggested_qty | decimal | 否 | 系统建议量 |
| ordered_qty | decimal | 是 | 实际下单量 |
| received_qty | decimal | 否 | 实收量 |
| unit_cost | decimal | 否 | 采购单价 |
| po_status | string | 是 | `created/ordered/received/canceled` |
| expected_arrival_time | datetime | 否 | 预计到货时间 |
| actual_arrival_time | datetime | 否 | 实际到货时间 |
| created_time | datetime | 是 | 创建时间 |
| updated_time | datetime | 是 | 更新时间 |

## 4.4 `fact_loss_report`

粒度：报损记录

| 字段名 | 类型 | 必填 | 说明 |
|---|---|---|---|
| loss_id | string | 是 | 报损 ID |
| store_id | string | 是 | 门店 ID |
| sku_id | string | 是 | SKU ID |
| loss_qty | decimal | 是 | 报损数量 |
| loss_amount | decimal | 是 | 报损金额 |
| loss_reason | string | 否 | 报损原因 |
| reported_by | string | 否 | 提报人 |
| report_time | datetime | 是 | 提报时间 |
| business_date | date | 是 | 营业日期 |

## 4.5 `fact_shift_plan`

粒度：门店 + 班次 + 日期

| 字段名 | 类型 | 必填 | 说明 |
|---|---|---|---|
| shift_id | string | 是 | 班次 ID |
| store_id | string | 是 | 门店 ID |
| business_date | date | 是 | 营业日期 |
| shift_slot | string | 是 | 如 `breakfast/lunch/dinner` |
| planned_headcount | int | 是 | 计划人数 |
| actual_headcount | int | 否 | 实际人数 |
| predicted_demand_index | decimal | 否 | 需求指数 |
| labor_cost_amount | decimal | 否 | 班次人工成本 |
| created_time | datetime | 是 | 创建时间 |
| updated_time | datetime | 是 | 更新时间 |

## 4.6 `fact_campaign_daily`

粒度：门店 + 活动 + 日

| 字段名 | 类型 | 必填 | 说明 |
|---|---|---|---|
| campaign_id | string | 是 | 活动 ID |
| store_id | string | 是 | 门店 ID |
| business_date | date | 是 | 营业日期 |
| coupon_cost | decimal | 否 | 优惠成本 |
| campaign_sales_amount | decimal | 否 | 活动销售额 |
| campaign_orders | int | 否 | 活动订单数 |
| incremental_sales_amount | decimal | 否 | 归因增量销售额 |
| campaign_roi | decimal | 否 | 活动 ROI |
| created_time | datetime | 是 | 创建时间 |
| updated_time | datetime | 是 | 更新时间 |

## 4.7 `fact_device_event`

粒度：设备事件

| 字段名 | 类型 | 必填 | 说明 |
|---|---|---|---|
| device_event_id | string | 是 | 设备事件 ID |
| device_id | string | 是 | 设备 ID |
| store_id | string | 是 | 门店 ID |
| event_type | string | 是 | `temperature_alert/offline/power_abnormal/door_open` |
| severity | string | 是 | `low/medium/high/critical` |
| event_value | string | 否 | 事件值 |
| event_time | datetime | 是 | 事件时间 |
| created_time | datetime | 是 | 创建时间 |

## 4.8 `fact_robot_delivery_event`

粒度：机器人任务事件

| 字段名 | 类型 | 必填 | 说明 |
|---|---|---|---|
| robot_event_id | string | 是 | 机器人事件 ID |
| device_id | string | 是 | 机器人设备 ID |
| store_id | string | 是 | 门店 ID |
| order_id | string | 否 | 关联订单 ID |
| task_stage | string | 是 | `pickup/delivery/arrival/complete/failed` |
| start_time | datetime | 否 | 开始时间 |
| end_time | datetime | 否 | 结束时间 |
| duration_seconds | int | 否 | 持续时长 |
| target_table_code | string | 否 | 目标桌台 |
| result_status | string | 否 | 任务结果 |
| created_time | datetime | 是 | 创建时间 |

---

## 5. 动作闭环层

## 5.1 `fact_recommendation`

粒度：建议

| 字段名 | 类型 | 必填 | 说明 |
|---|---|---|---|
| rec_id | string | 是 | 建议 ID |
| rec_type | string | 是 | `replenishment/expiry_disposal/staffing/promotion/process_optimization/device_maintenance` |
| store_id | string | 是 | 门店 ID |
| business_date | date | 是 | 业务日期 |
| title | string | 是 | 建议标题 |
| reason_json | json | 是 | 触发原因 |
| action_json | json | 是 | 建议动作 |
| expected_impact_json | json | 否 | 预估收益 |
| model_version | string | 否 | 建议模型或规则版本 |
| confidence_score | decimal | 否 | 建议置信度 |
| status | string | 是 | `new/adopted/rejected/expired` |
| adopted_by | string | 否 | 采纳人 |
| adopted_time | datetime | 否 | 采纳时间 |
| feedback_note | string | 否 | 反馈说明 |
| created_time | datetime | 是 | 创建时间 |
| updated_time | datetime | 是 | 更新时间 |

## 5.2 `fact_alert_event`

粒度：告警事件

| 字段名 | 类型 | 必填 | 说明 |
|---|---|---|---|
| alert_id | string | 是 | 告警 ID |
| alert_type | string | 是 | `inventory/operation/security/device/robot` |
| alert_code | string | 是 | 规则编码 |
| severity | string | 是 | `low/medium/high/critical` |
| store_id | string | 是 | 门店 ID |
| employee_id | string | 否 | 关联员工 |
| sku_id | string | 否 | 关联 SKU |
| trigger_time | datetime | 是 | 触发时间 |
| first_notify_time | datetime | 否 | 首次通知时间 |
| evidence_json | json | 否 | 证据链 |
| model_version | string | 否 | 识别模型或规则版本 |
| status | string | 是 | `open/processing/closed/false_positive` |
| assignee_id | string | 否 | 处理人 |
| close_time | datetime | 否 | 关闭时间 |
| close_reason | string | 否 | 关闭原因 |
| created_time | datetime | 是 | 创建时间 |
| updated_time | datetime | 是 | 更新时间 |

## 5.3 `fact_task`

粒度：任务

| 字段名 | 类型 | 必填 | 说明 |
|---|---|---|---|
| task_id | string | 是 | 任务 ID |
| task_type | string | 是 | `alert/recommendation/manual` |
| source_id | string | 是 | 来源 ID |
| source_type | string | 是 | `alert/recommendation/manual` |
| store_id | string | 是 | 门店 ID |
| owner_id | string | 是 | 责任人 ID |
| status | string | 是 | `todo/in_progress/done/overdue/canceled` |
| priority | string | 是 | `p1/p2/p3` |
| due_time | datetime | 否 | 截止时间 |
| start_time | datetime | 否 | 开始时间 |
| done_time | datetime | 否 | 完成时间 |
| result_note | string | 否 | 处理结果 |
| created_time | datetime | 是 | 创建时间 |
| updated_time | datetime | 是 | 更新时间 |

## 5.4 `fact_action_outcome`

粒度：动作结果

| 字段名 | 类型 | 必填 | 说明 |
|---|---|---|---|
| outcome_id | string | 是 | 结果 ID |
| source_type | string | 是 | `recommendation/task` |
| source_id | string | 是 | 来源 ID |
| store_id | string | 是 | 门店 ID |
| business_date | date | 是 | 营业日期 |
| metric_code | string | 是 | 如 `stockout_rate/loss_rate/campaign_roi` |
| baseline_value | decimal | 否 | 基线值 |
| current_value | decimal | 否 | 当前值 |
| improvement_value | decimal | 否 | 改善值 |
| attributed_profit_amount | decimal | 否 | 归因收益 |
| created_time | datetime | 是 | 创建时间 |

## 5.5 `fact_prediction_result`

粒度：预测结果

| 字段名 | 类型 | 必填 | 说明 |
|---|---|---|---|
| prediction_id | string | 是 | 预测 ID |
| prediction_type | string | 是 | `demand/expiry/staffing/campaign_roi` |
| store_id | string | 是 | 门店 ID |
| target_id | string | 否 | 目标对象 ID，如 SKU、活动、班次 |
| business_date | date | 是 | 业务日期 |
| predicted_value | decimal | 是 | 预测值 |
| actual_value | decimal | 否 | 实际值 |
| error_value | decimal | 否 | 偏差值 |
| model_version | string | 否 | 模型版本 |
| feature_snapshot_json | json | 否 | 关键特征快照 |
| created_time | datetime | 是 | 创建时间 |

---

## 6. 聚合层

## 6.1 `ads_store_daily_kpi`

粒度：门店日

| 字段名 | 类型 | 说明 |
|---|---|---|
| business_date | date | 营业日期 |
| store_id | string | 门店 ID |
| sales_amount | decimal | 销售额 |
| gross_profit_amount | decimal | 毛利额 |
| stockout_rate | decimal | 缺货率 |
| loss_rate | decimal | 报损率 |
| critical_alert_count | int | 严重告警数 |
| critical_alert_close_rate | decimal | 严重告警闭环率 |
| critical_alert_response_minutes | decimal | 严重告警响应时长 |
| staffing_efficiency_index | decimal | 排班效率指数 |
| campaign_roi | decimal | 活动 ROI |
| recommendation_adopt_rate | decimal | 建议采纳率 |
| task_ontime_rate | decimal | 任务按时完成率 |
| attributable_profit_amount | decimal | 归因利润改善额 |
| site_event_count | int | 现场事件数 |
| device_online_rate | decimal | 设备在线率 |
| robot_delivery_avg_seconds | decimal | 机器人平均送餐时长 |

## 6.2 `ads_region_weekly_roi`

粒度：区域周

| 字段名 | 类型 | 说明 |
|---|---|---|
| business_week | string | 业务周 |
| region_id | string | 区域 ID |
| pilot_store_count | int | 试点门店数 |
| stockout_profit_recovery | decimal | 缺货回补收益 |
| loss_reduction_profit | decimal | 报损降低收益 |
| staffing_profit_gain | decimal | 排班收益 |
| campaign_profit_gain | decimal | 活动收益 |
| risk_loss_reduction | decimal | 风险损失减少 |
| total_attributable_profit | decimal | 总归因收益 |
| project_cost_amount | decimal | 项目成本 |
| roi_value | decimal | ROI |

---

## 7. 核心口径

1. 缺货率 = 缺货 SKU 数 / 在售 SKU 总数
2. 报损率 = 报损金额 / 领用金额
3. 严重告警响应时长 = 首次处理时间 - 首次通知时间
4. 建议采纳率 = 已采纳建议数 / 总建议数
5. 任务按时完成率 = 按时完成任务数 / 已到期任务数
6. 归因利润改善额 = 增量毛利 + 降本收益 + 风险损失减少
7. 预测偏差 = |预测值 - 实际值| / 实际值
8. 设备在线率 = 在线设备数 / 应在线设备数

---

## 8. 刷新频率与 SLA

| 数据域 | 刷新频率 | 延迟 SLA | 责任团队 |
|---|---|---|---|
| POS 销售 | 实时或 5 分钟 | <= 5 分钟 | 数据平台 |
| 库存快照 | 15 分钟 | <= 15 分钟 | 供应链系统 |
| 采购订单 | 15 分钟 | <= 15 分钟 | 供应链系统 |
| 建议与任务 | 实时 | <= 1 分钟 | 运营引擎 |
| 告警事件 | 实时 | <= 1 分钟 | 告警服务 |
| 日级 KPI | 每日 | T+1 08:00 | 数仓团队 |
| 周度 ROI | 每周 | T+1 12:00 | 数据分析 |

---

## 9. 数据质量规则

1. 核心主键重复率必须为 0
2. 关键指标字段空值率 < 1%
3. 金额、数量不得出现非法负值
4. 付款时间不得早于下单时间
5. 销售变化与库存扣减趋势异常时必须触发校验

---

## 10. 权限与审计

1. 数据访问按品牌、区域、门店隔离
2. 建议采纳、任务完成、规则变更必须留审计日志
3. 员工信息按最小权限原则开放
4. ROI 相关指标需保留口径版本号，避免试点中途口径漂移
