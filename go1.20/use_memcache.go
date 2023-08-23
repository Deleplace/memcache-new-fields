package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	"google.golang.org/appengine/v2/memcache"
)

func useMemcache(ctx context.Context, w io.Writer) {

	assertf := func(ok bool, errmsg string, args ...interface{}) {
		if ok {
			fmt.Fprintln(w, "✓")
		} else {
			fmt.Fprintf(w, "✖ Failed: "+errmsg, args...)
		}
	}

	fmt.Fprintln(w, "\nExported func memcache.Peek:")
	fmt.Fprintf(w, "%T\n\n", memcache.Peek)

	fmt.Fprintln(w, "\nExported func memcache.PeekMulti:")
	fmt.Fprintf(w, "%T\n\n", memcache.PeekMulti)

	{
		// TRADITIONAL USAGE (EXISTING API)

		key := "NON-EXISTENT"
		fmt.Fprintf(w, "memcache.Get(ctx, %q)...\n", key)
		item, err := memcache.Get(ctx, key)
		assertf(err == memcache.ErrCacheMiss, "expected nil error, got %v", err)
		assertf(item == nil, "expected nil item, got %v", item)
	}

	{
		// TRADITIONAL USAGE (EXISTING API)

		key := "mykey"

		// Cleanup, because previous run may have set it
		err := memcache.Delete(ctx, key)
		if err != nil {
			fmt.Fprintf(w, "Error deleting entry for key %q: %v\n", key, err)
		}

		fmt.Fprintf(w, "\nmemcache.Get(ctx, %q)...\n", key)
		item, err := memcache.Get(ctx, key)
		assertf(err == memcache.ErrCacheMiss, "expected ErrCacheMiss, got %v", err)
		assertf(item == nil, "expected nil item, got %v", item)

		value := []byte("foobar")
		item = &memcache.Item{
			Key:   key,
			Value: value,
		}
		fmt.Fprintf(w, "\nmemcache.Set(ctx, %v)...\n", item)
		err = memcache.Set(ctx, item)
		assertf(err == nil, "expected nil error, got %v", err)
	}

	{
		// NORMAL USAGE, BUT CHECK THE NEW FIELDS RETURNED (DEFAULT VALUES)

		key := "mykey"
		expectedValue := []byte("foobar")

		fmt.Fprintf(w, "\nmemcache.Get(ctx, %q)...\n", key)
		item, err := memcache.Get(ctx, key)
		assertf(err == nil, "expected nil error, got %v", err)
		assertf(item != nil, "expected non-nil item")
		if item != nil {
			assertf(bytes.Equal(expectedValue, item.Value), "expected item with value %q, got %q", string(expectedValue), string(item.Value))

			// When calling Get, these new timestamps should always be nil
			assertf(item.Timestamps.Expiration == nil, "expected nil Expiration timestamp, got %v", item.Timestamps.Expiration)
			assertf(item.Timestamps.LastAccess == nil, "expected nil LastAccess timestamp, got %v", item.Timestamps.LastAccess)
		}
	}

	{
		// NEW API USAGE

		key := "mykey"
		expectedValue := []byte("foobar")

		fmt.Fprintf(w, "\nmemcache.Peek(ctx, %q)...\n", key)
		item, err := memcache.Peek(ctx, key)
		assertf(err == nil, "expected nil error, got %v", err)
		assertf(item != nil, "expected non-nil item")
		if item != nil {
			assertf(bytes.Equal(expectedValue, item.Value), "expected item with value %q, got %q", string(expectedValue), string(item.Value))

			// When calling Peek, at least LastAccess should be non-nil
			assertf(item.Timestamps.Expiration == nil, "expected nil Expiration timestamp, got %v", item.Timestamps.Expiration)
			assertf(item.Timestamps.LastAccess != nil, "expected non-nil LastAccess timestamp")
		}

		fmt.Fprintf(w, "Item LastAccess timestamp: %v\n", item.Timestamps.LastAccess)
	}

	{
		// NEW API USAGE: WITH EXPIRATION

		key := "mykey"
		value := []byte("foobar")
		item := &memcache.Item{
			Key:        key,
			Value:      value,
			Expiration: 20 * time.Minute,
		}
		fmt.Fprintf(w, "\nmemcache.Set(ctx, %v)...\n", item)
		err := memcache.Set(ctx, item)
		assertf(err == nil, "expected nil error, got %v", err)

		fmt.Fprintf(w, "\nmemcache.Peek(ctx, %q)...\n", key)
		item, err = memcache.Peek(ctx, key)
		assertf(err == nil, "expected nil error, got %v", err)
		assertf(item != nil, "expected non-nil item")
		if item != nil {
			assertf(bytes.Equal(value, item.Value), "expected item with value %q, got %q", string(value), string(item.Value))
			assertf(item.Timestamps.Expiration != nil, "expected non-nil Expiration timestamp")
			if expi := item.Timestamps.Expiration; expi != nil {
				// Expiration timestamp should be in more than 15mn and less than 21mn,
				// which is a very generous tolerance margin around 20mn.
				d := time.Until(*expi)
				assertf(d > 15*time.Minute, "Expiration timestamp is %v which is in %v, which is in less than 15mn", expi, d)
				assertf(d < 21*time.Minute, "Expiration timestamp is %v which is in %v, which is in more than 21mn", expi, d)
			}
			assertf(item.Timestamps.LastAccess != nil, "expected non-nil LastAccess timestamp")
			fmt.Fprintf(w, "Item Expiration timestamp: %v\n", item.Timestamps.Expiration)
			fmt.Fprintf(w, "Item LastAccess timestamp: %v\n", item.Timestamps.LastAccess)
		}
	}

	{
		// NORMAL USAGE AGAIN
		key := "mykey"

		fmt.Fprintf(w, "\nmemcache.Get(ctx, %q)...\n", key)
		item, err := memcache.Get(ctx, key)
		assertf(err == nil, "expected nil error, got %v", err)
		assertf(item != nil, "expected non-nil item")
		if item != nil {
			// When calling Get, these new timestamps should always be nil
			// Even after some unrelated calls to Peek
			assertf(item.Timestamps.Expiration == nil, "expected nil Expiration timestamp, got %v", item.Timestamps.Expiration)
			assertf(item.Timestamps.LastAccess == nil, "expected nil LastAccess timestamp, got %v", item.Timestamps.LastAccess)
		}
	}

	// Cleanup (teardown)
	key := "mykey"
	_ = memcache.Delete(ctx, key)
}
