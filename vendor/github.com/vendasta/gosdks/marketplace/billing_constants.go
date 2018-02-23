package marketplace

// BillingFrequency indicates how frequent the billing is...
type BillingFrequency string

const (
	// Once time billing
	Once BillingFrequency = "once"

	// Daily recurrent billing
	Daily BillingFrequency = "daily"

	// Weekly recurrent billing
	Weekly BillingFrequency = "weekly"

	// Monthly recurrent billing
	Monthly BillingFrequency = "monthly"

	// Yearly recurrent billing
	Yearly BillingFrequency = "yearly"

	// Other billing
	Other BillingFrequency = "other"
)

// Currency indicates the billing curreny
type Currency string

const (
	// CAD Canadian Dollars
	CAD Currency = "CAD"

	// USD United States Dollars
	USD Currency = "USD"

	// EUR Euro Dollars
	EUR Currency = "EUR"

	// GBP Great Britain Pounds
	GBP Currency = "GBP"

	// AUD Australian Dollars
	AUD Currency = "AUD"

	// NZD New Zealand Dollars
	NZD Currency = "NZD"
)
