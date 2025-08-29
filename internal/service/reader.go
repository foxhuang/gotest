package service

import (
	"context"
	"go-service/internal/gen"
)

func (r *queryResolver) GetReader(ctx context.Context, id int) (*gen.Reader, error) {
	return r.Service.GetReader(ctx, id)
}

func (svc *Service) GetReader(ctx context.Context, id int) (*gen.Reader, error) {
	var r gen.Reader
	err1 := svc.DB.Table("reader").Where("id = ?", id).First(&r).Error
	if err1 != nil {
		return &gen.Reader{ID: -1, Name: "error", Email: ""}, nil
	}

	reader := &gen.Reader{
		ID:    r.ID,
		Name:  r.Name,
		Email: r.Email,
	}
	return reader, nil

}
