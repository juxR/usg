package usg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResultOfOneSymbol(t *testing.T) {
	GetDefault := setDefaultSymbols()
	assert.Equal(t, GetDefault.Play, "▶", "Symbols <play> should be the same")
	assert.Equal(t, GetDefault.Square, "▇", "Symbols <square> should be the same")
	assert.Equal(t, GetDefault.Tick, "✔", "Symbols <tick> should be the same")
	assert.Equal(t, GetDefault.SevenEighths, "⅞", "Symbols <seven eighths> should be the same")
}

func TestResultOfMultipleSymbols(t *testing.T) {
	GetDefault := setDefaultSymbols()
	assert.Equal(t, GetDefault.Play+" test "+Get.Tick+Get.Tick, "▶ test ✔✔", "Should be the same")
}

func TestWindowsFallback(t *testing.T) {
	GetWindows := setWindowsSymbols()
	assert.Equal(t, GetWindows.Play, "►", "Symbols <play> should be the same")
	assert.Equal(t, GetWindows.Square, "█", "Symbols <square> should be the same")
	assert.Equal(t, GetWindows.Tick, "√", "Symbols <tick> should be the same")
	assert.Equal(t, GetWindows.SevenEighths, "7/8", "Symbols <seven eighths> should be the same")

}
