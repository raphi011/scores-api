package scores

// PlayerService exposes various operations on Players
type PlayerService struct {
	Repository PlayerRepository
}

// ByGroup returns players of a group
func (s *PlayerService) ByGroup(groupID uint) ([]Player, error) {
	return s.Repository.ByGroup(groupID)
}

// Get retrieves a player
func (s *PlayerService) Get(playerID uint) (*Player, error) {
	return s.Repository.Get(playerID)
}

// All retrieves all players
func (s *PlayerService) All() ([]Player, error) {
	return s.Repository.All()
}

// Delete deletes a player
func (s *PlayerService) Delete(playerID uint) error {
	return s.Repository.Delete(playerID)
}

// Create creates a player
func (s *PlayerService) Create(player *Player) (*Player, error) {
	return s.Repository.Create(player)
}
