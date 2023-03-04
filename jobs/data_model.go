package jobs

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/data.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Data (data)
// Endpoint 	        : Data (DataID)
// For Project          : github.com/mt1976/ebEstimates/
// ----------------------------------------------------------------
// Template Generator   : Einsteinium [r5-23.01.23]
// Date & Time		    : 04/03/2023 at 12:47:42
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// Data defines the datamodel for the Data object
type DataItemIO struct {
	// Dynamically generated 04/03/2023 by matttownsend (Matt Townsend) on silicon.local
	//
	// Field Definitions
	//
	DataID   string `csv:"dataID"`
	Class    string `csv:"class"`
	Field    string `csv:"field"`
	Value    string `csv:"value"`
	Category string `csv:"category"`
}
