package models

// DB structures for the scan result
type ScanHistory struct {
	ID         uint   `gorm:"column:id;primaryKey;autoincrement"`
	DatabaseID uint   `gorm:"column:database_id"`
	ScanStatus string `gorm:"column:scan_status"`
}

type ScannedTable struct {
	ID        uint   `gorm:"column:id;primaryKey;autoincrement"`
	ScanID    uint   `gorm:"column:scan_id"`
	TableName string `gorm:"column:table_name"`
}

type ScanResult struct {
	ID              uint   `gorm:"column:id;primaryKey;autoincrement"`
	TableID         uint   `gorm:"column:table_id"`
	ColumnName      string `gorm:"column:column_name"`
	InformationType string `gorm:"column:information_type"`
}

// Structure to represent a MySQL connection
type DatabaseConnection struct {
	ID       uint   `gorm:"column:id;primaryKey;autoincrement"`
	Host     string `gorm:"column:host"`
	Port     int    `gorm:"column:port"`
	Username string `gorm:"column:dbusername"`
	Password string `gorm:"column:dbpassword"`
	DbName   string `gorm:"column:dbname"`
}
