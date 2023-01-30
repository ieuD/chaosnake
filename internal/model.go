package internal

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type TickMsg time.Time

type Game struct {
	// Keys holds key bindings.
	Keys KeyMap
	// Obstacles is a data structure containing obstacles.
	Obstacles Obstacles
	// Cursor is the location of the cursor.
	Cursor SnakeBody
	// Score is the number of obstacles avoided.
	Score int
	// Help contains the Bubble Tea help model.
	Help help.Model
	// Viewport is the size of the game area.
	Viewport Location
	// Over is true when the player has lost.
	Over bool
	// Pressed is used to lock the cursor from moving until the next tick.
	Pressed bool
	// Layouted tracks whether the initial layout has been performed.
	Layouted bool
	Grid     [][]string

	snakeBody SnakeBody
}

func NewGame() Game {

	return Game{
		Keys: KeyMap{
			Up:    key.NewBinding(key.WithKeys("up", " ", "w"), key.WithHelp("↑/w/space", "Go Upward")),
			Down:  key.NewBinding(key.WithKeys("down", " ", "s"), key.WithHelp("↑/s", "Go Downward")),
			Left:  key.NewBinding(key.WithKeys("left", " ", "a"), key.WithHelp("↑/a", "Turn Right")),
			Right: key.NewBinding(key.WithKeys("right", " ", "d"), key.WithHelp("↑/d", "Turn Left")),
			Quit:  key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q/ctrl+c", "quit")),
		},
		Obstacles: NewObstacles(),
		Cursor:    SnakeBody{X: 5, Y: 10, Xspeed: 1, Yspeed: 0},
		Score:     0,
		Help:      help.New(),
		Over:      false,
		Viewport:  Location{},
		Pressed:   false,
		Layouted:  false,
	}
}

func (m Game) tick() tea.Cmd {
	return tea.Tick(time.Second/10, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (m Game) Init() tea.Cmd {

	M := make([][]string, m.Viewport.X)

	for i := 0; i < m.Viewport.X; i++ {
		row := make([]string, m.Viewport.Y)
		for j, _ := range row {
			row[j] = "X"
		}
		M = append(M, row)
	}
	m.Grid = M

	return m.tick()
}

func (m Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.Keys.Quit):
			fmt.Println()
			return m, tea.Quit
		case key.Matches(msg, m.Keys.Up):
			m.Cursor.ChangeDir(-1, 0)
			m.Pressed = true
		case key.Matches(msg, m.Keys.Down):
			m.Cursor.ChangeDir(1, 0)
			m.Pressed = true
		case key.Matches(msg, m.Keys.Left):
			m.Cursor.ChangeDir(0, -1)
			m.Pressed = true
		case key.Matches(msg, m.Keys.Right):
			m.Cursor.ChangeDir(0, 1)
			m.Pressed = true
			// Disable the key until the next tick.
			// Since the view does not update in real time, this prevents
			// hidden states in the game that are invisible to the user.

		}

	case tea.WindowSizeMsg:
		// Terminal resized.
		if !m.Layouted {
			m.Layouted = true
			m.Help.Width = msg.Width
			m.Viewport.X = 100
			m.Viewport.Y = 50
			m.Cursor = SnakeBody{X: m.Viewport.X / 2, Y: m.Viewport.Y / 2}
		}
	case TickMsg:
		m.Cursor.Update(m.Viewport.X, m.Viewport.Y)
		return m.Frame()
	}
	return m, nil
}

func (m Game) Frame() (tea.Model, tea.Cmd) {
	//if !m.Pressed {
	//	m.Cursor.Y++
	//}
	//m.Pressed = false

	return m, m.tick()
}

func (m Game) View() string {
	var sb strings.Builder
	sb.WriteString(TitleStyle.Render("Flapioca"))
	sb.WriteByte('\n')
	//viewport := make([]string, 0, m.Viewport.Y)

	fmt.Println(m.Grid)

	//for y := 0; y < m.Viewport.Y; y++ {
	//
	//	var line strings.Builder
	//	// Store the index of the leftmost obstacle encountered.
	//	// This is used to slice the obstacle list to avoid checking obstacles
	//	// we've already seen.
	//	for x := 0; x < m.Viewport.X; x++ {
	//		// Check if any obstacles collide with this cell.
	//		cellValue := ' '
	//
	//		if m.Cursor.X == x && m.Cursor.Y == y {
	//			cellValue = '*'
	//		}
	//		line.WriteRune(cellValue)
	//
	//	}
	//	viewport = append(viewport, line.String())
	//
	//}

	var line strings.Builder
	m.Grid[m.Cursor.X][m.Cursor.Y] = "*"

	sb.WriteString(ViewportStyle.Render(line.String()))
	sb.WriteString(fmt.Sprintf("\n%d point(s) ", m.Score))
	sb.WriteString(m.Help.View(m.Keys))

	if m.Over {
		sb.WriteString(GameOverStyle.Render("\n\n> Game over! <"))
	}

	// Send the UI for rendering
	return ViewStyle.Render(sb.String())
}
