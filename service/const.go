package service

// EngineType is the type of engine
type EngineType string

// WellbeingHistoricalType includes sleep, localization
type WellbeingHistoricalType string

const (
	SoftSecurity EngineType = "softSecurity"
	Location     EngineType = "location"
	Wellbeing    EngineType = "wellBeing"
)

const (
	WellbeingSleep        WellbeingHistoricalType = "sleep"
	WellbeingLocalization WellbeingHistoricalType = "localization"
	WellbeingDeletion     WellbeingHistoricalType = "deletion"
)
