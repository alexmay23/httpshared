package permissions

type Repository interface {
	GetPermissions(userId string)[]string
}