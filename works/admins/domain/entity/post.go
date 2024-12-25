package entity

// Post represents a post/position in the system
type Post struct {
	PostID     int    `gorm:"primaryKey;column:post_id;autoIncrement"`
	PostName   string `gorm:"column:post_name;size:100;not null"`
	PostCode   string `gorm:"column:post_code;size:50;not null"`
	PostSort   int    `gorm:"column:post_sort;default:0"`
	PostStatus int8   `gorm:"column:post_status;default:1"`
	BaseModel
}

// TableName returns the table name for the Post model
func (Post) TableName() string {
	return "sys_posts"
}
