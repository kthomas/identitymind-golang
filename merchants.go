package identitymind

import "fmt"

// Merchant aggregation

// CreateMerchant creates a merchant account
func (i *IdentityMindAPIClient) CreateMerchant(params map[string]interface{}) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Post("im/admin/jax/merchant", params, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to create merchant account via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// GetMerchant creates a merchant account
func (i *IdentityMindAPIClient) GetMerchant(merchantID string, params map[string]interface{}) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Get(fmt.Sprintf("im/admin/jax/merchant/%s", merchantID), params, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch merchant account via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// UpdateMerchant creates a merchant account
func (i *IdentityMindAPIClient) UpdateMerchant(merchantID string, params map[string]interface{}) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Post(fmt.Sprintf("im/admin/jax/merchant/%s", merchantID), params, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to update merchant account via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// Merchant KYB

// RejectMerchantBusinessApplication see https://edoc.identitymind.com/reference#feedback_1
func (i *IdentityMindAPIClient) RejectMerchantBusinessApplication(applicationID string, params map[string]interface{}) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Post(fmt.Sprintf("im/account/merchant/%s/rejected", applicationID), params, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to reject merchant KYB application via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// GetMerchantBusinessApplication see https://edoc.identitymind.com/reference#getmerchantkyc
func (i *IdentityMindAPIClient) GetMerchantBusinessApplication(applicationID string) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Get(fmt.Sprintf("im/account/merchant/%s", applicationID), map[string]interface{}{}, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve merchant KYB application via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// ReevaluateMerchantBusinessApplication see https://edoc.identitymind.com/reference#reevaluatemerchant
func (i *IdentityMindAPIClient) ReevaluateMerchantBusinessApplication(applicationID string) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Post(fmt.Sprintf("im/account/merchant/%s", applicationID), map[string]interface{}{}, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to reevaluate merchant KYB application via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// SubmitMerchantBusinessApplication see https://edoc.identitymind.com/reference#merchant
func (i *IdentityMindAPIClient) SubmitMerchantBusinessApplication(merchantID string, params map[string]interface{}) (interface{}, error) {
	var resp map[string]interface{}
	params["m"] = merchantID
	status, err := i.Post(fmt.Sprintf("im/account/merchant?graphScoreResponse=false"), params, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to reevaluate merchant KYB application via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// ListMerchantBusinessDocuments see https://edoc.identitymind.com/reference#getfilelistforapplicationformerchant
func (i *IdentityMindAPIClient) ListMerchantBusinessDocuments(applicationID string) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Get(fmt.Sprintf("im/account/merchant/%s/files", applicationID), map[string]interface{}{}, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to list merchant KYB documents via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// DownloadMerchantBusinessDocument see https://edoc.identitymind.com/reference#reevaluatemerchant
func (i *IdentityMindAPIClient) DownloadMerchantBusinessDocument(applicationID, documentID string) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Get(fmt.Sprintf("im/account/merchant/%s/files/%s", applicationID, documentID), map[string]interface{}{}, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to download merchant KYB document via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// UploadMerchantBusinessDocument see https://edoc.identitymind.com/reference#processfileuploadrequestformerchantkyc
func (i *IdentityMindAPIClient) UploadMerchantBusinessDocument(applicationID string, params map[string]interface{}) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Post(fmt.Sprintf("im/account/merchant/%s/files", applicationID), params, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to upload merchant KYB document via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// UploadMerchantBusinessDocumentVerificationImage see https://edoc.identitymind.com/reference#processfileuploadrequestformerchantkyc
func (i *IdentityMindAPIClient) UploadMerchantBusinessDocumentVerificationImage(applicationID string, params map[string]interface{}) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Post(fmt.Sprintf("im/account/merchant/%s/dv", applicationID), params, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to upload KYB document image for verification via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// ApproveMerchantBusinessApplication see https://edoc.identitymind.com/reference#feedback_1
func (i *IdentityMindAPIClient) ApproveMerchantBusinessApplication(applicationID string, params map[string]interface{}) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Post(fmt.Sprintf("im/account/merchant/%s/accepted", applicationID), params, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to acccept merchant KYB application via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// EvaluateMerchantFraud evaluates a transaction for payment fraud on behalf of a given merchant; see https://edoc.identitymind.com/reference#anti-fraud-1
func (i *IdentityMindAPIClient) EvaluateMerchantFraud(merchantID, params map[string]interface{}) (interface{}, error) {
	var resp map[string]interface{}
	params["m"] = merchantID
	status, err := i.Post(fmt.Sprintf("im/transaction?graphScoreResponse=false"), params, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to evaluate tx for payment fraud via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// ReportMerchantTransaction reports various kinds of transactions including deposits, withdrawals and internal transfer
func (i *IdentityMindAPIClient) ReportMerchantTransaction(merchantID, txType string, params map[string]interface{}) (interface{}, error) {
	if txType != IdentityMindTxTypeDeposit && txType != IdentityMindTxTypeWithdrawal && txType != IdentityMindTxTypeTransfer {
		return nil, fmt.Errorf("Invalid tx type provided: %s", txType)
	}
	params["m"] = merchantID
	var resp map[string]interface{}
	status, err := i.Post(fmt.Sprintf("im/account/%s?graphScoreResponse=false", txType), params, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to report tx via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}
