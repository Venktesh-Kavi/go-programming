package document

import (
	"bulkflow/utils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// UploadRequest represents the request body for the upload API
type UploadRequest struct {
	FileName string `json:"fileName"`
	MimeType string `json:"mimeType"`
}

// UploadResponse represents the response from the upload API
type UploadResponse struct {
	FileID       string `json:"fileId"`
	PresignedURL string `json:"presignedURL"`
}

// Tag represents a tag in the create document request
type Tag struct {
	TagType  string `json:"tagType"`
	TagValue string `json:"tagValue"`
}

// CreateDocumentRequest represents the request body for create document API
type CreateDocumentRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	FileID      string `json:"fileId"`
	DocType     string `json:"docType"`
	DocSubType  string `json:"docSubType"`
	Platform    string `json:"platform"`
	UserId      string `json:"userId"`
	Tags        []Tag  `json:"tags"`
}

const (
	baseURL = "https://dev-oci-loanos-scf-core-api.go-yubi.in"
)

func UploadDoc(ctx context.Context, fp string) (string, error) {
	// Get the filename from the path
	fileName := filepath.Base(fp)
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: &utils.ContextRoundTripper{},
	}

	// Step 1: Get upload URL
	uploadReq := UploadRequest{
		FileName: fileName,
		MimeType: "text/csv",
	}

	uploadResp, err := getUploadURL(ctx, client, uploadReq)
	if err != nil {
		log.Fatalf("Failed to get upload URL: %v", err)
	}

	log.Println("generated uploadable presigned url with dms fileId: ", uploadResp.FileID)

	// Step 2: Upload file to presigned URL
	err = uploadFileToPresignedURL(ctx, client, uploadResp.PresignedURL, fp)
	if err != nil {
		log.Fatalf("Failed to upload file: %v", err)
	}

	log.Println("successfully uploaded file to presigned url")

	// Step 3: Create document
	createDocReq := CreateDocumentRequest{
		Title:       "SCF VF DF Mini Invoice Dataset",
		Description: "vf df invoice dataset",
		FileID:      uploadResp.FileID,
		DocType:     "instrument",
		DocSubType:  "invoice",
		Platform:    "yubi",
		UserId:      "scf_loanos",
		Tags: []Tag{
			{
				TagType:  "instrument",
				TagValue: "100",
			},
		},
	}

	rj, _ := json.Marshal(createDocReq)
	fmt.Sprintln("Creating document payload: %s", rj)
	err = createDocument(ctx, client, createDocReq)
	if err != nil {
		log.Fatalf("Failed to create document: %v", err)
	}

	log.Println("document created successfully in DMS, fileId: ", uploadResp.FileID)

	dresp, err := createDownloadablePresignedUrl(ctx, client, uploadResp.FileID)
	log.Println("Downloadable presigned url generated successfully: ", dresp)
	fmt.Println("Document upload and creation completed successfully")
	return dresp, nil
}

func getUploadURL(ctx context.Context, client *http.Client, uploadReq UploadRequest) (*UploadResponse, error) {
	jsonData, err := json.Marshal(uploadReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal upload request: %v", err)
	}

	ru := fmt.Sprintf("%s/api/v1/document/upload", baseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ru, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create post request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make upload request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("upload request failed with status: %d", resp.StatusCode)
	}

	var uploadResp UploadResponse
	if err := json.NewDecoder(resp.Body).Decode(&uploadResp); err != nil {
		return nil, fmt.Errorf("failed to decode upload response: %v", err)
	}

	return &uploadResp, nil
}

func uploadFileToPresignedURL(ctx context.Context, client *http.Client, presignedURL, filePath string) error {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Get file info to set the correct Content-Length
	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %v", err)
	}

	// Create PUT request
	req, err := http.NewRequest(http.MethodPut, presignedURL, file)
	if err != nil {
		return fmt.Errorf("failed to create PUT request: %v", err)
	}

	// Set required headers
	req.Header.Set("Content-Type", "text/csv") // Update as needed; AWS typically accepts text/csv for CSV uploads
	req.ContentLength = fileInfo.Size()

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to upload file: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("file upload failed with status: %d, body: %s", resp.StatusCode, body)
	}

	return nil
}

func createDocument(ctx context.Context, client *http.Client, createDocReq CreateDocumentRequest) error {
	jsonData, err := json.Marshal(createDocReq)
	if err != nil {
		return fmt.Errorf("failed to marshal create document request: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL+"/api/v1/create-document", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create post request: %v", err)
	}

	req.Header.Set("X-Entity-Id", createDocReq.UserId)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make create document request: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("create document request failed with status: %d, received response: %s", resp.StatusCode, body)
	}

	return nil
}

func createDownloadablePresignedUrl(ctx context.Context, client *http.Client, fileId string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL+"/api/v1/document/download/"+fileId, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create download request: %v", err)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("download request failed with status: %d", resp.StatusCode)
	}

	var downloadResp struct {
		PresignedURL string `json:"presignedURL"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&downloadResp); err != nil {
		return "", fmt.Errorf("failed to decode download response: %v", err)
	}

	return downloadResp.PresignedURL, nil
}
