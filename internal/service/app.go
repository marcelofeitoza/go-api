package service

import (
	"github.com/jmoiron/sqlx"
)

type App struct {
	DB *sqlx.DB
}

func NewApp(db *sqlx.DB) *App {
	return &App{DB: db}
}
