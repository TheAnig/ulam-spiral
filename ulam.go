// ulam.go
// Usage ./ulam -r "resolution scale" -o "output file name"

package main

import(
    "fmt"
    "flag"
    "image"
    "image/color"
    "image/png"
    "image/draw"
    "os"
    "log"
)

func isPrime(n int) bool {
    
    if n == 2 {
        return true
    }
    if n == 3 {
        return true
    }
    if n%2 == 0 {
        return false
    }
    if n%3 == 0 {
        return false
    }
    
    i := 5
    w := 2
    
    for i*i <= n {
        if n%i == 0 {
            return false
        }
        i += w
        w = 6 - w
    }
    
    return true
    
}

func main() {
    
    // Initialize flags for command line
    out := flag.String("o", "out", "Output file name (without png extention)")
    res := flag.Int("r", 800, "Resolution (width), 800 gives an image of 800*800")
    flag.Parse()
    width := *res
    
    // Array to store the list of primes
    a := make([]bool, width*width)
    
    for c:= range a {
        a[c] = isPrime(c)
    }
    
    // Open up an empty image
    img := image.NewRGBA(image.Rect(0,0,width,width))
    
    // Draw the white background
    draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{255, 255, 255, 255}}, image.ZP, draw.Src)
    
    // Set the color to draw the dots for primes
    col := color.RGBA{0, 0, 0, 255}
    
    // Begin generating spiral
    var w int = 0
    var h int = 0
    var dh int = -1
    var dw int = 0
    
    for k := range a {
        
        if (-width/2 <= w) && (w <= width/2) && (-width/2 <= h) && (h <= width/2) {
            if(a[k]){
                img.Set(w+(width/2), h+(width/2), col)
            }
        }
        
        if ( w == h) || ( w<0 && w == -h ) || (w>0 && w == 1-h) {
            dw, dh = -dh, dw
        }
        w += dw
        h += dh
    }
    
    outf, err := os.Create(fmt.Sprintf(*out+".png"))
    if err != nil {
        log.Fatalln(err)
    }
    defer outf.Close()
    png.Encode(outf, img)
}
