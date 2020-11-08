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
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "X-Platform,X-Api-Key,Authorization,Access-Control-Allow-Credentials,Access-Control-Allow-Origin,Origin,Content-Type,Accept",
		AllowMethods: "GET,POST,PUT,PATCH,DELETE,HEAD,OPTIONS",
	}))
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
