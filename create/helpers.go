package create

import "perturabo/common"

func NewTimestamps() []*Column {
	return []*Column{
		NewTimestamp("created_at").Nullable().Default(common.Now()),
		NewTimestamp("updated_at").Nullable().Default(common.Now()),
	}
}
