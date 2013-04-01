package core

import (
	"os"
)

/**
*  Log System Errors. Also used for debugging log.
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
*/
/** Dont trace
 */    
	const TRACE_NONE = 0;
/** Trace initialization code
 */        
/** Trace message processing
 */        
	const TRACE_MESSAGES = 16;
/** Trace exception processing
 */        
	const TRACE_EXCEPTION = 17;
/** Debug trace level (all tracing enabled).
 */        
	const TRACE_DEBUG = 32;
	/** trace level
	 */        
	 

/** Name of the log file in which the trace is written out
 * (default is /tmp/sipserverlog.txt)
 */     
	//const logFileName="debuglog.txt";
	
type LogWriterObject struct{
	
/** Print writer that is used to write out the log file.
 */        
	printWriter *os.File;

/** Flag to indicate that logging is enabled. This needs to be
* static and public in order to globally turn logging on or off.
* This is static for efficiency reasons (the java compiler will not
* generate the logging code if this is set to false).
*/
// protected     static int traceLevel = TRACE_DEBUG;
	traceLevel int;
	
	logFileName string;
	
	needsLogging bool;

	lineCount int;
}

var LogWriter = LogWriterObject{nil, TRACE_NONE, "debuglog.txt", false, 0}


	/** log a stack trace..
	*/
	/*public static void logStackTrace() {
		if (needsLogging) {
		   checkLogFile();
		   if (printWriter != null)  {
		      println("------------ Traceback ------");
		      logException(new Exception());
		      println("----------- End Traceback ------");
		   }
		}
	}*/

	
	/*public static void logException(Exception ex) {
	    if (needsLogging)  {
	      StringWriter sw = new StringWriter();
	      PrintWriter pw = new PrintWriter( sw);
	      checkLogFile();
	      if (printWriter != null) ex.printStackTrace(pw);
	      println(sw.toString());
	    }
	}*/
		

	/** Log an excption. 
	* 1.4x Code contributed by Brad Templeton
	*
	*@param sframe - frame to log grace.
	*/
	/*public static void logTrace(Throwable sframe) {
	    if (needsLogging)  {
	     checkLogFile();
             logException(new Exception(sframe.getMessage()) );
	   }
	}*/


	/** Set the log file name 
	*@param name is the name of the log file to set. 
	*/
	func (this *LogWriterObject) SetLogFileName(name string) {
		this.logFileName = name;
	}

	func (this *LogWriterObject) LogMessageToFile(message, logFileName string) {
		var err error
		//try {
			this.printWriter, err = os.OpenFile(logFileName, os.O_APPEND, 0);
			if err!=nil{
				println("Can't open file in LogMessageToFile");
				return;
			}
			defer this.printWriter.Close();
			
			this.printWriter.WriteString(" ---------------------------------------------- ");
			this.printWriter.WriteString(message);
		//} catch (IOException ex) {
		//	ex.printStackTrace();
		//}
	}
	
	func (this *LogWriterObject) checkLogFile() {
		if this.printWriter != nil {
			return;
		}
		if this.logFileName == "" {
			return;
		}
		//try {
			var err error
			this.printWriter, err = os.OpenFile(this.logFileName, os.O_APPEND, 0);
			if err!=nil{
				println("Can't open file in checkLogFile");
				return;
			}
		//} catch (IOException ex) {
		//	ex.printStackTrace();
		//}
	}
 
	func (this *LogWriterObject) println(message string) {
	    for i := 0; i < len(message); i++ {
			if message[i] == '\n'{ 
				this.lineCount ++;
			}
	    }
	    this.checkLogFile();
	    // String tname = Thread.currentThread().getName();
	    if this.printWriter != nil {
		   	this.printWriter.WriteString( message);	
	    }
	    this.lineCount++;
	}

	/** Log a message into the log file.
         * @param message message to log into the log file.
         */
	/*func (this *LogWriter) LogMessage(int level, String message) {
		if (! needsLogging) return;
	}*/
	/** Log a message into the log file.
         * @param message message to log into the log file.
         */
	func (this *LogWriterObject) LogMessage(message string) {
        if !this.needsLogging {
         	return;
		}
		
		this.checkLogFile();
	    this.println(message);	
	}
	
    
	
        /** Set the trace level for the stack.
         */
        func (this *LogWriterObject) SetTraceLevel(level int) {
            this.traceLevel = level;
        }
        
        /** Get the trace level for the stack.
         */
        func (this *LogWriterObject) GetTraceLevel() int { 
        	return this.traceLevel; 
        }
