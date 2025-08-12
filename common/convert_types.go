package common

import "fmt"

type ConvertType struct {
	Field string
}

func String(value string) *ConvertType {
	return &ConvertType{
		Field: fmt.Sprintf("'%s'", value),
	}
}

func Int(value int) *ConvertType {
	return &ConvertType{
		Field: fmt.Sprintf("%d", value),
	}
}

func Numeric(value float64) *ConvertType {
	return &ConvertType{
		Field: fmt.Sprintf("%f", value),
	}
}

func Now() *ConvertType {
	return &ConvertType{
		Field: "NOW()",
	}
}
