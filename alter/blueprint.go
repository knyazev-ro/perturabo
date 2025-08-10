package alter

import (
	"strings"
)

type Column struct {
	Field     string
	TypeAlter string

	DefaultSet  string
	DefaultDrop string

	NotNullSet  string
	NotNullDrop string

	ColumnRename      string
	ColumnRenameUsing string

	ColumnAdd  string
	ColumnDrop string
}

func Col(name any) *Column {

	if name == nil {
		return &Column{}
	}
	return &Column{
		Field: name.(string),
	}
}

func (c *Column) ToSQL() string {

	sql := []string{}
	AlterColumn := "ALTER COLUMN "
	if c.TypeAlter != "" {
		sql = append(sql, AlterColumn+c.TypeAlter)
	}

	if c.DefaultSet != "" {
		sql = append(sql, AlterColumn+c.DefaultSet)
	}

	if c.DefaultDrop != "" {
		sql = append(sql, AlterColumn+c.DefaultDrop)
	}

	if c.NotNullSet != "" {
		sql = append(sql, AlterColumn+c.NotNullSet)
	}

	if c.NotNullDrop != "" {
		sql = append(sql, AlterColumn+c.NotNullDrop)
	}

	if c.ColumnRename != "" {
		sql = append(sql, "RENAME COLUMN "+c.ColumnRename)
	}

	if c.ColumnRenameUsing != "" {
		sql = append(sql, AlterColumn+c.ColumnRenameUsing)
	}

	if c.ColumnAdd != "" {
		sql = append(sql, c.ColumnAdd)
	}

	if c.ColumnDrop != "" {
		sql = append(sql, c.ColumnDrop)
	}

	sqlStr := strings.Join(sql, "")
	return sqlStr
}
