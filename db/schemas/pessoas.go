package schemas

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)
type Pessoas struct {
	ID         string       `gorm:"primaryKey"`
	Apelido    string       `gorm:"type:varchar(32)"`
	Nome       string       `gorm:"type:varchar(100)"`
	Nascimento string       `gorm:"type:varchar(10)"`
	Stack        []Stack    `gorm:"foreignKey:PessoaId"`
}

func (pessoa *Pessoas) BeforeCreate(tx *gorm.DB) (err error) {
	if err := tx.Where("Apelido = ?", pessoa.Apelido).First(&pessoa).Error; err == nil {
		return fmt.Errorf("422")
	} else if err != gorm.ErrRecordNotFound {
		return err
	}
	pessoa.ID = uuid.New().String()
	if pessoa.Apelido == "" {
		return fmt.Errorf("apelido não pode ser vazio")
	}
	if pessoa.Apelido == "" {
		return fmt.Errorf("422")
	}
	if pessoa.Nome == "" {
		return fmt.Errorf("422")
	}
	return nil
}