package service

import (
	"context"
	"go-service/internal/gen"
)

func (r *queryResolver) GetMuser(ctx context.Context, id int) (*gen.Muser, error) {
	return r.Service.GetMuser(ctx, id)
}

func (svc *Service) GetMuser(ctx context.Context, id int) (*gen.Muser, error) {
	var r gen.Muser
	err1 := svc.DB.Where("id = ?", id).First(&r).Error
	if err1 != nil {
		return &gen.Muser{ID: -1, Name: "error", Code: ""}, nil
	}

	muser := &gen.Muser{
		ID:   r.ID,
		Name: r.Name,
		Code: r.Code,
	}
	return muser, nil

}
