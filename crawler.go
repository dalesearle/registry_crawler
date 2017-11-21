package main

import (
	"errors"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func SearchUsersForKey(path string) (registry.Key, error) {
	key, err := registry.OpenKey(registry.USERS, "", registry.READ)
	if err != nil {
		return key, err
	}
	defer key.Close()
	subs, err := key.ReadSubKeyNames(0)
	if err != nil {
		return key, errors.New("sub key read failed")
	}
	for _, key := range subs {
		foundKey, verified := pathVerifies(key, path)
		if verified {
			return foundKey, nil
		}
	}
	return key, errors.New("path not found")
}
func pathVerifies(subKey string, path string) (registry.Key, bool) {
	keyPath := strings.Join([]string{subKey, path}, "\\")
	key, err := registry.OpenKey(registry.USERS, keyPath, registry.READ)
	if err != nil {
		return key, false
	}
	return key, true
}

