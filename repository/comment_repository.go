package repository

import (
	"final-project/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (cr *CommentRepository) Get() ([]model.Comment, error) {

	var comments []model.Comment

	tx := cr.db.Find(&comments)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return comments, nil
}

func (cr *CommentRepository) Save(newComment model.Comment) (model.Comment, error) {
	tx := cr.db.Create(&newComment)
	if tx.Error != nil {
		return model.Comment{}, tx.Error
	}

	return newComment, nil
}

func (cr *CommentRepository) Delete(deletedComment *model.Comment) error {
	tx := cr.db.Clauses(clause.Returning{}).Delete(&deletedComment)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (cr *CommentRepository) GetById(getComment *model.Comment) error {
	comment := &model.Comment{}
	tx := cr.db.First(comment, getComment)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (cr *CommentRepository) Update(updatedComment *model.Comment) (*model.Comment, error) {
	tx := cr.db.Clauses(clause.Returning{
		Columns: []clause.Column{
			{
				Name: "id",
			},
		}}).Where("id = ?", updatedComment.Id).Updates(&updatedComment)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return updatedComment, nil
}
