package dblayer

import (
	"errors"
	"yuting/daily-pg/2020-04-25/models"
)

type DBLayer interface {
	GetAllTeams() ([]models.Team, error)
	AddOneTeam(models.Team) (models.Team, error)
}

var ErrINVALIDPASSWORD = errors.New("Invalid password")
