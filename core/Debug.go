package core

import ()

/**
*   A class to do debug printfs
 */

type DebugObject struct {
    debug       bool
    parserDebug bool
}

var Debug = DebugObject{true, false}

func (this *DebugObject) print(s string) {
    if this.debug {
        LogWriter.LogMessage(s)
    }
}

func (this *DebugObject) println(s string) {
    if this.debug {
        LogWriter.LogMessage(s + "\n")
    }
}