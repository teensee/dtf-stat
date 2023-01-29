package model

import (
	"github.com/google/uuid"
	"recomends/internal/domain/user/dto"
	"strconv"
	"time"
)

type DtfUser struct {
	ID                 uuid.UUID `gorm:"primaryKey:idx_name"`
	DtfUid             int
	CreatedAtStamp     int
	CreatedAt          time.Time
	Name               string
	Description        string
	UserType           int
	IsVerified         bool
	IsMuted            bool
	CounterKarma       int
	CounterEntries     int
	CounterComments    int
	CounterFav         int
	CounterSubscribers int
	IsPlus             bool
}

func NewDtfUser(
	dtfUid int,
	createdAtStamp int,
	createdAt time.Time,
	name string,
	description string,
	userType int,
	isVerified bool,
	isMuted bool,
	counterKarma int,
	counterEntries int,
	counterComments int,
	counterSubscribers int,
	counterFav int,
	isPlus bool,
) *DtfUser {
	return &DtfUser{
		ID:                 uuid.New(),
		DtfUid:             dtfUid,
		CreatedAtStamp:     createdAtStamp,
		CreatedAt:          createdAt,
		Name:               name,
		Description:        description,
		UserType:           userType,
		IsVerified:         isVerified,
		IsMuted:            isMuted,
		CounterKarma:       counterKarma,
		CounterEntries:     counterEntries,
		CounterComments:    counterComments,
		CounterSubscribers: counterSubscribers,
		CounterFav:         counterFav,
		IsPlus:             isPlus,
	}
}

func NewDtfUserFromDto(user dto.DtfUser) *DtfUser {
	i, _ := strconv.ParseInt("1405544146", 10, 64)

	return NewDtfUser(
		user.Result.Id,
		user.Result.Created,
		time.Unix(i, 0),
		user.Result.Name,
		user.Result.Description,
		user.Result.Type,
		user.Result.IsVerified,
		user.Result.IsMuted,
		user.Result.Karma,
		user.Result.Counters.Entries,
		user.Result.Counters.Comments,
		user.Result.SubscribersCount,
		user.Result.Counters.Favourites,
		user.Result.IsPlus,
	)
}
