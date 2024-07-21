package service

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/jo-hoe/routriever/app/config"
)

type SqlService interface {
	Save(config.Route, int) error
}

type PostgresService struct {
	psqlInfo string
}

func NewPostgresService(host string, port int, dbname string, user string, password string) *PostgresService {
	return &PostgresService{
		psqlInfo: fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname),
	}
}

func (s *PostgresService) Save(config config.Route, valueInSeconds int) (err error) {
	db, err := sql.Open("postgres", s.psqlInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	return err
}
