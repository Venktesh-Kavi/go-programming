package main

import (
	"bufio"
	"bulkflow/bulk"
	"bulkflow/document"
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Enter path to CSV file")

	r := bufio.NewReader(os.Stdin)
	path, _ := r.ReadString('\n')
	filepath.Clean(path)

	//fmt.Println("Enter path to your CSV file: ")
	//reader := bufio.NewReader(os.Stdin)
	//inputPath, err := reader.ReadString('\n')
	//inputPath = filepath.Clean(inputPath)
	//
	//if _, err := os.Stat(inputPath); os.IsExist(err) {
	//	log.Fatalf("File does not exist: %v", err)
	//}
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	inputPath := filepath.Join(cwd, "SCF_VF_DF_INVOICE_MINI_DATASET.csv")
	ctx := context.WithValue(context.Background(), "X-Request-Id", uuid.New().String())

	// Call the UploadDocument function from the bulk package
	ps, err := document.UploadDoc(ctx, inputPath)
	log.Println("Document upload and creation completed successfully")
	fmt.Println("Creating bulk transaction with presigned document url")
	if err != nil {
		log.Fatalf("Document upload failed: %v", err)
	}

	btr, err := bulk.CreateBulkTransaction(ctx, ps)
	if err != nil {
		log.Fatalf("Failed to create bulk transaction: %v", err)
	}
	fmt.Println("Bulk transaction created successfully with jobId: ", btr.JobId)

	fmt.Println("Checking job status")
	_, nerr := bulk.LongPollJobStatus(ctx, btr.JobId)
	if nerr != nil {
		log.Fatalf("Failed to check job status: %v", nerr)
	}
}
