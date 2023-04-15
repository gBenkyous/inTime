package interfaces

import (
	"intimeServer/model"
)

type Book interface {
	Read(id int) (*model.Book, error)
    Create(book *model.Book) (*model.Book, error)
    Update(book *model.Book) error
    Delete(id int) error
}