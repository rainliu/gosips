package core

import (
    "os"
)

/**
*  Log System Errors. Also used for debugging log.
 */
/** Dont trace
 */
const TRACE_NONE = 0

/** Trace message processing
 */
const TRACE_MESSAGES = 16

/** Trace exception processing
 */
const TRACE_EXCEPTION = 17

/** Debug trace level (all tracing enabled).
 */
const TRACE_DEBUG = 32

/** Name of the log file in which the trace is written out
 * (default is /tmp/sipserverlog.txt)
 */
type LogWriterObject struct {

    /** Print writer that is used to write out the log file.
     */
    printWriter *os.File

    /** Flag to indicate that logging is enabled. This needs to be
    * static and public in order to globally turn logging on or off.
    * This is static for efficiency reasons (the java compiler will not
    * generate the logging code if this is set to false).
     */
    traceLevel int

    logFileName string

    needsLogging bool

    lineCount int
}

var LogWriter = LogWriterObject{nil, TRACE_NONE, "debuglog.txt", false, 0}

/** Set the log file name 
*@param name is the name of the log file to set. 
 */
func (this *LogWriterObject) SetLogFileName(name string) {
    this.logFileName = name
}

func (this *LogWriterObject) LogMessageToFile(message, logFileName string) {
    var err error
    //try {
    this.printWriter, err = os.OpenFile(logFileName, os.O_APPEND, 0)
    if err != nil {
        println("Can't open file in LogMessageToFile")
        return
    }
    defer this.printWriter.Close()

    this.printWriter.WriteString(" ---------------------------------------------- ")
    this.printWriter.WriteString(message)
    //} catch (IOException ex) {
    //	ex.printStackTrace();
    //}
}

func (this *LogWriterObject) checkLogFile() {
    if this.printWriter != nil {
        return
    }
    if this.logFileName == "" {
        return
    }
    //try {
    var err error
    this.printWriter, err = os.OpenFile(this.logFileName, os.O_APPEND, 0)
    if err != nil {
        println("Can't open file in checkLogFile")
        return
    }
    //} catch (IOException ex) {
    //	ex.printStackTrace();
    //}
}

func (this *LogWriterObject) println(message string) {
    for i := 0; i < len(message); i++ {
        if message[i] == '\n' {
            this.lineCount++
        }
    }
    this.checkLogFile()
    if this.printWriter != nil {
        this.printWriter.WriteString(message)
    }
    this.lineCount++
}

/** Log a message into the log file.
 * @param message message to log into the log file.
 */
func (this *LogWriterObject) LogMessage(message string) {
    if !this.needsLogging {
        return
    }

    this.checkLogFile()
    this.println(message)
}

/** Set the trace level for the stack.
 */
func (this *LogWriterObject) SetTraceLevel(level int) {
    this.traceLevel = level
}

/** Get the trace level for the stack.
 */
func (this *LogWriterObject) GetTraceLevel() int {
    return this.traceLevel
}
