package scores

import (
	"time"
)

// Model adds an auto assignable ID to models.
type Model interface {
	SetID(id int)
}

// M is an entity with a primary key `ID` that gets auto
// assigned by the repository when set.
type M struct {
	ID int `json:"id" db:"id"`
}

// SetID sets the ID on the model
func (m *M) SetID(id int) {
	m.ID = id
}

// Tracked adds Created / Updated / Deleted metadata to models.
type Tracked interface {
	Create(when time.Time) Tracked
	Update(when time.Time) Tracked
	Delete(when time.Time) Tracked
	MockUpdates(when *time.Time)
}

// Track adds timestamps `CreatedAt`, `UpdatedAt`,
// `DeletedAt` to the model.
type Track struct {
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"-" db:"deleted_at"`
	mockTime  *time.Time
}

// Create sets the `CreatedAt` and `UpdatedAt` fields.
func (t *Track) Create(when time.Time) Tracked {
	if t.mockTime != nil {
		when = *t.mockTime
	}

	t.CreatedAt = when
	t.UpdatedAt = when

	return t
}

// Update sets the `UpdatedAt` field.
func (t *Track) Update(when time.Time) Tracked {
	if t.mockTime != nil {
		when = *t.mockTime
	}
	t.UpdatedAt = when

	return t
}

// Delete sets the `DeletedAt` field.
func (t *Track) Delete(when time.Time) Tracked {
	if t.mockTime != nil {
		when = *t.mockTime
	}

	t.UpdatedAt = when
	t.DeletedAt = &when

	return t
}

/* --- METHODS FOR TESTING ONLY--- */

// MockUpdates sets mockTime which overrides the arguments
// to other Set* functions and allows us to test the structs
// by equality which otherwise would not be possible (because)
// CreatesAt, UpdatedAt and DeletedAt are set in Repository functions.
func (t *Track) MockUpdates(time *time.Time) {
	t.mockTime = time
}
