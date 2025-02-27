package slip2surego

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/Slip2Sure/slip2sure-go/errors"
	"github.com/Slip2Sure/slip2sure-go/model"
	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
)

type Slip2SureAPI struct {
	ApiToken string
}

var (
	api_url = "https://api.slip2sure.com/api/v0"
)

func getExtension(raw []byte) string {
	// Detect MIME type and get extension
	mtype := mimetype.Detect(raw)

	return mtype.Extension()
}

func (cfg *Slip2SureAPI) ScanTruemoneySlip(image []byte) (*model.Slip2SureTruemoney, error) {
	// Get extension
	extension := getExtension(image)
	if extension == "" {
		return nil, fmt.Errorf("Invalid extension MIME type.")
	}

	// Create UUID file request
	filename, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	// Create payload
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	// Create form file
	part, err := writer.CreateFormFile("file[]", fmt.Sprintf("%s.%s", filename.String(), extension))
	if err != nil {
		return nil, fmt.Errorf("Failed to create form file: %v", err)
	}
	part.Write(image)
	// Close writer
	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("Failed to close writer: %v", err)
	}

	// Create request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/truemoney/v1/verify", api_url), payload)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %w", err)
	}
	// Set headers
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Request
	body, statusCode, err := requestAPI(req, cfg.ApiToken)
	if err != nil {
		return nil, fmt.Errorf("Failed to request API: %w", err)
	}

	if statusCode > 400 {
		return nil, handleError(body)
	}

	// Parse JSON
	var resp model.Slip2SureTruemoney
	if err := json.Unmarshal(body.Bytes(), &resp); err != nil {
		return nil, fmt.Errorf("Parse JSON failed: %v", err)
	}

	return &resp, nil
}

func (cfg *Slip2SureAPI) ScanBankSlipByPayload(payload string) (*model.Slip2SureTruemoney, error) {
	// Create payload
	_payload := map[string]string{"payload": payload}
	// Create JSON value
	jsonMarsal, err := json.Marshal(_payload)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal JSON: %v", err)
	}

	// Create request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/bank/v0/verifyByPayload", api_url), bytes.NewBuffer(jsonMarsal))
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %w", err)
	}
	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Request
	body, statusCode, err := requestAPI(req, cfg.ApiToken)
	if err != nil {
		return nil, fmt.Errorf("Failed to request API: %w", err)
	}

	if statusCode > 400 {
		return nil, handleError(body)
	}

	// Parse JSON
	var resp model.Slip2SureTruemoney
	if err := json.Unmarshal(body.Bytes(), &resp); err != nil {
		return nil, fmt.Errorf("Parse JSON failed: %v", err)
	}

	return &resp, nil
}

func requestAPI(req *http.Request, token string) (*bytes.Buffer, int, error) {
	// Add token
	req.Header.Set("x-api-key", token)

	// HTTP Client
	client := http.Client{}
	// Execute request
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("request failed: %w", err)
	}

	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to read response: %w", err)
	}
	resp.Body.Close()

	return body, resp.StatusCode, nil
}

func handleError(raw *bytes.Buffer) error {
	var resp model.Slip2SureErrorResponse

	// Parse JSON
	if err := json.Unmarshal(raw.Bytes(), &resp); err != nil {
		return fmt.Errorf("Parse JSON failed: %v", err)
	}

	switch resp.Result {
	case "INVALID_HEADER":
		return errors.INVALID_HEADER
	case "UNAUTHORIZED":
		return errors.UNAUTHORIZED
	case "VALIDATE_ERROR":
		return errors.VALIDATE_ERROR
	case "CREDIT_INSUFFIENCT":
		return errors.CREDIT_INSUFFIENCT
	case "FILE_REQUIRED":
		return errors.FILE_REQUIRED
	case "FILE_TOO_LARGE":
		return errors.FILE_TOO_LARGE
	case "FILE_NOT_SUPPORTED":
		return errors.FILE_NOT_SUPPORTED
	case "SLIP_NOT_EXIST":
		return errors.SLIP_NOT_EXIST
	case "SERVER_ERROR":
		return errors.SERVER_ERROR
	case "SERVICE_ERROR":
		return errors.SERVICE_ERROR
	case "SERVICE_TIMEOUT":
		return errors.SERVICE_TIMEOUT
	default:
		return errors.UNKNOWN_ERROR
	}
}
