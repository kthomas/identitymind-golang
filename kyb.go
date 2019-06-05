package identitymind

import "fmt"

// KYB

// GetBusinessApplication see https://edoc.identitymind.com/reference#getmerchantkyc
func (i *IdentityMindAPIClient) GetBusinessApplication(applicationID string) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Get(fmt.Sprintf("im/account/merchant/%s", applicationID), map[string]interface{}{}, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve KYB application via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// ReevaluateBusinessApplication see https://edoc.identitymind.com/reference#reevaluatemerchant
func (i *IdentityMindAPIClient) ReevaluateBusinessApplication(applicationID string) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Post(fmt.Sprintf("im/account/merchant/%s", applicationID), map[string]interface{}{}, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to reevaluate KYB application via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// SubmitBusinessApplication see https://edoc.identitymind.com/reference#merchant
func (i *IdentityMindAPIClient) SubmitBusinessApplication(params map[string]interface{}) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Post("im/account/merchant?graphScoreResponse=false", params, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to reevaluate KYB application via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// ListBusinessApplicationDocuments see https://edoc.identitymind.com/reference#getfilelistforapplicationformerchant
func (i *IdentityMindAPIClient) ListBusinessApplicationDocuments(applicationID string) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Get(fmt.Sprintf("im/account/merchant/%s/files", applicationID), map[string]interface{}{}, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to list KYB documents via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// DownloadBusinessApplicationDocument see https://edoc.identitymind.com/reference#reevaluatemerchant
func (i *IdentityMindAPIClient) DownloadBusinessApplicationDocument(applicationID, documentID string) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Get(fmt.Sprintf("im/account/consumer/%s/files/%s", applicationID, documentID), map[string]interface{}{}, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to download KYB document via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// UploadBusinessApplicationDocument see https://edoc.identitymind.com/reference#processfileuploadrequestformerchantkyc
func (i *IdentityMindAPIClient) UploadBusinessApplicationDocument(applicationID string, params map[string]interface{}) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Post(fmt.Sprintf("im/account/merchant/%s/files", applicationID), params, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to upload KYB document via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// UploadBusinessApplicationDocumentVerificationImage see https://edoc.identitymind.com/reference#processfileuploadrequestformerchantkyc
func (i *IdentityMindAPIClient) UploadBusinessApplicationDocumentVerificationImage(applicationID string, params map[string]interface{}) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Post(fmt.Sprintf("im/account/merchant/%s/dv", applicationID), params, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to upload KYB document image for verification via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// ApproveBusinessApplication see https://edoc.identitymind.com/reference#feedback_1
func (i *IdentityMindAPIClient) ApproveBusinessApplication(applicationID string, params map[string]interface{}) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Post(fmt.Sprintf("im/account/merchant/%s/accepted", applicationID), params, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to acccept KYB application via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// RejectBusinessApplication see https://edoc.identitymind.com/reference#feedback_1
func (i *IdentityMindAPIClient) RejectBusinessApplication(applicationID string, params map[string]interface{}) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Post(fmt.Sprintf("im/account/merchant/%s/rejected", applicationID), params, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to reject KYB application via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}
