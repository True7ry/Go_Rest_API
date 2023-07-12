package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model

	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"not null;unique_index"` // "unique_index" en el caso de que sea null, te lanza error.
	Tasks     []Task `gorm:"foreignkey:UserID"`
	/*
		no es necesario importarlo, ya que estan en el mismo package.
		EL TUTORIAL me fallo:
			no puso la foreningkey, le tuve que preguntar a chatpgt.
	*/
}
