package helper

import (
	"database/sql"
	"time"
)

// NullString is a function to convert string to sql null string
func NullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

// NullInt64 is a function to convert int to sql null int64
func NullInt64(i int) sql.NullInt64 {
	return sql.NullInt64{Int64: int64(i), Valid: i != 0}
}

// NullTime is a function to convert time to sql null time
func NullTime(t time.Time) sql.NullTime {
	return sql.NullTime{Time: t, Valid: !t.IsZero()}
}

// IsSQLErrNotFound is a function to check if error is sql not found
func IsSQLErrNotFound(err error) bool {
	return err.Error() == "sql: no rows in result set"
}
