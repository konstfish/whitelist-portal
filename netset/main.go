package main

import (
	"log"

	"github.com/konstfish/whitelist-portal/netset/pkg/cache"
	"github.com/konstfish/whitelist-portal/netset/pkg/config"
	"github.com/konstfish/whitelist-portal/netset/pkg/mappings"
)

func main() {
	log.Println("loading config")
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	log.Println("setting up redis")
	err = cache.SetupCacheClient(cfg.RedisAddr)
	if err != nil {
		log.Fatalf("Failed to setup cache client: %v", err)
	}

	log.Println("setting up gin router")
	mappings.CreateUrlMappings()
	mappings.Router.Run(":8080")
}
