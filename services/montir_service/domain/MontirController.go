package domain

import (
	"context"
	"database/sql"
	"tublessin/common/model"
)

type MontirServer struct {
	MontirUsecase MontirUsecaseInterface
}

func NewMontirController(db *sql.DB) *MontirServer {
	return &MontirServer{NewMontirUsecase(db)}
}

// Disini adalah pusat Method2 dari Montir-Service
func (c MontirServer) Login(ctx context.Context, param *model.MontirAccount) (*model.MontirAccount, error) {
	result, err := c.MontirUsecase.Login(param)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c MontirServer) RegisterNewMontir(ctx context.Context, param *model.MontirAccount) (*model.MontirResponeMessage, error) {
	result, err := c.MontirUsecase.RegisterNewMontir(param)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Ini adalah method yang dimiliki oleh Montir-Service untuk mendapatkan data Montir Secara lengkap
func (c MontirServer) GetMontirProfileByID(ctx context.Context, param *model.MontirAccount) (*model.MontirResponeMessage, error) {
	montirResponeMessage, err := c.MontirUsecase.GetMontirProfileByID(param.Id)
	if err != nil {
		return nil, err
	}

	return montirResponeMessage, nil
}

func (c MontirServer) UpdateMontirProfilePicture(ctx context.Context, param *model.MontirProfile) (*model.MontirResponeMessage, error) {
	montirResponeMessage, err := c.MontirUsecase.UpdateMontirProfilePicture(param)
	if err != nil {
		return nil, err
	}

	return montirResponeMessage, nil
}

func (c MontirServer) UpdateMontirProfileByID(ctx context.Context, param *model.MontirProfile) (*model.MontirResponeMessage, error) {
	montirResponeMessage, err := c.MontirUsecase.UpdateMontirProfileByID(param)
	if err != nil {
		return nil, err
	}

	return montirResponeMessage, nil
}

func (c MontirServer) UpdateMontirStatusByID(ctx context.Context, param *model.MontirProfile) (*model.MontirResponeMessage, error) {
	montirResponeMessage, err := c.MontirUsecase.UpdateMontirStatusByID(param)
	if err != nil {
		return nil, err
	}

	return montirResponeMessage, nil
}

func (c MontirServer) UpdateMontirLocation(ctx context.Context, param *model.MontirProfile) (*model.MontirResponeMessage, error) {
	montirResponeMessage, err := c.MontirUsecase.UpdateMontirLocation(param)
	if err != nil {
		return nil, err
	}

	return montirResponeMessage, nil
}

func (c MontirServer) GetAllActiveMontirWithLocation(ctx context.Context, param *model.RequestActiveMontir) (*model.ListActiveMontirWithLocation, error) {
	result, err := c.MontirUsecase.GetAllActiveMontirWithLocation(param)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c MontirServer) DeleteMontirByID(ctx context.Context, param *model.MontirAccount) (*model.MontirResponeMessage, error) {
	montirResponeMessage, err := c.MontirUsecase.DeleteMontirByID(param)
	if err != nil {
		return nil, err
	}

	return montirResponeMessage, nil
}

func (c MontirServer) GetAllMontirSummary(ctx context.Context, param *model.MontirPagination) (*model.ListActiveMontirWithLocation, error) {
	result, err := c.MontirUsecase.GetAllMontirSummary(param)
	if err != nil {
		return nil, err
	}

	return result, nil
}
