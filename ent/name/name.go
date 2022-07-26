// Code generated by ent, DO NOT EDIT.

package name

const (
	// Label holds the string label denoting the name type in the database.
	Label = "name"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTconst holds the string denoting the tconst field in the database.
	FieldTconst = "tconst"
	// FieldPrimaryName holds the string denoting the primaryname field in the database.
	FieldPrimaryName = "primary_name"
	// FieldBirthYear holds the string denoting the birthyear field in the database.
	FieldBirthYear = "birth_year"
	// FieldDeathYear holds the string denoting the deathyear field in the database.
	FieldDeathYear = "death_year"
	// FieldPrimaryProfession holds the string denoting the primaryprofession field in the database.
	FieldPrimaryProfession = "primary_profession"
	// FieldKnownForTitles holds the string denoting the knownfortitles field in the database.
	FieldKnownForTitles = "known_for_titles"
	// Table holds the table name of the name in the database.
	Table = "names"
)

// Columns holds all SQL columns for name fields.
var Columns = []string{
	FieldID,
	FieldTconst,
	FieldPrimaryName,
	FieldBirthYear,
	FieldDeathYear,
	FieldPrimaryProfession,
	FieldKnownForTitles,
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

var (
	// BirthYearValidator is a validator for the "birthYear" field. It is called by the builders before save.
	BirthYearValidator func(int) error
	// DeathYearValidator is a validator for the "deathYear" field. It is called by the builders before save.
	DeathYearValidator func(int) error
)
