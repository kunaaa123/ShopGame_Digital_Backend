package models

func (g *Game) ToDB() *GameDB {
	if g == nil {
		return nil
	}

	return &GameDB{
		Name:        g.Name,
		Description: g.Description,
		Price:       g.Price,
		Image:       g.Image,
	}
}

func (g *GameDB) ToDomain() *Game {
	if g == nil {
		return nil
	}

	return &Game{
		ID:          g.ID,
		Name:        g.Name,
		Description: g.Description,
		Price:       g.Price,
		Image:       g.Image,
	}
}
