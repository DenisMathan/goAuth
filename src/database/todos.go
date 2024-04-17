package database

func (handler *SqlHandler) GetTodos(obj interface{}, ownerID uint) {
	handler.db.Where("ownerID = ?", ownerID).Find(obj)
}
