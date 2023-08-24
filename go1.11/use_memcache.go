package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	"google.golang.org/appengine/memcache"
)

func useMemcache(ctx context.Context, w io.Writer) {

	assertf := func(ok bool, errmsg string, args ...interface{}) {
		if ok {
			fmt.Fprintln(w, "✓")
		} else {
			fmt.Fprintf(w, "✖ Failed: "+errmsg+"\n", args...)
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
		assertf(err == memcache.ErrCacheMiss, "expected ErrCacheMiss, got %v", err)
		assertf(item == nil, "expected nil item, got %v", item)
	}

	{
		// TRADITIONAL USAGE (EXISTING API)

		key := "mykey"

		// Cleanup, because previous run may have set it
		err := memcache.Delete(ctx, key)
		if err != nil && err != memcache.ErrCacheMiss {
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

	{
		// NEW API USAGE: PEEK MULTIPLE ITEMS

		key1 := "mykey1"
		value1 := []byte("foobar1")
		item1 := &memcache.Item{
			Key:   key1,
			Value: value1,
			// No Expiration
		}

		key2 := "mykey2"
		value2 := []byte("foobar2")
		item2 := &memcache.Item{
			Key:        key2,
			Value:      value2,
			Expiration: 5 * time.Minute,
		}

		items := []*memcache.Item{item1, item2}

		fmt.Fprintf(w, "\nmemcache.SetMulti(ctx, %d items)...\n", len(items))
		err := memcache.SetMulti(ctx, items)
		assertf(err == nil, "expected nil error, got %v", err)

		keys := []string{key1, key2}

		fmt.Fprintf(w, "\nmemcache.GetMulti(ctx, %q)...\n", keys)
		itemsmap, err := memcache.GetMulti(ctx, keys)
		assertf(err == nil, "expected nil error, got %v", err)
		assertf(len(itemsmap) == len(keys), "expected %d entries, got a map with %d entries", len(keys), len(itemsmap))

		item1in := itemsmap[key1]
		assertf(item1in != nil, "entry %q not found in %v", key1, itemsmap)
		if item1in != nil {
			assertf(bytes.Equal(value1, item1in.Value), "expected item with value %q, got %q", string(value1), string(item1in.Value))
			assertf(item1in.Timestamps.Expiration == nil, "expected nil Expiration timestamp, got %v", item1in.Timestamps.Expiration)
			assertf(item1in.Timestamps.LastAccess == nil, "expected nil LastAccess timestamp, got %v", item1in.Timestamps.LastAccess)
		}

		item2in := itemsmap[key2]
		assertf(item2in != nil, "entry %q not found in %v", key2, itemsmap)
		if item2in != nil {
			assertf(bytes.Equal(value2, item2in.Value), "expected item with value %q, got %q", string(value2), string(item2in.Value))
			assertf(item2in.Timestamps.Expiration == nil, "expected nil Expiration timestamp, got %v", item2in.Timestamps.Expiration)
			assertf(item2in.Timestamps.LastAccess == nil, "expected nil LastAccess timestamp, got %v", item2in.Timestamps.LastAccess)
		}

		fmt.Fprintf(w, "\nmemcache.PeekMulti(ctx, %q)...\n", keys)
		itemsmap, err = memcache.PeekMulti(ctx, keys)
		assertf(err == nil, "expected nil error, got %v", err)
		assertf(len(itemsmap) == len(keys), "expected %d entries, got a map with %d entries", len(keys), len(itemsmap))

		item1in = itemsmap[key1]
		assertf(item1in != nil, "entry %q not found in %v", key1, itemsmap)
		if item1in != nil {
			assertf(bytes.Equal(value1, item1in.Value), "expected item with value %q, got %q", string(value1), string(item1in.Value))
			assertf(item1in.Timestamps.Expiration == nil, "expected nil Expiration timestamp, got %v", item1in.Timestamps.Expiration)
			assertf(item1in.Timestamps.LastAccess != nil, "expected non-nil LastAccess timestamp")
			fmt.Fprintf(w, "Item 1 Expiration timestamp: %v\n", item1in.Timestamps.Expiration)
			fmt.Fprintf(w, "Item 1 LastAccess timestamp: %v\n", item1in.Timestamps.LastAccess)
		}

		item2in = itemsmap[key2]
		assertf(item2in != nil, "entry %q not found in %v", key2, itemsmap)
		if item2in != nil {
			assertf(bytes.Equal(value2, item2in.Value), "expected item with value %q, got %q", string(value2), string(item2in.Value))
			assertf(item2in.Timestamps.Expiration != nil, "expected non-nil Expiration timestamp")
			if expi := item2in.Timestamps.Expiration; expi != nil {
				// Expiration timestamp should be in more than 3mn and less than 6mn
				d := time.Until(*expi)
				assertf(d > 3*time.Minute, "Expiration timestamp is %v which is in %v, which is in less than 3mn", expi, d)
				assertf(d < 6*time.Minute, "Expiration timestamp is %v which is in %v, which is in more than 6mn", expi, d)
			}
			assertf(item2in.Timestamps.LastAccess != nil, "expected non-nil LastAccess timestamp")
			fmt.Fprintf(w, "Item 2 Expiration timestamp: %v\n", item2in.Timestamps.Expiration)
			fmt.Fprintf(w, "Item 2 LastAccess timestamp: %v\n", item2in.Timestamps.LastAccess)
		}
	}

	// Cleanup (teardown)
	keys := []string{"mykey", "mykey1", "mykey2"}
	_ = memcache.DeleteMulti(ctx, keys)
}
