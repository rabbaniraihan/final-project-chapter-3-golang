package repository

import (
	"final-project/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PhotoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *PhotoRepository {
	return &PhotoRepository{
		db: db,
	}
}

func (pr *PhotoRepository) Get() ([]model.Photo, error) {

	var photo []model.Photo

	tx := pr.db.Find(&photo)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return photo, nil
}

func (pr *PhotoRepository) Save(newPhoto model.Photo) (model.Photo, error) {
	tx := pr.db.Create(&newPhoto)
	if tx.Error != nil {
		return model.Photo{}, tx.Error
	}

	return newPhoto, nil
}

func (pr *PhotoRepository) Delete(deletedPhoto *model.Photo) error {
	tx := pr.db.Clauses(clause.Returning{}).Delete(&deletedPhoto)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (pr *PhotoRepository) GetById(getPhoto *model.Photo) error {
	photo := &model.Photo{}
	tx := pr.db.First(photo, getPhoto)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (pr *PhotoRepository) Update(updatedPhoto *model.Photo) (*model.Photo, error) {
	tx := pr.db.Clauses(clause.Returning{
		Columns: []clause.Column{
			{
				Name: "id",
			},
		}}).Where("id = ?", updatedPhoto.Id).Updates(&updatedPhoto)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return updatedPhoto, nil
}
