package note

import (
	"context"
	"log"

	"github.com/olezhek28/code-base/6_unit_test_in_clean_arch/internal/converter"
	desc "github.com/olezhek28/code-base/6_unit_test_in_clean_arch/pkg/note_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	uuid, err := i.noteService.Create(ctx, converter.ToNoteInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted note with id: %s", uuid)

	return &desc.CreateResponse{
		Uuid: uuid,
	}, nil
}
