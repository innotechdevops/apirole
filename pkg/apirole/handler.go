package apirole

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Handler is the interface
type Handler interface {
	GetRoleByID(c *fiber.Ctx) error
	GetRoleList(c *fiber.Ctx) error
	CreateRole(c *fiber.Ctx) error
	UpdateRole(c *fiber.Ctx) error
	DeleteRole(c *fiber.Ctx) error

	GetRoleUserByID(c *fiber.Ctx) error
	GetRoleUserList(c *fiber.Ctx) error
	CreateRoleUser(c *fiber.Ctx) error
	UpdateRoleUser(c *fiber.Ctx) error
	DeleteRoleUser(c *fiber.Ctx) error

	GetPolicyByID(c *fiber.Ctx) error
	GetPolicyListByRoleId(c *fiber.Ctx) error
	GetPolicyList(c *fiber.Ctx) error
	CreatePolicy(c *fiber.Ctx) error
	UpdatePolicy(c *fiber.Ctx) error
	DeletePolicy(c *fiber.Ctx) error
}

type handler struct {
	Uc UseCase
}

func (h *handler) GetRoleByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if r, err := h.Uc.GetRoleById(id); err == nil {
		return c.Status(http.StatusOK).JSON(r)
	}
	return fiber.ErrBadRequest
}

func (h *handler) GetRoleList(c *fiber.Ctx) error {
	if r, err := h.Uc.GetRoleAll(); err == nil {
		return c.Status(http.StatusOK).JSON(r)
	}
	return fiber.ErrBadRequest
}

func (h *handler) CreateRole(c *fiber.Ctx) error {
	r := Roles{}
	if err := c.BodyParser(&r); err != nil || r.Display == "" || r.Description == "" {
		return fiber.ErrBadRequest
	}
	r.CreatedAt = time.Now()
	if err := h.Uc.AddRole(&r); err == nil {
		return c.Status(http.StatusCreated).JSON(r)
	}
	return fiber.ErrBadRequest
}

func (h *handler) UpdateRole(c *fiber.Ctx) error {
	r := Roles{}
	id := c.Params("id")
	if err := c.BodyParser(&r); err != nil || id == "" || r.Display == "" || r.Description == "" {
		return fiber.ErrBadRequest
	}
	if role, err := h.Uc.GetRoleById(id); err == nil {
		role.Display = r.Display
		role.Description = r.Description
		role.UpdatedAt = time.Now()
		if err := h.Uc.UpdateRole(role); err == nil {
			return c.Status(http.StatusOK).JSON(role)
		}
	}
	return fiber.ErrNotFound
}

func (h *handler) DeleteRole(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return fiber.ErrBadRequest
	}
	err := h.Uc.DeleteRole(id)
	if err == nil {
		return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Deleted"})
	}
	return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
}

func (h *handler) GetRoleUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if r, err := h.Uc.GetRoleUserById(id); err == nil {
		return c.Status(http.StatusOK).JSON(r)
	}
	return fiber.ErrBadRequest
}

func (h *handler) GetRoleUserList(c *fiber.Ctx) error {
	if r, err := h.Uc.GetRoleUserAll(); err == nil {
		return c.Status(http.StatusOK).JSON(r)
	}
	return fiber.ErrBadRequest
}

func (h *handler) CreateRoleUser(c *fiber.Ctx) error {
	r := RoleUser{}
	if err := c.BodyParser(&r); err != nil || r.RoleID == "" || r.UserID <= 0 {
		return fiber.ErrBadRequest
	}

	if isExist := h.Uc.RoleUserExist(r); isExist {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "User role is exist"})
	}

	if err := h.Uc.AddRoleUser(&r); err == nil {
		return c.Status(http.StatusCreated).JSON(r)
	}

	return fiber.ErrBadRequest
}

func (h *handler) UpdateRoleUser(c *fiber.Ctx) error {
	id := c.Params("id")
	r := RoleUser{}
	if err := c.BodyParser(&r); err != nil || id == "" || r.UserID <= 0 || r.RoleID == "" {
		return fiber.ErrBadRequest
	}

	if role, err := h.Uc.GetRoleUserById(id); err == nil {
		role.UserID = r.UserID
		role.RoleID = r.RoleID
		if err := h.Uc.UpdateRoleUser(role); err == nil {
			return c.Status(http.StatusOK).JSON(role)
		}
	}

	return fiber.ErrNotFound
}

func (h *handler) DeleteRoleUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return fiber.ErrBadRequest
	}
	if err := h.Uc.DeleteRoleUser(id); err == nil {
		return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Deleted"})
	}
	return fiber.ErrBadRequest
}

func (h *handler) GetPolicyByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if r, err := h.Uc.GetPolicyById(id); err == nil {
		return c.Status(http.StatusOK).JSON(r)
	}
	return fiber.ErrBadRequest
}

func (h *handler) GetPolicyListByRoleId(c *fiber.Ctx) error {
	id := c.Params("id")
	if r, err := h.Uc.GetPolicyListByRoleId(id); err == nil {
		return c.Status(http.StatusOK).JSON(r)
	}
	return fiber.ErrBadRequest
}

func (h *handler) GetPolicyList(c *fiber.Ctx) error {
	if r, err := h.Uc.GetPolicyAll(); err == nil {
		return c.Status(http.StatusOK).JSON(r)
	}
	return fiber.ErrBadRequest
}

func (h *handler) CreatePolicy(c *fiber.Ctx) error {
	p := Policy{}
	if err := c.BodyParser(&p); err != nil || p.RoleId == "" || p.Method == "" || p.Path == "" {
		return fiber.ErrBadRequest
	}

	if isExist := h.Uc.PolicyExist(p); isExist {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Policy is exist"})
	}

	if err := h.Uc.AddPolicy(&p); err == nil {
		return c.Status(http.StatusCreated).JSON(p)
	}
	return fiber.ErrBadRequest
}

func (h *handler) UpdatePolicy(c *fiber.Ctx) error {
	id := c.Params("id")
	r := Policy{}
	if err := c.BodyParser(&r); err != nil || id == "" || r.RoleId == "" || r.Path == "" || r.Method == "" {
		return fiber.ErrBadRequest
	}

	if role, err := h.Uc.GetPolicyById(id); err == nil {
		role.RoleId = r.RoleId
		role.Path = r.Path
		role.Method = r.Method
		if err := h.Uc.UpdatePolicy(role); err == nil {
			return c.Status(http.StatusOK).JSON(role)
		}
	}

	return fiber.ErrNotFound
}

func (h *handler) DeletePolicy(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return fiber.ErrBadRequest
	}
	if err := h.Uc.DeletePolicy(id); err == nil {
		return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Deleted"})
	}
	return fiber.ErrBadRequest
}

// NewHandler new instance
func NewHandler(uc UseCase) Handler {
	return &handler{
		Uc: uc,
	}
}
