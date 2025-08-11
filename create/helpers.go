package create

func NewTimestamps() []*Column {
	return []*Column{
		NewTimestamp("created_at").Nullable().Default("now()"),
		NewTimestamp("updated_at").Nullable().Default("now()"),
	}
}
