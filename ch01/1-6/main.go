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

// 1-6. 자유롭게 palette에 더 많은 값을 추가하고 SetColorIndex의 세 번째 인수를 변경해 여러 색상의 이미지를 생성하도록 리사주 프로그램을 수정하라

var palette = []color.Color{
	color.Black,
	color.RGBA{R: 0x00, G: 0xFF, B: 0x00, A: 0xff},  // 녹색
	color.RGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0xff},  // 빨간색
	color.RGBA{R: 0x00, G: 0x00, B: 0xFF, A: 0xff},  // 파란색
	color.RGBA{R: 0xFF, G: 0xFF, B: 0x00, A: 0xff},  // 노란색
	color.RGBA{R: 0xFF, G: 0x00, B: 0xFF, A: 0xff},  // 자홍색
	color.RGBA{R: 0x0f0, G: 0xFF, B: 0xFF, A: 0xff}, // 청록색
}

func main() {
	// os.Stdout은 표준 출력(화면)을 의미하는 io.Writer 인터페이스를 구현한 객체 !!
	// lissajous 함수는 io.Writer 인터페이스를 매개변수로 받으므로
	// os.Stdout을 전달할 수 있음
	lissajous(os.Stdout)

	// 현재 코드에서 실행 방법:
	// go run main.go > 파일명.gif
	// (> 기호는 표준 출력을 파일로 리다이렉션합니다)

	// 직접 파일에 쓰게하려면: (그냥 go run main.go 만 하면 됨)
	// file, _ := os.Create("out.gif")
	// lissajous(file)
	// file.Close()

	lissajous(os.Stdout)
}

// io.Writer는 데이터를 쓸 수 있는 모든 타입(파일, 화면, 메모리 버퍼 등)을
// 포함하는 인터페이스
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

	// 색상 수
	colorCount := len(palette) - 1

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		// 각 프레임마다 색상 변경
		colorIndex := (i % colorCount) + 1

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(colorIndex))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	// 최종 GIF를 출력 대상(out)에 인코딩
	gif.EncodeAll(out, &anim)
}
