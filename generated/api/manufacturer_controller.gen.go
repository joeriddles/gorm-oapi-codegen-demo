// Code generated by github.com/joeriddles/gorm-oapi-codegen DO NOT EDIT.
package api

import (
  "context"

  "github.com/joeriddles/gorm-oapi-codegen-demo/generated/repository"
)

type ManufacturerController interface {
  GetManufacturer(ctx context.Context, request GetManufacturerRequestObject) (GetManufacturerResponseObject, error)
  PostManufacturer(ctx context.Context, request PostManufacturerRequestObject) (PostManufacturerResponseObject, error)
  DeleteManufacturerId(ctx context.Context, request DeleteManufacturerIdRequestObject) (DeleteManufacturerIdResponseObject, error)
  GetManufacturerId(ctx context.Context, request GetManufacturerIdRequestObject) (GetManufacturerIdResponseObject, error)
  PutManufacturerId(ctx context.Context, request PutManufacturerIdRequestObject) (PutManufacturerIdResponseObject, error)
}

type _ManufacturerController struct {
  repository repository.ManufacturerRepository
  mapper ManufacturerMapper
}

func NewManufacturerController() ManufacturerController {
  return &_ManufacturerController{
    repository: repository.NewManufacturerRepository(),
    mapper: NewManufacturerMapper(),
  }
}

func (c *_ManufacturerController) GetManufacturer(ctx context.Context, request GetManufacturerRequestObject) (GetManufacturerResponseObject, error) {
  manufacturers, err := c.repository.List(ctx, map[string]interface{}{})
  if err != nil {
    return nil, err
  }

  result := []Manufacturer{}
  for _, manufacturer := range manufacturers {
    apiManufacturer, err := c.mapper.MapToApi(*manufacturer)
    if err != nil {
      return nil, err
    }
    result = append(result, *apiManufacturer)
  }

  return GetManufacturer200JSONResponse(result), nil
}

func (c *_ManufacturerController) PostManufacturer(ctx context.Context, request PostManufacturerRequestObject) (PostManufacturerResponseObject, error) {
  model, err := c.mapper.MapCreate(*request.Body)
  if err != nil {
    return nil, err
  }
  if err = c.repository.Create(ctx, *model); err != nil {
    return nil, err
  }
  apiModel, err := c.mapper.MapToApi(*model)
  if err != nil {
    return nil, err
  }
  return PostManufacturer201JSONResponse(*apiModel), nil
}

func (c *_ManufacturerController) DeleteManufacturerId(ctx context.Context, request DeleteManufacturerIdRequestObject) (DeleteManufacturerIdResponseObject, error) {
  if err := c.repository.Delete(ctx, request.Id); err != nil {
    return nil, err
  }
  return DeleteManufacturerId204Response{}, nil
}

func (c *_ManufacturerController) GetManufacturerId(ctx context.Context, request GetManufacturerIdRequestObject) (GetManufacturerIdResponseObject, error) {
  model, err := c.repository.Get(ctx, request.Id)
  apiModel, err := c.mapper.MapToApi(*model)
  if err != nil {
    return nil, err
  }
  return GetManufacturerId200JSONResponse(*apiModel), err
}

func (c *_ManufacturerController) PutManufacturerId(ctx context.Context, request PutManufacturerIdRequestObject) (PutManufacturerIdResponseObject, error) {
  update, err := c.mapper.MapUpdate(*request.Body)
  if err != nil {
    return nil, err
  }
  if err := c.repository.Update(ctx, update); err != nil {
    return nil, err
  }
  return nil, nil
}
