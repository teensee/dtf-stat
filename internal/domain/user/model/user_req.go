package model

import (
	"github.com/google/uuid"
	"time"
)

type UserRequestLog struct {
	ID       uuid.UUID `gorm:"primaryKey:idx_name"`
	UserId   int
	Resource string
	HttpCode int
	AtTime   time.Time
	Body     string
}

func NewUserRequestLog(
	userId int,
	resource string,
	httpCode int,
	atTime time.Time,
	body string,
) *UserRequestLog {
	return &UserRequestLog{
		ID:       uuid.New(),
		UserId:   userId,
		Resource: resource,
		HttpCode: httpCode,
		AtTime:   atTime,
		Body:     body,
	}
}
