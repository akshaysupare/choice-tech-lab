package api

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"choice-tech-project/internal/excel"
	"choice-tech-project/internal/model"
	"choice-tech-project/internal/service"

	"github.com/gin-gonic/gin"
)

// Handler wraps the service layer for API endpoints.
type Handler struct {
	Service *service.Service
}

// NewHandler creates a new Handler instance.
func NewHandler(svc *service.Service) *Handler {
	return &Handler{Service: svc}
}

// ImportExcel handles the Excel file upload, validates header, and imports data.
func (h *Handler) ImportExcel(c *gin.Context) {
    // 1. Retrieve the uploaded file from the request (expects form-data key "file")
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
        return
    }

    // 2. Save the uploaded file to a temporary location on the server
    tmpPath := filepath.Join(os.TempDir(), file.Filename)
    if err := c.SaveUploadedFile(file, tmpPath); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
        return
    }
    // Ensure the temporary file is deleted after processing
    defer os.Remove(tmpPath)

    // 3. Parse the Excel file and validate only the header row
    records, err := excel.ParseExcelFile(tmpPath)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 4. Save the parsed records into the database and cache (synchronously)
	err = h.Service.SaveRecords(context.Background(), records)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "import failed: " + err.Error()})
        return
    }

    // 5. Respond to the client after successful import
    c.JSON(http.StatusOK, gin.H{"message": "import completed"})
}

// GetRecords returns all imported records, using Redis cache if available.
func (h *Handler) GetRecords(c *gin.Context) {
	records, err := h.Service.GetRecords(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, records)
}

// UpdateRecord updates a specific record by ID.
func (h *Handler) UpdateRecord(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var rec model.Record
	if err := c.ShouldBindJSON(&rec); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}
	rec.ID = id
	if err := h.Service.UpdateRecord(c.Request.Context(), rec); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "record updated"})
}

// DeleteRecord deletes a specific record by ID.
func (h *Handler) DeleteRecord(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.Service.DeleteRecord(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "record deleted"})
}

// GetRecordByID returns a single record by its ID.
func (h *Handler) GetRecordByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	rec, err := h.Service.GetRecordByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}
	c.JSON(http.StatusOK, rec)
}
