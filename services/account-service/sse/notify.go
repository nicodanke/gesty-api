package sse

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/nicodanke/gesty-api/shared/token"
)

type eventMessage struct {
	EventName string
	Data      any
}

func NewEventMessage(eventName string, data any) eventMessage {
	return eventMessage{
		EventName: eventName,
		Data:      data,
	}
}

type Notifier interface {
	BroadcastMessage(msg eventMessage)
	BoadcastMessageToAccount(msg eventMessage, accountId int64, userId *int64)
	SendMessageToUser(msg eventMessage, userId int64)
}

type HandlerEvent struct {
	m                sync.Mutex
	clients          map[int64]*client
	clientsByAccount map[int64][]int64
}

func NewHandlerEvent() *HandlerEvent {
	return &HandlerEvent{
		clients:          make(map[int64]*client),
		clientsByAccount: make(map[int64][]int64),
	}
}

func (h *HandlerEvent) Handler(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	client := newClient(authPayload.UserID, authPayload.AccountID)

	h.registerClient(client)
	LogInfo(fmt.Sprintf("Connected client %d of account %d\n", authPayload.UserID, authPayload.AccountID))
	client.onLine(ctx)
	LogInfo(fmt.Sprintf("Disconnected client %d of account %d\n", authPayload.UserID, authPayload.AccountID))
	h.unregisterClient(authPayload.UserID)
}

func (h *HandlerEvent) registerClient(c *client) {
	h.m.Lock()
	defer h.m.Unlock()
	h.clients[c.UserID] = c
	clientsOfAccount := h.clientsByAccount[c.AccountID]
	clientsOfAccount = append(clientsOfAccount, c.UserID)
	h.clientsByAccount[c.AccountID] = clientsOfAccount
}

func (h *HandlerEvent) unregisterClient(id int64) {
	h.m.Lock()
	defer h.m.Unlock()
	delete(h.clients, id)
}

func (client *client) onLine(ctx *gin.Context) {
	ctx.Stream(func(w io.Writer) bool {
		if msg, ok := <-client.sendMessage; ok {
			data, err := json.Marshal(msg.Data)
			if err != nil {
				LogError(err, "")
			}

			ctx.SSEvent(msg.EventName, string(data))
			return true
		}
		if _, ok := <-ctx.Done(); ok {
			return false
		}
		return false
	})
}

func (h *HandlerEvent) BroadcastMessage(msg eventMessage) {
	h.m.Lock()
	defer h.m.Unlock()
	for _, client := range h.clients {
		client.sendMessage <- msg
	}
}

func (h *HandlerEvent) BoadcastMessageToAccount(msg eventMessage, accountId int64, userId *int64) {
	h.m.Lock()
	defer h.m.Unlock()
	clientsOfAccount := h.clientsByAccount[accountId]
	for _, userIdIt := range clientsOfAccount {
		client := h.clients[userIdIt]
		if client != nil {
			if userId == nil || *userId != userIdIt {
				client.sendMessage <- msg
			}
		}
	}
}

func (h *HandlerEvent) SendMessageToUser(msg eventMessage, userId int64) {
	h.m.Lock()
	defer h.m.Unlock()
	client := h.clients[userId]
	if client != nil {
		client.sendMessage <- msg
	}
}
