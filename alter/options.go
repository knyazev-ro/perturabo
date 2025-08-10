package alter

import (
	"fmt"
	"strings"
)

func (c *Column) Type(value string) *Column {
	c.TypeAlter = fmt.Sprintf("%s TYPE %s", c.Field, value)
	return c
}

func (c *Column) SetDefault(value string) *Column {
	c.DefaultSet = fmt.Sprintf("%s SET DEFAULT %s", c.Field, value)
	return c
}

func (c *Column) DropDefault() *Column {
	c.DefaultDrop = fmt.Sprintf("%s DROP DEFAULT", c.Field)
	return c
}

func (c *Column) SetNotNull() *Column {
	c.NotNullSet = fmt.Sprintf("%s SET NOT NULL", c.Field)
	return c
}

func (c *Column) DropNotNull() *Column {
	c.NotNullDrop = fmt.Sprintf("%s DROP NOT NULL", c.Field)
	return c
}

func (c *Column) Rename(newName string) *Column {
	c.ColumnRename = fmt.Sprintf("%s TO %s", c.Field, newName)
	return c
}

func (c *Column) RenameTypeUsing(newType string) *Column {
	c.ColumnRenameUsing = fmt.Sprintf("%s TYPE %s USING %s::%s", c.Field, strings.ToUpper(newType), c.Field, newType)
	return c
}

func (c *Column) Add(value string) *Column {
	c.ColumnAdd = fmt.Sprintf("ADD COLUMN %s", value)
	return c
}

func (c *Column) Drop() *Column {
	c.ColumnDrop = fmt.Sprintf("DROP COLUMN %s", c.Field)
	return c
}
