package yuque

import "time"

type DocsRes struct {
	Data []*Doc `json:"data"`
}
type Doc struct {
	ID                int         `json:"id"`
	Slug              string      `json:"slug"`
	Title             string      `json:"title"`
	Description       string      `json:"description"`
	UserID            int         `json:"user_id"`
	BookID            int         `json:"book_id"`
	Format            string      `json:"format"`
	Public            int         `json:"public"`
	Status            int         `json:"status"`
	LikesCount        int         `json:"likes_count"`
	CommentsCount     int         `json:"comments_count"`
	ContentUpdatedAt  time.Time   `json:"content_updated_at"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
	PublishedAt       time.Time   `json:"published_at"`
	FirstPublishedAt  time.Time   `json:"first_published_at"`
	DraftVersion      int         `json:"draft_version"`
	LastEditorID      int         `json:"last_editor_id"`
	WordCount         int         `json:"word_count"`
	Cover             string      `json:"cover"`
	CustomDescription string      `json:"custom_description"`
	LastEditor        LastEditor  `json:"last_editor"`
	Book              interface{} `json:"book"`
	Serializer        string      `json:"_serializer"`
}
type LastEditor struct {
	ID              int       `json:"id"`
	Type            string    `json:"type"`
	Login           string    `json:"login"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	AvatarURL       string    `json:"avatar_url"`
	LargeAvatarURL  string    `json:"large_avatar_url"`
	MediumAvatarURL string    `json:"medium_avatar_url"`
	SmallAvatarURL  string    `json:"small_avatar_url"`
	FollowersCount  int       `json:"followers_count"`
	FollowingCount  int       `json:"following_count"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Serializer      string    `json:"_serializer"`
}

type DocRes struct {
	Abilities Abilities `json:"abilities"`
	Data      Data      `json:"data"`
}
type Abilities struct {
	Update  bool `json:"update"`
	Destroy bool `json:"destroy"`
}
type User struct {
	ID               int       `json:"id"`
	Type             string    `json:"type"`
	Login            string    `json:"login"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	AvatarURL        string    `json:"avatar_url"`
	LargeAvatarURL   string    `json:"large_avatar_url"`
	MediumAvatarURL  string    `json:"medium_avatar_url"`
	SmallAvatarURL   string    `json:"small_avatar_url"`
	BooksCount       int       `json:"books_count"`
	PublicBooksCount int       `json:"public_books_count"`
	FollowersCount   int       `json:"followers_count"`
	FollowingCount   int       `json:"following_count"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Serializer       string    `json:"_serializer"`
}
type Book struct {
	ID               int       `json:"id"`
	Type             string    `json:"type"`
	Slug             string    `json:"slug"`
	Name             string    `json:"name"`
	UserID           int       `json:"user_id"`
	Description      string    `json:"description"`
	CreatorID        int       `json:"creator_id"`
	Public           int       `json:"public"`
	ItemsCount       int       `json:"items_count"`
	LikesCount       int       `json:"likes_count"`
	WatchesCount     int       `json:"watches_count"`
	ContentUpdatedAt time.Time `json:"content_updated_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	CreatedAt        time.Time `json:"created_at"`
	Namespace        string    `json:"namespace"`
	User             User      `json:"user"`
	Serializer       string    `json:"_serializer"`
}
type Creator struct {
	ID               int       `json:"id"`
	Type             string    `json:"type"`
	Login            string    `json:"login"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	AvatarURL        string    `json:"avatar_url"`
	LargeAvatarURL   string    `json:"large_avatar_url"`
	MediumAvatarURL  string    `json:"medium_avatar_url"`
	SmallAvatarURL   string    `json:"small_avatar_url"`
	BooksCount       int       `json:"books_count"`
	PublicBooksCount int       `json:"public_books_count"`
	FollowersCount   int       `json:"followers_count"`
	FollowingCount   int       `json:"following_count"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Serializer       string    `json:"_serializer"`
}
type Data struct {
	ID                int         `json:"id"`
	Slug              string      `json:"slug"`
	Title             string      `json:"title"`
	BookID            int         `json:"book_id"`
	Book              Book        `json:"book"`
	UserID            int         `json:"user_id"`
	Creator           Creator     `json:"creator"`
	Format            string      `json:"format"`
	Body              string      `json:"body"`
	BodyDraft         string      `json:"body_draft"`
	BodyHTML          string      `json:"body_html"`
	BodyLake          string      `json:"body_lake"`
	Public            int         `json:"public"`
	Status            interface{} `json:"status"`
	LikesCount        int         `json:"likes_count"`
	CommentsCount     int         `json:"comments_count"`
	ContentUpdatedAt  time.Time   `json:"content_updated_at"`
	DeletedAt         interface{} `json:"deleted_at"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
	PublishedAt       time.Time   `json:"published_at"`
	FirstPublishedAt  time.Time   `json:"first_published_at"`
	WordCount         int         `json:"word_count"`
	Cover             string      `json:"cover"`
	Description       string      `json:"description"`
	CustomDescription string      `json:"custom_description"`
	Serializer        string      `json:"_serializer"`
}
