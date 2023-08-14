# Simple Golang SQLC

## Generate Code

```bash
sqlc generate
```

## Edit Configuration

Open file `sqlc.yaml`

```yaml
version: "2"
sql:
  - engine: "postgresql"
    queries: "query.sql"
    schema: "schema.sql"
    gen:
      go:
        package: "tutorial"
        out: "tutorial"
```
