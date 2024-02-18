package server

import (
	"astrologist/internal/app/migrations"
	"astrologist/internal/app/store/sqlstore"
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"log"
)

func Start(ctx context.Context, config *Config) (*App, error) {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return nil, err
	}
	store := sqlstore.New(db)
	migrations.MakeMigrations(db, logrus.New())
	if err != nil {
		log.Fatalf("Ошибка инициализации Firebase API: %v. Проверьте, находится ли serviceAccountKey.json в папке configs.", err.Error())
	}
	srv := newApp(ctx, &store, config.BindAddr, *config)
	//return http.ListenAndServe(config.BindAddr, srv)
	return srv, nil
}

func newDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
