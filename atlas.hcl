data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./cmd/migration",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "docker://postgres/12/dev"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}