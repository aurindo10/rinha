package schemas

type Stack struct {
	ID           string `gorm:"primaryKey"`	
	Apelido	  	 string
	Nome         string
	Nascimento   string
	PessoaId	 string
	Pessoa 		 *Pessoas 
}