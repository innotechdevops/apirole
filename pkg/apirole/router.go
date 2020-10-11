package apirole

import "github.com/gofiber/fiber/v2"

// Router is interface
type Router interface {
	Initial(e *fiber.App)
}

type router struct {
	Handle Handler
}

func (r *router) Initial(e *fiber.App) {
	v1 := e.Group("/api/v1")
	{
		v1.Get("/roles", r.Handle.GetRoleList)
		v1.Get("/roles/:id", r.Handle.GetRoleByID)
		v1.Post("/roles", r.Handle.CreateRole)
		v1.Put("/roles/:id", r.Handle.UpdateRole)
		v1.Delete("/roles/:id", r.Handle.DeleteRole)

		v1.Get("/roleuser", r.Handle.GetRoleUserList)
		v1.Get("/roleuser/:id", r.Handle.GetRoleUserByID)
		v1.Post("/roleuser", r.Handle.CreateRoleUser)
		v1.Put("/roleuser/:id", r.Handle.UpdateRoleUser)
		v1.Delete("/roleuser/:id", r.Handle.DeleteRoleUser)

		v1.Get("/rolepolicy", r.Handle.GetPolicyList)
		v1.Get("/rolepolicy/:id", r.Handle.GetPolicyByID)
		v1.Post("/rolepolicy", r.Handle.CreatePolicy)
		v1.Put("/rolepolicy/:id", r.Handle.UpdatePolicy)
		v1.Delete("/rolepolicy/:id", r.Handle.DeletePolicy)
	}
}

// NewRouter new instance
func NewRouter(handle Handler) Router {
	return &router{Handle: handle}
}
