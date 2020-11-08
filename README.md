# apirole

API for Role Management

## Install

```shell script
$ go get github.com/innotechdevops/apirole
```

### Requirement

- Create Collection

```
- casbin_rule
- role_user
- roles
```

#### Roles

- Insert document role `Admin` in collection `roles`

```json
{
    "_id": {
        "$oid": "5f95a41f2e94e13067a087e0"
    },
    "display": "Admin",
    "description": "All API managements",
    "createdat": {
        "$date": "2020-10-02T01:11:18.965Z"
    },
    "updatedat": {
        "$date": "2020-10-02T01:11:18.965Z"
    }    
}
```

- Insert document role `Anonymous` in collection `roles`

```json
{
    "_id": {
        "$oid": "5f95a61d2e94e13067a087e1"
    },
    "display": "Anonymous",
    "description": "Access API without authorization",
    "createdat": {
        "$date": "2020-10-02T01:11:18.965Z"
    },
    "updatedat": {
        "$date": "2020-10-02T01:11:18.965Z"
    }    
}
```

#### User Role

- Insert document user role `Admin` in collection `role_user`

```json
{
  "_id": {
    "$oid": "5fa810e34b40271bd002c8b8"
  },
  "roleId": "5f95a41f2e94e13067a087e0",
  "userId": {
    "$numberLong": "1"
  }
}
```

### Policy

- Insert document policy `Admin` in collection `casbin_role`

```json
{
  "_id": {
    "$oid": "5f95aa6e88b79b68f8b18396"
  },
  "ptype": "p",
  "v0": "5fa810e34b40271bd002c8b8",
  "v1": "/*",
  "v2": "(GET)|(POST)|(PUT)|(DELETE)",
  "v3": "",
  "v4": "",
  "v5": ""
}
```

- Insert document policy `Annonymous` in collection `casbin_role`

```json
{
  "_id": {
    "$oid": "5f95aa6e88b79b68f8b1839b"
  },
  "ptype": "p",
  "v0": "anonymous",
  "v1": "/v1/login",
  "v2": "(POST)",
  "v3": "",
  "v4": "",
  "v5": ""
}
```

## How to use

```go
package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
	mongodbadapter "github.com/casbin/mongodb-adapter/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/innotechdevops/apirole/pkg/apirole"
	"github.com/innotechdevops/mgo-driver/pkg/mgodriver"
	fibercasbinrest "github.com/prongbang/fiber-casbinrest"
)

func main() {
	cfg := mgodriver.Config{
		User:         "root",
		Pass:         "admin",
		Host:         "127.0.0.1",
		DatabaseName: "roledb",
		Port:         mgodriver.DefaultPort,
	}
	connUrl := fmt.Sprintf("mongodb://%s:%s@%s:27017/%s?authSource=admin&ssl=false", cfg.User, cfg.Pass, cfg.Host, cfg.DatabaseName)
	a, _ := mongodbadapter.NewAdapter(connUrl)
	e, _ := casbin.NewEnforcer("model.conf", a)
	driver := mgodriver.New(cfg)

	_ = e.LoadPolicy()

	app := fiber.New()
	app.Use(cors.New())
	app.Use(fibercasbinrest.NewDefault(e, "secret"))

	// Router
	source := apirole.NewDataSource(driver.Connect())
	repo := apirole.NewRepository(e, source)
	uc := apirole.NewUseCase(repo)
	handle := apirole.NewHandler(uc)
	route := apirole.NewRouter(handle)
	route.Initial(app)

	log.Fatal(app.Listen(":3500"))
}
```