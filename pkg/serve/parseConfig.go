package serve

import (
	"encoding/json"
	"fmt"
	"github.com/tjololo/testtools-container/cmd/structs"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

func ReadConfigFile(configfile string) (*structs.TestsConfig, error) {
	stat, err := os.Stat(configfile)
	if err != nil {
		return nil, err
	}
	if stat.IsDir() {
		return nil, fmt.Errorf("%s is a directory pleas supply a file", configfile)
	}
	if stat.Size() == 0 {
		return nil, fmt.Errorf("configfile %s has size 0, please supply a file with some content", configfile)
	}
	bytes, err := os.ReadFile(configfile)
	if err != nil {
		return nil, fmt.Errorf("failed to read configfile %s, %v", configfile, err)
	}
	conf, err := parseConfigFile(bytes)
	if err != nil {
		return nil, err
	}
	for _, d := range conf.DnsTests {
		if d.Interval == nil {
			d.Interval = toDuartionPointer(10 * time.Second)
		}
	}
	for _, c := range conf.ConnectTests {
		if c.Port == nil {
			c.Port = toIntPointer(443)
		}
		if c.Timeout == nil {
			c.Timeout = toDuartionPointer(30 * time.Second)
		}
		if c.Interval == nil {
			c.Interval = toDuartionPointer(10 * time.Second)
		}
	}
	for _, a := range conf.ApiTests {
		if a.Interval == nil {
			a.Interval = toDuartionPointer(10 * time.Second)
		}
	}
	return conf, nil
}

func parseConfigFile(bytes []byte) (conf *structs.TestsConfig, err error) {

	err = json.Unmarshal(bytes, &conf)
	if err != nil {
		err = yaml.Unmarshal(bytes, &conf)
	}
	return
}

func toDuartionPointer(duration time.Duration) *time.Duration {
	return &duration
}

func toIntPointer(int int) *int {
	return &int
}
