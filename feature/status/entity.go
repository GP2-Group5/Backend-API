package status

import "time"

type StatusCore struct {
	ID         uint
	Name       string
	Created_at time.Time
	Updated_at time.Time
}

type IStatusRepository interface {
	GetAll() (data []StatusCore, err error)
}

type IStatusService interface {
	GetAll() (data []StatusCore, err error)
}
