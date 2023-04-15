package repository

import (
	"final-project/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SocialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *SocialMediaRepository {
	return &SocialMediaRepository{
		db: db,
	}
}

func (sr *SocialMediaRepository) Get() ([]model.SocialMedia, error) {

	var socialMedia []model.SocialMedia

	tx := sr.db.Find(&socialMedia)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return socialMedia, nil
}

func (sr *SocialMediaRepository) Save(newSocialMedia model.SocialMedia) (model.SocialMedia, error) {
	tx := sr.db.Create(&newSocialMedia)
	if tx.Error != nil {
		return model.SocialMedia{}, tx.Error
	}

	return newSocialMedia, nil
}

func (sr *SocialMediaRepository) Delete(deletedSocialMedia *model.SocialMedia) error {
	tx := sr.db.Clauses(clause.Returning{}).Delete(&deletedSocialMedia)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (sr *SocialMediaRepository) GetById(getSocialMedia *model.SocialMedia) error {
	socialMedia := &model.SocialMedia{}
	tx := sr.db.First(socialMedia, getSocialMedia)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (sr *SocialMediaRepository) Update(updatedSosmed *model.SocialMedia) (*model.SocialMedia, error) {
	tx := sr.db.Clauses(clause.Returning{
		Columns: []clause.Column{
			{
				Name: "id",
			},
		}}).Where("id = ?", updatedSosmed.Id).Updates(&updatedSosmed)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return updatedSosmed, nil
}
