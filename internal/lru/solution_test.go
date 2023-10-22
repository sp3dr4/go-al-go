package lru

import "testing"

func TestLRU(t *testing.T) {
	lru := NewLRU[string, int](3)

	k, v := "foo", 69
	expectGetErr(t, lru, k)
	lru.Update(k, v)
	expectGetRes(t, lru, k, v)

	k, v = "bar", 420
	lru.Update(k, v)
	expectGetRes(t, lru, k, v)

	k, v = "baz", 1337
	lru.Update(k, v)
	expectGetRes(t, lru, k, v)

	k, v = "ball", 69420
	lru.Update(k, v)
	expectGetRes(t, lru, k, v)

	expectGetErr(t, lru, "foo")

	expectGetRes(t, lru, "bar", 420)

	lru.Update("foo", 69)

	expectGetRes(t, lru, "bar", 420)

	expectGetRes(t, lru, "foo", 69)

	expectGetErr(t, lru, "baz")
}

func expectGetErr(t *testing.T, lru *LRU[string, int], key string) {
	res, err := lru.Get(key)
	if err == nil {
		t.Errorf("expected error, got %v", res)
	}
}

func expectGetRes(t *testing.T, lru *LRU[string, int], key string, expected int) {
	res, err := lru.Get(key)
	if err != nil {
		t.Errorf("expected %v, got %v", expected, err)
	} else if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

// func expectGet(t *testing.T, lru *LRU[string, int], key string, expected testOptions) {
// 	res, err := lru.Get(key)
// 	if expected.err == 1 {
// 		if err == nil {
// 			t.Error("expected error")
// 		}
// 	} else if expected.res != nil {
// 		if res != *expected.res {
// 			t.Errorf("expected %v, got %v", expected.res, res)
// 		}
// 	} else {
// 		t.Error("invalid testOptions, result or shouldErr must be set")
// 	}
// }
