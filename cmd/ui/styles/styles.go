package styles

import "github.com/charmbracelet/lipgloss"

var (
	PromptStyle           = lipgloss.NewStyle().Foreground(lipgloss.Color("#7FC8F8")).Bold(true)
	ContrastStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("#B088F9")).Bold(true)
	SelectedCheckboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#34C759")).Bold(true)
	NormalTextStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#F5F9FF")).Bold(true)
	BlurStyle             = lipgloss.NewStyle().Foreground(lipgloss.Color("#7D8590")).Bold(true)
	DimStyle              = lipgloss.NewStyle().Foreground(lipgloss.Color("#B0B9C6")).Bold(true)
	ErrorStyle            = lipgloss.NewStyle().Foreground(lipgloss.Color("#CE5D4C")).Bold(true)
)
