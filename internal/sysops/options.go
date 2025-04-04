package sysops

import (
	"sync"

	"github.com/demdxx/xtypes"
)

type options struct {
	values sync.Map
}

func (o *options) Has(key string) bool {
	_, ok := o.values.Load(key)
	return ok
}

func (o *options) Get(key string, def ...any) *xtypes.Any {
	v, _ := o.values.Load(key)
	if v == nil {
		if len(def) > 0 {
			return &xtypes.Any{Val: def[0]}
		}
		return nil
	}
	return &xtypes.Any{Val: v}
}

func (o *options) Set(key string, value any) {
	o.values.Store(key, value)
}

func (o *options) Delete(keys ...string) {
	for _, key := range keys {
		o.values.Delete(key)
	}
}
