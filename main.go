package main

import (
	c "mexis/camera"
	"mexis/utility"
	"mexis/world"
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// 상수 정의
const (
	WINDOW_TITLE        = "Merely Existing | 단지 존재할 뿐."

	// 화면 차원 상수
	SCREEN_WIDTH = 50*32;
	SCREEN_HEIGHT = 30*32;

    OUR_NPC_IMAGE_PATH = "./assets/image/our_npc.png"
)

var (
    velocityX, velocityY float32
    damping              float32 = 0.9  // 감쇠율
    sensitivity          float32 = 2  // 마우스 민감도
    mouseDragging        bool    = false
	isWorldOutCamera bool = false
)

func run() int {
	// SDL 초기화
	err := sdl.Init(sdl.INIT_EVERYTHING)
	utility.CheckError(err, "Failed to initialize SDL")
	defer sdl.Quit()
	
	utility.CheckError(ttf.Init(), "Failed to initialize TTF")
	defer ttf.Quit()

	// 창 생성
	window, err := sdl.CreateWindow(WINDOW_TITLE, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		SCREEN_WIDTH, SCREEN_HEIGHT, sdl.WINDOW_SHOWN)
        utility.CheckError(err, "Failed to create window")
	defer window.Destroy()

	// 렌더러 생성
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	utility.CheckError(err, "Failed to create renderer")
	defer renderer.Destroy()

	font := utility.LoadFont("assets/Galmuri11.ttf", 24)
	defer font.Close()

	// 텍스트 렌더링
	text := "기다리면 화면이 돌아올거에요."
	color := sdl.Color{R: 255, G: 255, B: 255, A: 255}
	texture := utility.RenderText(renderer, font, text, color)
	defer texture.Destroy()

	_, _, width, height, err := texture.Query()
	utility.CheckError(err, "Failed to query texture")
	dstRect := sdl.Rect{X: 100, Y: 100, W: int32(width), H: int32(height)}

	camera := c.NewCamera(0, 0, renderer)

	//resourceRatios := map[string]float64{
	//	"stone":  0.15,
	//	"iron":   0.15,
	//	"tree":   0.15,
	//	"grass":  0.35,
	//	"empty":  0.2,
	//}

	

	//resourceMap := NewResourceMap(SCREEN_WIDTH/32, SCREEN_HEIGHT/32, resourceRatios, renderer)
	groundMap := world.NewGroundMap(100, 100, renderer, 1000) //183066722346246775
	world.PrintMap(groundMap)

    prevMouseX, prevMouseY := int32(0), int32(0) // 이전 마우스 위치 저장

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
			case *sdl.MouseMotionEvent: // 마우스 이동 이벤트
				if e.State == 1 { // 왼쪽 버튼이 눌려진 상태
					mouseDragging = true
		
					// 마우스 이동량 계산 (민감도 적용)
					deltaX := float32(e.X-prevMouseX) * sensitivity
					deltaY := float32(e.Y-prevMouseY) * sensitivity
		
					// 속도 업데이트
					velocityX = deltaX
					velocityY = deltaY
		
					// 카메라 위치 갱신
					camera.XPos -= int32(deltaX)
					camera.YPos -= int32(deltaY)
		
					// 현재 마우스 위치를 저장
					prevMouseX, prevMouseY = e.X, e.Y
				}
			case *sdl.MouseButtonEvent: // 마우스 버튼 이벤트
				if e.Type == sdl.MOUSEBUTTONDOWN {
					// 마우스 버튼을 눌렀을 때 현재 위치 저장
					prevMouseX, prevMouseY = e.X, e.Y
					mouseDragging = true
				} else if e.Type == sdl.MOUSEBUTTONUP {
					mouseDragging = false
				}
			}
		}
		
		// 관성 처리
		if !mouseDragging {
			// 카메라 위치를 속도에 따라 업데이트
			camera.XPos -= int32(velocityX)
			camera.YPos -= int32(velocityY)
		
			// 속도 감쇠
			velocityX *= damping
			velocityY *= damping
		
			// 일정 임계값 이하로 줄어들면 속도를 0으로 설정
			if velocityX < 0.1 && velocityX > -0.1 {
				velocityX = 0
			}
			if velocityY < 0.1 && velocityY > -0.1 {
				velocityY = 0
			}
		}

		renderer.SetDrawColor(0, 0, 0, 255) // 배경색 검정
		renderer.Clear()

		camera.MapDraw(groundMap)

		if camera.XPos < -100 {
			camera.XPos += 10
			renderer.Copy(texture, nil, &dstRect)
			
		}
		if camera.YPos < -100 {
			camera.YPos += 10
			renderer.Copy(texture, nil, &dstRect)
		}

		if camera.XPos > 1600+100 {
			camera.XPos -= 10
			renderer.Copy(texture, nil, &dstRect)
			
		}
		if camera.YPos > 2241+100 {
			camera.YPos -= 10
			renderer.Copy(texture, nil, &dstRect)
		}
		println(camera.XPos, camera.YPos, isWorldOutCamera)

		// 화면 그리기

		renderer.Present() // 화면에 렌더링
	}

	return 0
}

func main() {
	os.Exit(run())
}