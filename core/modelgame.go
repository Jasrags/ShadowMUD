package core

import (
	"fmt"
	"strings"

	"github.com/Jasrags/ShadowMUD/common"
	"github.com/sirupsen/logrus"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish/bubbletea"
)

var (
	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.BorderStyle(b)
	}()
)

type (
	errMsg     error
	contentMsg struct {
		from    string
		content string
	}
)

type gameModel struct {
	*World
	id        string
	width     int
	height    int
	char      common.Character
	styles    map[string]lipgloss.Style
	viewport  viewport.Model
	textInput textinput.Model
	ready     bool
	content   []string
	err       error
}

func NewGameModel(s ssh.Session) gameModel {
	renderer := bubbletea.MakeRenderer(s)
	// Styles
	txtStyle := renderer.NewStyle().Foreground(lipgloss.Color("10"))
	quitStyle := renderer.NewStyle().Foreground(lipgloss.Color("8"))
	statusBarStyle := renderer.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
		Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})
	statusText := renderer.NewStyle().Inherit(statusBarStyle)
	senderStyle := renderer.NewStyle().Foreground(lipgloss.Color("5"))

	// Components
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	ti.Placeholder = "Type here..."

	// content := lipgloss.PlaceVertical(
	// 	m.styles["text"].Render(m.char.Room.Spec.Name),
	// 	fmt.Sprintf("Description: %s", m.char.Room.Spec.Description),
	// 	m.styles["quit"].Render("Press q to quit."))

	var sb strings.Builder
	if _, err := sb.WriteString(txtStyle.Render("Welcome to ShadowMUD\n")); err != nil {
		logrus.WithError(err).Error("Could not write to string builder")
	}

	return gameModel{
		styles: map[string]lipgloss.Style{
			"text":       txtStyle,
			"quit":       quitStyle,
			"statusBar":  statusBarStyle,
			"statusText": statusText,
			"sender":     senderStyle,
		},
		textInput: ti,
		char: common.Character{
			Name: "Test",
			Room: common.Room{
				ID:   common.CoreRooms[0].ID,
				Spec: &common.CoreRooms[0],
			},
		},
		content: []string{txtStyle.Render("Welcome to ShadowMUD")},
	}
}

func (m gameModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m gameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			m.World.send(contentMsg{content: m.textInput.Value()})
			m.textInput.SetValue("")
		}

	case contentMsg:
		m.content = append(m.content, m.styles["sender"].Render(msg.content))
		m.viewport.SetContent(strings.Join(m.content, "\n"))
		m.viewport.GotoBottom()

	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.viewport.SetContent(strings.Join(m.content, "\n"))
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height
		}
		m.height = msg.Height
		m.width = msg.Width

	case errMsg:
		m.err = msg
		return m, nil
	}

	// Handle keyboard and mouse events in the viewport
	m.textInput, cmd = m.textInput.Update(msg)
	cmds = append(cmds, cmd)
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m gameModel) headerView() string {
	title := titleStyle.Render("Mr. Pager")
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m gameModel) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func (m gameModel) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		m.headerView(),
		// m.styles["text"].Render(m.char.Room.Spec.Name),
		// fmt.Sprintf("Description: %s", m.char.Room.Spec.Description),
		// m.styles["quit"].Render("Press q to quit."),
		m.viewport.View(),
		m.footerView(),

		lipgloss.JoinHorizontal(
			lipgloss.Top,
			m.styles["statusBar"].Width(m.width).Render("Status: "+m.char.Name),
		),
		m.textInput.View(),
	)
}
