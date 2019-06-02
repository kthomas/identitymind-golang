package identitymind

import "fmt"

// GetCase see https://edoc.identitymind.com/reference#update
func (i *IdentityMindAPIClient) GetCase(caseID string) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Get(fmt.Sprintf("im/admin/jax/case/%s", caseID), map[string]interface{}{}, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to create case via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// CreateCase see https://edoc.identitymind.com/reference#createcase
func (i *IdentityMindAPIClient) CreateCase(params map[string]interface{}) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Post("im/admin/jax/case", params, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to create case via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// CloseCase see https://edoc.identitymind.com/reference#closecase
func (i *IdentityMindAPIClient) CloseCase(caseID string, params map[string]interface{}) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Post("im/admin/jax/case/close", params, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to close case via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// UpdateCase see https://edoc.identitymind.com/reference#updatecasecontent
func (i *IdentityMindAPIClient) UpdateCase(caseID string, params map[string]interface{}) (interface{}, error) {
	var resp map[string]interface{}
	status, err := i.Post(fmt.Sprintf("im/admin/jax/case/%s", caseID), params, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to update case via identitymind API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}
