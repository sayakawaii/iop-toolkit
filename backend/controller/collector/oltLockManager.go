package collector

import "sync"

type OLTLockManager struct {
	mu    sync.Mutex
	locks map[string]*sync.Mutex
}

func NewOLTLockManager() *OLTLockManager {
	return &OLTLockManager{
		locks: make(map[string]*sync.Mutex),
	}
}

func (m *OLTLockManager) Lock(oamIP string) {
	m.mu.Lock()
	l, ok := m.locks[oamIP]
	if !ok {
		l = &sync.Mutex{}
		m.locks[oamIP] = l
	}
	m.mu.Unlock()

	l.Lock()
}

func (m *OLTLockManager) Unlock(oamIP string) {
	m.mu.Lock()
	if l, ok := m.locks[oamIP]; ok {
		l.Unlock()
	}
	m.mu.Unlock()
}
