package database

import (
	"time"

	"github.com/denismathan/goAuth/src/entities"
)

func (handler *SqlHandler) FindTokenuser(obj interface{}, id string) {
	handler.db.Where("token = ?", id).Find(obj)
}

func (handler *SqlHandler) DeleteExpiredTokens(userID uint) {
	handler.db.Where("userId = ? AND expirationDate < ?", userID, time.Now()).Delete(entities.Token{})
}

func (handler *SqlHandler) FindUser(obj interface{}, email string) {
	handler.db.Where("email = ?", email).Find(obj)
}

func (handler *SqlHandler) FindUserById(obj interface{}, id uint) {
	handler.db.Where("id = ?", id).Find(obj)
}
