package constant

const (
	MsgText     = "文本"
	MsgLocation = "位置"
	MsgFace     = "表情"
	MsgSound    = "语音"
	MsgImage    = "图片"
	MsgFile     = "文件"
	MsgVideo    = "视频"
	MsgRelay    = "合并转发消息"
	MsgCustom   = "自定义"
)

const (
	TIMTextElem      = "TIMTextElem"
	TIMLocationElem  = "TIMLocationElem"
	TIMFaceElem      = "TIMFaceElem"
	TIMSoundElem     = "TIMSoundElem"
	TIMCustomElem    = "TIMCustomElem"
	TIMImageElem     = "TIMImageElem"
	TIMFileElem      = "TIMFileElem"
	TIMVideoFileElem = "TIMVideoFileElem"
	TIMRelayElem     = "TIMRelayElem"
)

var TIMElemIDToName = map[string]string{
	TIMTextElem:      MsgText,
	TIMLocationElem:  MsgLocation,
	TIMFaceElem:      MsgFace,
	TIMSoundElem:     MsgSound,
	TIMCustomElem:    MsgCustom,
	TIMImageElem:     MsgImage,
	TIMFileElem:      MsgFile,
	TIMVideoFileElem: MsgVideo,
	TIMRelayElem:     MsgRelay,
}
