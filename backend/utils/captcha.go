package utils

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"

	svg "github.com/ajstarks/svgo"
)

func GenerateSVG(width, height int) ([]byte, string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var svgContent bytes.Buffer
	canvas := svg.New(&svgContent)
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height, "fill:white")

	code := fmt.Sprintf("%04d", r.Intn(10000))
	canvas.Text(width/2, height/2, code, "text-anchor:middle; dominant-baseline:middle; font-size:30px; fill:black;")

	canvas.End()

	return svgContent.Bytes(), code
}
