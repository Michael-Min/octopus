package config

const (
	ConsulHost  = "http://192.168.199.160:8500"
	Host		= "192.168.199.160"
	ElasticHost = "http://192.168.199.160:9200"
	RedisHost   = "192.168.199.160:6379"
	EtcdHost1   = "http://etcd1:2379"
	EtcdHost2   = "http://etcd2:2379"
	EtcdHost3   = "http://etcd3:2379"
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
