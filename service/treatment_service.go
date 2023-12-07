package service

import (
	"annisa-salon/input"
	"annisa-salon/models"
	"annisa-salon/repository"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gosimple/slug"
)

type ServiceTreatments interface {
	CreateTreatment(inputUser input.InputTreatments) (*models.Treatments, error)
	FindAllTreatment()([]*models.Treatments, error)
	UpdateTreatment(slug string, inputUser input.InputTreatments) (*models.Treatments, error)
	DeleteTreatment(slug string) error
	FindTreatmentBySlug(slug string) (*models.Treatments, error)
}

type serviceTreatments struct {
	repositoryTreatments repository.TreatmentsRepository
}

func NewTreatmentService (repositoryTreatments repository.TreatmentsRepository) *serviceTreatments {
	return &serviceTreatments{repositoryTreatments}
}

func (s *serviceTreatments) DeleteTreatment(slug string) error {
	findTreatment, err := s.repositoryTreatments.FindByTreatmentSlug(slug)
	if err != nil {
		return err
	}

	err = s.repositoryTreatments.DeletedTreatment(findTreatment.Slug)
	if err != nil {
		return err
	}
	return nil

}

func (s *serviceTreatments) UpdateTreatment(slugs string, inputUser input.InputTreatments) (*models.Treatments, error) {
    treatment, err := s.repositoryTreatments.FindByTreatmentSlug(slugs)
    if err != nil {
        return nil, err
    }

    // Simpan slug lama ke variabel sementara
    oldSlug := treatment.Slug

    treatment.TreatmentName = inputUser.TreatmentName
    treatment.Description = inputUser.Description
    treatment.Price = inputUser.Price

    // Update nilai-nilai yang ingin diubah di dalam treatment

    // Buat slug baru
    var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
    slugTitle := strings.ToLower(inputUser.TreatmentName)
    mySlug := slug.Make(slugTitle)
    randomNumber := seededRand.Intn(1000000) // Angka acak 0-999999
    treatment.Slug = fmt.Sprintf("%s-%d", mySlug, randomNumber)

    // Ubah nilai slug kembali ke nilai slug lama untuk mencegah perubahan slug dalam database
    treatment.Slug = oldSlug

    updateTreatment, err := s.repositoryTreatments.UpdatedTreatment(treatment)
    if err != nil {
        return nil, err
    }
    return updateTreatment, nil
}

func (s *serviceTreatments) CreateTreatment(inputUser input.InputTreatments) (*models.Treatments, error) {
	treatment := &models.Treatments{}

	treatment.TreatmentName = inputUser.TreatmentName
    treatment.Description = inputUser.Description
    treatment.Price = inputUser.Price

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	
	slugTitle := strings.ToLower(inputUser.TreatmentName)
	
	mySlug := slug.Make(slugTitle)
	
	randomNumber := seededRand.Intn(1000000) // Angka acak 0-999999

	treatment.Slug = fmt.Sprintf("%s-%d", mySlug, randomNumber)

	createTreatment, err := s.repositoryTreatments.CreateTreatment(treatment)
	if err != nil {
		return createTreatment, err
	}
	return createTreatment, nil
	
}

func (s *serviceTreatments) FindTreatmentBySlug(slug string) (*models.Treatments, error) {
	treatment, err := s.repositoryTreatments.FindByTreatmentSlug(slug)
	if err != nil {
		return treatment, err
	}
	if treatment == nil {
		return nil, errors.New("data with that slug not found")
	}
	return treatment, nil
}

func (s *serviceTreatments) FindAllTreatment()([]*models.Treatments, error) {
	blog, err := s.repositoryTreatments.FindAllTreatments()
	if err != nil {
		return blog, err
	}

	return blog, nil

}