package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Bank struct {
	Base    `valid:"required"`
	Code    string     `json:"code" gorm:"type:varchar(20)" valid:"notnull"`
	Name    string     `json:"name" gorm:"type:varchar(255)" valid:"notnull"`
	Account []*Account `gorm:"ForeignKey:BankID" valid:"-"`
}

//método
func (bank *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(bank)
	if err != nil {
		return err
	}
	//retorna erroe em branco
	return nil
}

//função
//*Bank -> ponteiro
func NewBank(code string, name string) (*Bank, error) {

	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().string()
	bank.CreatedAt = time.Now()

	err := bank.isValid()
	if err != nil {
		return nil, err
	}

	// retirna a refer^rncia &bank e não o objeto
	return &bank, nil
}
