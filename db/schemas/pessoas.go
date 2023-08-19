package schemas

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)
type Pessoas struct {
	ID         string       `gorm:"primaryKey"`
	Apelido    string       `gorm:"type:varchar(32)"`
	Nome       string       `gorm:"type:varchar(100)"`
	Nascimento time.Time 	`gorm:"type:date"`
	Stack      pq.StringArray  `gorm:"type:varchar(32)[]"`
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
