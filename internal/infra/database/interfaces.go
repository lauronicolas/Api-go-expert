package database

import "github.com/lauronicolas/curso-go/Api/internal/entity"

type UserDBInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ProductDBInterface interface {
	Create(product *entity.Product) error
	FindByID(id string) (*entity.Product, error)
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
