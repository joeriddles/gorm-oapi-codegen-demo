package model

import (
	"gorm.io/gorm"
)

// A vehicle manufacturer, like Chevrolet
type Manufacturer struct {
	gorm.Model
	Name     string
	Vehicles []Model
}

// A vehicle model, like a Chevrolet Silverado
type Model struct {
	gorm.Model
	Name           string
	ManufacturerID uint
	Manufacturer   *Manufacturer
	Parts          []*Part `gorm:"many2many:vehicle_parts;"`
}

// An individual of a model, like Joe's Chevrolet Silverado
type Vehicle struct {
	gorm.Model
	Vin            string
	VehicleModelID uint
	VehicleModel   Model
	PersonID       *int
	Person         *Person
}

// A vehicle part for one or more models, like a muffler for all Chevrolet pickups
type Part struct {
	gorm.Model
	Name   string
	Cost   int
	Models []*Model `gorm:"many2many:model_parts;"`
}

// A person, who may drive a vehicle
type Person struct {
	gorm.Model
	Name string
}
