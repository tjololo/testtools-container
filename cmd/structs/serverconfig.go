package structs

import "time"

type TestsConfig struct {
	ConnectTests map[string]ConnectTest `json:"connectTests,omitempty" yaml:"connectTests,omitempty"`
	DnsTests     map[string]DnsTest     `json:"dnsTests,omitempty" yaml:"dnsTests,omitempty"`
	ApiTests     map[string]ApiTest     `json:"apiTests,omitempty" yaml:"apiTests,omitempty"`
}

type ConnectTest struct {
	Hostname string         `json:"hostname" yaml:"hostname"`
	Port     *int           `json:"port,omitempty" yaml:"port,omitempty"`
	Timeout  *time.Duration `json:"timeout,omitempty" yaml:"timeout,omitempty"`
	Interval *time.Duration `json:"interval,omitempty" yaml:"interval,omitempty"`
}

type DnsTest struct {
	Hostname string         `json:"hostname" yaml:"hostname"`
	Interval *time.Duration `json:"interval,omitempty" yaml:"interval,omitempty"`
}

type ApiTest struct {
	URI      string         `json:"uri" yaml:"uri"`
	Interval *time.Duration `json:"interval,omitempty" yaml:"interval,omitempty"`
	Timeout  *time.Duration `json:"timeout,omitempty" yaml:"timeout,omitempty"`
}
