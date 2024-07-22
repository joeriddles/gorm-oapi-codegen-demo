// Code generated by github.com/joeriddles/gorm-oapi-codegen DO NOT EDIT.
package repository

import (
  "context"

  model "github.com/joeriddles/gorm-oapi-codegen-demo/main"
)

type VehicleRepository interface {
  List(
    ctx context.Context,
    filters any,
  ) ([]*model.Vehicle, error)

  Get(
    ctx context.Context,
    id int64,
  ) (*model.Vehicle, error)
  
  Create(
    ctx context.Context,
    vehicle model.Vehicle,
  ) error
  
  Update(
    ctx context.Context,
    update map[string]interface{},
  ) error
  
  Delete(
    ctx context.Context,
    id int64,
  ) error
}

type _VehicleRepository struct {
}

func NewVehicleRepository() VehicleRepository {
  return &_VehicleRepository{}
}

func (r *_VehicleRepository) List(
  ctx context.Context,
  filters any,
) ([]*model.Vehicle, error) {
  return nil, nil
}

func (r *_VehicleRepository) Get(
  ctx context.Context,
  id int64,
) (*model.Vehicle, error) {
  return nil, nil
}

func (r *_VehicleRepository) Create(
  ctx context.Context,
  vehicle model.Vehicle,
) error {
  return nil
}

func (r *_VehicleRepository) Update(
  ctx context.Context,
  update map[string]interface{},
) error {
  return nil
}

func (r *_VehicleRepository) Delete(
  ctx context.Context,
  id int64,
) error {
  return nil
}
