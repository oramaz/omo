package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RenderSchoolConnecting() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "connect_school", gin.H{
			"title": "Connect School",
		})
	}
}

func (h *Handler) RenderIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{
			"title": "Index Page",
		})
	}
}

func (h *Handler) RenderLol() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{
			"title": "Lol Page",
		})
	}
}

func (h *Handler) RenderSignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		return_to := c.Request.URL.Query().Get(ReturnToKey)
		c.HTML(http.StatusOK, "sign_up", gin.H{
			"title":     "Sign Up Page",
			"return_to": return_to,
		})
	}
}

func (h *Handler) RenderLogin() gin.HandlerFunc {

	return func(c *gin.Context) {
		return_to := c.Request.URL.Query().Get(ReturnToKey)
		c.HTML(http.StatusOK, "login", gin.H{
			"title":     "Login Page",
			"return_to": return_to,
		})
	}
}
