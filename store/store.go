package store

import (
	"context"
	"eng-bot/db"
	"log"

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

func (s *Store) AddNewWord(word string, translate string, chatId int64) bool {
	ctx := context.Background()

	_, err := s.client.Word.CreateOne(
		db.Word.ChatID.Set(types.BigInt(chatId)),
		db.Word.Word.Set(word),
		db.Word.Translate.Set(translate),
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
