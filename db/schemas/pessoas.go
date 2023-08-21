package schemas

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)
type Pessoas struct {
	ID         string       `gorm:"primaryKey" json:"id"`
	Apelido    string       `gorm:"type:varchar(32);index" json:"apelido"`
	Nome       string       `gorm:"type:varchar(100);index" json:"nome"`
	Nascimento time.Time 	`gorm:"type:date" json:"nascimento"`
	Stack      pq.StringArray  `gorm:"type:varchar(32)[]" json:"stack"`
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
