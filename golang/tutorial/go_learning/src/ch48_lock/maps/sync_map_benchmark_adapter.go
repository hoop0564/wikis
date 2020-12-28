package maps

import "sync"

type SyncMapBenchmarkAdapter struct {
	m sync.Map // 读多写少的map
}

func CreateSyncMapBenchmarkAdapter() *SyncMapBenchmarkAdapter {
	return &SyncMapBenchmarkAdapter{}
}

func (s *SyncMapBenchmarkAdapter) Set(key interface{}, val interface{}) {

	s.m.Store(key, val)
}
func (s *SyncMapBenchmarkAdapter) Get(key interface{}) (interface{}, bool) {
	return s.m.Load(key)
}

func (s *SyncMapBenchmarkAdapter) Del(key interface{}) {
	s.m.Delete(key)
}
