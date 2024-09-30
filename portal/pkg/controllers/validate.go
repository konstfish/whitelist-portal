package controllers

import (
	"fmt"
	"net"

	"github.com/konstfish/whitelist-portal/portal/pkg/cache"
)

func validateAddress(entry cache.AddressListEntry) (bool, string) {
	if !ipAddressValid(entry.Address) {
		return false, "Invalid IP address"
	}
	if !descriptionValid(entry.Description) {
		return false, "Invalid description"
	}
	if !ttlValid(entry.TTL) {
		return false, "Invalid TTL"
	}
	return true, ""
}

func ipAddressValid(ip string) bool {
	if net.ParseIP(ip) == nil {
		return false
	}
	return true
}

func descriptionValid(description string) bool {
	if len(description) > 20 {
		return false
	}
	return true
}

func ttlValid(ttl int64) bool {
	fmt.Println(ttl)
	if ttl < 1 || ttl > 999 {
		return false
	}
	return true
}
