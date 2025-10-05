package db

import (
    "database/sql"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
    dsn := "bloguser:blogpass@tcp(blog-mysql:3306)/blogdb?parseTime=true"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }
    log.Println("Povezano na MySQL")

    // Kreiranje tabele ako ne postoji
    query := `
    CREATE TABLE IF NOT EXISTS blogs (
        id VARCHAR(50) PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        description TEXT,
        author_id VARCHAR(100),
        created_at DATETIME,
        images JSON
    );
    `
    _, err = db.Exec(query)
    if err != nil {
        log.Fatal("Greška pri kreiranju tabele:", err)
    }

    return db
}

