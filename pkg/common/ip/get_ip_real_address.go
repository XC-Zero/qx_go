package ip

import (
	"encoding/json"
	"fmt"
	"qx/internal/config"
	my_http "qx/pkg/common/http"
)

type IpAddressDataJson struct {
	Code int `json:"code"`
	Data struct {
		Province string `json:"province"`
		City     string `json:"city"`
		Isp      string `json:"isp"`
	} `json:"data"`
	Errors interface{} `json:"errors"`
}

func GetIpRealAddress(ip string) *IpAddressDataJson {
	fmt.Println(config.CONFIG.IpRealAddressAPIConfig)
	temp := make(map[string]string)
	value := make(map[string]string)
	temp["Authorization"] = "APPCODE " + config.CONFIG.IpRealAddressAPIConfig.AppCode
	value["ip"] = ip
	data := IpAddressDataJson{}
	err := json.Unmarshal([]byte(my_http.Get(config.CONFIG.IpRealAddressAPIConfig.URL, temp, value)), &data)
	if err != nil {
		return nil
	}
	fmt.Println(data.Data.City)
	return &data
}
