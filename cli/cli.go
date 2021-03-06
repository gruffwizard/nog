package cli

var Verbose bool


type CLI struct {
	IDEMode bool

	Action       string
	NogVersion   string
	NogCommit    string
	NogBuiltDate string
	NogBuiltBy   string
	ActiveID     string

	QuickStart     string
	QuickStartOnly bool
	Clone          string
	Convert        bool
	
	SrcDir string
	SrcVol string
	MvnVol string
	MvnDir string
}
