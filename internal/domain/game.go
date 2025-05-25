package domain

type GameService interface {
	GetAll() ([]Game, error)
	GetByID(id int) (Game, error)
	Create(game Game) (Game, error)
	Update(id int, game Game) (Game, error)
	Delete(id int) error
}

type GameRepository interface {
	GetAll() ([]Game, error)
	GetByID(id uint) (Game, error)
	Create(game *Game) error
	Update(game *Game) error
	Delete(id uint) error
}

type Game struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
}
