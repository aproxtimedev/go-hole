package main

import (
	"log"
	"flag"
)

func main() {
	var configFile string

	flag.StringVar(&configFile, "config", "config.yaml", "Path to configuration file")
	flag.Parse()

	log.Printf("Starting Go-hole %s...\n", AppVersion)

	GetConfig().ReadConfig(configFile)
	GetConfig().Print()
	initServer()
	initBlacklistRenewal()
	listenAndServe()
}

func initServer() {
	initLogging()
	GetUpstreamCache().Init()
	updateLocalRecords()
	updateBlacklistRecords()
	updateWhitelistRecords()
}
