package core

/**
*   A class to do debug printfs
 */

type Debugger struct {
	Debug       bool
	ParserDebug bool
}

var Debug = Debugger{true, false}

func (this *Debugger) print(s string) {
	if this.Debug {
		LogWrite.LogMessage(s)
	}
}

func (this *Debugger) println(s string) {
	if this.Debug {
		LogWrite.LogMessage(s + "\n")
	}
}
