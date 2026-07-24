package collector

import (
	"sync"
	"time"
)

// OltInfoCacheEntry 缓存条目，包含OltInfo及其元数据
type OltInfoCacheEntry struct {
	Data         *OltInfo  `json:"data"`
	CreatedAt    time.Time `json:"created_at"`
	LastUpdateAt time.Time `json:"last_update_at"`
}

// OltInfoCache OltInfo的内存缓存
type OltInfoCache struct {
	mu      sync.RWMutex
	entries map[string]*OltInfoCacheEntry // key: OamIP
	ttl     time.Duration                 // 过期时间
	gcStop  chan struct{}                 // 用于停止GC的通道
}

// NewOltInfoCache 创建一个新的OltInfo缓存
// ttl: 缓存条目的生存时间，超过此时间未更新的条目将被删除
// gcInterval: GC运行的时间间隔
func NewOltInfoCache(ttl time.Duration, gcInterval time.Duration) *OltInfoCache {
	cache := &OltInfoCache{
		entries: make(map[string]*OltInfoCacheEntry),
		ttl:     ttl,
		gcStop:  make(chan struct{}),
	}

	// 启动后台GC协程
	go cache.startGC(gcInterval)

	return cache
}

// Set 设置或创建一个缓存条目
func (c *OltInfoCache) Set(oltInfo *OltInfo) {
	if oltInfo == nil {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	if entry, exists := c.entries[oltInfo.OamIP]; exists {
		// 更新现有条目
		entry.Data = oltInfo
		entry.LastUpdateAt = now
	} else {
		// 创建新条目
		c.entries[oltInfo.OamIP] = &OltInfoCacheEntry{
			Data:         oltInfo,
			CreatedAt:    now,
			LastUpdateAt: now,
		}
	}
}

// Get 根据OamIP获取OltInfo
func (c *OltInfoCache) Get(oamIP string) (*OltInfo, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, exists := c.entries[oamIP]
	if !exists {
		return nil, false
	}

	return entry.Data, true
}

// GetEntry 根据OamIP获取完整的缓存条目（包含元数据）
func (c *OltInfoCache) GetEntry(oamIP string) (*OltInfoCacheEntry, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, exists := c.entries[oamIP]
	if !exists {
		return nil, false
	}

	// 返回副本以避免外部修改
	return &OltInfoCacheEntry{
		Data:         entry.Data,
		CreatedAt:    entry.CreatedAt,
		LastUpdateAt: entry.LastUpdateAt,
	}, true
}

// Update 更新已存在的OltInfo，如果不存在则创建
func (c *OltInfoCache) Update(oltInfo *OltInfo) {
	c.Set(oltInfo)
}

// UpdatePartial 部分更新OltInfo的某些字段
func (c *OltInfoCache) UpdatePartial(oamIP string, updateFunc func(*OltInfo)) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exists := c.entries[oamIP]
	if !exists {
		return false
	}

	updateFunc(entry.Data)
	entry.LastUpdateAt = time.Now()
	return true
}

// Delete 删除指定的缓存条目
func (c *OltInfoCache) Delete(oamIP string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.entries[oamIP]; exists {
		delete(c.entries, oamIP)
		return true
	}
	return false
}

// GetAll 获取所有缓存的OltInfo
func (c *OltInfoCache) GetAll() map[string]*OltInfo {
	c.mu.RLock()
	defer c.mu.RUnlock()

	result := make(map[string]*OltInfo, len(c.entries))
	for key, entry := range c.entries {
		result[key] = entry.Data
	}
	return result
}

// GetAllEntries 获取所有缓存条目（包含元数据）
func (c *OltInfoCache) GetAllEntries() map[string]*OltInfoCacheEntry {
	c.mu.RLock()
	defer c.mu.RUnlock()

	result := make(map[string]*OltInfoCacheEntry, len(c.entries))
	for key, entry := range c.entries {
		result[key] = &OltInfoCacheEntry{
			Data:         entry.Data,
			CreatedAt:    entry.CreatedAt,
			LastUpdateAt: entry.LastUpdateAt,
		}
	}
	return result
}

// Exists 检查指定的OamIP是否存在于缓存中
func (c *OltInfoCache) Exists(oamIP string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	_, exists := c.entries[oamIP]
	return exists
}

// Count 返回缓存中的条目数量
func (c *OltInfoCache) Count() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.entries)
}

// Clear 清空所有缓存条目
func (c *OltInfoCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries = make(map[string]*OltInfoCacheEntry)
}

// startGC 启动后台垃圾回收协程
func (c *OltInfoCache) startGC(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.gc()
		case <-c.gcStop:
			return
		}
	}
}

// gc 执行垃圾回收，删除过期的缓存条目
func (c *OltInfoCache) gc() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	expiredKeys := make([]string, 0)

	// 找出所有过期的key
	for key, entry := range c.entries {
		if now.Sub(entry.LastUpdateAt) > c.ttl {
			expiredKeys = append(expiredKeys, key)
		}
	}

	// 删除过期条目
	for _, key := range expiredKeys {
		delete(c.entries, key)
	}
}

// StopGC 停止垃圾回收协程
func (c *OltInfoCache) StopGC() {
	close(c.gcStop)
}

// GetExpiredEntries 获取所有已过期但尚未被GC清理的条目
func (c *OltInfoCache) GetExpiredEntries() map[string]*OltInfoCacheEntry {
	c.mu.RLock()
	defer c.mu.RUnlock()

	now := time.Now()
	expired := make(map[string]*OltInfoCacheEntry)

	for key, entry := range c.entries {
		if now.Sub(entry.LastUpdateAt) > c.ttl {
			expired[key] = &OltInfoCacheEntry{
				Data:         entry.Data,
				CreatedAt:    entry.CreatedAt,
				LastUpdateAt: entry.LastUpdateAt,
			}
		}
	}

	return expired
}

// SetTTL 动态设置TTL
func (c *OltInfoCache) SetTTL(ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.ttl = ttl
}

// GetTTL 获取当前TTL设置
func (c *OltInfoCache) GetTTL() time.Duration {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.ttl
}

var OltInfoCacheList *OltInfoCache

func GetOltInfoCacheList() *OltInfoCache {
	return OltInfoCacheList
}

func init() {
	OltInfoCacheList = NewOltInfoCache(24*time.Hour, 24*time.Hour)
}
