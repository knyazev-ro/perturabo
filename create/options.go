package create

func (column *Column) SetNullable() *Column {
	column.Nullable = true
	return column
}

func (column *Column) SetCascadeOnDelete() *Column {
	column.CascadeOnDelete = true
	return column
}

func (column *Column) SetCascadeOnUpdate() *Column {
	column.CascadeOnUpdate = true
	return column
}

func (column *Column) SetNullOnDelete() *Column {
	column.NullOnDelete = true
	return column
}

func (column *Column) SetNullOnUpdate() *Column {
	column.NullOnUpdate = true
	return column
}

func (column *Column) SetRestrictOnDelete() *Column {
	column.RestrictOnDelete = true
	return column
}

func (column *Column) SetRestrictOnUpdate() *Column {
	column.RestrictOnUpdate = true
	return column
}

func (column *Column) SetNoActionOnDelete() *Column {
	column.NoActionOnDelete = true
	return column
}

func (column *Column) SetNoActionOnUpdate() *Column {
	column.NoActionOnUpdate = true
	return column
}

func (column *Column) SetDefaultOnDelete() *Column {
	column.DefaultOnDelete = true
	return column
}

func (column *Column) SetDefaultOnUpdate() *Column {
	column.DefaultOnUpdate = true
	return column
}

func (column *Column) SetUnique() *Column {
	column.Unique = true
	return column
}

func (column *Column) SetDefault(value string) *Column {
	column.Default = value
	return column
}
