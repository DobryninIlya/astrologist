package models

import (
	"github.com/jackc/pgx/pgtype"
)

type TGUser struct {
	ID       int64
	Username string
	RegDate  pgtype.Date
}
