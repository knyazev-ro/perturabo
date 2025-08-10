# Perturabo

**Perturabo** — это минималистичный инструмент для управления миграциями баз данных на Go.
Он позволяет создавать, изменять, применять и откатывать миграции в декларативном виде прямо в Go-коде.

## Возможности

* Создание миграций с помощью Go-структур
* Поддержка `CREATE` и `ALTER` миграций
* Автоматическая генерация SQL
* Разделение на `Up` и `Down` действия
* Rollback последних изменений
* Валидация имён и структуры файлов миграций
* CI/CD готовность
* Возможность писать миграции без ORM (использует `Masterminds/squirrel` для SQL-конструкций)

## Установка

```bash
go get github.com/yourusername/perturabo
```

## Структура проекта

```
perturabo/
  commands/      # CLI-команды
  migrate/       # Логика применения миграций
  migrations/    # Ваши миграции (Go-файлы)
  registry/      # Регистрация миграций
  create/        # Описание CREATE миграций
  alter/         # Описание ALTER миграций
  utils/         # Утилиты
```

## Использование

### Создание миграции

Создание миграции для создания таблицы:

```bash
perturabo create:migration user_table users
```

Создание миграции для изменения таблицы:

```bash
perturabo alter:migration user_table users
```

### Применение миграций

```bash
perturabo migrate
```

### Откат миграций

```bash
perturabo migrate:rollback
```

## Пример миграции

```go
var UpCreateUserTable_0001 = registry.Register(
	registry.Action.Up,
	"0001_create_user_table",
	func() any {
		return &create.Table{
			Name: "users",
			Body: []*create.Column{
				create.NewId(),
				create.NewString("name", 255).SetNullable(),
				create.NewBigInteger("value"),
				create.NewDate("start_date"),
				create.NewTimestamp("created_at").SetDefault("now()").SetUnique(),
				create.NewBoolean("is_active").SetNullable(),
			},
		}
	},
)

var DownCreateUserTable_0001 = registry.Register(
	registry.Action.Down,
	"0001_create_user_table",
	func() any {
		return &create.Table{
			Name: "users",
			Drop: true,
		}
	},
)
```

## Валидация миграций

Perturabo проверяет:

* имя файла (должно содержать `create` или `alter`)
* корректное расширение `.go`
* уникальность имён миграций
