// Code generated by github.com/joeriddles/gorm-oapi-codegen DO NOT EDIT.
package repository

import (
  "context"

  model "github.com/joeriddles/gorm-oapi-codegen-demo/pkg/model"
)

type PartRepository interface {
  List(
    ctx context.Context,
    filters any,
  ) ([]*model.Part, error)

  Get(
    ctx context.Context,
    id int64,
  ) (*model.Part, error)
  
  Create(
    ctx context.Context,
    part model.Part,
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

type _PartRepository struct {
}

func NewPartRepository() PartRepository {
  return &_PartRepository{}
}

func (r *_PartRepository) List(
  ctx context.Context,
  filters any,
) ([]*model.Part, error) {
  return nil, nil
}

func (r *_PartRepository) Get(
  ctx context.Context,
  id int64,
) (*model.Part, error) {
  return nil, nil
}

func (r *_PartRepository) Create(
  ctx context.Context,
  part model.Part,
) error {
  return nil
}

func (r *_PartRepository) Update(
  ctx context.Context,
  update map[string]interface{},
) error {
  return nil
}

func (r *_PartRepository) Delete(
  ctx context.Context,
  id int64,
) error {
  return nil
}
