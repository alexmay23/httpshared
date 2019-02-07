package permissions

type DefaultUseCase struct{
	repository Repository
}

func (self *DefaultUseCase) CheckUserPermission(userId string, permission string) bool {
	for _,item := range self.GetUserPermissions(userId){
		if item == permission{
			return true
		}
	}
	return false
}

func (self *DefaultUseCase) GetUserPermissions(userId string) []string {
	return self.repository.GetPermissions(userId)
}

func NewDefaultUseCase(repository Repository)*DefaultUseCase{
	return &DefaultUseCase{repository:repository}
}