package create

import (
	"fmt"
	"perturabo/utils"
	"strings"
)

type Column struct {
	Field    string
	Default  string
	Nullable bool
	Unique   bool

	CascadeOnDelete bool
	CascadeOnUpdate bool

	NullOnDelete bool
	NullOnUpdate bool

	RestrictOnDelete bool
	RestrictOnUpdate bool

	NoActionOnDelete bool
	NoActionOnUpdate bool

	DefaultOnDelete bool
	DefaultOnUpdate bool
}

func (c *Column) ToSQL() string {

	sql := []string{}

	sql = append(sql, c.Field)

	if c.Default != "" {
		if utils.IsFloat(c.Default) || utils.IsInt(c.Default) {
			sql = append(sql, "DEFAULT "+c.Default)
		} else {
			sql = append(sql, fmt.Sprintf("DEFAULT '%s'", c.Default))
		}
	}

	if c.Nullable {
		sql = append(sql, "NULL")
	} else {
		sql = append(sql, "NOT NULL")
	}

	if c.Unique {
		sql = append(sql, "UNIQUE")
	}

	// ON DELETE
	switch {
	case c.CascadeOnDelete:
		sql = append(sql, "ON DELETE CASCADE")
	case c.NullOnDelete:
		sql = append(sql, "ON DELETE SET NULL")
	case c.RestrictOnDelete:
		sql = append(sql, "ON DELETE RESTRICT")
	case c.NoActionOnDelete:
		sql = append(sql, "ON DELETE NO ACTION")
	case c.DefaultOnDelete:
		sql = append(sql, "ON DELETE SET DEFAULT")

	}

	// ON UPDATE
	switch {
	case c.CascadeOnUpdate:
		sql = append(sql, "ON UPDATE CASCADE")
	case c.NullOnUpdate:
		sql = append(sql, "ON UPDATE SET NULL")
	case c.RestrictOnUpdate:
		sql = append(sql, "ON UPDATE RESTRICT")
	case c.NoActionOnUpdate:
		sql = append(sql, "ON UPDATE NO ACTION")
	case c.DefaultOnUpdate:
		sql = append(sql, "ON UPDATE SET DEFAULT")
	}

	sqlStr := strings.Join(sql, " ")
	return sqlStr
}
