package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"

	desc "github.com/almira-galeeva/testGrpc/pkg/note_v1"
)

const address = "localhost:50051"

func main() {
	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn't connect: %s", err.Error())
	}
	defer con.Close()

	client := desc.NewNoteV1Client(con)
	resCreate, err := client.CreateNote(context.Background(), &desc.CreateNoteRequest{
		Title:  "Wow",
		Text:   "I'm surprised",
		Author: "Almira",
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	log.Printf("Note With Id %d Was Created", resCreate.Id)
	fmt.Println()

	resGet, err := client.GetNote(context.Background(), &desc.GetNoteRequest{
		Id: 1,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	log.Println("Got This Note")
	log.Println("Title:", resGet.Title)
	log.Println("Text:", resGet.Text)
	log.Println("Author:", resGet.Author)
	fmt.Println()

	resGetList, err := client.GetListNote(context.Background(), &desc.GetListNoteRequest{
		Ids: []int64{1, 2},
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	log.Println("Got These Notes")
	for i := 0; i < len(resGetList.GetResults()); i++ {
		log.Println("Title:", resGetList.GetResults()[i].Title)
		log.Println("Text:", resGetList.GetResults()[i].Text)
		log.Println("Author:", resGetList.GetResults()[i].Author)
		fmt.Println()
	}

	resUpdate, err := client.UpdateNote(context.Background(), &desc.UpdateNoteRequest{
		Id:     1,
		Title:  "New Title",
		Text:   "New Text",
		Author: "Not Almira",
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	log.Printf("Note %d Was Updated. Status Code: %d", resUpdate.Id, resUpdate.Res)
	fmt.Println()

	resDelete, err := client.DeleteNote(context.Background(), &desc.DeleteNoteRequest{
		Id: 1,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	log.Printf("Note %d Was Deleted. Status Code: %d", resDelete.Id, resDelete.Res)
}
