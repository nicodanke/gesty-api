package sse

type client struct {
	UserID      int64
	AccountID   int64
	sendMessage chan eventMessage
}

func newClient(userId int64, accountId int64) *client {
	return &client{
		UserID:      userId,
		AccountID:   accountId,
		sendMessage: make(chan eventMessage),
	}
}
