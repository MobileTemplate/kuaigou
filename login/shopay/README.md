# 商城支付接口

# TODO

添加订单的详情:
misc 包含参数
支付通道名称
支付方式
appid 等信息

返回的参数

修改为本地配置文件，而不是数据库配置关键信息


# 特性

慢慢把所有支付接口，都移动到此目录
只需要支付实现相对应的接口，即可正常使用
参数由此模块管理

# 支付通道名称

	使用官方域名，作为通道名称
	ftakf	乾通支付
	zfy		掌付云
	sf		顺付
	alipay	支付宝
	

# 满足特性

appid， appid_name

对外提供统一交换接口
对内所有的支付统一管理
可以同一种支付方式，多个实例
具体支付实现，不依赖于外部存储接口和存储方式

	统一的管理地址，以 ftafk 为例，有四个参数，支付商家、支付方式、用户id，商品id
	创建订单		/shoppay/:payname/:paytype/users/:id/:gid
	订单效验		/shoppay/:payname/check/:tradeno
	异步通知		/shoppay/:payname/notify
	
	
	type 类型定义
	wxpay	原生微信支付
	wxh5	微信h5支付
	wxgzh	微信公众号支付
	alih5	支付宝h5
	qqh5	QQ支付
	
	
下单步骤：
1. 获取支付通道，用户，商品以及支付方式
2. 判断传入参数是否合法
3. 盘但支付通道参数是否配置正确
4. 平台下单，返回下单参数，如果下单失败，直接修改状态为 4
5. 支付通道下单
6. 等待回调或者通知

json 始终小写加下划线

misc 包含参数
支付通道名称
支付方式
appid 等信息


支付通用配置，是否开启


通用配置参数：
app_name:		应用名称
is_open:		是否打开支付
app_id:			应用 ID 标识
subject:		订单标题
app_key:		AppKey
sign:			签名
pay_url:		支付网关
notify_url:		通知地址
tongbu_url:		同步地址
pay_type:		支付方式(可能多个，不需要)
nnotify_srcip:	通知源 IP 地址，备用
