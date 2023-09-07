package constant

const (
	PullSingleChatHistoryMsgApi = "v4/openim/admin_getroammsg"
	PullGroupChatHistoryMsgApi  = "v4/group_open_http_svc/group_msg_get_simple"
	PullFriendListApi           = "v4/sns/friend_get"     //拉取用户好友列表接口
	AddFriendApi                = "v4/sns/friend_add"     //添加用户好友接口
	DeleteFriendApi             = "v4/sns/friend_delete"  //删除用户好友接口
	AddBlackListApi             = "v4/sns/black_list_add" //拉黑用户好友接口
	SendMsgToAllUserApi         = "v4/all_member_push/im_push"
	SendMsgToBatchUserApi       = "v4/openim/batchsendmsg"
	AddSensitiveApi             = "v4/im_msg_audit_mgr/add_cloud_audit_keywords"    //添加敏感词接口
	DeleteSensitiveApi          = "v4/im_msg_audit_mgr/delete_cloud_audit_keywords" //删除敏感词接口
)
