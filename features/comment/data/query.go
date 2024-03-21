package data

import (
	"21-api/features/comment"
	"errors"
	"time"

	"gorm.io/gorm"
)

type model struct {
	connection *gorm.DB
}

func New(db *gorm.DB) comment.CommentModel {
	return &model{
		connection: db,
	}
}

// Tambah Komentar
func (cm *model) AddComment(userID string, content string) error {
	// Implementasi fungsi AddComment dengan dua parameter
	var inputProcess = Comment{
		UserID:    userID,
		Content:   content,
		CreatedAt: time.Now().UTC(),
	}

	qry := cm.connection.Create(&inputProcess)
	if err := qry.Error; err != nil {
		return err
	}

	if qry.RowsAffected < 1 {
		return errors.New("no data affected")
	}
	return nil
}

// Delete Komentar
func (cm *model) DeleteComment(userID string, postID uint, commentID string) error {
	qry := cm.connection.Where("id = ? AND userid = ?", commentID, userID).Delete(&Comment{})

	if err := qry.Error; err != nil {
		return err
	}

	if qry.RowsAffected < 1 {
		return errors.New("no data affected")
	}

	return nil
}
