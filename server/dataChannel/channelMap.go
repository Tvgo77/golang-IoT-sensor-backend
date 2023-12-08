package dataChannel

import (
	"sync"
)

type ChannelMap struct {
	mu    sync.Mutex
	table map[int]chan int
}

func NewChannelMap() *ChannelMap {
	return &ChannelMap{
		table: make(map[int]chan int),
	}
}

func (chanMap *ChannelMap) Insert(serialNum int, ch chan int) {
	chanMap.mu.Lock()
	chanMap.table[serialNum] = ch
	chanMap.mu.Unlock()
}

func (chanMap *ChannelMap) Delete(serialNum int) {
	chanMap.mu.Lock()
	delete(chanMap.table, serialNum)
	chanMap.mu.Unlock()
}

func (chanMap *ChannelMap) GetChannel(serialNum int) chan int {
	chanMap.mu.Lock()
	ch := chanMap.table[serialNum]
	chanMap.mu.Unlock()
	return ch
}
