// Code generated by entc, DO NOT EDIT.

package withfields

const (
	// Label holds the string label denoting the withfields type in the database.
	Label = "with_fields"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldExisting holds the string denoting the existing field in the database.
	FieldExisting = "existing"
	// Table holds the table name of the withfields in the database.
	Table = "with_fields"
)

// Columns holds all SQL columns for withfields fields.
var Columns = []string{
	FieldID,
	FieldExisting,
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