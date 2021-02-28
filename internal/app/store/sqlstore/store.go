package sqlstore

import (
	"database/sql"

	"gitlab.com/oramaz/omo/internal/app/store"
)

type Store struct {
	db               *sql.DB
	userRepository   *UserRepository
	schoolRepository *SchoolRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (st *Store) User() store.UserRepository {
	if st.userRepository != nil {
		return st.userRepository
	}

	st.userRepository = &UserRepository{
		store: st,
	}

	return st.userRepository
}

func (st *Store) School() store.SchoolRepository {
	if st.schoolRepository != nil {
		return st.schoolRepository
	}

	st.schoolRepository = &SchoolRepository{
		store: st,
	}

	return st.schoolRepository
}
