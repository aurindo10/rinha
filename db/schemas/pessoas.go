package schemas
type Pessoas struct {
	ID         uint         `gorm:"primaryKey"`
	Apelido    string       `gorm:"type:varchar(32)"`
	Nome       string       `gorm:"type:varchar(100)"`
	Nascimento string       `gorm:"type:varchar(10)"`
	Stack        *[]Stack   `gorm:"foreignKey:PessoaId"`
}