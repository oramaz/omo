package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gitlab.com/oramaz/omo/internal/app/server/handler"
	"gitlab.com/oramaz/omo/internal/app/server/middleware"
	"gitlab.com/oramaz/omo/internal/app/store"
)

const (
	sessionName = "user_session"
)

func New(st *store.Store, sessionStore *sessions.Store) *gin.Engine {
	var h = handler.New(st)
	var mw = middleware.New(st)

	r := gin.Default()
	r.Use(cors.Default())
	r.Use(mw.SetRequestID())
	r.Use(sessions.Sessions(sessionName, *sessionStore))

	// Guest
	guest := r.Group("/")
	guest.Use(mw.Guest())
	{
		guest.GET("/login", h.RenderLogin())
		guest.GET("/signup", h.RenderSignUp())
		r.POST("/sessions", h.CreateSession()) // Login
		r.POST("/users", h.CreateUser())       // Sign up
	}

	// Private
	p := r.Group("/")
	p.Use(mw.Private())
	{
		p.GET("/", h.RenderIndex())
		p.GET("/lol", h.RenderLol())
		p.DELETE("/sessions", h.ClearSession()) // Logout

		p.GET("/connect/school/:school_id", h.RenderSchoolConnecting())
		p.POST("/connect/school", h.ConnectUserWithSchool())
	}

	return r
}
