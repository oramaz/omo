package sqlstore

import (
	"database/sql"

	"gitlab.com/oramaz/omo/internal/app/model"
	"gitlab.com/oramaz/omo/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) ConnectSchool(u *model.User, s *model.School) error {
	if err := u.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"UPDATE users SET school_id = $1 WHERE id = $2 RETURNING id",
		s.ID,
		u.ID,
	).Scan(&u.ID)
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO users (username, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Username,
		u.EncryptedPassword,
	).Scan(&u.ID)
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, username, encrypted_password FROM users WHERE id = $1",
		id,
	).Scan(
		&u.ID,
		&u.Username,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}

func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, username, encrypted_password FROM users WHERE username = $1",
		username,
	).Scan(
		&u.ID,
		&u.Username,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}
