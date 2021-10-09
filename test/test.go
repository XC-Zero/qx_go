package main

import (
	"flag"
	"fmt"
	"log"
	"qx/internal/config"
	"qx/pkg/common/ip"
	"strings"
)

var configFileName = flag.String("cfn", "config", "name of configs file")
var configFilePath = flag.String("cfp", "./configs/", "path of configs file")

func main() {
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Llongfile)
	err := config.InitConfig(*configFileName, strings.Split(*configFilePath, ","))
	if err != nil {
		panic(err)
	}
	fmt.Println(ip.GetIpRealAddress("218.17.136.143"))
}
