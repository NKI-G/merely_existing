package mapgen

import (
	"math"

	"golang.org/x/exp/rand"
)

// MapGenerator는 맵 생성과 관련된 구조체입니다.
type MapGenerator struct {
	width     int
	height    int
	grid      [][]int
	canvas    [][]rune
	processed [][]bool // 이미 처리된 위치를 기록
}

const (
	waterFrequency = 200 // 높을수록 적음
)

// NewMapGenerator는 MapGenerator를 초기화하는 공개 함수입니다.
func NewMapGenerator(width, height int, seed uint64) *MapGenerator {
	rand.Seed(seed)

	grid := make([][]int, height)
	canvas := make([][]rune, height)
	processed := make([][]bool, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]int, width)
		canvas[i] = make([]rune, width)
		processed[i] = make([]bool, width)
		for j := 0; j < width; j++ {
			grid[i][j] = rand.Intn(waterFrequency) // 0~10 랜덤 값
			canvas[i][j] = ' '        // 빈 공간으로 초기화
			processed[i][j] = false   // 아직 처리되지 않음
		}
	}

	return &MapGenerator{
		width:     width,
		height:    height,
		grid:      grid,
		canvas:    canvas,
		processed: processed,
	}
}

// GenerateAndGetCanvas는 맵을 생성하고 결과를 반환하는 공개 함수입니다.
func (mg *MapGenerator) GenerateAndGetCanvas() [][]rune {
	mg.generateMap()
	return mg.canvas
}

// PrintCanvas는 생성된 맵을 콘솔에 출력하는 공개 함수입니다.
func (mg *MapGenerator) PrintCanvas() {
	for _, row := range mg.canvas {
		for _, cell := range row {
			print(string(cell))
		}
		println()
	}
}

// 비공개 함수: 맵을 생성합니다.
func (mg *MapGenerator) generateMap() {
	for y, row := range mg.grid {
		for x, value := range row {
			// 이미 처리된 위치는 무시
			if value == 0 && !mg.processed[y][x] {
				mg.canvas[y][x] = '.' // 중심 점 찍기
				radius := rand.Intn(3) + 3 // 반지름: 3~5
				mg.drawCircle(x, y, radius)
			}
		}
	}
}

// 비공개 함수: 원을 그립니다.
func (mg *MapGenerator) drawCircle(cx, cy, radius int) {
	for y := cy - radius; y <= cy+radius; y++ {
		for x := cx - radius; x <= cx+radius; x++ {
			if x >= 0 && x < mg.width && y >= 0 && y < mg.height {
				// 유클리드 거리로 원 점 확인
				if math.Sqrt(float64((x-cx)*(x-cx)+(y-cy)*(y-cy))) <= float64(radius) {
					mg.canvas[y][x] = 'o' // 원의 픽셀
					mg.processed[y][x] = true // 처리된 위치로 기록
				}
			}
		}
	}
}
