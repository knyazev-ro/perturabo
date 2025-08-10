package create

import (
	"fmt"
	"strconv"
)

func NewId() *Column {
	return &Column{
		Field:    "id SERIAL PRIMARY KEY",
		Nullable: false,
	}
}

func NewForeignId(columnName string, references string) *Column {
	return &Column{
		Field:    fmt.Sprintf("%s INT REFERENCES %s(id)", columnName, references),
		Nullable: false,
	}
}

func NewString(columnName string, length int) *Column {
	return &Column{
		Field:    fmt.Sprintf("%s VARCHAR(%s)", columnName, strconv.Itoa(length)),
		Nullable: false,
	}
}

func NewText(columnName string) *Column {
	return &Column{
		Field:    fmt.Sprintf("%s TEXT", columnName),
		Nullable: false,
	}
}

func NewJsonb(columnName string) *Column {
	return &Column{
		Field:    fmt.Sprintf("%s JSONB", columnName),
		Nullable: false,
	}
}

func NewSmallInteger(columnName string) *Column {
	return &Column{
		Field:    fmt.Sprintf("%s SMALLINT", columnName),
		Nullable: false,
	}
}

func NewInteger(columnName string) *Column {
	return &Column{
		Field:    fmt.Sprintf("%s INTEGER", columnName),
		Nullable: false,
	}
}

func NewBigInteger(columnName string) *Column {
	return &Column{
		Field:    fmt.Sprintf("%s BIGINT", columnName),
		Nullable: false,
	}
}

func NewDate(columnName string) *Column {
	return &Column{
		Field:    fmt.Sprintf("%s DATE", columnName),
		Nullable: false,
	}
}

func NewTime(columnName string) *Column {
	return &Column{
		Field:    fmt.Sprintf("%s TIME", columnName),
		Nullable: false,
	}
}

func NewTimeWithTimezone(columnName string) *Column {
	return &Column{
		Field:    fmt.Sprintf("%s TIME WITH TIME ZONE", columnName),
		Nullable: false,
	}
}

func NewTimestamp(columnName string) *Column {
	return &Column{
		Field:    fmt.Sprintf("%s TIMESTAMP", columnName),
		Nullable: false,
	}
}

func NewTimestampWithTimezone(columnName string) *Column {
	return &Column{
		Field:    fmt.Sprintf("%s TIMESTAMP WITH TIME ZONE", columnName),
		Nullable: false,
	}
}

func NewBoolean(columnName string) *Column {
	return &Column{
		Field:    fmt.Sprintf("%s BOOLEAN", columnName),
		Nullable: false,
	}
}

func NewNumeric(columnName string, precision, scale int) *Column {
	return &Column{
		Field:    fmt.Sprintf("%s NUMERIC(%d, %d)", columnName, precision, scale),
		Nullable: false,
	}
}

func NewUUID(columnName string) *Column {
	return &Column{
		Field:    fmt.Sprintf("%s UUID", columnName),
		Nullable: false,
	}
}
