package handler

import (
	"database/sql"
	"log/slog"
	"restaurant/Storage/postgres"
	"restaurant/pkg/logger"
)

type Handler struct {
	logger   slog.Logger
	UserRepo *postgres.UserRepo
}

func NewHandler(db *sql.DB) *Handler {
	logger := logger.NewLogger()
	return &Handler{
		UserRepo: postgres.NewUserRepo(db),
		logger:   logger,
	}
}
