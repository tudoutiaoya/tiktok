package domain

type Relation struct {
	ID       int64 `gorm:"column:id;primarykey"`
	UserID   int64 `gorm:"column:user_id;not null"`
	ToUserID int64 `gorm:"column:to_user_id;not null"`
}

func (r Relation) TableName() string {
	return "relation"
}
