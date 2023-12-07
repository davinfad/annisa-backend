package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDb() (*sql.DB, error) {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/annisa")
	if err != nil {
		log.Fatal("Error Connection to db", err)
	}

	// Cek koneksi database
	err = db.Ping()
	if err != nil {
		log.Fatal("DB Ping Error:", err)
		return nil, err
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            ID INT AUTO_INCREMENT PRIMARY KEY,
            Email VARCHAR(255) UNIQUE NOT NULL,
            Role INT NOT NULL,
            Password VARCHAR(255) NOT NULL
		    )
    `)
    if err != nil {
        return nil, err
    }

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS blogs (
		ID INT AUTO_INCREMENT PRIMARY KEY,
		Title VARCHAR(255) NOT NULL,
		Description TEXT,
		FileName VARCHAR(255),
		Slug VARCHAR(255),
		CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	)
    `)
    if err != nil {
        return nil, err
    }

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS treatments (
        ID INT AUTO_INCREMENT PRIMARY KEY,
        Slug VARCHAR(255) NOT NULL,
        TreatmentName VARCHAR(255),
        Description TEXT,
        Price INT,
        CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    )
`)
if err != nil {
    return nil, err
}

	return db, nil
}