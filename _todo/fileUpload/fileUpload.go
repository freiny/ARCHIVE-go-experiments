package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/upload", onUpload)

	targetURL := "http://localhost:3000/upload"
	filename := "./test.txt"
	postFile(filename, targetURL)
}

func onUpload(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method)

	if r.Method == "GET" {

		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)

	} else {

		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")

		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()

		fmt.Fprintf(w, "%v", handler.Header)

		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0644)

		if err != nil {
			fmt.Println(err)
			return
		}

		defer f.Close()

		io.Copy(f, file)
	}
}

func postFile(filename string, targetURL string) error {

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		fmt.Println(filename)
		return err
	}

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetURL, contentType, bodyBuf)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	fmt.Println(resp.Status)
	fmt.Println(string(respBody))

	return nil
}
