package main

type SnakeBody struct {
	X      int
	Y      int
	XSpeed int
	YSpeed int
}

func (sb *SnakeBody) ChangeDir(vertical int, horizontal int) {
	sb.YSpeed = vertical
	sb.XSpeed = horizontal

}

func (sb *SnakeBody) Update(width int, height int) {
	sb.X = (sb.X + sb.XSpeed) % width
	if sb.X < 0 {
		sb.X += width
	}

	sb.Y = (sb.Y + sb.YSpeed) % height

	if sb.Y < 0 {
		sb.Y += height
	}
}
