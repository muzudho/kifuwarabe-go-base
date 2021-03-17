package stone

// Stone - 石
type Stone int

// state
const (
	// None - 空点
	None Stone = iota
	// Black - 黒石
	Black
	// White - 白石
	White
	// Wall - 壁
	Wall
)

// FlipColor - 白黒反転させます
func FlipColor(col int) int {
	return 3 - col
}
