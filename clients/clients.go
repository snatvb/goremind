package clients

import (
	"eng-bot/state"
	"eng-bot/store"

	"github.com/mymmrac/telego"
)

type Client struct {
	Id    int64
	State state.FSM
}

type Clients struct {
	store   *store.Store
	bot     *telego.Bot
	clients map[int64]*Client
}

func New(store *store.Store, bot *telego.Bot) *Clients {
	return &Clients{
		clients: make(map[int64]*Client),
		store:   store,
		bot:     bot,
	}
}

func (clients *Clients) Add(id int64, fsm state.FSM, bot *telego.Bot) {
	clients.clients[id] = &Client{
		Id:    id,
		State: fsm,
	}
}

func (clients *Clients) Get(id int64) *Client {
	return clients.clients[id]
}

func (clients *Clients) Remove(id int64) {
	delete(clients.clients, id)
}

func (clients *Clients) GetOrAdd(id int64) *Client {
	client := clients.Get(id)
	if client == nil {
		client = &Client{
			Id: id,
			State: *state.NewFSM(&state.Context{
				Store:  clients.store,
				Bot:    clients.bot,
				ChatId: id,
			}),
		}
		clients.clients[id] = client
	}
	return client
}
