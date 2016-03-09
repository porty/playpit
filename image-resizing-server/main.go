package main

import (
	"bytes"
	"log"
	"net/http"
	"strconv"

	"github.com/golang/groupcache"
	"github.com/nfnt/resize"

	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
)

type httpGetter struct {
}

func (g httpGetter) Get(ctx groupcache.Context, key string, dest groupcache.Sink) error {
	resp, err := http.Get(key)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	dest.SetBytes(bytes)
	log.Println("Downloaded " + key)
	return nil
}

func main() {

	cache := groupcache.NewGroup("source-files", 20*1024*1024, httpGetter{})

	http.HandleFunc("/from-net", func(w http.ResponseWriter, r *http.Request) {
		sourceURL := r.FormValue("url")
		percentStr := r.FormValue("percent")

		if sourceURL == "" {
			http.Error(w, "Missing URL", 400)
			return
		}

		percent, err := strconv.Atoi(percentStr)
		if err != nil || percent <= 0 || percent >= 1000 {
			http.Error(w, "Bad percentage specified", 400)
			return
		}

		log.Printf("Client asked for %d%% image of %s", percent, sourceURL)
		resp, err := http.Get(sourceURL)
		if err != nil {
			message := "Failed to get image: " + err.Error()
			log.Println(message)
			http.Error(w, message, 500)
			return
		}
		defer resp.Body.Close()

		img, imgType, err := image.Decode(resp.Body)
		if err != nil {
			message := "Failed to decode image: " + err.Error()
			log.Println(message)
			http.Error(w, message, 500)
			return
		}

		origSize := img.Bounds()
		fractional := float64(percent) / 100
		newWidth := uint(float64(origSize.Max.X-origSize.Min.X) * fractional)
		newHeight := uint(float64(origSize.Max.Y-origSize.Min.Y) * fractional)

		newImage := resize.Resize(newWidth, newHeight, img, resize.NearestNeighbor)
		if imgType == "jpeg" {
			w.Header().Set("Content-type", "image/jpeg")
			jpeg.Encode(w, newImage, nil)
			log.Println("Sent back as JPEG")
		} else if imgType == "png" {
			w.Header().Set("Content-type", "image/png")
			png.Encode(w, newImage)
			log.Println("Sent back as PNG")
		}
	})

	http.HandleFunc("/from-net-with-cache", func(w http.ResponseWriter, r *http.Request) {
		sourceURL := r.FormValue("url")
		percentStr := r.FormValue("percent")

		if sourceURL == "" {
			http.Error(w, "Missing URL", 400)
			return
		}

		percent, err := strconv.Atoi(percentStr)
		if err != nil || percent <= 0 || percent >= 1000 {
			http.Error(w, "Bad percentage specified", 400)
			return
		}

		log.Printf("Client asked for %d%% image of %s", percent, sourceURL)

		var b []byte
		sink := groupcache.AllocatingByteSliceSink(&b)
		err = cache.Get(nil, sourceURL, sink)
		if err != nil {
			message := "Failed to get image: " + err.Error()
			log.Println(message)
			http.Error(w, message, 500)
			return
		}

		reader := bytes.NewReader(b)

		img, imgType, err := image.Decode(reader)
		if err != nil {
			message := "Failed to decode image: " + err.Error()
			log.Println(message)
			http.Error(w, message, 500)
			return
		}

		origSize := img.Bounds()
		fractional := float64(percent) / 100
		newWidth := uint(float64(origSize.Max.X-origSize.Min.X) * fractional)
		newHeight := uint(float64(origSize.Max.Y-origSize.Min.Y) * fractional)

		newImage := resize.Resize(newWidth, newHeight, img, resize.NearestNeighbor)
		if imgType == "jpeg" {
			w.Header().Set("Content-type", "image/jpeg")
			jpeg.Encode(w, newImage, nil)
			log.Println("Sent back as JPEG")
		} else if imgType == "png" {
			w.Header().Set("Content-type", "image/png")
			png.Encode(w, newImage)
			log.Println("Sent back as PNG")
		}
	})

	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
