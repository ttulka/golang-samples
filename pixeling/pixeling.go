package main

import (
    "fmt"
	"image"
_	"image/png"
_	"image/gif"
_	"image/jpeg"
	"os"
    "strconv"
)

func main() {
    args := os.Args[1:]
    if len(args) < 1 {
        fmt.Fprintf(os.Stderr, "Usage: <image.png>\n")
        os.Exit(1)
    }
    file, err := os.Open(args[0])
	if err != nil {
        fmt.Fprintf(os.Stderr, "Error by reading the file: %v\n", err)
        os.Exit(1)
	}
	defer file.Close()

    img, _, err := image.Decode(file)
	if err != nil {
        fmt.Fprintf(os.Stderr, "Error by decoding the image: %v\n", err)
        os.Exit(1)
	}

    writeAsSvg(img, args[0] + ".svg")
}

func writeAsSvg(img image.Image, filename string) {
    svg, err := os.Create(filename)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error by opening the file: %v", err)
        os.Exit(1)
    }
    defer svg.Close()    

    bounds := img.Bounds()

    write(svg, "<svg version=\"1.1\" viewBox=\"0.0 0.0 " + strconv.Itoa(bounds.Max.X) + " " + strconv.Itoa(bounds.Max.Y) + "\" fill=\"none\" stroke=\"none\" stroke-linecap=\"square\" stroke-miterlimit=\"10\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xmlns=\"http://www.w3.org/2000/svg\">")

    for x := bounds.Min.X; x < bounds.Max.X; x++ {
        for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
            r, g, b, a := img.At(x, y).RGBA()
            write(svg, "<rect x=\"" + strconv.Itoa(x) + "\" y=\"" + strconv.Itoa(y) + "\" width=\"1\" height=\"1\" fill=\"" + rgbString(r, b, g) + "\" fill-opacity=\"" + strconv.Itoa(int(float32(a / 0x101) / 255)) + "\" />")
        }
    }
    
    write(svg, "</svg>")
}

func rgbString(r uint32, b uint32, g uint32) (string) {
    return "rgb(" + to255(r) + "," + to255(g) + "," + to255(b) + ")"
}

func to255(x uint32) (string) {
    return strconv.Itoa(int(x / 0x101))
}

func write(f *os.File, data string) {
    _, err := f.WriteString(data)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error by wrinting into the file: %v", err)
        os.Exit(1)
    }
}
