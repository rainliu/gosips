package header

import (
	"errors"
	"gosips/core"
	"strconv"
	"strings"
)

/**
 *  CSeq SIP Header.
 */

type CSeq struct { //implements javax.sip.header.CSeqHeader{
	SIPHeader

	/** seqno field
	 */
	seqno int

	/** method field
	 */
	method string
}

/**
 *Constructor.
 */
/*public CSeq() {
    super(CSEQ);
}*/

/**
 * Constructor given the sequence number and method.
 *
 *@param seqno is the sequence number to assign.
 *@param method is the method string.
 */
func NewCSeq(seqno int, method string) *CSeq {
	this := &CSeq{}

	this.SIPHeader.super(core.SIPHeaderNames_CSEQ)

	this.seqno = seqno
	this.method = method

	return this
}

func (this *CSeq) super(seqno int, method string) {
	this.SIPHeader.super(core.SIPHeaderNames_CSEQ)

	this.seqno = seqno
	this.method = method
}

/**
 * Compare two cseq headers for equality.
 * @param other Object to compare against.
 * @return true if the two cseq headers are equals, false
 * otherwise.
 */
/*public boolean equals( Object other) {
	try {
        CSeq that = (CSeq) other;
        if (! this.seqno.equals(that.seqno)) {
            return false;
        }
        if (this.method.compareToIgnoreCase(that.method) != 0) {
            return false;
        }
        return true;
	} catch (ClassCastException ex) {
		return false;
	}
    }*/

/**
 * Return canonical encoded header.
 * @return String with canonical encoded header.
 */
func (this *CSeq) String() string {
	return this.headerName + core.SIPSeparatorNames_COLON + core.SIPSeparatorNames_SP + this.EncodeBody() + core.SIPSeparatorNames_NEWLINE
}

/**
 * Return canonical header content. (encoded header except headerName:)
 *
 * @return encoded string.
 */
func (this *CSeq) EncodeBody() string {
	return strconv.Itoa(this.seqno) + core.SIPSeparatorNames_SP + strings.ToUpper(this.method)
}

/**
 * Get the method.
 * @return String the method.
 */
func (this *CSeq) GetMethod() string {
	return this.method
}

/** Sets the sequence number of this CSeqHeader. The sequence number
 * MUST be expressible as a 32-bit unsigned integer and MUST be less than
 * 2**31.
 *
 * @param sequenceNumber - the sequence number to set.
 * @throws InvalidArgumentException -- if the seq number is <= 0
 */
func (this *CSeq) SetSequenceNumber(sequenceNumber int) (InvalidArgumentException error) {
	if sequenceNumber < 0 {
		return errors.New("InvalidArgumentException: GoSIP Exception, CSeq, setSequenceNumber(), the sequence number parameter is < 0")
	}

	this.seqno = sequenceNumber
	return nil
}

/**
 * Set the method member
 *
 * @param meth -- String to set
 */
func (this *CSeq) SetMethod(meth string) (ParseException error) {
	if meth == "" {
		return errors.New("NullPointerException:GoSIP Exception, CSeq, setMethod(), the meth parameter is null")
	}

	this.method = meth
	return nil
}

/** Gets the sequence number of this CSeqHeader.
 *
 * @return sequence number of the CSeqHeader
 */

func (this *CSeq) GetSequenceNumber() int {
	return this.seqno
}
