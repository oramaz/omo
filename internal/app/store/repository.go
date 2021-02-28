package store

import (
	uuid "github.com/satori/go.uuid"
	"gitlab.com/oramaz/omo/internal/app/model"
)

type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByUsername(string) (*model.User, error)
	ConnectSchool(*model.User, *model.School) error
}

type SchoolRepository interface {
	Find(uuid.UUID) (*model.School, error)
}
