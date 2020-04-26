package dblayer

import (
	"yuting/daily-pg/2020-04-25/models"
)

func (db *DBORM) GetAllTeams() (team []models.Team, err error) {
	return team, db.Find(&team).Error
}

func (db *DBORM) AddOneTeam(team models.Team) (models.Team, error) {
	return team, db.Create(&team).Error
}
