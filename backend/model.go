package scores

import (
	"time"
)

// Model is an entity with a primary key `ID`.
type Model struct {
	ID        int       `json:"id" db:"id"`
}

// SetID sets the ID on the model
func (m *Model) SetID(id int) {
	m.ID = id
}

// TrackedModel adds timestamps `CreatedAt`, `UpdatedAt`,
// `DeletedAt` to the model.
type TrackedModel struct {
	Model
	CreatedAt time.Time  `json:"createdAt" db:"createdAt"`
	UpdatedAt time.Time  `json:"-" db:"updatedAt"`
	DeletedAt *time.Time `json:"-" db:"deletedAt"`

}