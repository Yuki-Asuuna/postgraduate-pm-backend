package zookeeper

import (
	"encoding/json"
	"fmt"
	"github.com/go-zookeeper/zk"
	"time"
)

const (
	zookeeper_server_ip = "localhost"
	port                = 2181
)

var client *zk.Conn

func GetZookeeperClient() *zk.Conn {
	return client
}

func ZookeeperInit() error {
	var err error
	target := fmt.Sprintf("%s:%d", zookeeper_server_ip, port)
	client, _, err = zk.Connect([]string{target}, time.Second*10)
	if err != nil {
		return err
	}
	return nil
}

func GetNodeData(path string) ([]byte, error) {
	data, _, err := client.Get(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetUtilsConfig(path string, v interface{}) error {
	data, err := GetNodeData(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	return nil
}
