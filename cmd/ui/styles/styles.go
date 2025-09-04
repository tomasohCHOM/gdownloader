package styles

import "github.com/charmbracelet/lipgloss"

var (
	LogoStyle             = lipgloss.NewStyle().Foreground(lipgloss.Color("#5AA9E6")).Bold(true)
	HeaderStyle           = lipgloss.NewStyle().Foreground(lipgloss.Color("#7FC8F8")).Bold(true)
	ContrastStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("#B088F9")).Bold(true)
	StatsStyle            = lipgloss.NewStyle().Foreground(lipgloss.Color("#4CC9F0")).Bold(true)
	SelectedCheckboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#34C759")).Bold(true)
	SelectedTextStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#F5F9FF")).Bold(true)
	BlurStyle             = lipgloss.NewStyle().Foreground(lipgloss.Color("#7D8590")).Bold(true)
	DimStyle              = lipgloss.NewStyle().Foreground(lipgloss.Color("#B0B9C6")).Bold(true)
)
