package sqlstore

import (
	"astrologist/internal/app/models"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create ...
func (r UserRepository) Create(u *models.TGUser) error {
	return r.store.db.QueryRow(
		"INSERT INTO tgusers (id, username, reg_date) VALUES ($1, $2, $3) RETURNING id",
		u.ID, u.Username, u.RegDate,
	).Scan(&u.ID)
}
