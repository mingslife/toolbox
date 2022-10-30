package conf

type Http struct {
	Host string `default:"127.0.0.1"`
	Port int    `default:"9000"`
}
