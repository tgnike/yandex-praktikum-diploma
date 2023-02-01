package authservice

import (
	"sync"

	"github.com/google/uuid"
)

type SimpleTokenRepository struct {
	mu   *sync.RWMutex
	repo map[string]string
}

func (s *SimpleTokenRepository) Add(userID string) string {

	id := uuid.New().String()

	s.mu.Lock()
	s.repo[id] = userID
	s.mu.Unlock()

	return id
}

func (s *SimpleTokenRepository) Get(token string) string {

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.repo[token]

}
