package service

import (
	"context"

	"github.com/olezhek28/code-base/6_unit_test_in_clean_arch/internal/model"
)

type NoteService interface {
	Create(ctx context.Context, info model.NoteInfo) (string, error)
	Get(ctx context.Context, uuid string) (model.Note, error)
}
