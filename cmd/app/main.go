package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/marcelofeitoza/track-co-api/internal/api/routes"
	"github.com/marcelofeitoza/track-co-api/internal/service"
	"log"
	"net/http"
)

func main() {
	//db, err := sqlx.Connect("postgres", "postgres://user:password@localhost:5432/db?sslmode=disable")
	db, err := sqlx.Connect("postgres", "postgres://postgres:postgres@postgres.cviycem2u89d.us-east-1.rds.amazonaws.com:5432/postgres")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`
CREATE TABLE IF NOT EXISTS widgets (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    link TEXT NOT NULL,
    question TEXT NOT NULL,
    color VARCHAR(50) NOT NULL DEFAULT '#e5e7eb',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS nps (
    id SERIAL PRIMARY KEY,
    widget_id INT NOT NULL,
    answer TEXT,
    rating INT CHECK (rating >= 0 AND rating <= 10) NOT NULL,
    FOREIGN KEY (widget_id) REFERENCES widgets (id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);`)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database tables created\nConnection established")
	}

	app := service.NewApp(db)

	router := routes.NewRouter(app)

	port := 8080
	addr := fmt.Sprintf(":%d", port)

	log.Printf("Listening on %s", addr)

	err = http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
