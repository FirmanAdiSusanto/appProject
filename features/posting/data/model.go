type Post struct {
	Id        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string
	Image     string
	Caption   string
	Comments  []data.Comment `gorm:"foreignKey:PostId;references:Id"`
}