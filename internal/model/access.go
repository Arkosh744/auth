package model

import (
	"time"
)

type AccessInfo struct {
	ID              int64  `db:"id"`
	EndpointAddress string `db:"endpoint_address"`
	Role            string `db:"role"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
