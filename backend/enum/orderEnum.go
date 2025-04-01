package enum

type OrderEnum int

const (
	_OrderEnum                = iota - 1
	OrderNoPayEnum            // 未支付
	OrderPaidEnum             // 已支付
	OrderPayErrorEnum         // 支付失败
	OrderCancelOrderEnum      // 取消订单
	OrderRefundEnum           // 退款成功
	OrderVerificationEnum     // 已核销
	OrderPendingEnum          // 等待支付中
	OrderApplyRefundEnum      // 退款申请中
	OrderRejectRefundEnum     // 退款关闭(拒绝申请退款)
	OrderFailRefundEnum       // 退款失败(微信端错误)
	OrderPartRefundEnum       // 部分退款成功
	OrderPartApplyRefundEnum  // 部分退款申请中
	OrderPartRejectRefundEnum // 部分退款关闭(拒绝申请退款)
	OrderPartFailRefundEnum   // 部分退款失败(微信端错误)
	OrderPendingRefundEnum    // 等待微信退款回调
)

func GetOrderMessageByStatus(status int64) string {
	orderStatusMap := make(map[int64]string)
	orderStatusMap[OrderNoPayEnum] = "未支付"
	orderStatusMap[OrderPaidEnum] = "已支付"
	orderStatusMap[OrderPayErrorEnum] = "支付失败"
	orderStatusMap[OrderCancelOrderEnum] = "取消订单"
	orderStatusMap[OrderRefundEnum] = "退款成功"
	orderStatusMap[OrderVerificationEnum] = "已核销"
	orderStatusMap[OrderPendingEnum] = "等待支付中"
	orderStatusMap[OrderApplyRefundEnum] = "退款申请中"
	orderStatusMap[OrderRejectRefundEnum] = "退款申请中"
	orderStatusMap[OrderRejectRefundEnum] = "退款关闭(拒绝申请退款)"
	orderStatusMap[OrderFailRefundEnum] = "退款失败(微信端错误)"
	orderStatusMap[OrderPartRefundEnum] = "部分退款成功"
	orderStatusMap[OrderPartApplyRefundEnum] = "部分退款申请中"
	orderStatusMap[OrderPartRejectRefundEnum] = "部分退款关闭(拒绝申请退款)"
	orderStatusMap[OrderPartFailRefundEnum] = "部分退款失败(微信端错误)"
	orderStatusMap[OrderPendingRefundEnum] = "等待微信退款回调"
	return orderStatusMap[status]
}
