package list

import (
	"clx/constants/style"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
	"time"
)

const (
	bullet   = "•"
	ellipsis = "…"
)

// Styles contains style definitions for this list component. By default, these
// values are generated by DefaultStyles.
type Styles struct {
	TitleBar     lipgloss.Style
	Title        lipgloss.Style
	Spinner      lipgloss.Style
	FilterPrompt lipgloss.Style
	FilterCursor lipgloss.Style

	// Default styling for matched characters in a filter. This can be
	// overridden by delegates.
	DefaultFilterCharacterMatch lipgloss.Style

	StatusBar             lipgloss.Style
	StatusEmpty           lipgloss.Style
	StatusBarActiveFilter lipgloss.Style
	StatusBarFilterCount  lipgloss.Style

	NoItems lipgloss.Style

	PaginationStyle lipgloss.Style
	HelpStyle       lipgloss.Style

	// Styled characters.
	ActivePaginationDot   lipgloss.Style
	InactivePaginationDot lipgloss.Style
	ArabicPagination      lipgloss.Style
	DividerDot            lipgloss.Style
}

// DefaultStyles returns a set of default style definitions for this list
// component.
func DefaultStyles() (s Styles) {
	verySubduedColor := lipgloss.AdaptiveColor{Light: "#DDDADA", Dark: "#3C3C3C"}
	subduedColor := lipgloss.AdaptiveColor{Light: "#9B9B9B", Dark: "#5C5C5C"}

	s.TitleBar = lipgloss.NewStyle().Padding(0, 0, 1, 2)

	//s.Title = lipgloss.NewStyle().
	//	Background(lipgloss.Color("62")).
	//	Foreground(lipgloss.Color("230")).
	//	Padding(0, 1)

	s.Spinner = lipgloss.NewStyle()
	//Foreground(lipgloss.Color("250")).
	//Background(lipgloss.Color("238")).
	//Faint(true)

	s.FilterPrompt = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#ECFD65"})

	s.FilterCursor = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"})

	s.DefaultFilterCharacterMatch = lipgloss.NewStyle().Underline(true)

	//s.StatusBar = lipgloss.NewStyle().
	//	Foreground(lipgloss.Color("16")).
	//	Background(lipgloss.Color("4"))

	s.StatusEmpty = lipgloss.NewStyle().Foreground(subduedColor)

	s.StatusBarActiveFilter = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#dddddd"})

	s.StatusBarFilterCount = lipgloss.NewStyle().Foreground(verySubduedColor)

	s.NoItems = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#909090", Dark: "#626262"})

	s.ArabicPagination = lipgloss.NewStyle().Foreground(subduedColor)

	s.PaginationStyle = lipgloss.NewStyle().PaddingLeft(2) //nolint:gomnd

	s.HelpStyle = lipgloss.NewStyle().Padding(1, 0, 0, 2)

	s.ActivePaginationDot = lipgloss.NewStyle().
		Foreground(style.GetSelectedPageForeground()).
		Background(style.GetHeaderBackground()).
		SetString(bullet)

	s.InactivePaginationDot = lipgloss.NewStyle().
		Foreground(style.GetUnselectedPageForeground()).
		Background(style.GetHeaderBackground()).
		SetString(bullet)

	s.DividerDot = lipgloss.NewStyle().
		Faint(true).
		SetString(" " + bullet + " ")

	return s
}

func getSpinner() spinner.Spinner {
	normal := lipgloss.NewStyle().
		Foreground(style.GetUnselectedItemForeground()).
		Background(style.GetLogoBackground()).
		Faint(true)

	color := normal.Copy()

	magenta := style.GetMagenta()
	yellow := style.GetYellow()
	blue := style.GetBlue()

	return spinner.Spinner{
		Frames: []string{
			normal.Render("fetching"),
			normal.Render("fetching"),
			normal.Render("fetching"),
			normal.Render("fetching"),
			normal.Render("fetching"),
			normal.Render("fetching"),
			color.Foreground(blue).Render("f") + normal.Render("etching"),
			color.Foreground(yellow).Render("f") + color.Foreground(blue).Render("e") + normal.Render("tching"),
			color.Foreground(magenta).Render("f") + color.Foreground(yellow).Render("e") + color.Foreground(blue).Render("t") + normal.Render("ching"),
			normal.Render("f") + color.Foreground(magenta).Render("e") + color.Foreground(yellow).Render("t") + color.Foreground(blue).Render("c") + normal.Render("hing"),
			normal.Render("fe") + color.Foreground(magenta).Render("t") + color.Foreground(yellow).Render("c") + color.Foreground(blue).Render("h") + normal.Render("ing"),
			normal.Render("fet") + color.Foreground(magenta).Render("c") + color.Foreground(yellow).Render("h") + color.Foreground(blue).Render("i") + normal.Render("ng"),
			normal.Render("fetc") + color.Foreground(magenta).Render("h") + color.Foreground(yellow).Render("i") + color.Foreground(blue).Render("n") + normal.Render("g"),
			normal.Render("fetch") + color.Foreground(magenta).Render("i") + color.Foreground(yellow).Render("n") + color.Foreground(blue).Render("g"),
			normal.Render("fetchi") + color.Foreground(magenta).Render("n") + color.Foreground(yellow).Render("g"),
			normal.Render("fetchin") + color.Foreground(magenta).Render("g"),
			normal.Render("fetching"),
		},
		FPS: 150 * time.Millisecond,
	}
}
