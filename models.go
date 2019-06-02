package identitymind

// KYCApplication represents a identitymind KYC application
type KYCApplication struct {
	EDNAScorecard map[string]interface{} `json:"ednaScoreCard"`
	MTID          *string                `json:"mtid"`
	TID           *string                `json:"tid"`
	RCD           *string                `json:"rcd"`
	State         *string                `json:"state"`
}

// IsAccepted returns true if the KYC application has been accepted
func (k *KYCApplication) IsAccepted() bool {
	return k.State != nil && *k.State == "A"
}

// IsRejected returns true if the KYC application has been rejected
func (k *KYCApplication) IsRejected() bool {
	return k.State != nil && *k.State == "D"
}

// IsUnderReview returns true if the KYC application is currently pending review
func (k *KYCApplication) IsUnderReview() bool {
	return k.State != nil && *k.State == "R"
}
