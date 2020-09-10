package storage

import (
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

func SaveFileToStorage(r *http.Request, fileNamePrefix, folderName string) (string, error) {
	r.ParseMultipartForm(1024)
	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		log.Println(`Error while parsing file`, err)
		return "", err
	}

	// Untuk dapetin Lokasi Project
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	// Generate random number
	rand.Seed(time.Now().UnixNano())
	min := 1111
	max := 9999

	// Nama File dan File Location untuk lokasi save file
	fileName := fileNamePrefix + "-" + folderName + "-" + strconv.Itoa(rand.Intn(max-min+1)+min) + filepath.Ext(handler.Filename)
	fileLocation := filepath.Join(dir, "fileServer", folderName, fileName)

	// Mempersiapkan temporary file, kalo error, bikin manual folder api_gateway/files/user nya
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(`Error while creating temporary file`, err)
		log.Println("Silahkan bikin folder manual dengan directory api_gateway/fileServer/" + folderName)
		return "", err
	}
	defer targetFile.Close()

	// Mengcopy uploaded File ke temporary file
	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		log.Println(`Error while coping file to Local Storage`, err)
		return "", err
	}

	return fileName, nil
}
