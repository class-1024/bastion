package database

import (
	"time"
)

// MComment [...]
type MComment struct {
	Model
	UserID  int    `json:"user_id"`
	Content string `json:"content"`
}

// MMovie [...]
type MMovie struct {
	Model
	Title       string `json:"title"`
	SubTitle    string `json:"sub_title"`
	Summary     string `json:"summary"`
	SmallCover  string `json:"small_cover"`
	BigCover    string `json:"big_cover"`
	MovieURL    string `json:"movie_url"`
	Type        string `json:"type"`
	Country     string `json:"country"`
	Director    string `json:"director"`
	Performer   string `json:"performer"`
	ReleaseDate string `json:"release_date"`
}

// MUser [...]
type MUser struct {
	Model
	Openid    string    `gorm:"unique" json:"openid"`
	NickName  string    `json:"nick_name"`
	AvatarURL string    `json:"avatar_url"`
	Gender    int       `json:"gender"`
	Country   string    `json:"country"`
	Province  string    `json:"province"`
	City      string    `json:"city"`
	Phone     int       `json:"phone"`
	IsActive  int       `json:"is_active"`
	LastLogin time.Time `json:"last_login"`
}

// StatAdmin [...]
type StatAdmin struct {
	Model
	Username string `gorm:"index" json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Phone    int    `json:"phone"`
}

// StatError [...]
type StatError struct {
	Model
	Appid       string `json:"appid"`
	Tag         string `json:"tag"`
	Path        string `json:"path"`
	ErrorMsg    string `json:"error_msg"`
	ErrorString string `json:"error_string"`
	FileURL     string `json:"file_url"`
	Lineno      string `json:"lineno"`
	Colno       string `json:"colno"`
	UId         string `json:"uid"`
	BizUserID   string `json:"biz_user_id"`
}

// MWatchLog [...]
type MWatchLog struct {
	Model
	UserID   int    `json:"user_id"`
	MovieID  int    `json:"movie_id"`
	Progress string `json:"progress"`
}

// StatBehavior [...]
type StatBehavior struct {
	Model
	Appid      string `json:"appid"`
	Path       string `json:"path"`
	Behavior   string `json:"behavior"`
	BehaviorZh string `json:"behavior_zh"`
	Channel    string `json:"channel"`
	Msg        string `json:"msg"`
	UId        string `json:"uid"`
	BizUserID  string `json:"biz_user_id"`
}

// StatDevice [...]
type StatDevice struct {
	Model
	UId           string `json:"uid"`
	BrowserUa     string `json:"browser_ua"`
	BrowserResult string `json:"browser_result"`
}

// StatProject [...]
type StatProject struct {
	Model
	Appid string `json:"appid"`
	Name  string `json:"name"`
	Type  string `json:"type"`
}
