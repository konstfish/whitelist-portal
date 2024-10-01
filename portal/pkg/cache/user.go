package cache

import (
	"context"
	"fmt"
	"sort"
	"time"
)

type AddressListEntry struct {
	Address     string `form:"address"`
	Description string `form:"description"`
	TTL         int64  `form:"expiry"`
}

func userKey(user string) string {
	return fmt.Sprintf("user:%s:ips", user)
}

func AddAddress(ctx context.Context, user string, entry AddressListEntry) error {
	if err := rdb.HSet(ctx, userKey(user), entry.Address, entry.Description).Err(); err != nil {
		return err
	}

	rdb.HExpire(ctx, userKey(user), time.Duration(entry.TTL)*time.Hour*24, entry.Address)

	return nil
}

func getAddresses(ctx context.Context, user string) (map[string]string, error) {
	return rdb.HGetAll(ctx, userKey(user)).Result()
}

func getAddressTTL(ctx context.Context, user string, address string) (int64, error) {
	ttl, err := rdb.HExpireTime(ctx, userKey(user), address).Result()
	if err != nil && len(ttl) != 1 {
		return 0, err
	}

	final := time.Unix(ttl[0], 0).Sub(time.Now()).Hours()

	return int64(final), nil
}

func GetAddressList(ctx context.Context, user string) ([]AddressListEntry, error) {
	entries, err := getAddresses(ctx, user)
	if err != nil {
		return nil, err
	}

	var result []AddressListEntry
	for addr, desc := range entries {
		ttl, _ := getAddressTTL(ctx, user, addr)
		result = append(result, AddressListEntry{
			Address:     addr,
			Description: desc,
			TTL:         ttl,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Address < result[j].Address
	})

	return result, nil
}

func DeleteAddress(ctx context.Context, user string, address string) error {
	return rdb.HDel(ctx, userKey(user), address).Err()
}

func GetAddressCount(ctx context.Context, user string) (int64, error) {
	return rdb.HLen(ctx, userKey(user)).Result()
}
