package main

import (
	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	//"image/jpeg"
	"log"
	"os"

	"github.com/oliamb/cutter"
	"github.com/urfave/cli"
	//    "github.com/oliamb/cutter"
)

func main() {
	app := cli.NewApp()
	app.Name = "L.E.R Qzone Image | Dark Magic"
	app.Usage = "Crop an image to nine parts"

	var imgurl string
	app.Action = func(c *cli.Context) error {
		fmt.Printf("Welcome to L.E.R Qzone Image Dark Magic\n ")
		imgurl = c.Args().Get(0)
		if imgurl == "" {
			fmt.Printf("Usage: DM-ImageGrid yourpicture.jpg")
			os.Exit(1)
		}
		//fmt.Println(imgurl)
		log.Print(imgurl)
		MagicCrop(imgurl)
		return nil
	}

	app.Run(os.Args)
}

func MagicCrop(imgurl string) {
	f, err := os.Open(imgurl)
	if err != nil {
		log.Fatal("Cannot open file", err)
	}
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal("Cannot decode image:", err)
	}
	//imgc, _, err := image.DecodeConfig(f)
	log.Print(img.Bounds())
	height := img.Bounds().Dy()
	width := img.Bounds().Dx()
	log.Printf("H: %d W: %d", height, width)
	//log.Printf("Image size: %d %d", imgc.Height, imgc.Width)
	imgname := imgurl[0 : len(imgurl)-4]
	err = os.MkdirAll(imgname, 0777)
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Print("Create Directory OK!")
	}
	CropAndSave(img, 0, 0, width/3, height/3, fmt.Sprintf("%s/%s%d.jpg", imgname, imgname, 1))
	CropAndSave(img, width/3, 0, width/3, height/3, fmt.Sprintf("%s/%s%d.jpg", imgname, imgname, 2))
	CropAndSave(img, width/3*2, 0, width/3, height/3, fmt.Sprintf("%s/%s%d.jpg", imgname, imgname, 3))
	CropAndSave(img, 0, width/3, width/3, height/3, fmt.Sprintf("%s/%s%d.jpg", imgname, imgname, 4))
	CropAndSave(img, width/3, width/3, width/3, height/3, fmt.Sprintf("%s/%s%d.jpg", imgname, imgname, 5))
	CropAndSave(img, width/3*2, width/3, width/3, height/3, fmt.Sprintf("%s/%s%d.jpg", imgname, imgname, 6))
	CropAndSave(img, 0, width/3*2, width/3, height/3, fmt.Sprintf("%s/%s%d.jpg", imgname, imgname, 7))
	CropAndSave(img, width/3, width/3*2, width/3, height/3, fmt.Sprintf("%s/%s%d.jpg", imgname, imgname, 8))
	CropAndSave(img, width/3*2, width/3*2, width/3, height/3, fmt.Sprintf("%s/%s%d.jpg", imgname, imgname, 9))
}

func CropAndSave(img image.Image, x int, y int, w int, h int, fn string) {
	cImg, err := cutter.Crop(img, cutter.Config{
		Height:  h,                 // height in pixel or Y ratio(see Ratio Option below)
		Width:   w,                 // width in pixel or X ratio
		Mode:    cutter.TopLeft,    // Accepted Mode: TopLeft, Centered
		Anchor:  image.Point{x, y}, // Position of the top left point
		Options: 0,                 // Accepted Option: Ratio
	})
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	out, err := os.Create(fn)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	var opt jpeg.Options
	opt.Quality = 80
	err = jpeg.Encode(out, cImg, &opt) // put quality to 80%
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	log.Printf("Generated image %s", fn)

}
