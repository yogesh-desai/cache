package cache

import "testing"

func TestLFUNGetSet(t *testing.T) {
	cache := NewLFU(2)
	testCacheGetSet(t, cache)
}

func TestLFUDelete(t *testing.T) {
	cache := NewLFU(2)
	testCacheDelete(t, cache)
}

func TestLFUNilValue(t *testing.T) {
	cache := NewLFU(2)
	testCacheNilValue(t, cache)
}

func TestLFUEviction(t *testing.T) {
	cache := NewLFU(2)
	testCacheGetSet(t, cache)

	_, err := cache.Get("test_key2")
	if err != nil {
		t.Fatal("test_key2 should be in the cache")
	}
	// get-> test_key should not be in cache after insert test_key3
	err = cache.Set("test_key3", "test_data3")
	if err != nil {
		t.Fatal("should not give err while setting item")
	}

	_, err = cache.Get("test_key")
	if err == nil {
		t.Fatal("test_key should not be in the cache")
	}
}
