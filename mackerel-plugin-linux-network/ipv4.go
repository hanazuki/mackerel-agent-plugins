package main

import (
	"os"

	mp "github.com/mackerelio/go-mackerel-plugin-helper"
)

var graphdef4 = map[string](mp.Graphs){
}

func ipv4FetchMetrics() (metrics map[string]interface{}, err error) {
	metrics = make(map[string]interface{})

	for _, statFile := range [...]string {"/proc/net/netstat", "/proc/net/snmp"} {
		file, err := os.Open(statFile)
		if err != nil {
			continue
		}

		for k, v := range parseCompact(file) {
			metrics[prefix + "ipv4." + k] = v
		}
	}

	return
}

func ipv4GraphDefinition() map[string]mp.Graphs {
	return graphdef4
}
