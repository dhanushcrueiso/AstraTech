package dtos

import "github.com/google/uuid"

type Req struct {
	Id   uuid.UUID `json:"id"`
	Data string    `json:"data"`
}
