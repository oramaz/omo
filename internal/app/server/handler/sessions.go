package handler

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h *Handler) CreateSession() gin.HandlerFunc {
	type request struct {
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
	}
	return func(c *gin.Context) {
		session := sessions.Default(c)
		req := request{}
		if err := c.MustBindWith(&req, binding.FormMultipart); err != nil {
			h.error(c, http.StatusBadRequest, err)
			return
		}
		u, err := st.User().FindByUsername(req.Username)
		if err != nil || !u.ComparePassword(req.Password) {
			h.error(c, http.StatusUnauthorized, errors.New("Record not found"))
			return
		}
		session.Set(SessionMasterKey, u.ID)
		session.Save()

		return_to := c.Request.URL.Query().Get(ReturnToKey)
		c.JSON(http.StatusOK, gin.H{
			ReturnToKey: return_to,
		})
	}
}

func (h *Handler) ClearSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()

		c.Status(http.StatusOK)
	}
}
