# Graphql on GO
[gqlgen getting started](https://gqlgen.com/getting-started/)

```bash
mkdir go-graphql-todoapi
cd go-graphql-todoapi
```

```bash
go mod init github.com/bryanpinheiro/go-graphql-todoapi
```

```bash
go get gorm.io/gorm
go get gorm.io/driver/sqlite
```

### Setup gqlgen

- create a tools.go file
  ```go
  //go:build tools
  // +build tools

  package tools

  import (
    _ "github.com/99designs/gqlgen"
  )
  ```

- add the dependency automatically:
  ```bash
  go mod tidy
  ```

### Create gqlgen project structure 
```bash
go run github.com/99designs/gqlgen init
```

### Change *schema.graphql* and *resolver.go* 

- schema.graphqls:
  ```go
  type Todo {
      id: Int!
      title: String!
      completed: Boolean!
  }

  type Query {
      todo(id: Int!): Todo
  }

  schema {
      query: Query
  }
  ```

- resolver.go:
  ```go
  package graph

  import (
    "context"

    "github.com/bryanpinheiro/go-graphql-todoapi/database"
    "github.com/bryanpinheiro/go-graphql-todoapi/models"
  )

  type Resolver struct{}

  func (r *Resolver) Todos(ctx context.Context) ([]*models.Todo, error) {
    // Access the DB instance from the database package
    db := database.DB

    var todos []*models.Todo
    if err := db.Find(&todos).Error; err != nil {
      return nil, err
    }

    return todos, nil
  }
  ```

- fix some dependencies:
  ```bash
  go get github.com/99designs/gqlgen/graphql/handler
  ```

### Generate auto code using gqlgen
```bash
go run github.com/99designs/gqlgen generate
```

### Run project

```bash
go run server.go
```