package service

import (
	"fmt"
	"strings"

	"github.com/kunaaa123/SHOPGAME/internal/domain"
)

type GameService struct {
	repo domain.GameRepository
}

func NewgameService(repo domain.GameRepository) *GameService {
	return &GameService{
		repo: repo,
	}
}

func (s *GameService) GetAll() ([]domain.Game, error) {
	return s.repo.GetAll()
}

func (s *GameService) GetByID(id int) (domain.Game, error) {
	return s.repo.GetByID(uint(id))
}

func (s *GameService) Create(game domain.Game) (domain.Game, error) {
	// Validate name
	if game.Name == "" {
		return domain.Game{}, fmt.Errorf("name cannot be empty")
	}

	// Validate price
	if game.Price <= 0 {
		return domain.Game{}, fmt.Errorf("price must be greater than 0")
	}

	// Validate image URL (เพิ่มการตรวจสอบรูปแบบ URL ตามต้องการ)
	if !strings.HasSuffix(game.Image, ".jpg") && !strings.HasSuffix(game.Image, ".png") {
		return domain.Game{}, fmt.Errorf("invalid image format")
	}

	err := s.repo.Create(&game)
	return game, err
}

func (s *GameService) Update(id int, game domain.Game) (domain.Game, error) {
	game.ID = uint(id)
	err := s.repo.Update(&game)
	return game, err
}

func (s *GameService) Delete(id int) error {
	return s.repo.Delete(uint(id))
}
