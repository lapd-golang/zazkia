package main

import (
	"errors"
	"log"
	"sync"
)

type linkManager struct {
	links map[int]*link
	mutex *sync.Mutex
}

func newLinkManager() *linkManager {
	return &linkManager{
		links: map[int]*link{},
		mutex: new(sync.Mutex),
	}
}

func (m *linkManager) add(l *link) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.links[l.ID] = l
}

func (m *linkManager) remove(l *link) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(m.links, l.ID)
}

func (m *linkManager) sortedLinks() []*link {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	sorted := []*link{}
	for _, each := range m.links {
		sorted = append(sorted, each)
	}
	return sorted
}

func (m *linkManager) get(id int) *link {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	l, _ := m.links[id]
	return l
}

func (m *linkManager) disconnectAndRemove(id int) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	link, ok := m.links[id]
	if !ok {
		return errors.New("no link with id")
	}
	log.Println("disconnecting", link.ID)
	link.disconnect()
	delete(m.links, link.ID)
	return nil
}

func (m *linkManager) close() {
	for _, each := range m.links {
		m.disconnectAndRemove(each.ID)
	}
}