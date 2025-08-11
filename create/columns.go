package create

import (
	"fmt"
	"strconv"
)

func NewId() *Column {
	return &Column{
		field:    "id SERIAL PRIMARY KEY",
		nullable: false,
	}
}

func NewForeignId(columnName string, references string) *Column {
	return &Column{
		field:    fmt.Sprintf("%s INT REFERENCES %s(id)", columnName, references),
		nullable: false,
	}
}

func NewString(columnName string, length int) *Column {
	return &Column{
		field:    fmt.Sprintf("%s VARCHAR(%s)", columnName, strconv.Itoa(length)),
		nullable: false,
	}
}

func NewText(columnName string) *Column {
	return &Column{
		field:    fmt.Sprintf("%s TEXT", columnName),
		nullable: false,
	}
}

func NewJsonb(columnName string) *Column {
	return &Column{
		field:    fmt.Sprintf("%s JSONB", columnName),
		nullable: false,
	}
}

func NewSmallInteger(columnName string) *Column {
	return &Column{
		field:    fmt.Sprintf("%s SMALLINT", columnName),
		nullable: false,
	}
}

func NewInteger(columnName string) *Column {
	return &Column{
		field:    fmt.Sprintf("%s INTEGER", columnName),
		nullable: false,
	}
}

func NewBigInteger(columnName string) *Column {
	return &Column{
		field:    fmt.Sprintf("%s BIGINT", columnName),
		nullable: false,
	}
}

func NewDate(columnName string) *Column {
	return &Column{
		field:    fmt.Sprintf("%s DATE", columnName),
		nullable: false,
	}
}

func NewTime(columnName string) *Column {
	return &Column{
		field:    fmt.Sprintf("%s TIME", columnName),
		nullable: false,
	}
}

func NewTimeWithTimezone(columnName string) *Column {
	return &Column{
		field:    fmt.Sprintf("%s TIME WITH TIME ZONE", columnName),
		nullable: false,
	}
}

func NewTimestamp(columnName string) *Column {
	return &Column{
		field:    fmt.Sprintf("%s TIMESTAMP", columnName),
		nullable: false,
	}
}

func NewTimestampWithTimezone(columnName string) *Column {
	return &Column{
		field:    fmt.Sprintf("%s TIMESTAMP WITH TIME ZONE", columnName),
		nullable: false,
	}
}

func NewBoolean(columnName string) *Column {
	return &Column{
		field:    fmt.Sprintf("%s BOOLEAN", columnName),
		nullable: false,
	}
}

func NewNumeric(columnName string, precision, scale int) *Column {
	return &Column{
		field:    fmt.Sprintf("%s NUMERIC(%d, %d)", columnName, precision, scale),
		nullable: false,
	}
}

func NewUUID(columnName string) *Column {
	return &Column{
		field:    fmt.Sprintf("%s UUID", columnName),
		nullable: false,
	}
}
