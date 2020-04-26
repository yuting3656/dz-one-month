package handler

import (
	"yuting/daily-pg/2020-04-25/dblayer"

	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	GetAllTeams(c *gin.Context)
	AddOneTeam(c *gin.Context)
}

type Handler struct {
	db dblayer.DBLayer
}

func NewHandler(dbName string, conn string) (*Handler, error) {
	db, err := dblayer.NewORM(dbName, conn)
	if err != nil {
		return nil, err
	}

	return &Handler{
		db: db,
	}, nil
}
