package config

const (
	// Parser names
	ParseCityList = "ParseCityList"
	ParseCity = "ParseCity"
	ParseProfile = "ParseProfile"
	NilParser = "NilParser"

	// elasticSearch的服务器地址
	ElasticServerUrl = "http://47.107.149.234:9200"
	// elasticSearch的数据库名字 (对应其Index字段)
	ElasticSearchIndex = "dating_profile"
	// elasticSearch的表名 (对应其Type字段)
	ElasticSearchTypeZhenai = "zhenai"

	// redis 服务地址 (前面不能加 http://)
	RedisServerUrl = "47.107.149.234:6379"

	// 爬虫的起始url
	SeedUrl = "http://www.zhenai.com/zhenghun"

	// rpc service host:port
	RpcServeHost = ":1234"

	WorkerHost0 = ":9000"
	WorkerHost1 = ":9001"
	WorkerHost2 = ":9002"


	// RPC EndPoints
	ItemSaverRpc = "ItemSaverService.Save"    // itemSaver rpc service
    CrawlServiceRpc = "CrawlService.Process"  // worker rpc service

    // Rate limiting
    Qps = 5            // 限制Fetcher发起请求的频率
)
