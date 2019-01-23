package mdconv

type Option int

const (
	EnableTables        Option = 1 << iota // https://github.github.com/gfm/#tables-extension-
	EnableTaskListItems                    // https://github.github.com/gfm/#task-list-items-extension-
	EnableStrikethrough                    // https://github.github.com/gfm/#strikethrough-extension-
	EnableAutolinks
	EnableDisallowedRawHTML
	EnableLatex
	EnableGFM               = EnableTables | EnableTaskListItems | EnableStrikethrough | EnableAutolinks | EnableDisallowedRawHTML
	EnableAll               = EnableGFM | EnableLatex
	DisableExtension Option = 0
)
