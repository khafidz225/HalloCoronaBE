package repositories

import (
	"server/models"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	Finduser() ([]models.Article, error)
	GetArticle(ID int) (models.Article, error)
	AddArticle(article models.Article) (models.Article, error)
	UpdateArticle(article models.Article) (models.Article, error)
	DeleteArticle(article models.Article) (models.Article, error)
}

func RepositoryArticle(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Finduser() ([]models.Article, error) {
	var articles []models.Article
	err := r.db.Preload("User").Find(&articles).Error
	return articles, err
}

func (r *repository) GetArticle(ID int) (models.Article, error) {
	var article models.Article
	err := r.db.Preload("User").First(&article, ID).Error

	return article, err
}

func (r *repository) AddArticle(article models.Article) (models.Article, error) {
	err := r.db.Preload("User").Create(&article).Error

	return article, err
}

func (r *repository) UpdateArticle(article models.Article) (models.Article, error) {
	err := r.db.Model(&article).Updates(article).Error

	return article, err
}

func (r *repository) DeleteArticle(article models.Article) (models.Article, error) {
	err := r.db.Delete(&article).Error

	return article, err
}
