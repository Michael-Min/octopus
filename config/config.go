package config

const (
	ConsulHost  = "http://127.0.0.1:8500"
	ElasticHost = "http://127.0.0.1:9200"
	RedisHost   = "127.0.0.1:6379"
	EtcdHost1   = "http://192.168.64.5:2380"
	EtcdHost2   = "http://192.168.64.6:2380"
	EtcdHost3   = "http://192.168.64.7:2380"
	MaxIdle     = 30
	MaxActive   = 1000
	IdleTimeout = 10

	// Parser names
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ParseProfile"

	ParseCarDetail = "ParseCarDetail"
	ParseCarList   = "ParseCarList"
	ParseCarModel  = "ParseCarModel"

	NilParser = "NilParser"

	// ElasticSearch
	ElasticIndex = "car_profile_latest"

	// Rate limiting
	Qps = 1
)
