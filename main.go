package main

import (
	"fmt"
	"os"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/exp/rand"
)

// 상수 정의
const (
	winTitle        = "Go-SDL2 Texture Example"
	winWidth, winHeight = 32*25, 32*20
	baseGroundImagePath = "./assets/image/ground"
	ourNpcImagePath = "./assets/image/our_npc.png"
)

// 맵 관련
// Tile 구조체
type Tile struct {
    Type     string 
    Texture  *sdl.Texture
	Width    int32
	Height   int32
	Angle    float64
}

// Map 구조체
type Map struct {
    Name   string  
    Width  int     
    Height int     
    Tiles  [][]Tile
}

// 새 흙 타일 생성 함수
func NewDirtTile(renderer *sdl.Renderer) Tile {
    texture, w, h := loadTexture(renderer, baseGroundImagePath+"/dirt.png")
    return Tile{
        Type:    "dirt", 
        Texture: texture,
        Width:   w,
        Height:  h,
		Angle: randomRotation(),
    }
}

// 새 물 타일 생성 함수
func NewWaterTile(renderer *sdl.Renderer) Tile {
    texture, w, h := loadTexture(renderer, baseGroundImagePath+"/water.png")
    return Tile{
        Type:    "water",
        Texture: texture,
        Width:   w,
        Height:  h,
    }
}

// 새 풀 타일 생성 함수
func NewGrassTile(renderer *sdl.Renderer) Tile {
    texture, w, h := loadTexture(renderer, "./assets/image/object/grass.png")
    return Tile{
        Type:    "grass",
        Texture: texture,
        Width:   w,
        Height:  h,
    }
}

// 새 나무 타일 생성 함수
func NewTreeTile(renderer *sdl.Renderer) Tile {
    texture, w, h := loadTexture(renderer, "./assets/image/object/tree.png")
    return Tile{
        Type:    "tree",
        Texture: texture,
        Width:   w,
        Height:  h,
    }
}

// 새 철 타일 생성 함수
func NewStoneTile(renderer *sdl.Renderer) Tile {
    texture, w, h := loadTexture(renderer, "./assets/image/object/stone.png")
    return Tile{
        Type:    "stone",
        Texture: texture,
        Width:   w,
        Height:  h,
    }
}

// 새 철 타일 생성 함수
func NewIronTile(renderer *sdl.Renderer) Tile {
    texture, w, h := loadTexture(renderer, "./assets/image/object/iron.png")
    return Tile{
        Type:    "iron",
        Texture: texture,
        Width:   w,
        Height:  h,
    }
}

// 빈 타일 생성 함수
func NewEmptyTile(renderer *sdl.Renderer) Tile {
    return Tile{
        Type:    "empty",
        Texture: nil, // 빈 타일은 텍스처가 없음
        Width:   0,    // 크기 없음
        Height:  0,    // 크기 없음
    }
}



// 새 땅 맵을 생성하고 초기화하는 함수
func NewGroundMap(width, height int, renderer *sdl.Renderer) Map {
    // 맵의 타일을 2D 슬라이스로 초기화
    tiles := make([][]Tile, height)
    for i := range tiles {
        tiles[i] = make([]Tile, width)
    }

    // 각 타일의 값을 기본값으로 설정 (여기서는 "dirt" 타입)
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            tiles[y][x] = NewDirtTile(renderer) // 각 타일을 "dirt"로 초기화
        }
    }

    // 새로운 맵 반환
    return Map{
        Name:   "ground",   // 맵 이름
        Width:  width,  // 맵의 가로 크기
        Height: height, // 맵의 세로 크기
        Tiles:  tiles,  // 타일들
    }
}

// 자원 타일맵 생성 함수
func NewResourceMap(width, height int, ratios map[string]float64, renderer *sdl.Renderer) Map {
	// 맵의 타일을 2D 슬라이스로 초기화
	tiles := make([][]Tile, height)
	for i := range tiles {
		tiles[i] = make([]Tile, width)
	}

	// 랜덤 시드 초기화
	rand.Seed(uint64(time.Now().UnixNano())) // int64 -> uint64 변환

	// 자원 타일 배치
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// 랜덤 값 생성
			randomValue := rand.Float64()

			// 각 자원 타입을 비율에 맞게 배치
			cumulativeRatio := 0.0
			for resource, ratio := range ratios {
				cumulativeRatio += ratio
				if randomValue < cumulativeRatio {
					// 자원 타입에 맞는 타일을 배치
					switch resource {
					case "stone":
						tiles[y][x] = NewStoneTile(renderer)
					case "iron":
						tiles[y][x] = NewIronTile(renderer)
					case "tree":
						tiles[y][x] = NewTreeTile(renderer)
					case "grass":
						tiles[y][x] = NewGrassTile(renderer)
					default:
						tiles[y][x] = NewEmptyTile(renderer) // 기본적으로 빈 타일
					}
					break
				}
			}
		}
	}

	// 새 맵 반환
	return Map{
		Name:   "resource",
		Width:  width,
		Height: height,
		Tiles:  tiles,
	}
}


// 맵 타일 출력 함수
func PrintMap(m Map) {
    for y := 0; y < m.Height; y++ {
        for x := 0; x < m.Width; x++ {
            fmt.Printf("%s ", m.Tiles[y][x].Type) // 각 타일의 Type 출력
        }
        fmt.Println()
    }
}


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

// 타일 회전 랜덤 함수 (0, 90, 180, 270 중 하나를 랜덤으로 선택)
func randomRotation() float64 {
	// 0, 90, 180, 270 중에서 랜덤으로 선택
	rotationAngles := []float64{0, 90, 180, 270}
	return rotationAngles[rand.Intn(4)] // 랜덤 인덱스 선택
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

	//resourceRatios := map[string]float64{
	//	"stone":  0.15,
	//	"iron":   0.15,
	//	"tree":   0.15,
	//	"grass":  0.35,
	//	"empty":  0.2,
	//}

	

	//resourceMap := NewResourceMap(winWidth/32, winHeight/32, resourceRatios, renderer)
	groundMap := NewGroundMap(winWidth/32, winHeight/32, renderer)
	PrintMap(groundMap)

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
					
				}
			}
		}

		// 화면 그리기
		renderer.SetDrawColor(0, 0, 0, 255) // 배경색 검정
		renderer.Clear()

		renderer.Copy(texture, &src, &dst) // 텍스처 복사 및 그리기

		// 타일 그리기 (랜덤 회전 적용)
		for y, i := range groundMap.Tiles {
			for x, j := range i {
				// 텍스처 소스 영역
				tsrc := sdl.Rect{X: 0, Y: 0, W: j.Width, H: j.Height}

				// 대상 위치
				tdst := sdl.Rect{X: int32(x * 32), Y: int32(y * 32), W: j.Width, H: j.Height}

				// 회전하여 타일 그리기
				err := renderer.CopyEx(j.Texture, &tsrc, &tdst, j.Angle, nil, sdl.FLIP_NONE)
				if err != nil {
					fmt.Printf("Failed to copy texture: %v\n", err)
					return 1
				}
			}
		}
		

		renderer.Present() // 화면에 렌더링
	}

	return 0
}

func main() {
	os.Exit(run())
}