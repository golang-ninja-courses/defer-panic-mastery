package smsrepository

import (
	"errors"
	"sync"
)

var (
	ErrMsgAlreadyExists       = errors.New("message already exists")
	ErrMsgNotFound            = errors.New("message not found")
	ErrInvalidMsgStatusChange = errors.New("invalid message status change")
)

type (
	MessageID     string
	MessageStatus string
)

const (
	// MessageStatusAccepted – сообщение принято внешней системой.
	MessageStatusAccepted = MessageStatus("ACCEPTED")

	// MessageStatusConfirmed - сообщение подтверждено к отправке внешней системой.
	MessageStatusConfirmed = MessageStatus("CONFIRMED")

	// MessageStatusFailed - сообщение не доставлено по какой-то причине.
	MessageStatusFailed = MessageStatus("FAILED")

	// MessageStatusDelivered - сообщение доставлено.
	MessageStatusDelivered = MessageStatus("DELIVERED")
)

// Repo представляет собой конкурентнобезопасное
// in-memory хранилище статусов SMS-сообщений.
type Repo struct {
	db   map[MessageID]MessageStatus
	lock sync.RWMutex
}

func NewRepo() *Repo {
	// Реализуй меня.
	return nil
}

// Save сохраняет сообщение со статусом MessageStatusAccepted.
// Если сообщение с таким идентификатором уже присутствует в базе,
// то возвращается ErrMsgAlreadyExists.
func (r *Repo) Save(id MessageID) error {
	// Реализуй меня.
	return nil
}

// Get возвращает статус сообщения по его идентификатору
// или ошибку ErrMsgNotFound в случае отсутствия id в базе.
func (r *Repo) Get(id MessageID) (MessageStatus, error) {
	// Реализуй меня.
	return "", nil
}

// Update обновляет статус сообщения по его идентификатору.
// Возвращает ошибку ErrMsgNotFound, если сообщения нет в базе.
// Возвращает ошибку ErrInvalidMsgStatusChange, если из текущего статуса нельзя перейти в новый.
func (r *Repo) Update(id MessageID, newStatus MessageStatus) error {
	// Реализуй меня.
	return nil
}
