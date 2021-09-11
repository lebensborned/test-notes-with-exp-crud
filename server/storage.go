package server

import (
	"kode-task/models"
	"sync"
	"time"
)

type Storage struct {
	sync.RWMutex
	defaultExpiration time.Duration
	cleanupInterval   time.Duration
	items             []models.Note
}

func NewStorage(defaultExpiration, cleanupInterval time.Duration) *Storage {

	items := []models.Note{}

	storage := Storage{
		items:             items,
		defaultExpiration: defaultExpiration,
		cleanupInterval:   cleanupInterval,
	}

	storage.StartCleanup()

	return &storage
}
func (s *Storage) Add(value string, duration time.Duration) {

	var expiration int64

	if duration == 0 {
		duration = s.defaultExpiration
	}

	if duration > 0 {
		expiration = time.Now().Add(duration).UnixNano()
	}
	s.Lock()
	item := models.Note{
		ID:         pseudo_uuid(),
		Value:      value,
		Expiration: expiration,
	}
	s.items = append(s.items, item)
	s.Unlock()
}
func (s *Storage) GetAll() []models.Note {
	s.RLock()
	defer s.RUnlock()
	return s.items
}
func (s *Storage) DeleteNoteByID(ID string) {
	s.Lock()
	defer s.Unlock()
	filteredItems := s.items[:0]
	for _, v := range s.items {
		if ID != v.ID {
			filteredItems = append(filteredItems, v)
		}
	}
	s.items = filteredItems
}
func (s *Storage) GetFirstNote() *models.Note {
	s.RLock()
	defer s.RUnlock()
	if len(s.items) == 0 {
		return nil
	}
	return &s.items[0]
}
func (s *Storage) GetLastNote() *models.Note {
	s.RLock()
	defer s.RUnlock()
	if len(s.items) == 0 {
		return nil
	}
	return &s.items[len(s.items)-1]
}
func (s *Storage) StartCleanup() {
	go s.GC()
}

func (s *Storage) GC() {

	for {
		<-time.After(s.cleanupInterval)
		filteredItems := s.items[:0]
		s.Lock()
		for _, v := range s.items {
			if time.Now().UnixNano() < v.Expiration || v.Expiration == 0 {
				filteredItems = append(filteredItems, v)
			}
		}
		s.items = filteredItems
		s.Unlock()
	}

}
