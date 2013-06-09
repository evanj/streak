package main

import (
	"fmt"
	"github.com/evanj/streak"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Stderr.WriteString("Usage: streak_example (API key)\n")
		os.Exit(1)
	}
	client := streak.New(os.Args[1])

	fmt.Println("Pipelines:")
	pipelines, err := client.GetPipelines()
	if err != nil {
		log.Fatal("Failed to get pipelines: ", err)
	}
	for _, pipeline := range pipelines {
		fmt.Printf("  %s:\n", pipeline.Name)
		fmt.Printf("  .Key: %s:\n", pipeline.Key)
		fmt.Printf("  .Description: %s:\n\n", pipeline.Description)
	}

	fmt.Printf("Boxes in pipeline \"%s\":\n", pipelines[0].Name)
	boxes, err := client.GetBoxes(&pipelines[0])
	if err != nil {
		log.Fatal("Failed to get boxes: ", err)
	}
	var gmailThreadBox *streak.Box
	for i, box := range boxes {
		fmt.Printf("  %s:\n", box.Name)
		fmt.Printf("  .Key: %s:\n", box.Key)
		fmt.Printf("  .LastUpdatedTimestamp: %v:\n", streak.TimestampToTime(box.LastUpdatedTimestamp))
		fmt.Printf("  .GmailThreadCount: %d:\n\n", box.GmailThreadCount)
		if gmailThreadBox == nil && box.GmailThreadCount > 0 {
			gmailThreadBox = &boxes[i]
			fmt.Println("wtf?", gmailThreadBox)
		}
	}

	if gmailThreadBox != nil {
		fmt.Printf("Threads in box \"%s\":\n", gmailThreadBox.Name)
		threads, err := client.GetThreads(gmailThreadBox)
		if err != nil {
			log.Fatal("Failed to get threads: ", err)
		}
		for _, thread := range threads {
			fmt.Printf("  %s:\n", thread.Subject)
			fmt.Printf("  .Key: %s\n", thread.Key)
			fmt.Printf("  .CreationTimestamp: %v\n", streak.TimestampToTime(thread.CreationTimestamp))
			fmt.Printf("  .LastEmailTimestamp: %v\n", streak.TimestampToTime(thread.LastEmailTimestamp))
			fmt.Printf("  .LastUpdatedTimestamap: %v\n", streak.TimestampToTime(thread.LastUpdatedTimestamp))
		}
	}
}
