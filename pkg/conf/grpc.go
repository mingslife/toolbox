package conf

type Grpc struct {
	Host string `default:"127.0.0.1"`
	Port int    `default:"31000"`
}
