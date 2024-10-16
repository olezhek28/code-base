package note

import (
	"context"
	"log"

	"github.com/olezhek28/code-base/6_unit_test_in_clean_arch/internal/converter"
	desc "github.com/olezhek28/code-base/6_unit_test_in_clean_arch/pkg/note_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	noteObj, err := i.noteService.Get(ctx, req.GetUuid())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %s, title: %s, body: %s, created_at: %v, updated_at: %v\n", noteObj.UUID, noteObj.Info.Title, noteObj.Info.Content, noteObj.CreatedAt, noteObj.UpdatedAt)

	return &desc.GetResponse{
		Note: converter.ToNoteFromService(noteObj),
	}, nil
}
