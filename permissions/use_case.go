package permissions

type UseCase interface {
	CheckUserPermission(userId string, permission string)bool
	GetUserPermissions(userId string)[]string
}