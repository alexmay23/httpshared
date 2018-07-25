package user





type Repository interface {
	GetByPhone(phone string) *Model
	GetByFBId(id string) *Model
	GetById(id string) *Model
	CreateWithPhone(phone string, code int)*Model
	CreateWithFB(id string, name string, avatar string, phone *string)*Model
	GetByIdList(idList []string)[]Model
	Update(model *Model)
}
