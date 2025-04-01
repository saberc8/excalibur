package enum

//0表示系统通知,1表示消息通知,2表示组队消息

const (
	SystemMessageEnum  = iota // 系统通知
	InfoMessageEnum           //  消息通知
	PairingMessageEnum        //  组对消息
)
