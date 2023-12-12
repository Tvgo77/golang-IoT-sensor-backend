package dataChannel

import (
	"sync"
)

type ChannelMap struct {
	mu    sync.Mutex
	table map[string]chan int32
}

func NewChannelMap() *ChannelMap {
	return &ChannelMap{
		table: make(map[string]chan int32),
	}
}

func (chanMap *ChannelMap) Insert(serialNum string, ch chan int32) {
	chanMap.mu.Lock()
	chanMap.table[serialNum] = ch
	chanMap.mu.Unlock()
}

func (chanMap *ChannelMap) Delete(serialNum string) {
	chanMap.mu.Lock()
	delete(chanMap.table, serialNum)
	chanMap.mu.Unlock()
}

func (chanMap *ChannelMap) GetChannel(serialNum string) chan int32 {
	chanMap.mu.Lock()
	ch := chanMap.table[serialNum]
	chanMap.mu.Unlock()
	return ch
}
