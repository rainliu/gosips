package core

import ()

/**
*   A class to do debug printfs
 */

type DebugObject struct {
    Debug       bool
    ParserDebug bool
}

var Debug = DebugObject{true, false}

func (this *DebugObject) print(s string) {
    if this.Debug {
        LogWriter.LogMessage(s)
    }
}

func (this *DebugObject) println(s string) {
    if this.Debug {
        LogWriter.LogMessage(s + "\n")
    }
}