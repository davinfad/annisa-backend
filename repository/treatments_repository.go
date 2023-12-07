package repository

import (
	"annisa-salon/models"
	"database/sql"
	"time"
)

type TreatmentsRepository interface {
	CreateTreatment(treatments *models.Treatments) (*models.Treatments, error)
	FindByTreatmentSlug(slug string) (*models.Treatments, error)
	FindAllTreatments() ([]*models.Treatments, error)
    UpdatedTreatment(treatments *models.Treatments) (*models.Treatments, error)
	DeletedTreatment(slug string) error
}

type treatmentsRepository struct {
	db *sql.DB
}

func NewTreatmentsRepository (db *sql.DB) *treatmentsRepository {
	return &treatmentsRepository{db}
}

func (r *treatmentsRepository) CreateTreatment(treatments *models.Treatments) (*models.Treatments, error) {
	    // currentTime := time.Now() // Mendapatkan waktu saat ini
		query := "INSERT INTO treatments (Slug, TreatmentName, Description, Price ) VALUES (?, ?, ?, ?)"
		result, err := r.db.Exec(query, treatments.Slug, treatments.TreatmentName, treatments.Description, treatments.Price)
		if err != nil {
			return treatments, err
		}
	
		treatmentsID, _ := result.LastInsertId()
		treatments.ID = int(treatmentsID)
		// treatments.CreatedAt = currentTime
	
		return treatments, nil
} 

func (r *treatmentsRepository) UpdatedTreatment(treatments *models.Treatments) (*models.Treatments, error) {
	query := "UPDATE treatments SET TreatmentName=?, Description=?, Price=?, UpdatedAt=? WHERE Slug=?"
    currentTime := time.Now().Format("2006-01-02 15:04:05")

    _, err := r.db.Exec(
        query,
        treatments.Slug,
        treatments.TreatmentName,
        treatments.Description,
        treatments.Price,
        currentTime,
    )

    if err != nil {
        return nil, err
    }

    return treatments, nil
}

func (r *treatmentsRepository) FindByTreatmentSlug(slug string) (*models.Treatments, error) {
    query := "SELECT ID, Slug, TreatmentName, Description, Price FROM treatments WHERE Slug = ?"

    row := r.db.QueryRow(query, slug)

    treatments := &models.Treatments{}
    err := row.Scan(&treatments.ID, &treatments.Slug, &treatments.TreatmentName, &treatments.Description, &treatments.Price)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil // treatments tidak ditemukan
        }
        return nil, err // Kesalahan lainnya
    }

    return treatments, nil
}

func (r *treatmentsRepository) DeletedTreatment(slug string) error {
    query := "DELETE FROM treatments WHERE Slug = ?"

    _ , err := r.db.Exec(query, slug)
    if err != nil {
        return err
    }

    return nil
}

func (r *treatmentsRepository) FindAllTreatments() ([]*models.Treatments, error) {
    var treatments []*models.Treatments

    query := "SELECT ID, Slug, TreatmentName, Description, Price FROM treatments"
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        newTreatments := &models.Treatments{}
         err := rows.Scan(&newTreatments.ID, &newTreatments.Slug, &newTreatments.TreatmentName, &newTreatments.Description, &newTreatments.Price)
		 if err != nil {
            return nil, err
        }
        treatments = append(treatments, newTreatments)
    }

     err = rows.Err()
	 if err != nil {
        return nil, err
    }

    return treatments, nil
}