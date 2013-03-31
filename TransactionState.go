/**
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 * Module Name   : GoSIP Specification
 * File Name     : TransactionState.go
 * Author        : Rain Liu
 *~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 */

package gosip

import (
	"errors"
)

/**

 * This class contains the enumerations that define the underlying state of an 
 * existing transaction. SIP defines four types of 

 * transactions, these are Invite Client transactions, Invite Server transactions,

 * Non-Invite Client transactions and Non-Invite Server transactions.

 *

 * There are six explicit states for the various transactions, namely:

 * <ul>

 * <li> <b>Calling:</b> 

 * <ul>

 * <li> Invite Client transaction: The initial state, "calling", MUST be entered 

 * when the application initiates a new client transaction with an INVITE request.

 * </ul>

 * <li> <b>Trying:</b> 

 * <ul>

 * <li> Non-Invite Client transaction: The initial state "Trying" is entered 

 * when the application initiates a new client transaction with a request.  

 * <li> Non-Invite Server transaction: The initial state "Trying" is entered 

 * when the application is passed a request other than INVITE or ACK. 

 * </ul>

 * <li> <b>Proceeding:</b> 

 * <ul>

 * <li> Invite Client transaction: If the client transaction receives a 

 * provisional response while in the "Calling" state, it transitions to the 

 * "Proceeding" state. 

 * <li> Non-Invite Client transaction: If a provisional response is received 

 * while in the "Trying" state, the client transaction SHOULD move to the 

 * "Proceeding" state.  

 * <li> Invite Server transaction: When a server transaction is constructed for 

 * a request, it enters the initial state "Proceeding".  

 * <li> Non-Invite Server transaction: While in the "Trying" state, if the 

 * application passes a provisional response to the server transaction, the 

 * server transaction MUST enter the "Proceeding" state.

 * </ul>

 * <li> <b>Completed:</b> The "Completed" state exists to buffer any additional 

 * response retransmissions that may be received, which is why the client 

 * transaction remains there only for unreliable transports.

 * <ul>

 * <li> Invite Client transaction: When in either the "Calling" or "Proceeding" 

 * states, reception of a response with status code from 300-699 MUST cause the 

 * client transaction to transition to "Completed".

 * <li> Non-Invite Client transaction: If a final response (status codes 

 * 200-699) is received while in the "Trying" or "Proceeding" state, the client 

 * transaction MUST transition to the "Completed" state.

 * <li> Invite Server transaction: While in the "Proceeding" state, if the 

 * application passes a response with status code from 300 to 699 to the server 

 * transaction, the state machine MUST enter the "Completed" state. 

 * <li>Non-Invite Server transaction: If the application passes a final response 

 * (status codes 200-699) to the server while in the "Proceeding" state, the

 * transaction MUST enter the "Completed" state.

 * </ul>

 * <li> <b>Confirmed:</b> The purpose of the "Confirmed" state is to absorb any 

 * additional ACK messages that arrive, triggered from retransmissions of the 

 * final response. Once this time expires the server MUST transition to the 

 * "Terminated" state.

 * <ul>

 * <li> Invite Server transaction: If an ACK is received while the server 

 * transaction is in the "Completed" state, the server transaction MUST 

 * transition to the "Confirmed" state.

 * </ul>

 * <li> <b>Terminated:</b> The transaction MUST be available for garbage collection 

 * the instant it enters the "Terminated" state.

 * <ul>

 * <li> Invite Client transaction:  When in either the "Calling" or "Proceeding" 

 * states, reception of a 2xx response MUST cause the client transaction to 

 * enter the "Terminated" state. If amount of time that the server transaction 

 * can remain in the "Completed" state when unreliable transports are used 

 * expires while the client transaction is in the "Completed" state, the client 

 * transaction MUST move to the "Terminated" state. 

 * <li> Non-Invite Client transaction: If the transaction times out while the 

 * client transaction is still in the "Trying" or "Proceeding" state, the client 

 * transaction SHOULD inform the application about the timeout, and then it 

 * SHOULD enter the "Terminated" state. If the response retransmissions buffer 

 * expires while in the "Completed" state, the client transaction MUST transition 

 * to the "Terminated" state.

 * <li> Invite Server transaction: If in the "Proceeding" state, and the application 

 * passes a 2xx response to the server transaction, the server transaction MUST 

 * transition to the "Terminated" state. When the server transaction abandons 

 * retransmitting the response while in the "Completed" state, it implies that 

 * the ACK was never received.  In this case, the server transaction MUST 

 * transition to the "Terminated" state, and MUST indicate to the TU that a 

 * transaction failure has occurred.

 * <li> Non-Invite Server transaction: If the request retransmissions buffer 

 * expires while in the "Completed" state, the server transaction MUST transition 

 * to the "Terminated" state.

 * </ul>

 * </ul>

 * 

 * For each specific transaction state machine, refer to 

 * <a href = "http://www.ietf.org/rfc/rfc3261.txt">RFC3261</a>.

 */

const (
	_TRANSACTIONSTATE_CALLING = iota //0
	_TRANSACTIONSTATE_TRYING			//1
	_TRANSACTIONSTATE_PROCEEDING		//2
	_TRANSACTIONSTATE_COMPLETED		//3
	_TRANSACTIONSTATE_CONFIRMED		//4
	_TRANSACTIONSTATE_TERMINATED		//5
)

const m_transStateSize = 6

var m_transStateArray = []*TransactionState{&TransactionState{_TRANSACTIONSTATE_CALLING},
											&TransactionState{_TRANSACTIONSTATE_TRYING},
											&TransactionState{_TRANSACTIONSTATE_PROCEEDING},
											&TransactionState{_TRANSACTIONSTATE_COMPLETED},
											&TransactionState{_TRANSACTIONSTATE_CONFIRMED},
											&TransactionState{_TRANSACTIONSTATE_TERMINATED},
};    
    
    /**
     * This constant value indicates that the transaction state is "Calling".
     */    
	var TRANSACTIONSTATE_CALLING = m_transStateArray[_TRANSACTIONSTATE_CALLING];     
  
    /**
     * This constant value indicates that the transaction state is "Trying".
     */   
    var TRANSACTIONSTATE_TRYING = m_transStateArray[_TRANSACTIONSTATE_TRYING];   
 
    /**
     * This constant value indicates that the transaction state is "Proceeding".
     */        
    var TRANSACTIONSTATE_PROCEEDING = m_transStateArray[_TRANSACTIONSTATE_PROCEEDING];    

    /**
     * This constant value indicates that the transaction state is "Completed".
     */    
    var TRANSACTIONSTATE_COMPLETED = m_transStateArray[_TRANSACTIONSTATE_COMPLETED];    

    /**
     * This constant value indicates that the transaction state is "Confirmed".
     */    
    var TRANSACTIONSTATE_CONFIRMED = m_transStateArray[_TRANSACTIONSTATE_CONFIRMED];

    /**
     * This constant value indicates that the transaction state is "Terminated".
     */    
    var TRANSACTIONSTATE_TERMINATED = m_transStateArray[_TRANSACTIONSTATE_TERMINATED];


type TransactionState struct{ //implements Serializable{   
    // internal private variables
    m_transactionState int;
}


    /**

     * This method returns the object value of the TransactionState

     *

     * @return  The TransactionState Object

     * @param timeout The integer value of the TransactionState

     */

    func GetTransactionState(transactionState int) (*TransactionState, error){
        if (transactionState >= 0 && transactionState < m_transStateSize) {
            return m_transStateArray[transactionState], nil;
        } 
        //else {
            return nil, errors.New("IllegalArgumentException: Invalid transactionState value");
        //}
    }



    /**

     * This method returns the integer value of the TransactionState

     *

     * @return The integer value of the TransactionState

     */

    func (this *TransactionState) GetValue() int {
        return this.m_transactionState;
    }



    /**

     * Returns the designated type as an alternative object to be used when

     * writing an object to a stream.

     *

     * This method would be used when for example serializing TransactionState.EARLY

     * and deserializing it afterwards results again in TransactionState.EARLY.

     * If you do not implement readResolve(), you would not get

     * TransactionState.EARLY but an instance with similar content.

     *

     * @return the TransactionState

     * @exception ObjectStreamException

     */

    //func (this *TransactionState) ReadResolve() *TransactionState {
    //    return m_transStateArray[m_transactionState];
    //}



    /**

    
     * This method returns a string version of this class.
     * 
     * @return The string version of the TransactionState

     */

    func (this *TransactionState) ToString() string{

        var text string;

        switch this.m_transactionState {

            case _TRANSACTIONSTATE_CALLING:

                text = "Calling Transaction";

                //break;

            case _TRANSACTIONSTATE_TRYING:

                text = "Trying Transaction";

                //break;                

            case _TRANSACTIONSTATE_PROCEEDING:

                text = "Proceeding Transaction";

                //break;

            case _TRANSACTIONSTATE_COMPLETED:

                text = "Completed Transaction";

                //break;                 

            case _TRANSACTIONSTATE_CONFIRMED:

                text = "Confirmed Transaction";

                //break; 

            case _TRANSACTIONSTATE_TERMINATED:

                text = "Terminated Transaction";

                //break;                

            default:

                text = "Error while printing Transaction State";

                //break;

        }

        return text;

    }


