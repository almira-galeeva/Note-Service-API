package main

import (
	"context"
	"fmt"
	"log"

	desc "github.com/almira-galeeva/note-service-api/pkg/note_v1"
	"google.golang.org/grpc"
)

const address = "localhost:50051"

func main() {
	ctx := context.Background()

	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn't connect: %s", err.Error())
	}
	defer con.Close()

	client := desc.NewNoteV1Client(con)
	resCreate, err := client.CreateNote(ctx, &desc.CreateNoteRequest{
		Title:  "Wow",
		Text:   "I'm surprised",
		Author: "Almira",
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	log.Printf("Note With Id %d Was Created", resCreate.GetId())
	fmt.Println()

	resGet, err := client.GetNote(ctx, &desc.GetNoteRequest{
		Id: 1,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	log.Printf("Got Note With Id %d", resGet.GetId())
	log.Println("Title:", resGet.GetTitle())
	log.Println("Text:", resGet.GetText())
	log.Println("Author:", resGet.GetAuthor())
	fmt.Println()

	resGetList, err := client.GetListNote(ctx, &desc.GetListNoteRequest{
		Ids: []int64{1, 2},
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	log.Println("Got These Notes")
	for i := 0; i < len(resGetList.GetResults()); i++ {
		log.Println("Note Id:", resGetList.GetResults()[i].GetId())
		log.Println("Title:", resGetList.GetResults()[i].GetTitle())
		log.Println("Text:", resGetList.GetResults()[i].GetText())
		log.Println("Author:", resGetList.GetResults()[i].GetAuthor())
		fmt.Println()
	}

	resUpdate, err := client.UpdateNote(ctx, &desc.UpdateNoteRequest{
		Id:     1,
		Title:  "New Title",
		Text:   "New Text",
		Author: "Not Almira",
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	log.Printf("Note %d Was Updated. Status Code: %d", resUpdate.GetId(), resUpdate.GetRes())
	fmt.Println()

	_, err = client.DeleteNote(ctx, &desc.DeleteNoteRequest{
		Id: 1,
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	log.Println("Note Was Successfully Deleted")

}
