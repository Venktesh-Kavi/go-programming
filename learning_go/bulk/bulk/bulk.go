package bulk

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
	"time"
)

type CreateBulkTransactionRequest struct {
	Sector         string `json:"sector"`
	InstrumentType string `json:"instrumentType"`
	InvestorId     string `json:"investorId"`
	FileTemplate   string `json:"fileTemplateName"`
	FileUrl        string `json:"fileUrl"`
}

type CreateBulkTransactionResponse struct {
	JobId string `json:"jobId"`
}

type JobStatusResponse struct {
	status string           `json:"status"`
	error  []JobStatusError `json:"errors"`
}

type JobStatusError struct {
	message string `json:"message"`
	code    string `json:"code"`
}

type JobStatsResponse struct {
	Total             int `json:"total"`
	Success           int `json:"success"`
	Failure           int `json:"failure"`
	CustomAggregation int `json:"customAggregation"`
}

const (
	baseUrl = "https://dev-oci-loanos-scf-core-api.go-yubi.in"
)

func getBulkTxnReqPayload(presignedurl string) CreateBulkTransactionRequest {
	createBulkTxnReq := CreateBulkTransactionRequest{
		Sector:         "Dealer",
		InstrumentType: "INVOICE",
		InvestorId:     "636b968a5990f700402d2638",
		FileTemplate:   "scf_vf_df_invoice_template",
		FileUrl:        presignedurl,
	}

	return createBulkTxnReq
}

func setCommonHeaders(req *http.Request, entityId string, userId string, groupId string) error {
	req.Header.Set("X-User-Id", userId)
	req.Header.Set("X-Entity-Id", entityId)
	req.Header.Set("X-Group-Id", groupId)
	return nil
}
func CreateBulkTransaction(ctx context.Context, ps string) (*CreateBulkTransactionResponse, error) {
	client, err := utils.HttpClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create http client: %v", err)
	}

	var createBulkTxnReq CreateBulkTransactionRequest = getBulkTxnReqPayload(ps)
	jsonData, err := json.Marshal(createBulkTxnReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal create bulk transaction request: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseUrl+"/api/v1/transaction/bulk", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create bulk transaction request: %v", err)
	}

	setCommonHeaders(req, "65dc6b2b42407b004e962c48", "65dc6cd642407b004e962c49", "vendor")
	res, err := client.Do(req)

	defer res.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %v", err)
		}
		return nil, fmt.Errorf("create bulk transaction failed with status: %d, %v", res.StatusCode, string(bodyBytes))
	}

	var txnResp CreateBulkTransactionResponse
	if err := json.NewDecoder(res.Body).Decode(&txnResp); err != nil {
		return nil, fmt.Errorf("failed to decode upload response: %v", err)
	}

	return &txnResp, nil
}

func LongPollJobStatus(ctx context.Context, jobId string) (string, error) {
	client, err := utils.HttpClient()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseUrl+"/api/v1/bulk/result/"+jobId, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create long poll request: %v", err)
	}

	setCommonHeaders(req, "65dc6b2b42407b004e962c48", "65dc6cd642407b004e962c49", "vendor")

	res, err := client.Do(req)

	defer res.Body.Close()

	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return "", fmt.Errorf("failed to read response body: %v", err)
		}
		return "", fmt.Errorf("create bulk transaction failed with status: %d, %v", res.StatusCode, string(bodyBytes))
	}

	var sr JobStatusResponse
	if err := json.NewDecoder(res.Body).Decode(&sr); err != nil {
		return "", fmt.Errorf("failed to decode upload response: %v", err)
	}

	if sr.status != "SUCCESS" || sr.status != "ERROR" {
		// repoll the api
		fmt.Println("Repolling job status api after 10ms, received status: ", sr.status)
		time.Sleep(10 * time.Millisecond)
		LongPollJobStatus(ctx, jobId)
	}
	log.Println("Job status: ", sr.status)
	return sr.status, nil
}

func FetchJobStats(ctx context.Context, jobId string) (JobStatsResponse, error) {
	client, err := utils.HttpClient()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseUrl+"/api/v1/bulk/stats/"+jobId, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create long poll request: %v", err)
	}

	res, err := client.Do(req)

	setCommonHeaders(req, "65dc6b2b42407b004e962c48", "65dc6cd642407b004e962c49", "vendor")
	return nil, nil
}
