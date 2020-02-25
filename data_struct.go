package main

type TpData struct {
	V uint32 `json:"v,omitempty"`
}
type PowerData struct {
	V float32 `json:"v,omitempty"`
}

type MsgId = int
const (
	UNKNOWN MsgId = iota
	LOGIN
	LOGOUT
	DP
	CMD_REPLY
	IMAGE_UPDATE
	IMAGE_GET
	LOGIN_ACCEPTED_RESPONSE
	LOGIN_REJECTED_RESPONSE
	LOGOUT_ACCEPTED_RESPONSE
	LOGOUT_NOTIFY_RESPONSE
	DP_ACCEPTED_RESPONSE
	DP_REJECTED_RESPONSE
	DOWN_LINK_CMD
	CMD_REPLY_ACCEPTED_RESPONSE
	CMD_REPLY_REJECTED_RESPONSE
	IMAGE_UPDATE_ACCEPTED_RESPONSE
	IMAGE_UPDATE_REJECTED_RESPONSE
	IMAGE_GET_ACCEPTED_RESPONSE
	IMAGE_GET_REJECTED_RESPONSE
	IMAGE_DELTA
)
type MessageType struct {
	V MsgId `json:"v,omitempty"`
}

type DpType struct {
	Tp      []TpData `json:"temperatrue,omitempty"`
	Power   []PowerData `json:"power,omitempty"`
	MsgType []MessageType `json:"MessageType,omitempty"`
}
type date struct {
	Id uint32 `json:"id"`
	Dp DpType `json:"dp,omitempty"`
}


type TopicType = string

const(
	//TopicPublicSection = "$sys/{pid}/{device-name}/%s"
	TopicPublicSection = "$sys/%s/%s/%s"
	//Subscribe
	TopicSubDpAcc    TopicType = "dp/post/json/accepted"
	TopicSubDpRej    TopicType = "dp/post/json/rejected"
	TopicSubDpAllEv  TopicType = "dp/post/json/+"
	TopicSubCmdReq   TopicType = "cmd/request/+"
	TopicSubCmdRsp   TopicType = "cmd/response/+/+"
	TopicSubCmdAllEv TopicType = "cmd/#"
	TopicSubAllEv    TopicType = "#"
	//Data
	TopicDpUplink TopicType = "dp/post/json"
	TopicDpAcc    TopicType = "dp/post/json/accepted"
	TopicDpRej    TopicType = "dp/post/json/rejected"
	//CMD
	TopicCmdReq TopicType = "cmd/request/{cmdid}"
	TopicCmdRsp TopicType = "cmd/response/{cmdid}"
	TopicCmdAcc TopicType = "cmd/response/{cmdid}/accepted"
	TopicCmdRej TopicType = "cmd/response/{cmdid}/rejected"
	//Image
	TopicImageUpdate      TopicType = "image/update"
	TopicImageUpdateAcc   TopicType = "image/update/accepted"
	TopicImageUpdateRej   TopicType = "image/update/rejected"
	TopicImageUpdateDelta TopicType = "image/update/delta"
	TopicImageGet         TopicType = "image/get"
	TopicImageGetAcc      TopicType = "image/get/accepted"
	TopicImageGetRej      TopicType = "image/get/rejected"
)