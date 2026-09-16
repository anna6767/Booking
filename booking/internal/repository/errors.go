package repository

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrRoomBusy = errors.New("room is busy")
)
