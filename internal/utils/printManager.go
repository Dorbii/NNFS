package utils

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout)
	styles := log.DefaultStyles()
	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
		Padding(0, 1, 0, 1).
		Foreground(lipgloss.Color("204"))
	logger.SetStyles(styles)
}
func Error(etype string, args ...any) {
	msgFlag := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("204")).
		SetString("ERROR:")
	switch etype {
	case "dotProduct":
		msg := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("204")).
			SetString(fmt.Sprintf("finding dot product of \n\n%v \n\nand\n\n%v\n\n", args[0], args[1]))
		logger.Error(msgFlag, "msg", msg)
	case "sum":
		msg := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("204")).
			SetString(fmt.Sprintf("adding \n\n%v \n\nand\n\n%v", args[0], args[1]))
		logger.Error(msgFlag, "msg", msg)
	case "mulScalar":
		msg := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("204")).
			SetString(fmt.Sprintf("when multiplying scalar \n\n%v\n\n", args[0]))
		logger.Error(msgFlag, "msg", msg)
	}
}
func Output(output ...any) {
	fmt.Println(lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("10")).
		SetString(fmt.Sprintf("%v", output)))
}
func Header(s string) {
	println()
	fmt.Println(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12")).SetString(s))
}
