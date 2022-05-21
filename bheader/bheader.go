package bheader

import (
	"clx/constants/category"
	"clx/constants/style"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

func GetHeader(selectedSubHeader int, width int) string {
	bg := lipgloss.AdaptiveColor{Light: style.LogoBackgroundLight, Dark: style.LogoBackgroundDark}

	c := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: style.MagentaLight, Dark: style.MagentaDark}).
		Background(bg)

	l := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: style.YellowLight, Dark: style.YellowDark}).
		Background(bg)

	x := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: style.BlueLight, Dark: style.BlueDark}).
		Background(bg)

	title := c.Render("  c") + l.Render("l") + x.Render("x  ")

	categories := getCategories(selectedSubHeader)
	filler := getFiller(title, categories, width)
	return title + categories + filler
}

func getFiller(title string, categories string, width int) string {
	availableSpace := width - lipgloss.Width(title+categories)

	if availableSpace < 0 {
		return ""
	}

	filler := strings.Repeat(" ", availableSpace)

	return lipgloss.NewStyle().
		Background(lipgloss.AdaptiveColor{Light: style.HeaderBackgroundLight, Dark: style.HeaderBackgroundDark}).
		Render(filler)
}

func getCategories(selectedSubHeader int) string {
	subHeaders := []string{"new", "ask", "show"}
	fg := lipgloss.AdaptiveColor{Light: style.UnselectedItemLight, Dark: style.UnselectedItemDark}
	bg := lipgloss.AdaptiveColor{Light: style.HeaderBackgroundLight, Dark: style.HeaderBackgroundDark}

	categories := lipgloss.NewStyle().
		Background(bg).
		Render("   ")

	separator := lipgloss.NewStyle().
		Foreground(fg).
		Background(bg).
		Render(" • ")

	for i, subHeader := range subHeaders {
		isOnLastItem := i == len(subHeaders)-1
		selectedCatColor, isSelected := getColor(i, selectedSubHeader)

		categories += lipgloss.NewStyle().
			Foreground(selectedCatColor).
			Background(bg).
			Bold(isSelected).
			Render(subHeader)

		if !isOnLastItem {
			categories += separator
		}

	}

	return categories
}

func getColor(i int, selectedSubHeader int) (lipgloss.TerminalColor, bool) {
	if i+1 == selectedSubHeader {
		return getSelectedCategoryColor(i + 1)
	}

	return lipgloss.AdaptiveColor{Light: style.UnselectedItemLight, Dark: style.UnselectedItemDark}, false
}

func getSelectedCategoryColor(selectedSubHeader int) (lipgloss.TerminalColor, bool) {
	switch selectedSubHeader {
	case category.New:
		return lipgloss.AdaptiveColor{Light: style.MagentaLight, Dark: style.MagentaDark}, true
	case category.Ask:
		return lipgloss.AdaptiveColor{Light: style.YellowLight, Dark: style.YellowDark}, true
	case category.Show:
		return lipgloss.AdaptiveColor{Light: style.BlueLight, Dark: style.BlueDark}, true
	case category.Favorites:
		return lipgloss.AdaptiveColor{Light: style.PinkLight, Dark: style.PinkDark}, true
	default:
		return lipgloss.AdaptiveColor{Light: style.UnselectedItemLight, Dark: style.UnselectedItemDark}, false
	}
}
