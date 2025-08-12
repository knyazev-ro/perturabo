package alter

import (
	"strings"
)

type Column struct {
	field     string
	typeAlter string

	defaultSet  string
	defaultDrop string

	notNullSet  string
	notNullDrop string

	columnRename      string
	columnRenameUsing string

	columnAdd  string
	columnDrop string
	statistics string
}

func Col(name any) *Column {

	if name == nil {
		return &Column{}
	}
	return &Column{
		field: name.(string),
	}
}

func (c *Column) ToSQL() string {

	sql := []string{}
	AlterColumn := "ALTER COLUMN "
	if c.typeAlter != "" {
		sql = append(sql, AlterColumn+c.typeAlter)
	}

	if c.defaultSet != "" {
		sql = append(sql, AlterColumn+c.defaultSet)
	}

	if c.defaultDrop != "" {
		sql = append(sql, AlterColumn+c.defaultDrop)
	}

	if c.notNullSet != "" {
		sql = append(sql, AlterColumn+c.notNullSet)
	}

	if c.notNullDrop != "" {
		sql = append(sql, AlterColumn+c.notNullDrop)
	}

	if c.columnRename != "" {
		sql = append(sql, c.columnRename)
	}

	if c.columnRenameUsing != "" {
		sql = append(sql, AlterColumn+c.columnRenameUsing)
	}

	if c.columnAdd != "" {
		sql = append(sql, c.columnAdd)
	}

	if c.columnDrop != "" {
		sql = append(sql, c.columnDrop)
	}

	if c.statistics != "" {
		sql = append(sql, c.statistics)
	}

	sqlStr := strings.Join(sql, "")
	return sqlStr
}
