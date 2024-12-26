package entity

import "time"

// Post represents a post/position in the system
type Post struct {
	PostID          int       `gorm:"column:post_id;primaryKey;autoIncrement"`
	PostName        string    `gorm:"column:post_name;size:50;not null"`
	PostCode        string    `gorm:"column:post_code;size:50;not null;uniqueIndex"`
	PostSort        int       `gorm:"column:post_sort;default:0"`
	PostStatus      int8      `gorm:"column:post_status;default:1"`
	PostDescription string    `gorm:"column:post_description;size:200"`
	CreatedAt       time.Time `gorm:"column:created_at;not null"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
	DeletedAt       time.Time `gorm:"column:deleted_at;index"`
}

// TableName returns the table name for the Post model
func (Post) TableName() string {
	return "sys_posts"
}
