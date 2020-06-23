package database

import "time"

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"delete_at"`
}

type CommentBiz struct {
	MComment
	User *MUser `json:"user"`
}

type UserMovieLogBiz struct {
	MWatchLog
	User  *MUser  `json:"user"`
	Movie *MMovie `json:"movie"`
}

func (m *MMovie) TableName() string {
	return "m_movie"
}

func (m *MUser) TableName() string {
	return "m_user"
}

func (m *MWatchLog) TableName() string {
	return "m_watch_log"
}

func (m *MComment) TableName() string {
	return "m_comment"
}

func (s *StatAdmin) TableName() string {
	return "stat_admin"
}

func (s *StatBehavior) TableName() string {
	return "stat_behavior"
}

func (s *StatError) TableName() string {
	return "stat_error"
}

func (s *StatProject) TableName() string {
	return "stat_project"
}
