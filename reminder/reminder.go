package reminder

import (
	"goreminder/clients"
	"goreminder/db"
	"goreminder/state/events"
	"goreminder/store"
	"log"
	"math/rand"
	"sync"
	"time"
)

type Options struct {
	Interval time.Duration
}

type Reminder struct {
	Options Options
	Store   *store.Store
	Clients *clients.Clients

	channel chan *db.WordModel
}

func (r *Reminder) remind(word *db.WordModel) {
	r.Clients.GetOrAdd(int64(word.ChatID)).State.Handle(events.Remind, word)
}

func (r *Reminder) takeRandWord(chatId int64, timeAt time.Time) *db.WordModel {
	words := r.Store.GetWordsToTime(chatId, timeAt)
	if len(words) == 0 {
		log.Printf("No words to remind for %d", chatId)
		return nil
	}

	rand.Seed(time.Now().Unix())
	return &words[rand.Intn(len(words))]
}

func (r *Reminder) run() {
	for {
		log.Printf("Remind...")
		users := r.Store.GetUsers()

		var wg sync.WaitGroup

		for _, user := range users {
			wg.Add(1)
			log.Printf("Remind to %d", user.ChatID)
			go func(chatId int64) {
				word := r.takeRandWord(chatId, time.Now())
				wg.Done()
				if word != nil {
					r.channel <- word
				}
			}(int64(user.ChatID))
		}

		wg.Wait()

		log.Printf("Reminded")
		time.Sleep(r.Options.Interval)
	}
}

func (r *Reminder) Start() {
	r.channel = make(chan *db.WordModel)
	go r.run()
	go func() {
		for word := range r.channel {
			r.remind(word)
		}
	}()
}

func (r *Reminder) Stop() {
	close(r.channel)
}
