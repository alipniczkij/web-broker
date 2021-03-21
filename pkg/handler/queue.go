package handler

import (
	"net/http"

	broker "github.com/alipniczkij/web-broker"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getValue(c *gin.Context) {
	value, err := h.services.QueueValue.Get()
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
	}
	c.JSON(http.StatusOK, value)
}

func (h *Handler) putValue(c *gin.Context) {
	var input broker.QueueValue
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.QueueValue.Put(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, true)
}
