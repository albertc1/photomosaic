package main

import (
    "bufio"
    "fmt"
    "log"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "strconv"
    "image"
)

const IMG_LIST string = "/Users/albertc/tmp/thumbnail_urls2.csv"
// const TILE_FILE_PATH string = "/Users/albertc/tmp/tiles/"
const TILE_FILE_PATH string = "/Users/udaysaraf/gocode/src/github.com/udaysaraf/photomosaic/tiles/"
const IMAGE_FILE_PATH string = "/Users/albertc/tmp/images/"

// FetchFromUrlFile reads a text file containing one img URL per line, fetches each URL, and
// writes the fetched jpeg image to a file.
func FetchFromUrlFile() int {
    file, err := os.Open(IMG_LIST)

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    defer file.Close()

    reader := bufio.NewReader(file)
    scanner := bufio.NewScanner(reader)

    imgCount := 50000
    fetched := 0
    for scanner.Scan() {
        url := scanner.Text()
        //TODO: don't hardcode filenames and file formats
        filename := TILE_FILE_PATH + strconv.Itoa(imgCount) + ".jpg"
        if FetchImageFromUrl(url, filename) {
            fetched++
        }
        imgCount++
    }
    log.Printf("Fetched %d images!", fetched)
    return fetched
}

//FetchImageFromUrl fetches an image from a URL and write it to the specified filename
func FetchImageFromUrl(url string, filename string) bool {
    file, err := os.Create(filename)

    if err != nil {
        fmt.Println(err)
        panic(err)
    }
    defer file.Close()

    client := http.Client{}

    resp, err := client.Get(url)

    if err != nil {
        log.Printf("Failed to fetch from %q: %q", url, err)
        return false
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        log.Printf("Failed to fetch from %q: %q", url, resp.Status)
        return false
    }

    size, err := io.Copy(file, resp.Body)

    if err != nil {
        panic(err)
    }
    log.Printf("Fetched image %q from %q size %d", filename, url, size)
    return true
}

// ListSavedImages returns a list of all files in the image directory
func ListSavedImages() []string {
    fileInfos, err := ioutil.ReadDir(TILE_FILE_PATH)
    if err != nil {
        panic(err)
    }
    filenames := []string{}
    for _, fileInfo := range fileInfos {
        if !fileInfo.IsDir() {
            filenames = append(filenames, fileInfo.Name())
        }
    }
    return filenames
}

// ReadImageFromFile reads an image file from disk and returns an Image object
func ReadImageFromFile(filename string) image.Image {
    filepath := TILE_FILE_PATH+filename
    reader, err := os.Open(filepath)
    if err != nil {
        panic(err)
    }
    // lots of files have errors for some reason, so instead of killing the program, just ignore those files
    img, _, err := image.Decode(reader)
    if err != nil {
        fmt.Println(filename)
        return nil
    }
    return img
}
