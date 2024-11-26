package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// 상수 정의
const (
	winTitle        = "Go-SDL2 Texture Example"
	winWidth, winHeight = 800, 600
	ourNpcImagePath = "./assets/image/our_npc.png"
)

// 에러 체크 함수
func checkError(err error, message string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", message, err)
		os.Exit(1)
	}
}

// 텍스처 로드 함수
func loadTexture(renderer *sdl.Renderer, filePath string) (*sdl.Texture, int32, int32) {
	image, err := img.Load(filePath)
	checkError(err, "Failed to load image")
	defer image.Free()

	texWidth := image.W
	texHeight := image.H

	texture, err := renderer.CreateTextureFromSurface(image)
	checkError(err, "Failed to create texture")

	return texture, texWidth, texHeight
}

func run() int {
	// SDL 초기화
	err := sdl.Init(sdl.INIT_EVERYTHING)
	checkError(err, "Failed to initialize SDL")
	defer sdl.Quit()

	// 창 생성
	window, err := sdl.CreateWindow(winTitle, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	checkError(err, "Failed to create window")
	defer window.Destroy()

	// 렌더러 생성
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	checkError(err, "Failed to create renderer")
	defer renderer.Destroy()

	// 텍스처 로드
	texture, texWidth, texHeight := loadTexture(renderer, ourNpcImagePath)
	defer texture.Destroy()

	// 텍스처 렌더링 영역 정의
	src := sdl.Rect{X: 0, Y: 0, W: texWidth, H: texHeight}
	dst := sdl.Rect{X: (winWidth - texWidth) / 2, Y: (winHeight - texHeight) / 2, W: texWidth, H: texHeight}

	// 이벤트 루프
	running := true
	for running {
		// 이벤트 처리
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent: // 창 닫기 이벤트
				running = false
			case *sdl.KeyboardEvent: // 키보드 입력
				if e.Type == sdl.KEYDOWN && e.Keysym.Sym == sdl.K_ESCAPE {
					running = false
				}
			}
		}

		// 화면 그리기
		renderer.SetDrawColor(0, 0, 0, 255) // 배경색 검정
		renderer.Clear()
		renderer.Copy(texture, &src, &dst) // 텍스처 복사 및 그리기
		renderer.Present()                // 화면에 렌더링
	}

	return 0
}

func main() {
	os.Exit(run())
}
