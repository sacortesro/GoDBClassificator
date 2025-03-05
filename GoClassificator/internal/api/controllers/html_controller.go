package controllers

import (
	"GoClassificator/internal/database/repository/webrepository"
	"GoClassificator/internal/logger"
	"bytes"
	"html/template"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

/****************************************************
 * 													*
 * Controller that handles the html page 			*
 * 													*
 ****************************************************/

// RenderScanReport renders the scan report template
func RenderScanReport(c echo.Context) error {

	logger.Info("Rendering scan report")

	id := c.Param("id")

	// Convert id to uint
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	data, err := webrepository.GetScanReportData(uint(uintID))
	if err != nil {
		logger.Error("Failed to get scan report data", err)
		return c.String(http.StatusInternalServerError, "Failed to get scan report data")
	}

	tmpl, err := template.ParseFiles("web/templates/scan_report.html")
	if err != nil {
		logger.Error("Failed to parse template", err)
		return c.String(http.StatusInternalServerError, "Failed to parse template")
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return err
	}

	logger.Info("Scan report rendered successfully")

	return c.HTMLBlob(http.StatusOK, buf.Bytes())
}
