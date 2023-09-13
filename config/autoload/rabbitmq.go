package autoload

import "strconv"

/*
Created by 斑斑砖 on 2023/9/13.
	Description：
		rabbitmq配置
*/

type RabbitMQ struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Vhost    string `mapstructure:"vhost" json:"vhost" yaml:"vhost"`
}

func (rabbit RabbitMQ) Dsn() string {
	return "amqp://" + rabbit.Username + ":" + rabbit.Password + "@" + rabbit.Host + ":" + strconv.Itoa(rabbit.Port) + "/" + rabbit.Vhost
}
