package testPkg

import (
	"intimeServer/model"
	"intimeServer/interfaces"
)

type MockBooks struct {
	interfaces.Book
	FakeCreate func(book *model.Book) (*model.Book, error)
}

func (m *MockBooks) Create(book *model.Book) (*model.Book, error) {
    return m.FakeCreate(book)
}