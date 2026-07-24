package collector

import (
	"testing"
	"time"
)

func TestOltInfoCache_SetAndGet(t *testing.T) {
	cache := NewOltInfoCache(5*time.Minute, 1*time.Minute)
	defer cache.StopGC()

	oltInfo := &OltInfo{
		OamIP: "192.168.1.1",
		NTInfo: NTInfo{
			Component: Component{
				Name: "NT-1",
				Type: "NetworkTermination",
			},
		},
	}

	// 测试Set
	cache.Set(oltInfo)

	// 测试Get
	retrieved, exists := cache.Get("192.168.1.1")
	if !exists {
		t.Fatal("Expected OltInfo to exist in cache")
	}

	if retrieved.OamIP != "192.168.1.1" {
		t.Errorf("Expected OamIP to be 192.168.1.1, got %s", retrieved.OamIP)
	}
}

func TestOltInfoCache_Update(t *testing.T) {
	cache := NewOltInfoCache(5*time.Minute, 1*time.Minute)
	defer cache.StopGC()

	oltInfo := &OltInfo{
		OamIP: "192.168.1.1",
		NTInfo: NTInfo{
			Component: Component{
				Name: "NT-1",
				Type: "NetworkTermination",
			},
		},
	}

	cache.Set(oltInfo)
	time.Sleep(10 * time.Millisecond)

	// 获取原始条目的时间
	entry1, _ := cache.GetEntry("192.168.1.1")

	// 更新
	oltInfo.NTInfo.Component.Name = "NT-2"
	cache.Update(oltInfo)

	// 验证更新
	retrieved, _ := cache.Get("192.168.1.1")
	if retrieved.NTInfo.Component.Name != "NT-2" {
		t.Errorf("Expected Name to be NT-2, got %s", retrieved.NTInfo.Component.Name)
	}

	// 验证LastUpdateAt已更新
	entry2, _ := cache.GetEntry("192.168.1.1")
	if !entry2.LastUpdateAt.After(entry1.LastUpdateAt) {
		t.Error("Expected LastUpdateAt to be updated")
	}
}

func TestOltInfoCache_UpdatePartial(t *testing.T) {
	cache := NewOltInfoCache(5*time.Minute, 1*time.Minute)
	defer cache.StopGC()

	oltInfo := &OltInfo{
		OamIP: "192.168.1.1",
		NTInfo: NTInfo{
			Component: Component{
				Name: "NT-1",
				Type: "NetworkTermination",
			},
		},
	}

	cache.Set(oltInfo)

	// 部分更新
	success := cache.UpdatePartial("192.168.1.1", func(info *OltInfo) {
		info.NTInfo.Component.Name = "NT-Updated"
	})

	if !success {
		t.Fatal("Expected UpdatePartial to succeed")
	}

	// 验证更新
	retrieved, _ := cache.Get("192.168.1.1")
	if retrieved.NTInfo.Component.Name != "NT-Updated" {
		t.Errorf("Expected Name to be NT-Updated, got %s", retrieved.NTInfo.Component.Name)
	}

	// 测试不存在的key
	success = cache.UpdatePartial("192.168.1.99", func(info *OltInfo) {
		info.NTInfo.Component.Name = "Should not update"
	})

	if success {
		t.Error("Expected UpdatePartial to fail for non-existent key")
	}
}

func TestOltInfoCache_Delete(t *testing.T) {
	cache := NewOltInfoCache(5*time.Minute, 1*time.Minute)
	defer cache.StopGC()

	oltInfo := &OltInfo{
		OamIP: "192.168.1.1",
	}

	cache.Set(oltInfo)

	// 删除
	deleted := cache.Delete("192.168.1.1")
	if !deleted {
		t.Error("Expected Delete to return true")
	}

	// 验证已删除
	_, exists := cache.Get("192.168.1.1")
	if exists {
		t.Error("Expected OltInfo to be deleted from cache")
	}

	// 删除不存在的key
	deleted = cache.Delete("192.168.1.99")
	if deleted {
		t.Error("Expected Delete to return false for non-existent key")
	}
}

func TestOltInfoCache_GetAll(t *testing.T) {
	cache := NewOltInfoCache(5*time.Minute, 1*time.Minute)
	defer cache.StopGC()

	cache.Set(&OltInfo{OamIP: "192.168.1.1"})
	cache.Set(&OltInfo{OamIP: "192.168.1.2"})
	cache.Set(&OltInfo{OamIP: "192.168.1.3"})

	all := cache.GetAll()
	if len(all) != 3 {
		t.Errorf("Expected 3 entries, got %d", len(all))
	}

	if _, exists := all["192.168.1.1"]; !exists {
		t.Error("Expected 192.168.1.1 to be in result")
	}
}

func TestOltInfoCache_Count(t *testing.T) {
	cache := NewOltInfoCache(5*time.Minute, 1*time.Minute)
	defer cache.StopGC()

	if cache.Count() != 0 {
		t.Errorf("Expected count to be 0, got %d", cache.Count())
	}

	cache.Set(&OltInfo{OamIP: "192.168.1.1"})
	cache.Set(&OltInfo{OamIP: "192.168.1.2"})

	if cache.Count() != 2 {
		t.Errorf("Expected count to be 2, got %d", cache.Count())
	}
}

func TestOltInfoCache_Clear(t *testing.T) {
	cache := NewOltInfoCache(5*time.Minute, 1*time.Minute)
	defer cache.StopGC()

	cache.Set(&OltInfo{OamIP: "192.168.1.1"})
	cache.Set(&OltInfo{OamIP: "192.168.1.2"})

	cache.Clear()

	if cache.Count() != 0 {
		t.Errorf("Expected count to be 0 after Clear, got %d", cache.Count())
	}
}

func TestOltInfoCache_GC(t *testing.T) {
	// 使用短TTL和GC间隔进行测试
	cache := NewOltInfoCache(100*time.Millisecond, 50*time.Millisecond)
	defer cache.StopGC()

	cache.Set(&OltInfo{OamIP: "192.168.1.1"})
	cache.Set(&OltInfo{OamIP: "192.168.1.2"})

	// 等待条目过期并被GC
	time.Sleep(200 * time.Millisecond)

	if cache.Count() != 0 {
		t.Errorf("Expected all entries to be GC'd, but found %d entries", cache.Count())
	}
}

func TestOltInfoCache_GetExpiredEntries(t *testing.T) {
	cache := NewOltInfoCache(100*time.Millisecond, 1*time.Hour) // 长GC间隔，手动检查过期
	defer cache.StopGC()

	cache.Set(&OltInfo{OamIP: "192.168.1.1"})
	cache.Set(&OltInfo{OamIP: "192.168.1.2"})

	// 等待条目过期
	time.Sleep(150 * time.Millisecond)

	expired := cache.GetExpiredEntries()
	if len(expired) != 2 {
		t.Errorf("Expected 2 expired entries, got %d", len(expired))
	}
}

func TestOltInfoCache_SetTTL(t *testing.T) {
	cache := NewOltInfoCache(5*time.Minute, 1*time.Minute)
	defer cache.StopGC()

	newTTL := 10 * time.Minute
	cache.SetTTL(newTTL)

	if cache.GetTTL() != newTTL {
		t.Errorf("Expected TTL to be %v, got %v", newTTL, cache.GetTTL())
	}
}

func TestOltInfoCache_Exists(t *testing.T) {
	cache := NewOltInfoCache(5*time.Minute, 1*time.Minute)
	defer cache.StopGC()

	cache.Set(&OltInfo{OamIP: "192.168.1.1"})

	if !cache.Exists("192.168.1.1") {
		t.Error("Expected entry to exist")
	}

	if cache.Exists("192.168.1.99") {
		t.Error("Expected entry not to exist")
	}
}

func TestOltInfoCache_GetEntry(t *testing.T) {
	cache := NewOltInfoCache(5*time.Minute, 1*time.Minute)
	defer cache.StopGC()

	cache.Set(&OltInfo{OamIP: "192.168.1.1"})

	entry, exists := cache.GetEntry("192.168.1.1")
	if !exists {
		t.Fatal("Expected entry to exist")
	}

	if entry.Data.OamIP != "192.168.1.1" {
		t.Errorf("Expected OamIP to be 192.168.1.1, got %s", entry.Data.OamIP)
	}

	if entry.CreatedAt.IsZero() {
		t.Error("Expected CreatedAt to be set")
	}

	if entry.LastUpdateAt.IsZero() {
		t.Error("Expected LastUpdateAt to be set")
	}
}

func TestOltInfoCache_ConcurrentAccess(t *testing.T) {
	cache := NewOltInfoCache(5*time.Minute, 1*time.Minute)
	defer cache.StopGC()

	// 并发写入
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(id int) {
			for j := 0; j < 100; j++ {
				cache.Set(&OltInfo{OamIP: "192.168.1.1"})
			}
			done <- true
		}(i)
	}

	// 并发读取
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				cache.Get("192.168.1.1")
			}
			done <- true
		}()
	}

	// 等待所有goroutine完成
	for i := 0; i < 20; i++ {
		<-done
	}

	// 验证数据一致性
	if !cache.Exists("192.168.1.1") {
		t.Error("Expected entry to exist after concurrent access")
	}
}
