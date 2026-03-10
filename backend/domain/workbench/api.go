package workbench

import (
	"backend/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (d *Domain) ApiOverview(c *gin.Context) {
	role := c.DefaultQuery("role", "store_manager")
	utils.RespSuccess(c, d.store.Overview(role))
}

func (d *Domain) ApiTransition(c *gin.Context) {
	cardID := c.Param("id")
	var req TransitionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespError(c, http.StatusBadRequest, "invalid request body")
		return
	}

	card, err := d.store.Transition(cardID, req.ToState, req.Reason)
	if err != nil {
		if errors.Is(err, ErrCardNotFound) {
			utils.RespError(c, http.StatusNotFound, err.Error())
			return
		}
		if errors.Is(err, ErrInvalidTransition) {
			utils.RespError(c, http.StatusBadRequest, err.Error())
			return
		}
		utils.RespError(c, http.StatusInternalServerError, "failed to transition action")
		return
	}

	utils.RespSuccess(c, card)
}

func (d *Domain) ApiEvents(c *gin.Context) {
	cardID := c.Param("id")
	events, err := d.store.Events(cardID)
	if err != nil {
		utils.RespError(c, http.StatusInternalServerError, "failed to load events")
		return
	}
	utils.RespSuccess(c, events)
}
