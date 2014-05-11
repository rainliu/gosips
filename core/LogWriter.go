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
type LogWriter struct {

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

var LogWrite = LogWriter{nil, TRACE_NONE, "debug.log", false, 0}

/** Set the log file name
*@param name is the name of the log file to set.
 */
func (this *LogWriter) SetLogFileName(name string) {
	this.logFileName = name
}

func (this *LogWriter) LogMessageToFile(message, logFileName string) {
	var err error
	this.printWriter, err = os.OpenFile(logFileName, os.O_APPEND, 0)
	if err != nil {
		println("Can't open file in LogMessageToFile")
		return
	}
	defer this.printWriter.Close()

	this.printWriter.WriteString(" ---------------------------------------------- ")
	this.printWriter.WriteString(message)
}

func (this *LogWriter) checkLogFile() {
	if this.printWriter != nil {
		return
	}
	if this.logFileName == "" {
		return
	}

	var err error
	this.printWriter, err = os.OpenFile(this.logFileName, os.O_APPEND, 0)
	if err != nil {
		println("Can't open file in checkLogFile")
		return
	}
}

func (this *LogWriter) println(message string) {
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
func (this *LogWriter) LogMessage(message string) {
	if !this.needsLogging {
		return
	}

	this.checkLogFile()
	this.println(message)
}

/** Set the trace level for the stack.
 */
func (this *LogWriter) SetTraceLevel(level int) {
	this.traceLevel = level
}

/** Get the trace level for the stack.
 */
func (this *LogWriter) GetTraceLevel() int {
	return this.traceLevel
}
