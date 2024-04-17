package interfaces

type SqlHandler interface {
	Create(obj interface{})
	FindAll(obj interface{})
	GetTodos(obj interface{}, ownerID uint)
	DeleteById(obj interface{}, id string)
	Update(obj interface{})
}
