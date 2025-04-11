package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand/v2"
	"net/http"
	"os"
	"strconv"
)

// 1-12 URL에서 파라미터 값을 읽도록 리사주 서버를 수정하라.
// 예를 들어 http://localhost:8000/?cycles=20 과 같은
// URL이 반복횐수로 기본 값 5 대신 20을 지정하게 할 수 있을 것이다.
// strconv.Atoi 함수를 사용해 문자열 파라미터를 정수로 변환하라.
// go doc strconv.Atoi로 관련 문서를 볼 수 있다.

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	params := make(map[string]int, 4)
	for k, v := range r.Form {
		if value, err := strconv.Atoi(v[0]); err != nil {
			fmt.Fprint(os.Stderr, "invalid parameter")
			fmt.Fprint(w, "invalid parameter")
		} else {
			params[k] = value
		}
	}

	lissajous(w, params)

}

func lissajous(out io.Writer, params map[string]int) {
	var palette = []color.Color{
		color.Black,
		color.RGBA{R: 0x00, G: 0xFF, B: 0x00, A: 0xff},  // 녹색
		color.RGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0xff},  // 빨간색
		color.RGBA{R: 0x00, G: 0x00, B: 0xFF, A: 0xff},  // 파란색
		color.RGBA{R: 0xFF, G: 0xFF, B: 0x00, A: 0xff},  // 노란색
		color.RGBA{R: 0xFF, G: 0x00, B: 0xFF, A: 0xff},  // 자홍색
		color.RGBA{R: 0x0f0, G: 0xFF, B: 0xFF, A: 0xff}, // 청록색
	}

	cycles := 5       // x 진동자의 회전수
	const res = 0.001 // 회전각
	size := 100       // 이미지 캔버스 크기 [-size..+size]
	nframes := 64     // 애니메이션 프레임 수
	delay := 8        // 10ms 단위의 프레임 간 지연

	for k, v := range params {
		if k == "cycles" {
			cycles = v
		} else if k == "size" {
			size = v
		} else if k == "nframes" {
			nframes = v
		} else if k == "delay" {
			delay = v
		} else {
			fmt.Printf("쿼리파라미터 %s 는 무시됨\n", k)
		}
	}

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

		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), uint8(colorIndex))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	// 최종 GIF를 출력 대상(out)에 인코딩
	gif.EncodeAll(out, &anim)
}
