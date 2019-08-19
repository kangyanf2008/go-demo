package service_map

import "sync"

type Safe2Map struct {
	sync.RWMutex
	Map map[string]int
}

func NewSafe2Map() *Safe2Map {
	sm := new(Safe2Map)
	sm.Map = make(map[string]int)
	return sm
}

func (sm *Safe2Map) ReadMap(key string) int {
	value := sm.Map[key]
	return value
}

func (sm *Safe2Map) WriteMap(key string, value int) {
	sm.Map[key] = value
}