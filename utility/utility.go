package utility

import (
	"fmt"
	"math"
	"os"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"golang.org/x/exp/rand"
)

func UnitVectorNormalize(values []float64) []float64 {
    sumSquares := 0.0
    for _, v := range values {
        sumSquares += v * v
    }
    magnitude := math.Sqrt(sumSquares)

    normalized := make([]float64, len(values))
    if magnitude == 0 {
        return normalized // 크기가 0이면 0으로 처리
    }
    for i, v := range values {
        normalized[i] = v / magnitude
    }
    return normalized
}

// 에러 체크 함수
func CheckError(err error, message string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", message, err)
		os.Exit(1)
	}
}

// 텍스처 로드 함수
func LoadTexture(renderer *sdl.Renderer, filePath string) (*sdl.Texture, int32, int32) {
	image, err := img.Load(filePath)
	CheckError(err, "Failed to load image")
	defer image.Free()

	texWidth := image.W
	texHeight := image.H

	texture, err := renderer.CreateTextureFromSurface(image)
	CheckError(err, "Failed to create texture")

	return texture, texWidth, texHeight
}

// 타일 회전 랜덤 함수 (0, 90, 180, 270 중 하나를 랜덤으로 선택)
func RandomRotation() float64 {
	// 0, 90, 180, 270 중에서 랜덤으로 선택
	rotationAngles := []float64{0, 90, 180, 270}
	return rotationAngles[rand.Intn(4)] // 랜덤 인덱스 선택
}

// LoadFont 함수: 폰트를 로드하여 반환합니다.
func LoadFont(fontPath string, fontSize int) *ttf.Font {
	font, err := ttf.OpenFont(fontPath, fontSize)
	CheckError(err, "Failed to load font")
	return font
}

// RenderText 함수: 텍스트를 렌더링하여 SDL Texture로 반환합니다.
func RenderText(renderer *sdl.Renderer, font *ttf.Font, text string, color sdl.Color) *sdl.Texture {
	surface, err := font.RenderUTF8Solid(text, color)
	CheckError(err, "Failed to render text")
	defer surface.Free()

	texture, err := renderer.CreateTextureFromSurface(surface)
	CheckError(err, "Failed to create texture")
	return texture
}