package main

import (
	"log"

	"github.com/konstfish/whitelist-portal/portal/pkg/cache"
	"github.com/konstfish/whitelist-portal/portal/pkg/config"
	"github.com/konstfish/whitelist-portal/portal/pkg/controllers"
	"github.com/konstfish/whitelist-portal/portal/pkg/mappings"
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
	// very temporary
	controllers.AuthHeader = cfg.AuthHeader
	controllers.UserAddressLimit = cfg.UserAddressLimit
	mappings.CreateUrlMappings()
	mappings.Router.Run(":8080")
}
