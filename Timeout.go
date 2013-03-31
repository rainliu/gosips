/**
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 * Module Name   : GoSIP Specification
 * File Name     : Timeout.go
 * Author        : Rain Liu
 *~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 */
 
package gosip

import (
	"errors"
)


/**
 * This class contains the enumerations that define whether a timeout has 
 * occured in the underlying implementation. The application gets
 * informed on whether a retransmission or transaction timer has expired.
 *
 * There are two types of Timeout, namely:
 * <ul>
 * <li> {@link Timeout#RETRANSMIT} - This type is used to alert an application that
 * the underlying retransmit timer has expired so an application can
 * resend the message specific to a transaction. This timer is defaulted to
 * 500ms and is doubled each time it is fired until the transaction expires
 * {@link Timeout#TRANSACTION}. The default retransmission value can be changed
 * per transaction using {@link Transaction#setRetransmitTimer(int)}. The 
 * RETRANSMIT type is exposed to the following applications as follows:
 * <UL>
 * <li><b>User Agent</b> - Retransmissions on Invite transactions are the
 * responsibility of the application. This is due to the three way handshake
 * for an INVITE Request. All other retransmissions are handled by the underlying
 * implementation. Therefore the application will only receive this Timeout type 
 * specific to Invite transactions.
 * <li><b>User Agent with No Retransmission</b> - an application can configure an
 * implementation to handle all retransmissions using property characteristics  
 * of the {@link SipStack}. Therefore a Timeout
 * of this type will not be passed to the application. The application
 * will only get notified when the {@link Timeout#TRANSACTION} occurs.
 * <li><b>Stateful Proxy</b> - a stateful proxy remembers transaction state about
 * each incoming request and any requests it sends as a result of
 * processing the incoming request. The underlying implementation 
 * will handle the retransmissions of all messages it sends and the application 
 * will not receive {@link Timeout#RETRANSMIT} events, however the application 
 * will get notified of {@link Timeout#TRANSACTION} events. As an Invite 
 * transaction is end to end a stateful proxy will not handle the 
 * retransmissions of messages on an Invite transaction, unless it decides to
 * respond to the Invite transaction, in essence becoming an User Agent Server 
 * and as such should behave as described by the User Agent semantics above 
 * bearing in mind the retranmission property of the underlying implementation.
 * <li><b>Stateless Proxy</b> - as a stateless proxy acts as a simple forwarding
 * agent, i.e. it simply forwards every message it receives upstream, it
 * keeps no transaction state for messages. The implementation does not retransmit
 * messages, therefore an application will not receive {@link Timeout#RETRANSMIT}
 * events on a message handled statelessly. If retransmission semantics are
 * required by an application using a stateless method, it is the responsibility
 * of the application to provide this feature, not the underlying implementation.
 * </UL>
 * <li>{@link Timeout#TRANSACTION} - This type is used to alert an application 
 * that the underlying transaction has expired. A transaction timeout typically
 * occurs at a time 64*T1 were T1 is the initial value of the
 * {@link Timeout#RETRANSMIT}, usually defaulted to 500ms. The 
 * TRANSACTION type is exposed to the following applications as follows:
 * <UL>
 * <li><b>User Agent</b> - All retransmissions except retransmissions on Invite
 * transactions are the responsibility of the underlying implementation, i.e. 
 * Invite transactions are the responsibility of the application. Therefore the 
 * application will only recieve TRANSACTION Timeout events on transactions that 
 * are not Invite transactions.
 * <li><b>User Agent with No Retransmission</b> - an application can configure an
 * implementation to handle all retransmissions using property characteristics  
 * of the {@link SipStack}. Therefore a TRANSACTION Timeout will be fired to 
 * the application on any transaction that expires including an Invite 
 * transaction.
 * <li><b>Stateful Proxy</b> - a stateful proxy remembers transaction state about
 * each incoming request and any requests it sends as a result of
 * processing the incoming request. The underlying implementation handles
 * the retransmissions of all messages it sends and will notify the application
 * of {@link Timeout#TRANSACTION} events on any of its transactions. As an Invite
 * transaction is end to end a stateful proxy will not handle transaction
 * timeouts on an Invite transaction, unless it decides to respond to the Invite
 * transaction, in essence becoming an User Agent Server and as such should
 * behave as described by the User Agent semantics above bearing in mind
 * the retransmission property of the underlying implementation.
 * <li><b>Stateless Proxy</b> - as a stateless proxy acts as a simple forwarding
 * agent, i.e. it simply forwards every message it receives upstream, it
 * keeps no transaction state of the messages. The implementation does not 
 * maintain transaction state, therefore an application will not receive 
 * {@link Timeout#TRANSACTION} events on a message handled statelessly. 
 * If transaction timeout semantics are required by an application using a 
 * stateless method, it the responsibility of the application to provide this 
 * feature, not the underlying implementation.
 * </ul>
 * </ul>
 *
 * @author Sun Microsystems
 * @since 1.1
 */
 
const (
	_TIMEOUT_RETRANSMIT = iota 	//0
	_TIMEOUT_TRANSACTION		//1
)

const m_timeoutSize = 2;

var m_timeoutArray = []*Timeout{ &Timeout{_TIMEOUT_RETRANSMIT}, &Timeout{_TIMEOUT_TRANSACTION} }

    /**
     * This constant value indicates the "Retransmit" timeout.
     */ 
    var TIMEOUT_RETRANSMIT = m_timeoutArray[_TIMEOUT_RETRANSMIT];

    
    /**
     * This constant value indicates the "Transaction" timeout.
     */ 
    var TIMEOUT_TRANSACTION = m_timeoutArray[_TIMEOUT_TRANSACTION];
 
 
type Timeout struct{//implements Serializable{
    // internal private variables
    m_timeout int;
}
    /**
     * Constructor for the Timeout
     *
     * @param  timeout the integer value for the Timeout
     */
    /*private Timeout(int timeout) {
        m_timeout = timeout;
        m_timeoutArray[m_timeout] = this;
    }*/

    /**
     * This method returns the object value of the Timeout
     *
     * @return  The Timeout Object
     * @param timeout The integer value of the Timeout
     */
    func GetTimeout(timeout int) (*Timeout, error){
        if (timeout >= 0 && timeout < m_timeoutSize) {
            return m_timeoutArray[timeout], nil;
        } 
        //else {
            return nil, errors.New("IllegalArgumentException: Invalid timeout value");
        //}
    }

    /**
     * This method returns the integer value of the Timeout
     *
     * @return The integer value of the Timeout
     */
    func (this *Timeout) GetValue() int{
        return this.m_timeout;
    }

    /**
     * Returns the designated type as an alternative object to be used when
     * writing an object to a stream.
     *
     * This method would be used when for example serializing Timeout.RETRANSMIT
     * and deserializing it afterwards results again in Timeout.RETRANSMIT.
     * If you do not implement readResolve(), you would not get
     * Timeout.RETRANSMIT but an instance with similar content.
     *
     * @return the Timeout
     * @exception ObjectStreamException
     */
    /*private Object readResolve() throws ObjectStreamException {
        return m_timeoutArray[m_timeout];
    }*/

    /**
     * This method returns a string version of this class.
     * 
     * @return The string version of the Timeout
     */
    func (this *Timeout) ToString() string {
        var text string;
        switch this.m_timeout {
            case _TIMEOUT_RETRANSMIT:
                text = "Retransmission Timeout";
                //break;
            case _TIMEOUT_TRANSACTION:
                text = "Transaction Timeout";
                //break;
            default:
                text = "Error while printing Timeout";
                //break;
        }
        return text;
    }























