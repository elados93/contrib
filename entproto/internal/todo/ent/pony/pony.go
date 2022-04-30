// Code generated by entc, DO NOT EDIT.

package pony

const (
	// Label holds the string label denoting the pony type in the database.
	Label = "pony"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// Table holds the table name of the pony in the database.
	Table = "ponies"
)

// Columns holds all SQL columns for pony fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
