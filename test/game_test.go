package test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/kunaaa123/SHOPGAME/internal/domain"
	"github.com/kunaaa123/SHOPGAME/internal/service"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type mockGameRepository struct {
	games map[uint]domain.Game
}

func newMockGameRepository() *mockGameRepository {
	return &mockGameRepository{
		games: make(map[uint]domain.Game),
	}
}

func (m *mockGameRepository) GetAll() ([]domain.Game, error) {
	var games []domain.Game
	for _, game := range m.games {
		games = append(games, game)
	}
	return games, nil
}

func (m *mockGameRepository) GetByID(id uint) (domain.Game, error) {
	game, exists := m.games[id]
	if !exists {
		return domain.Game{}, gorm.ErrRecordNotFound
	}
	return game, nil
}

func (m *mockGameRepository) Create(game *domain.Game) error {
	game.ID = uint(len(m.games) + 1)
	m.games[game.ID] = *game
	return nil
}

func (m *mockGameRepository) Update(game *domain.Game) error {
	if _, exists := m.games[game.ID]; !exists {
		return gorm.ErrRecordNotFound
	}
	m.games[game.ID] = *game
	return nil
}

func (m *mockGameRepository) Delete(id uint) error {
	if _, exists := m.games[id]; !exists {
		return gorm.ErrRecordNotFound
	}
	delete(m.games, id)
	return nil
}

func TestGameService_GetAll(t *testing.T) {
	repo := newMockGameRepository()
	service := service.NewgameService(repo)

	testGame := domain.Game{
		Name:        "Test Game",
		Description: "Test Description",
		Price:       59.99,
		Image:       "test.jpg",
	}
	repo.Create(&testGame)

	games, err := service.GetAll()

	assert.NoError(t, err)
	assert.Len(t, games, 1)
	assert.Equal(t, testGame.Name, games[0].Name)
}

func TestGameService_GetByID(t *testing.T) {
	repo := newMockGameRepository()
	service := service.NewgameService(repo)

	testGame := domain.Game{
		Name:        "Test Game",
		Description: "Test Description",
		Price:       59.99,
		Image:       "test.jpg",
	}
	repo.Create(&testGame)

	game, err := service.GetByID(int(testGame.ID))
	assert.NoError(t, err)
	assert.Equal(t, testGame.Name, game.Name)

	_, err = service.GetByID(999)
	assert.Error(t, err)
}

func TestGameService_Create(t *testing.T) {
	repo := newMockGameRepository()
	service := service.NewgameService(repo)

	newGame := domain.Game{
		Name:        "New Game",
		Description: "New Description",
		Price:       29.99,
		Image:       "new.jpg",
	}

	created, err := service.Create(newGame)
	assert.NoError(t, err)
	assert.NotZero(t, created.ID)
	assert.Equal(t, newGame.Name, created.Name)
}

func TestGameService_Update(t *testing.T) {
	repo := newMockGameRepository()
	service := service.NewgameService(repo)

	testGame := domain.Game{
		Name:        "Test Game",
		Description: "Test Description",
		Price:       59.99,
		Image:       "test.jpg",
	}
	repo.Create(&testGame)

	updatedGame := domain.Game{
		Name:        "Updated Game",
		Description: "Updated Description",
		Price:       69.99,
		Image:       "updated.jpg",
	}

	result, err := service.Update(int(testGame.ID), updatedGame)
	assert.NoError(t, err)
	assert.Equal(t, updatedGame.Name, result.Name)

	_, err = service.Update(999, updatedGame)
	assert.Error(t, err)
}

func TestGameService_Delete(t *testing.T) {
	repo := newMockGameRepository()
	service := service.NewgameService(repo)

	testGame := domain.Game{
		Name:        "Test Game",
		Description: "Test Description",
		Price:       59.99,
		Image:       "test.jpg",
	}
	repo.Create(&testGame)

	err := service.Delete(int(testGame.ID))
	assert.NoError(t, err)

	err = service.Delete(999)
	assert.Error(t, err)
}

func TestGameService_CreateWithInvalidData(t *testing.T) {
	repo := newMockGameRepository()
	service := service.NewgameService(repo)

	testCases := []struct {
		name    string
		game    domain.Game
		wantErr bool
	}{
		{
			name: "Empty name",
			game: domain.Game{
				Name:        "",
				Description: "Test Description",
				Price:       59.99,
				Image:       "test.jpg",
			},
			wantErr: true,
		},
		{
			name: "Negative price",
			game: domain.Game{
				Name:        "Test Game",
				Description: "Test Description",
				Price:       -10.99,
				Image:       "test.jpg",
			},
			wantErr: true,
		},
		{
			name: "Zero price",
			game: domain.Game{
				Name:        "Test Game",
				Description: "Test Description",
				Price:       0,
				Image:       "test.jpg",
			},
			wantErr: true,
		},
		{
			name: "Invalid image URL",
			game: domain.Game{
				Name:        "Test Game",
				Description: "Test Description",
				Price:       59.99,
				Image:       "invalid-url",
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := service.Create(tc.game)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGameService_GetAllPerformance(t *testing.T) {
	repo := newMockGameRepository()
	service := service.NewgameService(repo)

	for i := 0; i < 1000; i++ {
		game := domain.Game{
			Name:        fmt.Sprintf("Game %d", i),
			Description: "Test Description",
			Price:       59.99,
			Image:       "test.jpg",
		}
		repo.Create(&game)
	}

	start := time.Now()
	games, err := service.GetAll()
	duration := time.Since(start)

	assert.NoError(t, err)
	assert.Len(t, games, 1000)
	assert.Less(t, duration, 2*time.Second)
}

func TestGameService_ConcurrentAccess(t *testing.T) {
	repo := newMockGameRepository()
	service := service.NewgameService(repo)

	testGame := domain.Game{
		Name:        "Test Game",
		Description: "Test Description",
		Price:       59.99,
		Image:       "test.jpg",
	}
	repo.Create(&testGame)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := service.GetByID(int(testGame.ID))
			assert.NoError(t, err)
		}()
	}
	wg.Wait()
}
