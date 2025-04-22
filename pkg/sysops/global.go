package sysops

import "github.com/demdxx/xtypes"

var _globalOpts options

// Has checks if the option with the given key exists
func Has(key string) bool {
	return _globalOpts.Has(key)
}

// Get option value by key
// If the key does not exist, it returns the default value if provided
// Otherwise, it returns nil
func Get(key string, def ...any) *xtypes.Any {
	return _globalOpts.Get(key, def...)
}

// First returns the first option value for the given keys
func First(keys ...string) *xtypes.Any {
	for _, key := range keys {
		if value := _globalOpts.Get(key); value != nil {
			return value
		}
	}
	return nil
}

// Set option value by key
func Set(key string, value any) {
	_globalOpts.Set(key, value)
}

// Delete option by key
func Delete(keys ...string) {
	_globalOpts.Delete(keys...)
}
