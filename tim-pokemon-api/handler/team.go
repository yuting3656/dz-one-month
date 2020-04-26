package handler

import (
	"net/http"
	"yuting/daily-pg/2020-04-25/models"

	"time"

	"github.com/gin-gonic/gin"
)

// GetAllTeams godoc
// @Summary get team
// @Description get all teams
// @Tags Teams
// @Produce json
// @Success 200 {object} models.Team "ok"
// @Router /teams [get]
func (h *Handler) GetAllTeams(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}
	team, err := h.db.GetAllTeams()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, team)
}

// GetAllTeams godoc
// @Summary create team
// @Description create one teams
// @Tags Teams
// @Produce json
// @Param team body models.Team true "Add team"
// @Success 200
// @Router /teams [post]
func (h *Handler) AddOneTeam(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}
	var team models.Team
	// check from request
	err := c.ShouldBindJSON(&team)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// interactive with db
	team.CreateTime = time.Now().UTC()
	team.UpdateTime = time.Now().UTC()
	team, err = h.db.AddOneTeam(team)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}
	c.JSON(http.StatusOK, team)
}
