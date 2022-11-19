package main

import (
	"context"
	"fmt"
	"log"

	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const address = "localhost:50051"

func main() {
	ctx := context.Background()

	con, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("didn't connect: %s", err.Error())
	}
	defer con.Close()

	client := desc.NewNoteV1Client(con)
	resCreate, err := client.CreateNote(ctx, &desc.CreateNoteRequest{
		NoteBody: &desc.NoteBody{
			Title:  "Wow",
			Text:   "I'm surprised",
			Author: "Almira",
			Email:  "lalala@mail.ru",
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	log.Printf("Note With Id %d Was Created", resCreate.GetId())
	fmt.Println()

	resGet, err := client.GetNote(ctx, &desc.GetNoteRequest{
		Id: 2,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	log.Printf("Got Note With Id %d", resGet.GetWholeNote().GetId())
	log.Println("Title:", resGet.GetWholeNote().GetNoteBody().GetTitle())
	log.Println("Text:", resGet.GetWholeNote().GetNoteBody().GetText())
	log.Println("Author:", resGet.GetWholeNote().GetNoteBody().GetAuthor())
	log.Println("Email:", resGet.GetWholeNote().GetNoteBody().GetEmail())
	log.Println("Created At:", resGet.GetWholeNote().GetCreatedAt().AsTime())
	if resGet.GetWholeNote().GetUpdatedAt().GetSeconds() == 0 &&
		resGet.GetWholeNote().GetUpdatedAt().GetNanos() == 0 {
		log.Println("Updated At:", nil)
	} else {
		log.Println("Updated At:", resGet.GetWholeNote().GetUpdatedAt().AsTime())
	}
	fmt.Println()

	resGetList, err := client.GetListNote(ctx, &desc.GetListNoteRequest{
		Ids: []int64{2, 3},
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	log.Println("Got These Notes")
	for i := 0; i < len(resGetList.GetResults()); i++ {
		log.Println("Note Id:", resGetList.GetResults()[i].GetId())
		log.Println("Title:", resGetList.GetResults()[i].GetNoteBody().GetTitle())
		log.Println("Text:", resGetList.GetResults()[i].GetNoteBody().GetText())
		log.Println("Author:", resGetList.GetResults()[i].GetNoteBody().GetAuthor())
		log.Println("Email:", resGetList.GetResults()[i].GetNoteBody().GetEmail())
		log.Println("Created At:", resGetList.GetResults()[i].GetCreatedAt().AsTime())
		if resGetList.GetResults()[i].GetUpdatedAt().GetSeconds() == 0 &&
			resGetList.GetResults()[i].GetUpdatedAt().GetNanos() == 0 {
			log.Println("Updated At:", nil)
		} else {
			log.Println("Updated At:", resGetList.GetResults()[i].GetUpdatedAt().AsTime())
		}
		fmt.Println()
	}

	resUpdate, err := client.UpdateNote(ctx, &desc.UpdateNoteRequest{
		Id: 3,
		NoteBody: &desc.NoteBody{
			Title:  "New Title",
			Text:   "New Text",
			Author: "Not Almira",
			Email:  "example@mail.com",
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	log.Printf("Note %d Was Updated", resUpdate.GetId())
	fmt.Println()

	_, err = client.DeleteNote(ctx, &desc.DeleteNoteRequest{
		Id: 5,
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	log.Println("Note Was Successfully Deleted")
}
