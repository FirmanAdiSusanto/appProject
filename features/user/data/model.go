package data

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Password string `json:"password" form:"password" validate:"required"`
	Hp       string `gorm:"type:varchar(13);uniqueIndex;primaryKey" json:"hp" form:"hp" validate:"required,max=13,min=10"`
	// Posts    []Post
	//Comments []data.Comment `gorm:"foreignKey:userid;references:Hp"`
}
