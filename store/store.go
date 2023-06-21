package store

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Word struct {
	gorm.Model
	ID            uint `gorm:"primaryKey"`
	ChatId        int64
	Word          string
	Translate     string
	rememberLevel uint
	RemindAt      time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time      `gorm:"autoUpdateTime:milli"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type Options struct {
	DbPath string
}

type Store struct {
	db *gorm.DB
}

func New(options *Options) *Store {
	db, err := gorm.Open(sqlite.Open(options.DbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Word{})
	return &Store{
		db,
	}
}

func (s *Store) Save(word *Word) {
	s.db.Create(word)
}

func (s *Store) NewWord(word string, translate string, chatId int64) *Word {
	return &Word{
		Word:      word,
		Translate: translate,
		ChatId:    chatId,
	}
}

func (s *Store) LastWord() *Word {
	var word Word
	s.db.Last(&word)
	return &word
}
