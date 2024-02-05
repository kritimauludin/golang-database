package repository

import (
	"context"
	"golang-database/model"
)

type CommentRepository interface {
	Insert(ctx context.Context, comment model.Comment) (model.Comment, error)
	FindById(ctx context.Context, id int32) (model.Comment, error)
	FindAll(ctx context.Context) ([]model.Comment, error)
}
