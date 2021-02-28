package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gitlab.com/oramaz/omo/internal/app/model"
)

func (h *Handler) CreateUser() gin.HandlerFunc {
	type request struct {
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
	}
	return func(c *gin.Context) {
		req := request{}
		if err := c.MustBindWith(&req, binding.FormMultipart); err != nil {
			h.error(c, http.StatusBadRequest, err)
			return
		}
		u := &model.User{
			Username: req.Username,
			Password: req.Password,
		}
		if err := st.User().Create(u); err != nil {
			h.error(c, http.StatusUnprocessableEntity, err)
			return
		}

		u.Sanitize()

		return_to := c.Request.URL.Query().Get(ReturnToKey)
		redirect_url := "/sessions"
		if return_to != "" {
			redirect_url += "?return_to=" + return_to
		}

		c.Redirect(http.StatusTemporaryRedirect, redirect_url)
		c.Abort()
	}
}
