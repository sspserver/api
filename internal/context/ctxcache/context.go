package ctxcache

import "context"

var (
	ctxCacheKey = struct{ s string }{"ctxCache"}
)

func WithCacheBlock(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxCacheKey, newCacheBlock())
}

func ReleaseCacheBlock(ctx context.Context) {
	cacheBlock, ok := ctx.Value(ctxCacheKey).(*cacheBlock)
	if ok {
		cacheBlock.Release()
	}
}

func GetCache(ctx context.Context, key any) Cacher {
	cacheBlock, ok := ctx.Value(ctxCacheKey).(*cacheBlock)
	if !ok {
		return nil
	}
	return cacheBlock.Get(key)
}
