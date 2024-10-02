package cache

import (
	"context"
	"fmt"
	"slices"
)

func userKey(user string) string {
	return fmt.Sprintf("user:%s:ips", user)
}

func getUserIPAddresses(ctx context.Context, key string) ([]string, error) {
	return rdb.HKeys(ctx, key).Result()
}

func GetUserIPAddresses(ctx context.Context, user string) ([]string, error) {
	ips, err := getUserIPAddresses(ctx, userKey(user))
	if err != nil {
		return nil, err
	}

	slices.Sort(ips)

	return ips, nil
}

func GetAllIPAddresses(ctx context.Context) ([]string, error) {
	matchingKeys, err := rdb.Keys(ctx, userKey("*")).Result()
	if err != nil {
		return nil, err
	}

	seen := make(map[string]bool)

	for _, key := range matchingKeys {
		ips, err := getUserIPAddresses(ctx, key)
		if err != nil {
			return nil, err
		}

		// ipList = append(ipList, ips...)

		for _, ip := range ips {
			seen[ip] = true
		}
	}

	ipList := make([]string, 0, len(seen))
	for ip := range seen {
		ipList = append(ipList, ip)
	}

	slices.Sort(ipList)

	return ipList, nil
}
