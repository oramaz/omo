package middleware

import (
	"gitlab.com/oramaz/omo/internal/app/store"
)

var st store.Store

type Middleware struct{}

func New(store *store.Store) *Middleware {
	st = *store
	return &Middleware{}
}
