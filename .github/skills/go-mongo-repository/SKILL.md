---
name: go-mongo-repository
description: 'Generate Go MongoDB repository boilerplate for this project. Use when adding a new domain entity and its repository, scaffolding CRUD for a new collection, or writing a new repository method. Covers entity struct, repository struct, constructor, and custom query methods.'
argument-hint: 'Domain name to scaffold, e.g. "flights" or "alerts"'
---

# Go MongoDB Repository Boilerplate

## When to Use
- Adding a new domain package with a MongoDB-backed repository
- Adding a custom query method to an existing repository
- Unsure of the project's repository / entity conventions

## Module Path
All imports use `aerowatch.com/api/repository`.

---

## Step 1 — Create the Entity (`<domain>/<domain>_entity.go`)

Embed `repository.DBEntity` to inherit `_id`, `createdAt`, `updatedAt` and the `Entity` interface:

```go
package <domain>

import "aerowatch.com/api/repository"

type <Domain>Entity struct {
    repository.DBEntity
    Field1 string `json:"field1" bson:"field1"`
    // add domain-specific fields
}
```

**Rules**
- All fields need both `json` and `bson` struct tags.
- Do **not** redeclare `ID`, `CreatedAt`, `UpdatedAt` — they come from `DBEntity`.

---

## Step 2 — Create the Repository (`<domain>/<domain>_repository.go`)

Wrap the generic `MongoRepository` and add domain-specific methods:

```go
package <domain>

import (
    "context"

    "aerowatch.com/api/repository"
    "go.mongodb.org/mongo-driver/v2/bson"
    "go.mongodb.org/mongo-driver/v2/mongo"
)

const collectionName = "<domain_collection>"

type <Domain>Repository struct {
    *repository.MongoRepository[<Domain>Entity, *<Domain>Entity]
}

func New<Domain>Repository(db *mongo.Database) (*<Domain>Repository, error) {
    if db == nil {
        return nil, repository.ErrDbRequired
    }
    mongoRepo, err := repository.NewMongoRepository[<Domain>Entity, *<Domain>Entity](db, collectionName)
    if err != nil {
        return nil, err
    }
    return &<Domain>Repository{MongoRepository: mongoRepo}, nil
}
```

**Built-in methods** (from `MongoRepository`):

| Method | Signature |
|--------|-----------|
| Save | `Save(ctx, *T) (*T, error)` — insert if ID is zero, replace otherwise |
| Find | `Find(ctx, bson.ObjectID) (*T, error)` |
| FindOne | `FindOne(ctx, bson.M, ...opts) (*T, error)` |
| FindMany | `FindMany(ctx, bson.M, ...opts) ([]*T, error)` |
| Patch | `Patch(ctx, bson.ObjectID, map[string]any) (*T, error)` — partial update via `$set` |
| Delete | `Delete(ctx, bson.ObjectID) error` |
| ToID | `ToID(hexStr string) bson.ObjectID` — hex string → ObjectID |

---

## Step 3 — Add Custom Query Methods

Use `FindOne`, `FindMany`, `Patch`, and `ToID` from the embedded repository:

```go
// Single field lookup
func (r *<Domain>Repository) FindByXxx(ctx context.Context, value string) (*<Domain>Entity, error) {
    return r.FindOne(ctx, bson.M{"fieldName": value})
}

// Partial update
func (r *<Domain>Repository) UpdateXxx(ctx context.Context, id string, value any) (*<Domain>Entity, error) {
    return r.Patch(ctx, r.ToID(id), map[string]any{"fieldName": value})
}
```

---

## Step 4 — Wire Up in `app/main.go`

```go
db := mongoClient.Database()

<domain>Repo, err := <domain>.New<Domain>Repository(db)
if err != nil {
    log.Fatal(err)
}
```

---

## Checklist

- [ ] Entity embeds `repository.DBEntity`
- [ ] All struct fields have `json` + `bson` tags
- [ ] `collectionName` constant matches the MongoDB collection name
- [ ] Constructor guards `db == nil` and returns `repository.ErrDbRequired`
- [ ] Generic type parameters are `[<Domain>Entity, *<Domain>Entity]`
- [ ] Custom methods delegate to `r.FindOne`, `r.FindMany`, or `r.Patch`
- [ ] Repository initialized in `app/main.go` with the shared `*mongo.Database`

---

## Reference Files

- [repository/dbentity.go](../../repository/dbentity.go) — `Entity` interface and `DBEntity`
- [repository/repository.go](../../repository/repository.go) — generic `MongoRepository` and `Repository` interface
- [repository/mongoclient.go](../../repository/mongoclient.go) — client singleton setup
- [users/user_entity.go](../../users/user_entity.go) — canonical entity example
- [users/user_repository.go](../../users/user_repository.go) — canonical repository example
