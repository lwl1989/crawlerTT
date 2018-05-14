package user

const  (
	UserListUrl string = "http://ttservice.toplogis.com/ttService/queryUser.action"
	UserInfoUrl  = "http://ttservice.toplogis.com/ttService/queryUserDataSet!editInput.action"
	UserCoinUrl  = "https://glcp.toplogis.com/product/queryCoinExpire"
	UserCoinLogUrl = "https://glcp.toplogis.com/product/queryCoinLog"

	DepartmentUrl = "http://ttservice.toplogis.com/ttService/queryDepartMent.action"

	MessageUrl = "http://ttservice.toplogis.com/ttService/queryPushMsgSetting.action"
	MessageContent = "http://ttservice.toplogis.com/ttService/editPushMsgSetting.action?msDto.msg_id=%s" //get
	MessageAlreadyUrl = "http://ttservice.toplogis.com/ttService/queryPushNotification.action"


	ActivityUrl = "http://ttservice.toplogis.com/ttService/queryActivity.action"
	ActivityUrlContent =  "http://ttservice.toplogis.com/ttService/editActivity.action"  //checkedActivityId => id

	AdminUrl = "http://ttservice.toplogis.com/ttService/queryAccount.action"
	AdminContentUrl = "http://ttservice.toplogis.com/ttService/queryAccountDataSet!editInput.action" //editAccountId  2

	ShopUrl = "http://ttservice.toplogis.com/ttService/queryStoreAccount.action"
	ShopContentUrl = "http://ttservice.toplogis.com/ttService/queryStoreAccountDataSet!editInput.action" //editAccountId 2
)