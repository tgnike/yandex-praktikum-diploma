package authservice

import (
	"sync"

	"github.com/google/uuid"
	"github.com/tgnike/yandex-praktikum-diploma/internal/models"
)

type SimpleTokenRepository struct {
	mu   *sync.RWMutex
	repo map[models.Token]models.UserID
}

func (s *SimpleTokenRepository) Add(userID models.UserID) models.Token {

	id := models.Token(uuid.New().String())

	s.mu.Lock()
	s.repo[id] = userID
	s.mu.Unlock()

	return id
}

func (s *SimpleTokenRepository) Get(token models.Token) models.UserID {

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.repo[token]

}
