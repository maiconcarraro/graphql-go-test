# graphql-go (schema first, gqlgen)

## run

```bash
go run .\server\server.go
```

## update generated models

```bash
go run github.com/99designs/gqlgen -v
```

## playground

### query

```graphql
query {
  users {
    id
    name
  }
  
  tasks {
    id
    text
    user {
      id
      name
    }
  }
}
```

### mutation

```graphql
mutation {
  createTask(input: {
    text: "testing"
    userId: "1"
  }) {
    id
    text
    user {
      id
      name
    }
  }
}
```

