package note

import (
	"github.com/olezhek28/code-base/6_unit_test_in_clean_arch/internal/service"
	desc "github.com/olezhek28/code-base/6_unit_test_in_clean_arch/pkg/note_v1"
)

type Implementation struct {
	desc.UnimplementedNoteV1Server
	noteService service.NoteService
}

func NewImplementation(noteService service.NoteService) *Implementation {
	return &Implementation{
		noteService: noteService,
	}
}
