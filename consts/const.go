package consts

const (
	Usage = "" +
		"Type `asciicount-test generate` to fill directory with test data\n" +
		"Type `asciicount-test count` to get a histogram"
	CmdGenerate = "generate"
	CmdCount    = "count"
)

const (
	Directory      = "data"
	FilesCount     = 1000
	SymbolsPerFile = 1000
	CounterWorkers = 5
)

const (
	ASCIIMinWithSpace = 32
	ASCIIMin          = ASCIIMinWithSpace + 1
	ASCIIMax          = 126

	ASCIIRange          = ASCIIMax - ASCIIMin
	ASCIIRangeWithSpace = ASCIIMax - ASCIIMinWithSpace
)

const (
	CounterBarLength = 100
)
