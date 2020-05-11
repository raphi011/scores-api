package scores

// NameValue is a tuple that contains a name and value.
type NameValue struct {
	Name  string `db:"name"`
	Value string `db:"value"`
}
