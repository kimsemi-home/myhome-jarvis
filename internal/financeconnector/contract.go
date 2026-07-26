package financeconnector

var CanonicalFieldContract = []ParityField{
	{Name: "transaction_id", SourceField: "transaction_id", CanonicalPath: "transaction_id", Required: true},
	{Name: "source", SourceField: "source", CanonicalPath: "source", Required: true},
	{Name: "owner", SourceField: "owner", CanonicalPath: "owner", Required: true},
	{Name: "occurred_at", SourceField: "occurred_at", CanonicalPath: "occurred_at", Required: true},
	{Name: "posted_at", SourceField: "posted_at", CanonicalPath: "posted_at", Required: false},
	{Name: "amount", SourceField: "amount", CanonicalPath: "amount.minor_units", Required: true},
	{Name: "currency", SourceField: "currency", CanonicalPath: "amount.currency", Required: true},
	{Name: "direction", SourceField: "direction", CanonicalPath: "direction", Required: true},
	{Name: "merchant_name", SourceField: "merchant_name", CanonicalPath: "merchant_name", Required: false},
	{Name: "category", SourceField: "category", CanonicalPath: "category", Required: false},
	{Name: "account_id", SourceField: "account_id", CanonicalPath: "account_id", Required: false},
	{Name: "card_id", SourceField: "card_id", CanonicalPath: "card_id", Required: false},
	{Name: "raw_ref", SourceField: "raw_ref", CanonicalPath: "raw_ref", Required: true},
	{Name: "tags", SourceField: "tags", CanonicalPath: "tags", Required: false},
}

type ParityReport struct {
	Provider                    string        `json:"provider"`
	FixturePath                 string        `json:"fixture_path"`
	Records                     int           `json:"records"`
	RecordsWithRequiredFields   int           `json:"records_with_required_fields"`
	MappedFields                int           `json:"mapped_fields"`
	ExpectedFields              int           `json:"expected_fields"`
	ParityPercent               float64       `json:"parity_percent"`
	ThresholdPercent            float64       `json:"threshold_percent"`
	Pass                        bool          `json:"pass"`
	MissingRequiredFields       []string      `json:"missing_required_fields"`
	UnsupportedProviderFeatures []string      `json:"unsupported_provider_features"`
	FieldContract               []ParityField `json:"field_contract"`
}
