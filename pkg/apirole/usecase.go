package apirole

// UseCase is the interface
type UseCase interface {
	GetRoleAll() ([]Roles, error)
	GetRoleById(id string) (Roles, error)
	AddRole(data *Roles) error
	UpdateRole(data Roles) error
	DeleteRole(id string) error
	CheckRoleDisplayExist(display string) Roles

	GetRoleUserAll() ([]RoleUser, error)
	GetRoleUserById(id string) (RoleUser, error)
	AddRoleUser(data *RoleUser) error
	UpdateRoleUser(data RoleUser) error
	DeleteRoleUser(id string) error
	RoleUserExist(data RoleUser) bool

	GetPolicyAll() ([]Policy, error)
	GetPolicyById(id string) (Policy, error)
	GetPolicyListByRoleId(id string) ([]Policy, error)
	AddPolicy(data *Policy) error
	UpdatePolicy(data Policy) error
	DeletePolicy(id string) error
	PolicyExist(data Policy) bool
}

type useCase struct {
	Repo Repository
}

func (u *useCase) PolicyExist(data Policy) bool {
	return u.Repo.PolicyExist(data)
}

func (u *useCase) GetRoleAll() ([]Roles, error) {
	return u.Repo.GetRoleAll()
}

func (u *useCase) GetRoleById(id string) (Roles, error) {
	return u.Repo.GetRoleById(id)
}

func (u *useCase) AddRole(data *Roles) error {
	return u.Repo.AddRole(data)
}

func (u *useCase) UpdateRole(data Roles) error {
	return u.Repo.UpdateRole(data)
}

func (u *useCase) DeleteRole(id string) error {
	return u.Repo.DeleteRole(id)
}

func (u *useCase) CheckRoleDisplayExist(display string) Roles {
	return u.Repo.CheckRoleDisplayExist(display)
}

func (u *useCase) GetRoleUserAll() ([]RoleUser, error) {
	return u.Repo.GetRoleUserAll()
}

func (u *useCase) GetRoleUserById(id string) (RoleUser, error) {
	return u.Repo.GetRoleUserById(id)
}

func (u *useCase) AddRoleUser(data *RoleUser) error {
	return u.Repo.AddRoleUser(data)
}

func (u *useCase) UpdateRoleUser(data RoleUser) error {
	return u.Repo.UpdateRoleUser(data)
}

func (u *useCase) DeleteRoleUser(id string) error {
	return u.Repo.DeleteRoleUser(id)
}

func (u *useCase) RoleUserExist(data RoleUser) bool {
	return u.Repo.RoleUserExist(data)
}

func (u *useCase) GetPolicyAll() ([]Policy, error) {
	return u.Repo.GetPolicyAll()
}

func (u *useCase) GetPolicyById(id string) (Policy, error) {
	return u.Repo.GetPolicyById(id)
}

func (u *useCase) GetPolicyListByRoleId(id string) ([]Policy, error) {
	return u.Repo.GetPolicyListByRoleId(id)
}

func (u *useCase) AddPolicy(data *Policy) error {
	return u.Repo.AddPolicy(data)
}

func (u *useCase) UpdatePolicy(data Policy) error {
	return u.Repo.UpdatePolicy(data)
}

func (u *useCase) DeletePolicy(id string) error {
	return u.Repo.DeletePolicy(id)
}

// NewUseCase new instance
func NewUseCase(repo Repository) UseCase {
	return &useCase{
		Repo: repo,
	}
}
