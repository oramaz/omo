package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/oramaz/omo/internal/app/model"
)

func (h *Handler) ConnectUserWithSchool() gin.HandlerFunc {
	type request struct {
		SchoolID string `json:"schoolID" binding:"required"`
	}
	return func(c *gin.Context) {
		req := &request{}
		if err := c.ShouldBindJSON(req); err != nil {
			h.error(c, http.StatusBadRequest, err)
			return
		}
		u := c.Value(CtxUserKey).(*model.User)
		school_id_uuid, err := uuid.FromString(req.SchoolID)
		if err != nil {
			h.error(c, http.StatusBadRequest, errors.New("School id have incorrect type"))
			return
		}

		if u.School.ID == school_id_uuid {
			c.Status(http.StatusOK)
			return
		}
		school, err := st.School().Find(school_id_uuid)
		if err != nil {
			h.error(c, http.StatusNotFound, errors.New("School id not found"))
			return
		}
		u.School = *school
		if err := st.User().ConnectSchool(u, school); err != nil {
			h.error(c, http.StatusUnprocessableEntity, err)
			return
		}

		c.Status(http.StatusOK)
	}
}
