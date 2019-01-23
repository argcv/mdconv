package mdconv

import (
	"testing"
)

func TestOption(t *testing.T) {
	t.Logf("DisableExtension: %v", DisableExtension)
	t.Logf("EnableTables: %v", EnableTables)
	t.Logf("EnableTaskListItems: %v", EnableTaskListItems)
	t.Logf("EnableStrikethrough: %v", EnableStrikethrough)
	t.Logf("EnableAutolinks: %v", EnableAutolinks)
	t.Logf("EnableDisallowedRawHTML: %v", EnableDisallowedRawHTML)
	t.Logf("EnableLatex: %v", EnableLatex)
	t.Logf("EnableGFM: %v", EnableGFM)
	t.Logf("EnableAll: %v", EnableAll)

}
