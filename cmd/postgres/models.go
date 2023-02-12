package postgres

import (
	"time"

	"github.com/lib/pq"
)

var SENTENCE_STATUS = map[string]uint{
	"WinnerByVote":     1,
	"RejectedByAuthor": 2,
	"WinnerByAuthor":   3,
	"Locked":           4,
	"NextSentence":     5,
}

var USER_STORY_STATUS = map[string]uint{
	"Creator":  1,
	"Bookmark": 2,
}

var ENTITY_STATUS = map[string]uint{
	"Active":   1,
	"Inactive": 2,
}

var ENTITY_TYPE = map[string]uint{
	"Sentence": 1,
	"Story":    2,
}

// ID, createdAt, updatedAt are automatically added by gorm
type User struct {
	Id        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
	Name      string
	Email     string `gorm:"uniqueIndex"`
	Status    uint8  `gorm:"default:1"`
}

type Story struct {
	Id                 uint      `gorm:"primaryKey"`
	CreatedAt          time.Time `gorm:"index"`
	UpdatedAt          time.Time
	Heading            string
	LikesCount         uint8
	CollaboratorsCount uint8
	Tags               pq.Int32Array `gorm:"type:integer[]"`
	Preview            string
	AuthorUserName     string
	Image              string
	Status             uint8 `gorm:"default:1"`
}

type Sentence struct {
	Id         uint      `gorm:"primaryKey"`
	CreatedAt  time.Time `gorm:"index"`
	UpdatedAt  time.Time
	StoryId    uint
	Position   uint8
	Rank       uint8
	Status     uint8 `gorm:"default:1"`
	LikesCount uint8
	Text       string
	ImageUrl   string
}

type Tags struct {
	Id        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
	Name      string
}

type UserStory struct {
	Id        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
	StoryId   uint
	UserId    uint
	Status    uint8 `gorm:"default:1"`
}

type EntityLikes struct {
	Id         uint      `gorm:"primaryKey"`
	CreatedAt  time.Time `gorm:"index"`
	UpdatedAt  time.Time
	Type       uint8 `gorm:"default:1"`
	SentenceId uint
	StoryId    uint
	UserId     uint
}
