package scores

import (
	"time"
)

// Model is an entity with a primary key `ID` that gets auto
// assigned by the repository when set.
type Model struct {
	ID        int       `json:"id" db:"id"`
}
	
// SetID sets the ID on the model
func (m *Model) SetID(id int) {
	m.ID = id
}

// Tracked adds timestamps `CreatedAt`, `UpdatedAt`,
// `DeletedAt` to the model.
type Tracked struct {
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"-" db:"updated_at"`
	DeletedAt *time.Time `json:"-" db:"deleted_at"`
	testTime  *time.Time /*`db:"-"`*/
}

// SetCreatedAt sets the `CreatedAt` field.
func (t *Tracked) SetCreatedAt(created time.Time) {
	if t.testTime != nil {
		created = *t.testTime
	}
	t.CreatedAt = created
}

// SetUpdatedAt sets the `UpdatedAt` field.
func (t *Tracked) SetUpdatedAt(updated time.Time) {
	if t.testTime != nil {
		updated = *t.testTime
	}
	t.UpdatedAt = updated
}

// SetDeletedAt sets the `DeletedAt` field.
func (t *Tracked) SetDeletedAt(deleted *time.Time) {
	if t.testTime != nil {
		deleted = t.testTime
	}
	t.DeletedAt = deleted
}

// SetTestTime sets testTime which overrides the arguments
// to other Set* functions, which allows us to test the structs
// by equality which otherwise would not be possible.
func (t *Tracked) SetTestTime(time *time.Time) {
	t.testTime = time
}