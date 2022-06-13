package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const authorizationHeader = "Authorization"
const userCtx = "userId"

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		c.AbortWithError(http.StatusUnauthorized, errors.New("empty auth header"))
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.AbortWithError(http.StatusUnauthorized, errors.New("invalid auth header"))
		return
	}

	userId, err := h.services.ParseToken(headerParts[1])
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return 0, errors.New("user not found")
	}

	idInt, ok := id.(int)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return 0, errors.New("user not found")
	}

	return idInt, nil
}
