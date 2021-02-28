package app

import (
	"database/sql"

	"github.com/gin-contrib/sessions"

	"github.com/gin-contrib/sessions/cookie"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"gitlab.com/oramaz/omo/internal/app/config"
	templates "gitlab.com/oramaz/omo/internal/app/server"
	"gitlab.com/oramaz/omo/internal/app/server/router"
	"gitlab.com/oramaz/omo/internal/app/store"
	"gitlab.com/oramaz/omo/internal/app/store/sqlstore"
)

func Start(config *config.Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	var st store.Store = sqlstore.New(db)
	var sessionStore sessions.Store = cookie.NewStore([]byte(config.SessionKey))
	sessionStore.Options(sessions.Options{
		Secure: true,
	})
	r := router.New(&st, &sessionStore)
	templates.PrepareRouter(r)

	return r.Run(config.BindAddr)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
