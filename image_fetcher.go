package main

import (
    "bufio"
    "fmt"
    "log"
    "image"
    // "image/draw"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "strconv"
)

const IMG_LIST string = "/Users/albertc/tmp/thumbnail_urls2.csv"
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
        filename := IMAGE_FILE_PATH + strconv.Itoa(imgCount) + ".jpg"
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
    fileInfos, err := ioutil.ReadDir(IMAGE_FILE_PATH)
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

func InitDb() {
    filenames := ListSavedImages()
    for _, filename := range filenames {
        //get image from file
        //decode image
        //scale image
        //encode image
        //store image into new file in new directory
        //index scaled image
    }
}

func Draw(dst image.Image, dstRect image.Rectangle, srcPath string) {
    // src := ReadImageFromPath
    // draw.Draw(dst, dstRect, src, image.Image.ZP, draw.Src)
}