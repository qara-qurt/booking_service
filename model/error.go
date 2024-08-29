package model

import (
	"errors"
)

var (
	ErrRoomAlreadyReserved = errors.New("room is already reserved for the selected time slot")
)
