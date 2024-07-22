package api

import (
	"context"
)

var _ StrictServerInterface = (*Server)(nil)

type Server struct {
  ManufacturerController
  ModelController
  VehicleController
  PartController
  PersonController
  
}

func NewServer() *Server {
	return &Server{
    ManufacturerController: NewManufacturerController(),
    ModelController: NewModelController(),
    VehicleController: NewVehicleController(),
    PartController: NewPartController(),
    PersonController: NewPersonController(),
    
  }
}

func (s *Server) GetManufacturer(ctx context.Context, request GetManufacturerRequestObject) (GetManufacturerResponseObject, error) {
  return s.ManufacturerController.GetManufacturer(ctx, request)
}

func (s *Server) PostManufacturer(ctx context.Context, request PostManufacturerRequestObject) (PostManufacturerResponseObject, error) {
  return s.ManufacturerController.PostManufacturer(ctx, request)
}

func (s *Server) DeleteManufacturerId(ctx context.Context, request DeleteManufacturerIdRequestObject) (DeleteManufacturerIdResponseObject, error) {
  return s.ManufacturerController.DeleteManufacturerId(ctx, request)
}

func (s *Server) GetManufacturerId(ctx context.Context, request GetManufacturerIdRequestObject) (GetManufacturerIdResponseObject, error) {
  return s.ManufacturerController.GetManufacturerId(ctx, request)
}

func (s *Server) PutManufacturerId(ctx context.Context, request PutManufacturerIdRequestObject) (PutManufacturerIdResponseObject, error) {
  return s.ManufacturerController.PutManufacturerId(ctx, request)
}
func (s *Server) GetModel(ctx context.Context, request GetModelRequestObject) (GetModelResponseObject, error) {
  return s.ModelController.GetModel(ctx, request)
}

func (s *Server) PostModel(ctx context.Context, request PostModelRequestObject) (PostModelResponseObject, error) {
  return s.ModelController.PostModel(ctx, request)
}

func (s *Server) DeleteModelId(ctx context.Context, request DeleteModelIdRequestObject) (DeleteModelIdResponseObject, error) {
  return s.ModelController.DeleteModelId(ctx, request)
}

func (s *Server) GetModelId(ctx context.Context, request GetModelIdRequestObject) (GetModelIdResponseObject, error) {
  return s.ModelController.GetModelId(ctx, request)
}

func (s *Server) PutModelId(ctx context.Context, request PutModelIdRequestObject) (PutModelIdResponseObject, error) {
  return s.ModelController.PutModelId(ctx, request)
}
func (s *Server) GetVehicle(ctx context.Context, request GetVehicleRequestObject) (GetVehicleResponseObject, error) {
  return s.VehicleController.GetVehicle(ctx, request)
}

func (s *Server) PostVehicle(ctx context.Context, request PostVehicleRequestObject) (PostVehicleResponseObject, error) {
  return s.VehicleController.PostVehicle(ctx, request)
}

func (s *Server) DeleteVehicleId(ctx context.Context, request DeleteVehicleIdRequestObject) (DeleteVehicleIdResponseObject, error) {
  return s.VehicleController.DeleteVehicleId(ctx, request)
}

func (s *Server) GetVehicleId(ctx context.Context, request GetVehicleIdRequestObject) (GetVehicleIdResponseObject, error) {
  return s.VehicleController.GetVehicleId(ctx, request)
}

func (s *Server) PutVehicleId(ctx context.Context, request PutVehicleIdRequestObject) (PutVehicleIdResponseObject, error) {
  return s.VehicleController.PutVehicleId(ctx, request)
}
func (s *Server) GetPart(ctx context.Context, request GetPartRequestObject) (GetPartResponseObject, error) {
  return s.PartController.GetPart(ctx, request)
}

func (s *Server) PostPart(ctx context.Context, request PostPartRequestObject) (PostPartResponseObject, error) {
  return s.PartController.PostPart(ctx, request)
}

func (s *Server) DeletePartId(ctx context.Context, request DeletePartIdRequestObject) (DeletePartIdResponseObject, error) {
  return s.PartController.DeletePartId(ctx, request)
}

func (s *Server) GetPartId(ctx context.Context, request GetPartIdRequestObject) (GetPartIdResponseObject, error) {
  return s.PartController.GetPartId(ctx, request)
}

func (s *Server) PutPartId(ctx context.Context, request PutPartIdRequestObject) (PutPartIdResponseObject, error) {
  return s.PartController.PutPartId(ctx, request)
}
func (s *Server) GetPerson(ctx context.Context, request GetPersonRequestObject) (GetPersonResponseObject, error) {
  return s.PersonController.GetPerson(ctx, request)
}

func (s *Server) PostPerson(ctx context.Context, request PostPersonRequestObject) (PostPersonResponseObject, error) {
  return s.PersonController.PostPerson(ctx, request)
}

func (s *Server) DeletePersonId(ctx context.Context, request DeletePersonIdRequestObject) (DeletePersonIdResponseObject, error) {
  return s.PersonController.DeletePersonId(ctx, request)
}

func (s *Server) GetPersonId(ctx context.Context, request GetPersonIdRequestObject) (GetPersonIdResponseObject, error) {
  return s.PersonController.GetPersonId(ctx, request)
}

func (s *Server) PutPersonId(ctx context.Context, request PutPersonIdRequestObject) (PutPersonIdResponseObject, error) {
  return s.PersonController.PutPersonId(ctx, request)
}

