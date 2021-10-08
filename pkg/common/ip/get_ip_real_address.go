package ip

import (
	"qx/internal/config"
	my_http "qx/pkg/common/http"
)

func GetIpRealAddress(ip string) string {
	temp := make(map[string]string)
	value := make(map[string]string)
	temp["Authorization"] = "APPCODE " + config.CONFIG.IpRealAddressAPIConfig.AppCode
	value["ip"] = ip
	return my_http.Get(config.CONFIG.IpRealAddressAPIConfig.URL, temp, value)
}
