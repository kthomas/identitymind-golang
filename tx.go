package identitymind

import "fmt"

// IdentityMindTxTypeDeposit maps to 'transferin' URI; see https://edoc.identitymind.com/reference#transferin
const IdentityMindTxTypeDeposit = "transferin"

// IdentityMindTxTypeWithdrawal maps to 'transferout' URI; see https://edoc.identitymind.com/reference#transferout
const IdentityMindTxTypeWithdrawal = "transferout"

// IdentityMindTxTypeTransfer maps to 'transfer' URI; see https://edoc.identitymind.com/reference#transfer
const IdentityMindTxTypeTransfer = "transfer"

// EvaluateFraud evaluates a transaction for payment fraud; see https://edoc.identitymind.com/reference#anti-fraud-1
func (i *IdentityMindAPIClient) EvaluateFraud(params map[string]interface{}) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Post(fmt.Sprintf("im/transaction?graphScoreResponse=false"), params, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to evaluate tx for payment fraud via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// ReportFraud reports a fraud event; see https://edoc.identitymind.com/reference#event
func (i *IdentityMindAPIClient) ReportFraud(params map[string]interface{}) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Post("im/admin/jax/feg", params, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to report fraud event via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// ReportTransaction reports various kinds of transactions including deposits, withdrawals and internal transfer
func (i *IdentityMindAPIClient) ReportTransaction(txType string, params map[string]interface{}) (interface{}, error) {
	if txType != IdentityMindTxTypeDeposit && txType != IdentityMindTxTypeWithdrawal && txType != IdentityMindTxTypeTransfer {
		return nil, fmt.Errorf("Invalid tx type provided: %s", txType)
	}
	var resp map[string]interface{}
	status, err := i.Post(fmt.Sprintf("im/account/%s?graphScoreResponse=false", txType), params, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to report tx via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}
