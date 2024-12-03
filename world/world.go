package world

import (
	"fmt"
	"mexis/mapgen"
	"mexis/utility"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/exp/rand"
)

const (
	BASE_GROUND_IMAGE_PATH = "./assets/image/ground"
)

// 맵 관련
// Map 구조체
type Map struct {
    Name   string  
    Width  int     
    Height int     
    Tiles  [][]Tile
}

// Tile 구조체 개선
type Tile struct {
    Type     string        // 타일 종류
    Texture  *sdl.Texture  // 타일의 텍스처
    Width    int32         // 타일의 너비
    Height   int32         // 타일의 높이
    Angle    float64       // 회전 각도
    X        int32         // 타일의 X 위치
    Y        int32         // 타일의 Y 위치
}

// 새 흙 타일 생성 함수
func NewDirtTile(renderer *sdl.Renderer, x, y int32) Tile {
    texture, w, h := utility.LoadTexture(renderer, BASE_GROUND_IMAGE_PATH+"/dirt.png")
    return Tile{
        Type:    "dirt",
        Texture: texture,
        Width:   w,
        Height:  h,
        X:       x,
        Y:       y,
        Angle:   utility.RandomRotation(), // 임의 회전 각도
    }
}

// 새 물 타일 생성 함수
func NewWaterTile(renderer *sdl.Renderer, x, y int32) Tile {
    texture, w, h := utility.LoadTexture(renderer, BASE_GROUND_IMAGE_PATH+"/water.png")
    return Tile{
        Type:    "water",
        Texture: texture,
        Width:   w,
        Height:  h,
        X:       x,
        Y:       y,
    }
}

// 새 풀 타일 생성 함수
func NewGrassTile(renderer *sdl.Renderer, x, y int32) Tile {
    texture, w, h := utility.LoadTexture(renderer, "./assets/image/object/grass.png")
    return Tile{
        Type:    "grass",
        Texture: texture,
        Width:   w,
        Height:  h,
        X:       x,
        Y:       y,
    }
}

// 새 나무 타일 생성 함수
func NewTreeTile(renderer *sdl.Renderer, x, y int32) Tile {
    texture, w, h := utility.LoadTexture(renderer, "./assets/image/object/tree.png")
    return Tile{
        Type:    "tree",
        Texture: texture,
        Width:   w,
        Height:  h,
        X:       x,
        Y:       y,
    }
}

// 새 철 타일 생성 함수
func NewStoneTile(renderer *sdl.Renderer, x, y int32) Tile {
    texture, w, h := utility.LoadTexture(renderer, "./assets/image/object/stone.png")
    return Tile{
        Type:    "stone",
        Texture: texture,
        Width:   w,
        Height:  h,
        X:       x,
        Y:       y,
    }
}

// 새 철 타일 생성 함수
func NewIronTile(renderer *sdl.Renderer, x, y int32) Tile {
    texture, w, h := utility.LoadTexture(renderer, "./assets/image/object/iron.png")
    return Tile{
        Type:    "iron",
        Texture: texture,
        Width:   w,
        Height:  h,
        X:       x,
        Y:       y,
    }
}

// 빈 타일 생성 함수
func NewEmptyTile(renderer *sdl.Renderer, x, y int32) Tile {
    return Tile{
        Type:    "empty",
        Texture: nil, // 빈 타일은 텍스처가 없음
        Width:   0,    // 크기 없음
        Height:  0,    // 크기 없음
        X:       x,    // 위치 설정
        Y:       y,    // 위치 설정
    }
}

// 새 땅 맵을 생성하고 초기화하는 함수
func NewGroundMap(width, height int, renderer *sdl.Renderer, groundSeed uint64) Map {
    // MapGenerator 초기화 (맵 크기와 시드 설정)
    mapGen := mapgen.NewMapGenerator(width, height, groundSeed)

    // 맵 생성
    canvas := mapGen.GenerateAndGetCanvas()

    // 맵의 타일을 2D 슬라이스로 초기화
    tiles := make([][]Tile, height)
    for i := range tiles {
        tiles[i] = make([]Tile, width)
    }

    // MapGenerator에서 생성된 맵을 기반으로 타일 설정
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            var tile Tile
            switch canvas[y][x] {
            case '.':
                tile = NewWaterTile(renderer, int32(x*32), int32(y*32))
            case 'o':
                tile = NewWaterTile(renderer, int32(x*32), int32(y*32))
            default:
                tile = NewDirtTile(renderer, int32(x*32), int32(y*32))
            }
            tiles[y][x] = tile
        }
    }

    // 새로운 맵 반환
    return Map{
        Name:   "ground",
        Width:  width,
        Height: height,
        Tiles:  tiles,
    }
}

// 자원 타일맵 생성 함수
func NewResourceMap(width, height int, ratios map[string]float64, renderer *sdl.Renderer, groundMap Map) Map {
    // 맵의 타일을 2D 슬라이스로 초기화
    tiles := make([][]Tile, height)
    for i := range tiles {
        tiles[i] = make([]Tile, width)
    }

    // 랜덤 시드 초기화
    rand.Seed(uint64(time.Now().UnixNano()))

    // 자원 타일 배치
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            randomValue := rand.Float64()
            cumulativeRatio := 0.0
            var tile Tile
            for resource, ratio := range ratios {
                cumulativeRatio += ratio
                if randomValue < cumulativeRatio && groundMap.Tiles[y][x] == NewDirtTile(renderer, int32(x), int32(y)) {
                    switch resource {
                    case "stone":
                        tile = NewStoneTile(renderer, int32(x*32), int32(y*32))
                    case "iron":
                        tile = NewIronTile(renderer, int32(x*32), int32(y*32))
                    case "tree":
                        tile = NewTreeTile(renderer, int32(x*32), int32(y*32))
                    case "grass":
                        tile = NewGrassTile(renderer, int32(x*32), int32(y*32))
                    default:
                        tile = NewEmptyTile(renderer, int32(x*32), int32(y*32))
                    }
                    break
                }
            }
            tiles[y][x] = tile
        }
    }

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
            fmt.Printf("%s ", m.Tiles[y][x].Type)
        }
        fmt.Println()
    }
}
