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

type ServiceBlog interface {
	CreateBlog(inputUser input.InputBlog, FileName string) (*models.Blog, error)
	FindAllBlog(limit, offset int) ([]*models.Blog, error)
	UpdateBlog(slug string, inputUser input.InputBlog, FileName string) (*models.Blog, error)
	DeleteBlog(slug string) error
	FindBlogBySlug(slug string) (*models.Blog, error)
}

type serviceBlog struct {
	repositoryBlog repository.BlogRepository
}

func NewBlogService (repositoryBlog repository.BlogRepository) *serviceBlog {
	return &serviceBlog{repositoryBlog}
}

func (s *serviceBlog) DeleteBlog(slug string) error {

	findBlog, err := s.repositoryBlog.FindByBlogSlug(slug)
	if err != nil {
		return errors.New("blog with that slug not found")
	}

	err = s.repositoryBlog.DeleteBlog(findBlog.Slug)
	if err != nil {
		return errors.New("fail to delete blog")
	}
	return nil

}

func (s *serviceBlog) UpdateBlog(slugs string, inputUser input.InputBlog, FileName string) (*models.Blog, error) {
    blog, err := s.repositoryBlog.FindByBlogSlug(slugs)
    if err != nil {
        return nil, err
    }

    // Simpan slug lama ke variabel sementara
    oldSlug := blog.Slug

    blog.Title = inputUser.Title
    blog.Description = inputUser.Description
    blog.FileName = FileName

    // Update nilai-nilai yang ingin diubah di dalam blog

    // Buat slug baru
    var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
    slugTitle := strings.ToLower(inputUser.Title)
    mySlug := slug.Make(slugTitle)
    randomNumber := seededRand.Intn(1000000) // Angka acak 0-999999
    blog.Slug = fmt.Sprintf("%s-%d", mySlug, randomNumber)

    // Ubah nilai slug kembali ke nilai slug lama untuk mencegah perubahan slug dalam database
    blog.Slug = oldSlug

    updateBlog, err := s.repositoryBlog.UpdateBlog(blog)
    if err != nil {
        return nil, err
    }
    return updateBlog, nil
}

func (s *serviceBlog) CreateBlog (inputUser input.InputBlog, FileName string) (*models.Blog, error) {
	blog := &models.Blog{}

	blog.Title = inputUser.Title 
	blog.Description = inputUser.Description
	blog.FileName = FileName

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	
	slugTitle := strings.ToLower(inputUser.Title)
	
	mySlug := slug.Make(slugTitle)
	
	randomNumber := seededRand.Intn(1000000) // Angka acak 0-999999

	blog.Slug = fmt.Sprintf("%s-%d", mySlug, randomNumber)

	createBlog, err := s.repositoryBlog.CreateBlog(blog)
	if err != nil {
		return createBlog, err
	}
	return createBlog, nil
	
}

func (s *serviceBlog) FindBlogBySlug(slug string) (*models.Blog, error) {
	blog, err := s.repositoryBlog.FindByBlogSlug(slug)
	if err != nil {
		return blog, err
	}
	if blog == nil {
		return nil, errors.New("blog with that slug not found")
	}
	return blog, nil
}

func (s *serviceBlog) FindAllBlog(page, limit int) ([]*models.Blog, error) {
	blog, err := s.repositoryBlog.FindAllBlog(page, limit)
	if err != nil {
		return blog, err
	}

    return blog, nil
}