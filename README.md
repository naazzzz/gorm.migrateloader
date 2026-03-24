# GORM Migrate Loader

GORM Migrate Loader is a tool that helps you generate migrations for gorm-models in Go.
It uses the 'gorm.Model' to migrate based on the compositing structures in your project.
It is assumed that migrations will be located in the following directory:
`internal/infrastructure/database/migrations/{driver_name}`

## Installation

Copy atlas.hcl for your PWD of project.

```sh
go get github.com/naazzzz/gorm-migrateloader
go install github.com/naazzzz/gorm-migrateloader

gorm-migrateloader install
```

## Usage
1. Add gorm.Model in your models. For example:
```go
type ExampleModel struct {
	gorm.Model
	Data string `gorm:"column:data;not null"`
}
```

2. Call the 'generate' command to generate migrations (assuming this will be the root of the project). For example:
```sh
gorm-migrateloader generate .
```

The generated migrations will look like golang-migrate, so they are fully compatible with all golang-migrate commands.
![img.png](docs/img.png)

## Avalibility drivers

- `postgres`
- `sqlite`

## Help

```shell
gorm-migrateloader --help
Usage:
  gorm-migrateloader [command]

Available Commands:
  generate    Generate models
  help        Help about any command
  install     Install dependencies

Flags:
  -h, --help   help for gorm-migrateloader

Use "gorm-migrateloader [command] --help" for more information about a command.
```

## [License](LICENSE)
This project is licensed under the MIT.
