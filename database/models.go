package database

import "gorm.io/gorm"

type CategoriaEvento struct {
	gorm.Model
	Nombre                string       `json:"nombre"`
	CategoriaEnfocadaPara string       `json:"categoria_enfocada_para"`
	Publicada             bool         `json:"publicada"`
	Planes                []PlanEvento `json:"planes"`
}

type PlanEvento struct {
	gorm.Model
	Nombre            string          `json:"nombre"`
	Precio            float64         `json:"precio"`
	Decoracion        bool            `json:"decoracion"`
	Pintucaritas      bool            `json:"pintucaritas"`
	HoraLoca          bool            `json:"hora_loca"`
	Magia             bool            `json:"magia"`
	Titeres           bool            `json:"titeres"`
	Recreacion        bool            `json:"recreacion"`
	Comida            bool            `json:"comida"`
	Disfraces         string          `json:"disfraces"`
	Estado            bool            `json:"estado"`
	CategoriaEvento   CategoriaEvento `json:"categoria_evento"`
	CategoriaEventoID uint            `json:"categoria_evento_ID" gorm:"not null"`
}
