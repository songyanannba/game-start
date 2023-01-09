package conf

import "github.com/jinzhu/configor"

const HOST = "127.0.0.1:8765"
const PATH = "/ktpd/room"

var CliConf struct {
	Host string `default:"127.0.0.1:8765"`
	Port string `default:"8765"`
}

func CliConfInit() {

	configor.Load(&CliConf, "")
}
