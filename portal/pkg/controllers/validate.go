package controllers

import (
	"log"
	"net"
	"unicode"

	"github.com/konstfish/whitelist-portal/portal/pkg/cache"
)

var maxDescriptionLength int = 20
var minTTL int64 = 1
var maxTTL int64 = 999

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
	if len(description) > maxDescriptionLength || !isAlphanumeric(description) {
		return false
	}
	return true
}

func isAlphanumeric(s string) bool {
	log.Println("asdf")
	for _, r := range s {
		log.Println(!unicode.IsLetter(r) && !unicode.IsNumber(r))
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return false
		}
	}
	return true
}

func ttlValid(ttl int64) bool {
	if ttl < minTTL || ttl > maxTTL {
		return false
	}
	return true
}
