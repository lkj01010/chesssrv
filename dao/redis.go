package dao

//keys for user db
const (
// account
	k_account_user_ = "account:user:"    // + account = id [string] 账号和id的映射
	k_account_count = "account:count"        // [string] 当前总账号数
	k_account_userlist = "account:userlist"    // [set] 全部账号id
	k_psw = "psw"    // [string]

// user
	k_user_ = "user:"    // + id [hash] 用户档案
	k_gold = "gold" //
	k_nickname = "nickname"    //

)