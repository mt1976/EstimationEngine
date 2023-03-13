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
	Migrate  string `csv:"migrate"`
}

type OriginStateIO struct {
	OriginStateID string `csv:"originStateID"`
	Code          string `csv:"code"`
	Name          string `csv:"name"`
	IsLocked      string `csv:"isLocked"`
	Notify        string `csv:"notify"`
	Migrate       string `csv:"migrate"`
}

type DocumentTypeIO struct {
	DocTypeID string `csv:"docTypeID"`
	Code      string `csv:"code"`
	Comments  string `csv:"comments"`
	Name      string `csv:"name"`
	Migrate   string `csv:"migrate"`
}

type ProjectStateIO struct {
	ProjectStateID string `csv:"projectStateID"`
	Code           string `csv:"code"`
	Name           string `csv:"name"`
	IsLocked       string `csv:"isLocked"`
	Notify         string `csv:"notify"`
	Migrate        string `csv:"migrate"`
}

type EstimationStateIO struct {
	EstimationStateID string `csv:"estimationStateID"`
	Code              string `csv:"code"`
	Name              string `csv:"name"`
	IsLocked          string `csv:"isLocked"`
	Notify            string `csv:"notify"`
	Migrate           string `csv:"migrate"`
}

type TranslationIO struct {
	TranslationID string `csv:"translationID"`
	Class         string `csv:"class"`
	Message       string `csv:"message"`
	Translation   string `csv:"translation"`
	Migrate       string `csv:"migrate"`
}
