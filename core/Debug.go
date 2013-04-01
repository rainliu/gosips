package core

import (
)

/**
*   A class to do debug printfs
*/

type DebugObject struct{
	debug bool;
	parserDebug bool;
}

var Debug = DebugObject{true,false};

		    
func (this *DebugObject) print (s string) {
	if this.debug {
		LogWriter.LogMessage(s);
	}
}
	
func (this *DebugObject) println (s string) {
	if this.debug {
		LogWriter.LogMessage(s+"\n");
	}
}

/*
func (this *DebugObject) printStackTrace(ex string) {
	if this.debug {
		log.Panic(ex);//LogWriter.logException(ex);
	}
}
	protected static void Abort(Exception e) {
	    System.out.println("Fatal error");
	     e.printStackTrace();
	     if (debug) {
		    LogWriter.logException(e);
	     }
	     System.exit(0);
	}

	protected static void Assert(boolean b) {
		if ( ! b) {
		   System.out.println("Assertion failure !");
		    new Exception().printStackTrace();
		    if (debug) LogWriter.logStackTrace();
		    System.exit(0);
		}
	}
         */

