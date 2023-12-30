package repositories

import (
	db "api/src/config"
	"api/src/config/logger"
)

type IpRepository struct{}

func (ir *IpRepository) Create(ip string, agent string) (err error) {
	query := `
		INSERT INTO death_note (ip, user_agent)
		VALUES ($1, $2);
	`
	_, err = db.Inst.Exec(query, ip, agent)
	if err != nil {
		logger.Error("Error Query:", err)
		return
	}

	return
}
