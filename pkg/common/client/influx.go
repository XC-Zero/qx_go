package client

import (
	"context"
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	cfg "qx/pkg/common/config"
)

func InitInflux(config cfg.InfluxConfiguration) (influxdb2.Client, error) {

	url := fmt.Sprintf("http://%s:%s", config.Host, config.Port)

	client := influxdb2.NewClient(url, config.Token)

	_, err := client.Ready(context.TODO())
	if err != nil {
		return nil, err
	}
	return client, nil
}
