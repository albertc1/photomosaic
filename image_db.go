package main

type ImageRecord struct {
    r int
    g int
    b int
    path string
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
        score := SumSquares(r, g, b, img.r, img.g, img.b)
        if score < bestScore {
            bestScore = score
            bestMatch = img
        }
    }

    return bestMatch.path
}

func SumSquares(r1 int, g1 int, b1 int, r2 int, g2 int, b2 int) int {
    rdelta := r1 - r2
    gdelta := g1 - g2
    bdelta := b1 - b2
    return rdelta * rdelta + gdelta * gdelta + bdelta * bdelta
}