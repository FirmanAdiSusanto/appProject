package data

import (
	"21-api/features/comment"

	"github.com/golang-jwt/jwt/v5"
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

func (cm *model) InsertComment(PostID string, contentBaru comment.Comment) (comment.Comment, error) {
	var inputProcess = Comment{Content: contentBaru.Content, PostID: PostID}
	if err := cm.connection.Create(&inputProcess).Error; err != nil {
		return comment.Comment{}, err
	}

	return comment.Comment{Content: inputProcess.Content}, nil
}

func (cm *model) GetComment(userID string) ([]comment.Comment, error) {
	var result []comment.Comment
	if err := cm.connection.Where("userid = ?", userID).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

// Delete Komentar
func (cm *model) DeleteComment(commentID uint) error {
	// Membuat objek komentar dengan ID yang diberikan
	var commentToDelete comment.Comment
	commentToDelete.ID = commentID

	// Menghapus komentar dari database
	if err := cm.connection.Delete(&commentToDelete).Error; err != nil {
		return err
	}

	return nil
}

// Tambah Komentar
func (cm *model) AddComment(userID *jwt.Token, contentBaru comment.Comment) (comment.Comment, error) {
	// Membuat objek komentar dengan data yang diberikan
	newComment := comment.Comment{
		Content: contentBaru.Content,
	}

	// Memasukkan komentar baru ke dalam database
	if err := cm.connection.Create(&newComment).Error; err != nil {
		return comment.Comment{}, err
	}

	return newComment, nil
}
