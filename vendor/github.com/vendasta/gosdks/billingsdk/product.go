package billingsdk

// BillingFrequency is a string enum for the product
type BillingFrequency int64

// Enums for BillingFrequency
const (
	Monthly BillingFrequency = iota
	Yearly
	OneTime
)

func (b BillingFrequency) String() string {
	switch b {
	case Monthly:
		return "Monthly"
	case Yearly:
		return "Yearly"
	case OneTime:
		return "One Time"
	default:
		return ""
	}
}

// Product args used for Product Create and Update
type Product struct {
	ProductName      string           `json:"product_name"`
	ProductID        string           `json:"product_id"`
	Icon             string           `json:"icon"`
	BillingFrequency BillingFrequency `json:"billing_frequency"`
	Price            int64            `json:"price"`
}
