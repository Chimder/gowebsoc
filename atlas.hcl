data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./model",
    "--dialect", "postgres"
  ]
}

env "lol" {
  src = data.external_schema.gorm.url
  dev = "docker://postgres/15"

  migration {
    dir = "file://migrations?format=goose"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}