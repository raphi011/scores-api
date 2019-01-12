package scores

import (
	"time"
)

// Model is an entity with a primary key `ID`.
type Model struct {
	ID        uint       `json:"id"`
}

// SetID sets the ID on the model
func (m *Model) SetID(id int) {
	m.ID = id
}

// TrackedModel adds timestamps `CreatedAt`, `UpdatedAt`,
// `DeletedAt` to the model.
type TrackedModel struct {
	Model
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`

}