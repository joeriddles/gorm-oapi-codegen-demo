// Code generated by github.com/joeriddles/gorm-oapi-codegen DO NOT EDIT.
package repository

import (
  "context"

  model "github.com/joeriddles/gorm-oapi-codegen-demo/pkg/model"
)

type ManufacturerRepository interface {
  List(
    ctx context.Context,
    filters any,
  ) ([]*model.Manufacturer, error)

  Get(
    ctx context.Context,
    id int64,
  ) (*model.Manufacturer, error)
  
  Create(
    ctx context.Context,
    manufacturer model.Manufacturer,
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

type _ManufacturerRepository struct {
}

func NewManufacturerRepository() ManufacturerRepository {
  return &_ManufacturerRepository{}
}

func (r *_ManufacturerRepository) List(
  ctx context.Context,
  filters any,
) ([]*model.Manufacturer, error) {
  return nil, nil
}

func (r *_ManufacturerRepository) Get(
  ctx context.Context,
  id int64,
) (*model.Manufacturer, error) {
  return nil, nil
}

func (r *_ManufacturerRepository) Create(
  ctx context.Context,
  manufacturer model.Manufacturer,
) error {
  return nil
}

func (r *_ManufacturerRepository) Update(
  ctx context.Context,
  update map[string]interface{},
) error {
  return nil
}

func (r *_ManufacturerRepository) Delete(
  ctx context.Context,
  id int64,
) error {
  return nil
}
