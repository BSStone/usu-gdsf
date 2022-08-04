package db

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jak103/usu-gdsf/log"
	"github.com/jak103/usu-gdsf/models"
)

var _ Database = (*Mock)(nil)

type Mock struct {
	games map[uuid.UUID]models.Game
	users map[uuid.UUID]models.User
}

func (db *Mock) GetGameByID(id uuid.UUID) (*models.Game, error) {
	if id.String() != "" {
		if game, ok := db.games[id]; ok {
			return &game, nil
		}
	}
	return nil, errors.New("mockdb: game not found")
}

func (db *Mock) GetGamesByTags(tags []string) ([]models.Game, error) {
	games := make(map[string]models.Game)

	for _, game := range db.games {
		for _, tag := range tags {
			valid := map[string]bool{}
			game_tags := make([]string, 0)
			for _, v := range game_tags {
				valid[v] = true
			}
			if valid[tag] {
				games[game.ID.String()] = game
			}
		}
	}
	game_slice := []models.Game{}
	for _, game := range games {
		game_slice = append(game_slice, game)
	}

	return game_slice, nil
}

func (db *Mock) GetAllGames() ([]models.Game, error) {
	games := make([]models.Game, 0)

	for _, game := range db.games {
		games = append(games, game)
	}

	return games, nil
}

func (d *Mock) GetGamesByPublishDate(startRange string, endRange string) ([]models.Game, error) {
	panic("not implemented") // TODO: Implement
}

func (d *Mock) CreateGame(newGame models.Game) error {
	panic("not implemented") // TODO: Implement
}

func (d *Mock) DeleteGame(id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

func (d *Mock) UpdateGame(updatedGame models.Game) error {
	panic("not implemented") // TODO: Implement
}

// Users
func (d *Mock) GetAllUsers() ([]models.User, error) {
	panic("not implemented") // TODO: Implement
}

func (db *Mock) GetUserByID(id uuid.UUID) (*models.User, error) {
	if id.String() != "" {
		if user, ok := db.users[id]; ok {
			return &user, nil
		}
	}
	return nil, errors.New("mockdb: game not found")
}

func (d *Mock) GetUsersByRole(role int64) ([]models.User, error) {
	panic("not implemented") // TODO: Implement
}

func (db *Mock) CreateUser(user models.User) error {
	log.Info("mock connect")
	db.users[user.ID] = user

	return nil
}

func (d *Mock) DeleteUser(id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

func (d *Mock) UpdateUser(updatedUser models.User) error {
	panic("not implemented") // TODO: Implement
}

// Ratings
func (d *Mock) GetRatingByID(id uuid.UUID) (*models.GameRating, error) {
	panic("not implemented") // TODO: Implement
}

func (d *Mock) GetRatingsByGame(gameID uuid.UUID) ([]models.GameRating, error) {
	panic("not implemented") // TODO: Implement
}

func (d *Mock) GetRatingsByUser(userID uuid.UUID) ([]models.GameRating, error) {
	panic("not implemented") // TODO: Implement
}

func (d *Mock) CreateRating(newRating models.GameRating) error {
	panic("not implemented") // TODO: Implement
}

func (d *Mock) DeleteRating(id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

func (d *Mock) DeleteRatingsByGame(gameID uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

func (d *Mock) updateRating(updatedRating models.GameRating) error {
	panic("not implemented") // TODO: Implement
}

func (db *Mock) Connect() error {
	log.Info("mock connect")

	if len(db.games) == 0 {
		games := make(map[uuid.UUID]models.Game)
		for _, v := range CreateGamesFromJson() {
			games[v.ID] = v
		}
		db.games = games
	}

	if len(db.users) == 0 {
		users := make(map[uuid.UUID]models.User)
		// for _, v := range CreateGamesFromJson() {
		// 	games[v.ID.String()] = v
		// }
		db.users = users
	}

	log.Debug("Sample Database Initialized")

	return nil
}

func (db Mock) Disconnect() error {
	return nil
}

// func init(){

// }
