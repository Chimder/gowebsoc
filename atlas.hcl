variable "path" {
  type        = string
  description = "A path to the template directory"
}

data "template_dir" "schema" {
  path = var.path
  vars = {
    key = "value"
    // Pass the --env value as a template variable.
    env  = atlas.env
  }
}

env "dev" {
  url = var.url
  src = data.template_dir.schema.url
  exclude = [
    "auth",
    "extensions",
    "graphql",
    "graphql_public",
    "pgbouncer",
    "pgsodium",
    "pgsodium_masks",
    "realtime",
    "storage",
    "vault",
    "atlas_schema_revisions"
  ]
}