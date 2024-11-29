package mariuspg

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	pgDriver = "postgres"
)

var (
	pgDSNTemplate = `host=%s port=%d dbname=%s user=%s password=%s sslmode=disable`
)

type ConfigPG struct {
	Host string `yaml:"host"`
	Port int32  `yaml:"port"`
	Base string `yaml:"base"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
}

type PGAgent struct {
	db *sql.DB
}

func NewPGAgent(config ConfigPG) (*PGAgent, error) {
	dsn := fmt.Sprintf(
		pgDSNTemplate,
		config.Host,
		config.Port,
		config.Base,
		config.User,
		config.Pass,
	)
	db, err := sql.Open(pgDriver, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed create database connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	agent := &PGAgent{
		db: db,
	}

	err = agent.initDB()
	if err != nil {
		return nil, fmt.Errorf("failed init database: %v", err)
	}

	return agent, nil
}
