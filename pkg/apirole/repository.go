package apirole

import (
	"fmt"
	"github.com/casbin/casbin/v2"
)

// Repository is a interface
type Repository interface {
	GetRoleAll() ([]Roles, error)
	GetRoleById(id string) (Roles, error)
	AddRole(data *Roles) error
	UpdateRole(data Roles) error
	DeleteRole(id string) error

	GetRoleUserAll() ([]RoleUser, error)
	GetRoleUserById(id string) (RoleUser, error)
	AddRoleUser(data *RoleUser) error
	UpdateRoleUser(data RoleUser) error
	DeleteRoleUser(id string) error
	RoleUserExist(data RoleUser) bool

	GetPolicyAll() ([]Policy, error)
	GetPolicyById(id string) (Policy, error)
	AddPolicy(data *Policy) error
	UpdatePolicy(data Policy) error
	DeletePolicy(id string) error
	PolicyExist(data Policy) bool
}

type repository struct {
	Enforcer *casbin.Enforcer
	Source   DataSource
}

func (r *repository) PolicyExist(data Policy) bool {
	return r.Source.PolicyExist(data)
}

func (r *repository) GetRoleAll() ([]Roles, error) {
	return r.Source.GetRoleAll()
}

func (r *repository) GetRoleById(id string) (Roles, error) {
	return r.Source.GetRoleById(id)
}

func (r *repository) AddRole(data *Roles) error {
	return r.Source.AddRole(data)
}

func (r *repository) UpdateRole(data Roles) error {
	return r.Source.UpdateRole(data)
}

func (r *repository) DeleteRole(id string) error {
	_, uErr := r.Source.GetRoleUserByRoleId(id)
	_, pErr := r.Source.GetPolicyByRoleId(id)
	if uErr != nil && pErr != nil {
		return r.Source.DeleteRole(id)
	}
	return fmt.Errorf("%s", "Has row reference in policy or role user collection")
}

func (r *repository) GetRoleUserAll() ([]RoleUser, error) {
	return r.Source.GetRoleUserAll()
}

func (r *repository) GetRoleUserById(id string) (RoleUser, error) {
	return r.Source.GetRoleUserById(id)
}

func (r *repository) AddRoleUser(data *RoleUser) error {
	return r.Source.AddRoleUser(data)
}

func (r *repository) UpdateRoleUser(data RoleUser) error {
	return r.Source.UpdateRoleUser(data)
}

func (r *repository) DeleteRoleUser(id string) error {
	return r.Source.DeleteRoleUser(id)
}

func (r *repository) RoleUserExist(data RoleUser) bool {
	return r.Source.RoleUserExist(data)
}

func (r *repository) GetPolicyAll() ([]Policy, error) {
	return r.Source.GetPolicyAll()
}

func (r *repository) GetPolicyById(id string) (Policy, error) {
	return r.Source.GetPolicyById(id)
}

func (r *repository) AddPolicy(data *Policy) error {
	rs, err := r.Enforcer.AddPolicy(data.RoleId, data.Path, data.Method)
	if err == nil || rs {
		_ = r.Enforcer.SavePolicy()
		return r.Enforcer.LoadPolicy()
	}
	return err
}

func (r *repository) UpdatePolicy(data Policy) error {
	err := r.Source.UpdatePolicy(data)
	if err == nil {
		_ = r.Enforcer.SavePolicy()
		return r.Enforcer.LoadPolicy()
	}
	return err
}

func (r *repository) DeletePolicy(id string) error {
	p, err := r.Source.GetPolicyById(id)
	if err == nil {
		_, _ = r.Enforcer.RemovePolicy(p.RoleId, p.Path, p.Method)
		_ = r.Enforcer.SavePolicy()
		return r.Enforcer.LoadPolicy()
	}
	return err
}

// NewRepository new instance
func NewRepository(enforcer *casbin.Enforcer, source DataSource) Repository {
	return &repository{
		Enforcer: enforcer,
		Source:   source,
	}
}
