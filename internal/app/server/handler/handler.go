package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/oramaz/omo/internal/app/store"
)

var st store.Store

const (
	ReturnToKey      = "return_to"
	CtxUserKey       = "user"
	SessionMasterKey = "user_id"
)

type Handler struct{}

func New(store *store.Store) *Handler {
	st = *store
	return &Handler{}
}

func (h *Handler) error(c *gin.Context, code int, e error) {
	c.AbortWithStatusJSON(code, gin.H{
		"error": e.Error(),
	})
}
