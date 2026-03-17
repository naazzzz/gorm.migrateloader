data "external_schema" "gorm_postgres" {
  program = [
    "env",
    "ENCORERUNTIME_NOPANIC=1",
    "go",
    "run",
    "-mod=mod",
    ".",
  ]
}

env "gorm_postgres" {
  src = data.external_schema.gorm_postgres.url
  dev = "docker://postgres/15/dev?search_path=public"

  migration {
    dir = "file://internal/infrastructure/database/migrations/postgres"
    format = golang-migrate
  }
}

data "external_schema" "gorm_sqlite" {
  program = [
    "env",
    "ENCORERUNTIME_NOPANIC=1",
    "go",
    "run",
    "-mod=mod",
    ".",
  ]
}

env "gorm_sqlite" {
  src = data.external_schema.gorm_sqlite.url
  dev = "sqlite://tests?cache=shared&_fk=1"

  migration {
    dir = "file://internal/infrastructure/database/migrations/sqlite"
    format = golang-migrate
  }
}