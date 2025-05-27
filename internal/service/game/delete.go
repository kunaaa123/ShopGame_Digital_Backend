package game

func (g *gameServiceServerImpl) Delete(id uint) error {
	return g.repo.Delete(id)
}
