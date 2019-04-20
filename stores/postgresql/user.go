package postgresql

import "github.com/jmoiron/sqlx"

type UserStore struct {
	db *sqlx.DB
}

func NewUserStore(db *sqlx.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (u *UserStore) migrate() {
	u.db.MustExec(userCreateTable)
}
