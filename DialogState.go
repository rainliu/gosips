/**
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 * Module Name   : GoSIP Specification
 * File Name     : DialogState.go
 * Author        : Rain Liu
 *~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 */
 
package gosip

import (
	"errors"
)

/**
 * This class contains the enumerations that define the underlying state of an 
 * existing dialog. 
 *
 * There are four explicit states for a dialog, namely:
 * <ul>
 * <li> Early - A dialog is in the "early" state, which occurs when it is 
 * created when a provisional response is recieved to the INVITE Request.
 * <li> Confirmed - A dialog transitions to the "confirmed" state when a 2xx 
 * final response is received to the INVITE Request.
 * <li> Completed - A dialog transitions to the "completed" state when a BYE 
 * request is sent or received by the User Agent Client.
 * <li> Terminated - A dialog transitions to the "terminated" state when it is 
 * completed and ready for garbage collection.
 * </ul>
 */
 
const (
	_DIALOGSTATE_EARLY = iota 	//0
	_DIALOGSTATE_CONFIRMED		//1
	_DIALOGSTATE_COMPLETED		//2
	_DIALOGSTATE_TERMINATED		//3
)
 
const m_dialogStateSize = 4;

var m_dialogStateArray = []*DialogState{&DialogState{_DIALOGSTATE_EARLY},
										&DialogState{_DIALOGSTATE_CONFIRMED},
										&DialogState{_DIALOGSTATE_COMPLETED},
										&DialogState{_DIALOGSTATE_TERMINATED},
};


    /**
     * This constant value indicates that the dialog state is "Early".
     */        
    var DIALOGSTATE_EARLY = m_dialogStateArray[_DIALOGSTATE_EARLY];


    /**
     * This constant value indicates that the dialog state is "Confirmed".
     */        
    var DIALOGSTATE_CONFIRMED = m_dialogStateArray[_DIALOGSTATE_CONFIRMED];
   
    
    /**
     * This constant value indicates that the dialog state is "Completed".
     */        
    var DIALOGSTATE_COMPLETED = m_dialogStateArray[_DIALOGSTATE_COMPLETED];

    
    
    /**
     * This constant value indicates that the dialog state is "Terminated".
     */        
    var DIALOGSTATE_TERMINATED = m_dialogStateArray[_DIALOGSTATE_TERMINATED];    
    
 
 
type DialogState struct {//implements Serializable{
    // internal private variables
    m_dialogState int;
}

    /**
     * This method returns the object value of the DialogState
     *
     * @return  The DialogState Object
     * @param timeout The integer value of the DialogState
     */
    func GetDialogState(dialogState int) (*DialogState, error){
        if (dialogState >= 0 && dialogState < m_dialogStateSize) {
            return m_dialogStateArray[dialogState], nil;
        } 
        //else {
            return nil, errors.New("IllegalArgumentException: Invalid dialogState value");
        //}
    }

    /**
     * This method returns the integer value of the DialogState
     *
     * @return The integer value of the DialogState
     */
    func (this *DialogState) GetValue() int {
        return this.m_dialogState;
    }


    /**
     * Returns the designated type as an alternative object to be used when
     * writing an object to a stream.
     *
     * This method would be used when for example serializing DialogState.EARLY
     * and deserializing it afterwards results again in DialogState.EARLY.
     * If you do not implement readResolve(), you would not get
     * DialogState.EARLY but an instance with similar content.
     *
     * @return the DialogState
     * @exception ObjectStreamException
     */
    //private Object readResolve() throws ObjectStreamException {
    //    return m_dialogStateArray[m_dialogState];
    //}

    /**
     * This method returns a string version of this class.
     * @return The string version of the DialogState
     */
    func (this *DialogState) ToString() string {
        var text string;
        switch this.m_dialogState {
            case _DIALOGSTATE_EARLY:
                text = "Early Dialog";
                //break;
            case _DIALOGSTATE_CONFIRMED:
                text = "Confirmed Dialog";
                //break;
            case _DIALOGSTATE_COMPLETED:
                text = "Completed Dialog";
                //break;    
            case _DIALOGSTATE_TERMINATED:
                text = "Terminated Dialog";
                //break;                  
            default:
                text = "Error while printing Dialog State";
                //break;
        }
        return text;
    }
