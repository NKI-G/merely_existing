package camera

import (
	"mexis/world"

	"github.com/veandco/go-sdl2/sdl"
)

// 카메라 관련
type Camera struct {
	Renderer *sdl.Renderer
	XPos int32
	YPos int32
}

// Camera 생성자 함수
func NewCamera(x, y int32, renderer *sdl.Renderer) *Camera {
	return &Camera{XPos: x, YPos: y, Renderer: renderer}
}

// 객체 구조체
type Object struct {
	Texture *sdl.Texture
	X int32
	Y int32
	W int32
	H int32
}

func (c *Camera) ObjectDraw(obj Object) {
    // 화면에 그릴 좌표 계산
    screenX := obj.X - c.XPos
    screenY := obj.Y - c.YPos

    // 객체의 크기
    width := obj.W
    height := obj.H

    // 렌더링할 사각형 정의
    dstRect := &sdl.Rect{
        X: int32(screenX),
        Y: int32(screenY),
        W: int32(width),
        H: int32(height),
    }

    // CopyEx 함수 호출
    c.Renderer.CopyEx(
        obj.Texture, // 텍스처
        nil,         // 소스 사각형(nil은 전체 텍스처를 사용)
        dstRect,     // 대상 사각형
        0,           // 회전 각도 (0은 회전 없음)
        nil,         // 회전 중심(nil은 기본값 사용)
        sdl.FLIP_NONE, // 플립 옵션 (기본값은 플립 없음)
    )
}

// Draw 함수: 카메라 위치를 기준으로 타일 맵을 그리는 역할
func (c *Camera) MapDraw(tilemap world.Map) {
    for _, row := range tilemap.Tiles {
        for _, tile := range row {
            screenX := tile.X - c.XPos
            screenY := tile.Y - c.YPos

            width := tile.Width
            height := tile.Height

            dstRect := &sdl.Rect{
                X: int32(screenX),
                Y: int32(screenY),
                W: int32(width),
                H: int32(height),
            }

            // 회전 적용
            c.Renderer.CopyEx(tile.Texture, nil, dstRect, tile.Angle, nil, sdl.FLIP_NONE)
        }
    }
}
