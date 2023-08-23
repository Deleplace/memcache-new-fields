package main

import (
	"context"
	"fmt"
	"io"

	"google.golang.org/appengine/v2/memcache"
)

func useMemcache(ctx context.Context, w io.Writer) {

	assertf := func(ok bool, errmsg string, args ...interface{}) {
		if ok {
			fmt.Fprintln(w, "✓")
		} else {
			fmt.Fprintf(w, "Failed: "+errmsg, args...)
		}
	}

	fmt.Fprintln(w, "\nExported func memcache.Peek:")
	fmt.Fprintf(w, "%T\n\n", memcache.Peek)

	fmt.Fprintln(w, "\nExported func memcache.PeekMulti:")
	fmt.Fprintf(w, "%T\n\n", memcache.PeekMulti)

	{
		key := "NON-EXISTENT"
		fmt.Fprintf(w, "memcache.Get(ctx, %q)...\n", key)
		item, err := memcache.Get(ctx, key)
		assertf(item == nil, "expected nil item, got %v", item)
		assertf(err == memcache.ErrCacheMiss, "expected nil item, got %v", err)
	}
}
