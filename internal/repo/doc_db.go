package repo

import (
	"github.com/esuEdu/aws-go-backend-example/internal/models"
	"gorm.io/gorm"
)

type DocRepo interface {
	Save(d *models.Doc) error
	List(d *models.Doc) ([]models.Doc, error)
}

type docRepo struct {
	db *gorm.DB
}

func NewDocRepo(db *gorm.DB) DocRepo {
	return &docRepo{db: db}
}

func (repo *docRepo) Save(d *models.Doc) error {
	return repo.db.Save(d).Error
}

func (repo *docRepo) List(d *models.Doc) ([]models.Doc, error) {

	var docs []models.Doc

	query := repo.db.Model(models.Doc{})

	if err := query.Find(&docs).Error; err != nil {
		return nil, err
	}

	return docs, nil

}
