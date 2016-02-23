package main

import (
	"os"

	mp "github.com/mackerelio/go-mackerel-plugin-helper"
)

var graphdef6 = map[string](mp.Graphs){
}

func ipv6FetchMetrics() (metrics map[string]interface{}, err error) {
	metrics = make(map[string]interface{})

	for _, statFile := range [...]string {"/proc/net/snmp6"} {
		file, err := os.Open(statFile)
		if err != nil {
			continue
		}

		for k, v := range parseTable(file) {
			metrics[prefix + "ipv6." + k] = v
		}
	}

	return
}

func ipv6GraphDefinition() map[string]mp.Graphs {
	return graphdef6
}
