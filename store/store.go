package store

import (
	"context"
	"goreminder/db"
	"log"
	"time"

	"github.com/steebchen/prisma-client-go/runtime/types"
)

type Options struct {
	DbPath string
}

type Store struct {
	client *db.PrismaClient
}

func New() *Store {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic("failed to connect database")
	}
	return &Store{client}
}

func (s *Store) getOrAddUser(chatId int64) (*db.UserModel, error) {
	ctx := context.Background()
	user, err := s.client.User.FindFirst(
		db.User.ChatID.Equals(types.BigInt(chatId)),
	).Exec(ctx)
	if err != nil {
		_, err := s.client.User.CreateOne(
			db.User.ChatID.Set(types.BigInt(chatId)),
		).Exec(ctx)

		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return user, nil
}

func (s *Store) AddNewWord(word string, translate string, chatId int64, nextRemindAt time.Time) bool {
	ctx := context.Background()

	_, err := s.getOrAddUser(chatId)
	if err != nil {
		return false
	}

	_, err = s.client.Word.CreateOne(
		db.Word.ChatID.Set(types.BigInt(chatId)),
		db.Word.Word.Set(word),
		db.Word.Translate.Set(translate),
		db.Word.NextRemindAt.Set(types.DateTime(nextRemindAt)),
		db.Word.Users.Link(db.User.ChatID.Equals(types.BigInt(chatId))),
	).Exec(ctx)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func (s *Store) HasWord(word string, chatId int64) bool {
	ctx := context.Background()
	_, err := s.client.Word.FindFirst(
		db.Word.ChatID.Equals(types.BigInt(chatId)),
		db.Word.Word.Equals(word),
	).Exec(ctx)
	if err != nil {
		return false
	}
	return true
}

// get words list from db with limit and offset
func (s *Store) GetWords(chatId int64, offset int, limit int) []db.WordModel {
	ctx := context.Background()
	words, err := s.client.Word.FindMany(
		db.Word.ChatID.Equals(types.BigInt(chatId)),
	).Skip(offset).Take(limit).Exec(ctx)
	if err != nil {
		log.Println(err)
		return nil
	}
	return words
}

// remove word from db
func (s *Store) RemoveWord(word string, chatId int64) bool {
	ctx := context.Background()
	_, err := s.client.Word.FindMany(
		db.Word.ChatID.Equals(types.BigInt(chatId)),
		db.Word.Word.Equals(word),
	).Delete().Exec(ctx)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

const max_words_by_time = 100

// get words by time nextRemindAt
func (s *Store) GetWordsToTime(chatId int64, timeAt time.Time) []db.WordModel {
	ctx := context.Background()
	words, err := s.client.Word.FindMany(
		db.Word.NextRemindAt.Lte(types.DateTime(timeAt)),
		db.Word.ChatID.Equals(types.BigInt(chatId)),
	).Take(max_words_by_time).Exec(ctx)
	if err != nil {
		log.Println(err)
		return make([]db.WordModel, 0)
	}
	return words
}

func (s *Store) GetUsers() []db.UserModel {
	ctx := context.Background()
	users, err := s.client.User.FindMany().Exec(ctx)
	if err != nil {
		log.Println(err)
		return make([]db.UserModel, 0)
	}
	return users
}
