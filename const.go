package yunda

// OrderType order_type 可以为空，为空时默认值为 common
type OrderType string

var (
	// OrderTypeCommon 普通
	OrderTypeCommon OrderType = "common"
	// OrderTypeSupportValue 保价
	OrderTypeSupportValue OrderType = "support_value"
	// OrderTypeCod COD
	OrderTypeCod OrderType = "cod"
	// OrderTypeDf 到付
	OrderTypeDf OrderType = "df"
)

// Request 定义字典表 说明:大客户订单系统(A)与韵达订单平台(B)
type Request string

var (
	// RequestData A->B 下单到韵达平台
	RequestData Request = "data"

	// RequestCancelOrder A->B 取消订单
	RequestCancelOrder Request = "cancel_order"

	// RequestAcceptOrder A->B 韵达向客户方推送是否接单
	RequestAcceptOrder Request = "accept_order"
)
