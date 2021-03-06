package lru4go

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	cache, err := New(50)
	if err != nil || cache == nil {
		t.Fatal("create cache failed!")
	}
}

func TestLrucache_Set(t *testing.T) {
	cache, _ := New(50)
	if cache == nil {
		t.Fatal("create cache failed!")
	}
	cache.Set("test1", 123)
	v, err := cache.Get("test1")
	if err != nil {
		t.Fatal("get failed,err:", err)
	}
	if v == nil {
		t.Fatal("get failed, value is nil")
	}
}

func TestLrucache_Set_TTL(t *testing.T) {
	cache, _ := New(50)
	if cache == nil {
		t.Fatal("create cache failed!")
	}
	cache.Set("test2", 123, 5)
	v, err := cache.Get("test2")
	if err != nil {
		t.Fatal("get failed,err:", err)
	}
	if v == nil {
		t.Fatal("get failed, value is nil")
	}
	time.Sleep(time.Second * 6)
	v, err = cache.Get("test2")
	if v != nil {
		t.Fatal("ttl test is failed, value != nil, value:", v)
	}
}

func TestLrucache_Keys(t *testing.T) {
	cache, _ := New(50)
	cache.Set("t1",1)
	cache.Set("t2",2)
	cache.Set("t3",2)

	keys := cache.Keys()
	if nil == keys {
		t.Fatal("keys is nil")
	}
	for _,v :=  range keys {
		if "t1" != v && "t2" != v && "t3" != v {
			t.Error("keys wrong,key=",v)
		}
	}
}

func TestLrucache_Keys_TTL(t *testing.T) {
	cache, _ := New(50)
	cache.Set("t1",1)
	cache.Set("t2",2)
	cache.Set("t3",2, 5)

	keys := cache.Keys()
	if nil == keys {
		t.Fatal("keys is nil")
	}
	for _,v :=  range keys {
		if "t1" != v && "t2" != v && "t3" != v {
			t.Error("keys wrong,key=",v)
		}
	}
	time.Sleep(time.Second * 6)
	keys = cache.Keys()

	for _,v :=  range keys {
		if "t1" != v && "t2" != v {
			t.Error("keys wrong,key=",v)
		}
		if "t3" == v {
			t.Error("ttl failed!")
		}
	}
}

func TestLrucache_Delete(t *testing.T) {
	cache, _ := New(50)
	cache.Set("t1",1)
	cache.Set("t2",2)
	cache.Set("t3",2, 5)

	cache.Delete("t1")

	v, _ := cache.Get("t1")
	 if v != nil {
	 	t.Error("delete failed!")
	 }

	cache.Delete("t3")

	v, _ = cache.Get("t3")
	if v != nil {
		t.Error("delete ttl failed!")
	}
}

func TestLrucache_Reset(t *testing.T) {
	cache, _ := New(50)
	cache.Set("t1",1)
	cache.Set("t2",2)
	cache.Set("t3",2, 5)

	cache.Reset()
	v, _ := cache.Get("t1")
	if v != nil {
		t.Error("delete failed!")
	}

	keys := cache.Keys()
	if len(keys) != 0 {
		t.Error("failed: len(key) != 0")
	}

}