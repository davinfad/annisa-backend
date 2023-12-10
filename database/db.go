package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDb() (*sql.DB, error) {

	if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); exists == false {
		if err := godotenv.Load(); err != nil {
			log.Fatal("error loading ../.env file:", err)
		}
	}

	dbUsername := os.Getenv("MYSQLUSER")
	dbPassword := os.Getenv("MYSQLPASSWORD")
	dbHost := os.Getenv("MYSQLHOST")
	dbPort := os.Getenv("MYSQLPORT")
	dbName := os.Getenv("MYSQLDATABASE")
	
	// Gunakan nilai variabel lingkungan untuk koneksi database
	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

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