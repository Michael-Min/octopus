package config

const (
	ConsulHost  = "http://192.168.199.159:8500"
	Host		= "192.168.199.159"
	ElasticHost = "http://192.168.199.159:9200"
	RedisHost   = "192.168.199.159:6379"
	MqUrl 		= "amqp://admin:admin@192.168.199.159:5672/"
	EtcdHost1   = "http://192.168.199.159:23800"
	EtcdHost2   = "http://192.168.199.159:23801"
	EtcdHost3   = "http://192.168.199.159:23802"
	//EtcdHost1   = "http://etcd1:2379"
	//EtcdHost2   = "http://etcd2:2379"
	//EtcdHost3   = "http://etcd3:2379"
	MaxIdle     = 30
	MaxActive   = 1000
	IdleTimeout = 10

	// Parser names
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ParseProfile"

	ParseCarDetail = "ParseCarDetail"
	ParseCarDetailFake = "ParseCarDetailFake"
	ParseCarList   = "ParseCarList"
	ParseCarModel  = "ParseCarModel"

	NilParser = "NilParser"

	// ElasticSearch
	ElasticIndex = "car_profile_latest"

	// Rate limiting
	Qps = 1
)
