package main

import (
	. "ants/conf"
	"ants/node"
	"flag"
	"log"
	"os"
)

const (
	CONF_FILE = "/../conf/conf.json"
)

func initFlag(settings *Settings) {
	flag.IntVar(&settings.TcpPort, "tcp", 8200, "tcp port")
	flag.IntVar(&settings.HttpPort, "http", 8300, "http port")
}
func MakeSettings() *Settings {
	pwd, _ := os.Getwd()
	settings := LoadSettingFromFile(pwd + CONF_FILE)
	initFlag(settings)
	flag.Parse()
	return settings
}
func main() {
	log.Println("let us go shipping")
	setting := MakeSettings()
	Node := node.NewNode(setting)
	Node.Init()
	log.Println("finish init")
	Node.Start()
}