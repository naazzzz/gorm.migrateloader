# GORM Migrate Loader

GORM Migrate Loader is a tool that helps you generate migrations for gorm-models in Go.
It uses the 'gorm.Model' to migrate based on the compositing structures in your project.
It is assumed that migrations will be located in the following directory:
`internal/infrastructure/database/migrations/{driver_name}`

## Installation

```sh
go get github.com/naazzzz/gorm-migrateloader
go install github.com/naazzzz/gorm-migrateloader
```

## Usage
It is assumed that it will be used on the root directory of the project.

```sh
gorm-migrateloader install
gorm-migrateloader generate .
```

The generated migrations will look like golang-migrate, so they are fully compatible with all golang-migrate commands.
![img.png](docs/img.png)

## Avalibility drivers

- `postgres`
- `sqlite3`

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