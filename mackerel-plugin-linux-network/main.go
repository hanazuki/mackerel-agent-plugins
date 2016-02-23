package main

import (
	"os"
	"fmt"

	"github.com/codegangsta/cli"
	mp "github.com/mackerelio/go-mackerel-plugin-helper"
)

const prefix = "network."

type opts struct {
	ShowIpv4 bool
	ShowIpv6 bool
}

func (p *opts) metricsGenerators() (gens []func() (map[string]interface{}, error)) {
	if p.ShowIpv4 {
		gens = append(gens, ipv4FetchMetrics)
	}
	if p.ShowIpv6 {
		gens = append(gens, ipv6FetchMetrics)
	}
	return
}

func(p *opts) graphdefGenerators() (gens []func() map[string]mp.Graphs) {
	if p.ShowIpv4 {
		gens = append(gens, ipv4GraphDefinition)
	}
	if p.ShowIpv6 {
		gens = append(gens, ipv6GraphDefinition)
	}
	return
}

func (p *opts) FetchMetrics() (metrics map[string]interface{}, err error) {
	metrics = make(map[string]interface{})

	for _, gen := range p.metricsGenerators() {
		m, _ := gen()
		for k, v := range m {
			metrics[k] = v
		}
	}

	fmt.Printf("%v\n", metrics)

	return
}

func (p *opts) GraphDefinition() (graphdef map[string]mp.Graphs) {
	graphdef = make(map[string]mp.Graphs)

	for _, gen := range p.graphdefGenerators() {
		g := gen()
		for k, v := range g {
			graphdef[k] = v
		}
	}

	return
}

func doMain(c *cli.Context) {
	var opts opts

	opts.ShowIpv4 = c.Bool("4")
	opts.ShowIpv6 = c.Bool("6")
	if !opts.ShowIpv4 && !opts.ShowIpv6 {
		opts.ShowIpv4 = true
	}

	helper := mp.NewMackerelPlugin(&opts)
	helper.Tempfile = c.String("tempfile")

	if os.Getenv("MACKEREL_AGENT_PLUGIN_META") != "" {
		helper.OutputDefinitions()
	} else {
		helper.OutputValues()
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "mackerel-plugin-linux-network"
	app.Version = version
	app.Usage = "Get network metrics from Linux kernel."
	app.Action = doMain
	app.Flags = cliFlags

	app.Run(os.Args)
}
