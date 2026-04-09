package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func UploadRoomImage(w http.ResponseWriter, r *http.Request) {
	// Limit upload size (10 MB)
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "File too large or invalid", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Missing image file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Validate extension
	ext := filepath.Ext(header.Filename)
	allowed := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true}
	if !allowed[ext] {
		http.Error(w, "Invalid file type", http.StatusBadRequest)
		return
	}

	// Create unique filename
	timestamp := time.Now().UnixNano()
	newFileName := fmt.Sprintf("room_%d%s", timestamp, ext)
	savePath := filepath.Join("./uploads/rooms", newFileName)

	// Save file
	dst, err := os.Create(savePath)
	if err != nil {
		http.Error(w, "Cannot save file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	// Return the public URL
	fileURL := fmt.Sprintf("/uploads/rooms/%s", newFileName)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"path":"%s"}`, fileURL)
}
