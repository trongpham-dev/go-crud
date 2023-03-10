package common

import "time"

type SQLModel struct {
	Id        int        `json:"-" db:"id"`
	FakeId    *UID       `json:"id"`
	Status    int        `json:"status" db:"status"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

func (m *SQLModel) GenUID(dbType int) {
	uid := NewUID(uint32(m.Id), dbType, 1)
	m.FakeId = &uid
}
