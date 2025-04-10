package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

// Lissajous는 임의의 리사주 형태의 애니메이션 GIF를 생성한다

// 1-5. 가독성을 높이기 위해 검은색 위에 녹색을 칠하도록 리사주 프로그램의 색상 팔레트를 변경하라
// 웹 색상 #RRGGBB를 만들려면 color.RGBA(0xRR, 0xGG, 0xBB, 0xff)를 사용하며,
// 각 16진 숫자의 쌍은 픽셀에서 적색, 녹색, 청색의 세기를 나타낸다.

var pallette = []color.Color{color.Black, color.RGBA{R: 0x00, G: 0xFF, B: 0x00, A: 0xff}}

const (
	blackIndex = 0 // 팔레트의 첫 번째 색상
	greenIndex = 1 // 팔레트의 두 번째 색상

)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // x 진동자의 회전수
		res     = 0.001 // 회전각
		size    = 100   // 이미지 캔버스 크기 [-size..+size]
		nframes = 64    // 애니메이션 프레임 수
		delay   = 8     // 10ms 단위의 프레임 간 지연
	)
	freq := rand.Float64() * 3.0 // y 진동자의 상대적 진동수
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 위상 차이
	for range nframes {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
