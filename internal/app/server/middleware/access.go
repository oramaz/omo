package middleware

import (
	"net/http"
	"net/url"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	h "gitlab.com/oramaz/omo/internal/app/server/handler"
)

func goAwayFromPrivate(c *gin.Context) {
	q := url.Values{}
	p := c.Request.URL.Path
	if c.Request.Method == "GET" && p != "/" {
		q.Set(h.ReturnToKey, p)
	}
	endpoint := url.URL{Path: "/login", RawQuery: q.Encode()}
	c.Redirect(http.StatusFound, endpoint.String())
	c.Abort()
}

func (mw *Middleware) Private() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		id := session.Get(h.SessionMasterKey)
		if id == nil {
			goAwayFromPrivate(c)
			return
		}

		u, err := st.User().Find(id.(int))
		if err != nil {
			goAwayFromPrivate(c)
			return
		}

		c.Set(h.CtxUserKey, u)
		c.Next()
	}
}

func (mw *Middleware) Guest() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		id := session.Get(h.SessionMasterKey)
		if id == nil {
			c.Next()
			return
		}
		u, err := st.User().Find(id.(int))
		if err == nil {
			c.Redirect(http.StatusFound, "/")
			c.Abort()
			return
		}

		c.Set(h.CtxUserKey, u)
		c.Next()
	}
}
