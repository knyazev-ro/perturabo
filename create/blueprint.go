package create

import (
	"strings"
)

type Column struct {
	field    string
	dDefault string
	nullable bool
	unique   bool

	cascadeOnDelete bool
	cascadeOnUpdate bool

	nullOnDelete bool
	nullOnUpdate bool

	restrictOnDelete bool
	restrictOnUpdate bool

	noActionOnDelete bool
	noActionOnUpdate bool

	defaultOnDelete bool
	defaultOnUpdate bool
}

func (c *Column) ToSQL() string {

	sql := []string{}
	sql = append(sql, c.field)

	if c.dDefault != "" {
		sql = append(sql, "DEFAULT "+c.dDefault)
	}

	if c.nullable {
		sql = append(sql, "NULL")
	} else {
		sql = append(sql, "NOT NULL")
	}

	if c.unique {
		sql = append(sql, "UNIQUE")
	}

	// ON DELETE
	switch {
	case c.cascadeOnDelete:
		sql = append(sql, "ON DELETE CASCADE")
	case c.nullOnDelete:
		sql = append(sql, "ON DELETE SET NULL")
	case c.restrictOnDelete:
		sql = append(sql, "ON DELETE RESTRICT")
	case c.noActionOnDelete:
		sql = append(sql, "ON DELETE NO ACTION")
	case c.defaultOnDelete:
		sql = append(sql, "ON DELETE SET DEFAULT")

	}

	// ON UPDATE
	switch {
	case c.cascadeOnUpdate:
		sql = append(sql, "ON UPDATE CASCADE")
	case c.nullOnUpdate:
		sql = append(sql, "ON UPDATE SET NULL")
	case c.restrictOnUpdate:
		sql = append(sql, "ON UPDATE RESTRICT")
	case c.noActionOnUpdate:
		sql = append(sql, "ON UPDATE NO ACTION")
	case c.defaultOnUpdate:
		sql = append(sql, "ON UPDATE SET DEFAULT")
	}

	sqlStr := strings.Join(sql, " ")
	return sqlStr
}
