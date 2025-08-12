package common

import "fmt"

type DatabaseType struct {
	Field string
}

func StringType(length int) *DatabaseType {
	return &DatabaseType{
		Field: fmt.Sprintf("VARCHAR(%d)", length),
	}
}

func IntType() *DatabaseType {
	return &DatabaseType{
		Field: "INTEGER",
	}
}
