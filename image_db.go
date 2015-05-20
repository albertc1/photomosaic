package main

import (
    "encoding/json"
    "image"
    "log"
    "os"
)

type ImageRecord struct {
    R int
    G int
    B int
    Path string
}

type ImageDb struct {
    images []ImageRecord
}

func (ip *ImageDb) Add(path string, r int, g int, b int) {
    ip.images = append(ip.images, ImageRecord{r, g, b, path})
}

func (ip *ImageDb) Find(r int, g int, b int) string {
    var bestMatch ImageRecord
    var bestScore int = 255 * 255 * 3 + 1  //bigger than max calculatable score (lower is better)

    for _, img := range ip.images {
        score := SumSquares(r, g, b, img.R, img.G, img.B)
        if score < bestScore {
            bestScore = score
            bestMatch = img
        }
    }

    return bestMatch.Path
}

func SumSquares(r1 int, g1 int, b1 int, r2 int, g2 int, b2 int) int {
    rdelta := r1 - r2
    gdelta := g1 - g2
    bdelta := b1 - b2
    return rdelta * rdelta + gdelta * gdelta + bdelta * bdelta
}

const INDEX_FILE string = TILE_FILE_PATH + "index.json"

//Loads index of images from a json file
func (ip *ImageDb) loadIndex() bool {
    indexFile, err := os.Open(INDEX_FILE)
    if err != nil {
        log.Print("Failed to load image db index from file.", err)
        return false
    }
    defer indexFile.Close()
    jsonParser := json.NewDecoder(indexFile)
    if err = jsonParser.Decode(&ip.images); err != nil {
        log.Print("Failed to parse image db index from file.", err)
        return false
    }
    return len(ip.images) > 0
}

//Writes current image index to a json file
func (ip *ImageDb) writeIndex() {
    indexFile, err := os.Create(INDEX_FILE)
    if err != nil {
        panic(err)
    }
    defer indexFile.Close()

    encoder := json.NewEncoder(indexFile)
    if err = encoder.Encode(ip.images); err != nil {
        panic(err)
    }
}

//Processes image files from a directory and indexes them
func (ip *ImageDb) reindex() {
    filenames := ListSavedImages()
    indexed := 0
    for _, filename := range filenames {
        img := ReadImageFromFile(filename)
        if img == nil {
            continue
        }
        bounds := img.Bounds()
        r, g, b := Avg_color(img, image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y))
        ip.Add(filename, r, g, b)
        indexed++
        log.Printf("Indexed %d images", indexed)
    }
}

//Create instance of ImageDb populated with indexed images from a directory
func InitDb() *ImageDb {
    imgDb := ImageDb{}
    if !imgDb.loadIndex() {
        imgDb.reindex()
        imgDb.writeIndex()
    }
    return &imgDb
}
