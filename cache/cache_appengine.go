// +build appengine

package cache

import (
	"github.com/rubanbydesign/context"
	"github.com/rubanbydesign/gocache/cache"
	"github.com/rubanbydesign/gocache/codec"
	"github.com/rubanbydesign/gocache/drivers/appengine/memcache"
)

var (
	Codec = codec.Gob
)

func New(ctx context.Context) cache.Cache {
	memcache.Codec = Codec
	return memcache.New(ctx)
}
