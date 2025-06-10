package constant

const (
	Dji_Topic_Thing_Product_dsn_osd   = "thing/product/%s/osd"   // 设备定频上报设备属性
	Dji_Topic_Thing_Product_dsn_state = "thing/product/%s/state" // 设备按需上报设备属性

	Dji_Topic_Thing_product_gsn_services      = "thing/product/%s/services"       // 向设备发送的服务
	Dji_Topic_Thing_product_gsn_servicesReply = "thing/product/%s/services_reply" // 设备对服务的回复、处理结果

	Dji_Topic_Thing_product_gsn_events      = "thing/product/%s/events"       // 设备端向云平台发送的，需要关注和处理的事件
	Dji_Topic_Thing_product_gsn_eventsReply = "thing/product/%s/events_reply" // 云平台对设备端事件的回复、处理结果

	Dji_Topic_Thing_product_gsn_requests      = "thing/product/%s/requests"       // 设备端向云平台发送的请求，为了获取一些信息，比如上传的临时凭证
	Dji_Topic_Thing_product_gsn_requestsReply = "thing/product/%s/requests_reply" // 云平台对设备端请求的回复

	Dji_Topic_Thing_product_gsn_status      = "thing/product/%s/status"       // 设备上下线、更新拓扑
	Dji_Topic_Thing_product_gsn_statusReply = "thing/product/%s/status_reply" // 云平台对设备状态的回复

	Dji_Topic_Thing_product_gsn_property_set      = "thing/product/%s/property/set"       // 设备属性设置
	Dji_Topic_Thing_product_gsn_property_setReply = "thing/product/%s/property/set_reply" // 设备属性设置的响应

	Dji_Topic_Thing_product_gsn_drc_up   = "thing/product/%s/drc/up"   // DRC 协议上行
	Dji_Topic_Thing_product_gsn_drc_down = "thing/product/%s/drc/down" // DRC 协议下行
)
