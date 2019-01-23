# 登录服务器说明

## TODO

登录请求参数更新
微信支付后，刷新数据

刷新房卡和删除缓存可以做在一起

## API 接口


// 注册参数 post 请求，请求地址 /users/register
// 请求参数使用 body 数据，数据是个 json
{
	"channel": "channel_name",
	"acttype": "acttype",
	"sdkid": "sdkid",
	"nickname": "nickname",
	"sex": 1,
	"icon": "icon_url"
};

// 登录参数 get 请求地址 /users/login

// 包含以下数据，请求参数用 url 传递

"username": "username",
"userpass": "userpass",
"nickname": "nickname",
"sex": 1,
"icon": "icon_url"


## 管理员接口说明

1. 登录接口
   地址：https://scmj.ttdapai.com/login/users/login
   请求：POST
   参数：body  (使用 json)
	   {"username": "10000", "userpass": ""}
   说明：写入用户名和密码，只有 10000 号用户拥有接口调用权限
   
2. 订单查看
   地址：https://scmj.ttdapai.com/login/admin/orders/[orderid]
   请求：GET
   参数： orderid 订单编号
   说明：返回 Json 数据，内容以微信为准
   
3. 订单人工效验
   地址：https://scmj.ttdapai.com/login/admin/pay/check/[orderid]
   请求：GET
   参数：orderid 订单编号
   说明：返回说明内容，内部自己处理房卡数量的管理，此接口只用于人工检测，不需要根据结果做房卡数量的修改
   
4. 修改房卡
   地址：https://scmj.ttdapai.com/login/admin/users/[uid]/fangka/[addcount]
   请求：GET
   参数：uid 用户 id，addcount 添加房卡数量 （减少的话就写复数 如  -10）
   说明：修改指定用户的房卡数量变化，如果变化后小于 0 则等于 0
   
   登录接口，返回 josn 数据，提取 token 字段用于其它接口的请求权限
2，3，4 接口需要带入 登录返回的 token 作为请求头部。
Authorization:Bearer [token]
所有 http 请求 200 表示请求成功，4XX 表示请求失败
