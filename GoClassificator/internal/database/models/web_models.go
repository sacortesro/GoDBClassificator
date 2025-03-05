package models

// ScanReportData contains all the data needed for the scan report template
type ScanReportData struct {
	Date             string
	DatabaseName     string
	Host             string
	ScanCount        int
	TotalTables      int
	TotalColumns     int
	DataTypesSummary []DataTypeSummary
	Tables           []TableInfo
}

// DataTypeSummary represents summary information for each data type
type DataTypeSummary struct {
	Type  string
	Count int
}

// TableInfo represents information about a scanned table
type TableInfo struct {
	Name        string
	ColumnCount int
}
