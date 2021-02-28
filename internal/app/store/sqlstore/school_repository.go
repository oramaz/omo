package sqlstore

import (
	"database/sql"

	uuid "github.com/satori/go.uuid"
	"gitlab.com/oramaz/omo/internal/app/model"
	"gitlab.com/oramaz/omo/internal/app/store"
)

type SchoolRepository struct {
	store *Store
}

func (r *SchoolRepository) Find(id uuid.UUID) (*model.School, error) {
	s := &model.School{}
	if err := r.store.db.QueryRow(
		"SELECT id, name, city FROM schools WHERE id = $1",
		id,
	).Scan(
		&s.ID,
		&s.Name,
		&s.City,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return s, nil
}
