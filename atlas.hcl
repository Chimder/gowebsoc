env "dev" {
  url     = "postgres://postgres.brewtfaclndisuqiieoy:7DsShp1,l8Gx@aws-0-eu-central-1.pooler.supabase.com:6543/postgres"
  dev-url = "docker://postgres"

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }

  exclude = ["auth", "extensions", "graphql", "graphql_public", "pgbouncer", "pgsodium", "pgsodium_masks", "realtime", "storage", "vault", "atlas_schema_revisions"]
}