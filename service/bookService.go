package service

import (
	"ch6/model"
	"errors"

	"github.com/google/uuid"
)

func GetAllBook() []model.Book {
	return model.BookDatas
}

func GetBookById(id string) (model.Book, error) {

	for i, v := range model.BookDatas {
		if v.Id == id {
			return model.BookDatas[i], nil
		} else {
			return model.Book{}, errors.New("ID NOT FOUND")
		}
	}
	return model.Book{}, nil
}
func AddBook(data model.Book) error {

	if data.Title == "" || data.Author == "" || data.Desc == "" {
		return errors.New("DATA CANNOT EMPTY")
	}

	var book = model.Book{
		Id:     uuid.New().String(),
		Title:  data.Title,
		Author: data.Author,
		Desc:   data.Desc,
	}

	model.BookDatas = append(model.BookDatas, book)
	return nil
}
func UpdateBook(id string, data model.Book) error {

	if data.Title == "" || data.Author == "" || data.Desc == "" {
		return errors.New("DATA CANNOT EMPTY")
	}

	for i, v := range model.BookDatas {
		if v.Id == id {
			model.BookDatas[i].Title = data.Title
			model.BookDatas[i].Author = data.Author
			model.BookDatas[i].Desc = data.Desc
			return nil
		} else {
			return errors.New("ID NOT FOUND")
		}

	}
	return nil
}
func DeleteBook(id string) error {
	for i, v := range model.BookDatas {
		if v.Id == id {
			model.BookDatas[i] = model.BookDatas[len(model.BookDatas)-1]
			model.BookDatas = model.BookDatas[:len(model.BookDatas)-1]
			return nil
		} else {
			return errors.New("ID NOT FOUND")
		}

	}
	return nil
}
