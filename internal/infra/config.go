package internal

type Config struct {
	Httpd    httpd    `yaml:"httpd"`
	RabbitMQ rabbitMQ `yaml:"rabbitMQ"`
	Scylla   scylla   `yaml:"scylla"`
}
type httpd struct {
	Port string `yaml:"port"`
}
type rabbitMQ struct {
	ConnectionUrl string `yaml:"connectionUrl"`
}
type scylla struct {
	ConnectionUrl string `yaml:"connectionUrl"`
}
