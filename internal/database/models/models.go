package models

import (
	"time"

	"github.com/google/uuid"
)

type DataDao struct {
	Id        uuid.UUID
	Data      string
	CreatedAt time.Time
}
