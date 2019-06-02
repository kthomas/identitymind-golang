package identitymind

import "fmt"

// KYC

// GetApplication see https://edoc.identitymind.com/reference#getv2
func (i *IdentityMindAPIClient) GetApplication(applicationID string) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Get(fmt.Sprintf("im/account/consumer/v2/%s", applicationID), map[string]interface{}{}, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve KYC application via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// SubmitApplication see https://edoc.identitymind.com/reference#create
func (i *IdentityMindAPIClient) SubmitApplication(params map[string]interface{}) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Post("im/account/consumer?graphScoreResponse=false", params, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to upload consumer KYC document via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// ListDocuments see https://edoc.identitymind.com/reference#getfilelistforapplication
func (i *IdentityMindAPIClient) ListDocuments(applicationID string) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Get(fmt.Sprintf("im/account/consumer/%s/files", applicationID), map[string]interface{}{}, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to list KYC documents via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// DownloadDocument see https://edoc.identitymind.com/reference#reevaluatemerchant
func (i *IdentityMindAPIClient) DownloadDocument(applicationID, documentID string) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Get(fmt.Sprintf("im/account/consumer/%s/files/%s", applicationID, documentID), map[string]interface{}{}, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to download KYC document via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// UploadDocument see https://edoc.identitymind.com/reference#processfileuploadrequest
func (i *IdentityMindAPIClient) UploadDocument(applicationID string, params map[string]interface{}) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Post(fmt.Sprintf("im/account/consumer/%s/files", applicationID), params, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to upload consumer KYC document via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// UploadDocumentVerificationImage see https://edoc.identitymind.com/reference#processimageuploadrequest
func (i *IdentityMindAPIClient) UploadDocumentVerificationImage(applicationID string, params map[string]interface{}) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Post(fmt.Sprintf("im/account/consumer/%s/dv", applicationID), params, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to upload consumer KYC document image for verification via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// ApproveApplication see https://edoc.identitymind.com/reference#feedback
func (i *IdentityMindAPIClient) ApproveApplication(applicationID string, params map[string]interface{}) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Post(fmt.Sprintf("im/account/consumer/%s/accepted", applicationID), params, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to approve KYC application via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// RejectApplication see https://edoc.identitymind.com/reference#feedback
func (i *IdentityMindAPIClient) RejectApplication(applicationID string, params map[string]interface{}) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Post(fmt.Sprintf("im/account/consumer/%s/rejected", applicationID), params, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to reject KYC application via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}
