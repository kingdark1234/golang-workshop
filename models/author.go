package models

import "workshop-service/configs"

// Author represent the author model
type Author struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GetAllAuthor(au *[]Author) (err error) {
	if err = configs.DB.Find(au).Error; err != nil {
		return err
	}
	return nil
}

func AddNewAuthor(au *Author) (err error) {
	if err = configs.DB.Create(au).Error; err != nil {
		return err
	}
	return nil
}
