package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/olezhek28/code-base/6_unit_test_in_clean_arch/internal/api/note"
	noteRepository "github.com/olezhek28/code-base/6_unit_test_in_clean_arch/internal/repository/note"
	noteService "github.com/olezhek28/code-base/6_unit_test_in_clean_arch/internal/service/note"
	desc "github.com/olezhek28/code-base/6_unit_test_in_clean_arch/pkg/note_v1"
)

const (
	grpcAddress = ":50051"
)

func main() {
	noteRepo := noteRepository.NewRepository()
	noteSrv := noteService.NewService(noteRepo)
	noteImpl := note.NewImplementation(noteSrv)

	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	reflection.Register(grpcServer)
	desc.RegisterNoteV1Server(grpcServer, noteImpl)

	list, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatalf("failed to listen: %s", err.Error())
	}

	err = grpcServer.Serve(list)
	if err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}
