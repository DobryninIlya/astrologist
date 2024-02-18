package sqlstore

import (
	"database/sql"
)

// Store ...
type Store struct {
	db                   *sql.DB
	userRepository       *UserRepository
	natalChartRepository *NatalChartRepository
}

// New ...
func New(db *sql.DB) Store {
	return Store{
		db: db,
	}
}

// User ...
func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

func (s *Store) NatalChart() *NatalChartRepository {
	if s.natalChartRepository != nil {
		return s.natalChartRepository
	}

	s.natalChartRepository = &NatalChartRepository{
		store: s,
	}

	return s.natalChartRepository
}
