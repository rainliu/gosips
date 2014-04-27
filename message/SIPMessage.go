package message

import (
	"bytes"
	"container/list"
	"gosip/core"
	"gosip/header"
	"strings"
)

/**
 * This is the main SIP Message structure.
 *
 * @see StringMsgParser
 * @see PipelinedMsgParser
 *
 *
 *@version  JAIN-SIP-1.1
 *
 *@author M. Ranganathan <mranga@nist.gov>  <br/>
 *
 *<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
 *
 */

type SIPMessage struct { // MessageObject
	//implements javax.sip.message.Message {

	//protected static final String header.SIPConstants_DEFAULT_ENCODING = "UTF-8";

	/** unparsed headers
	 */
	unrecognizedHeaders *list.List
	/** List of parsed headers (in the order they were added)
	 */
	headers *list.List

	/** Direct accessors for frequently accessed headers  **/
	fromHeader          *header.From
	toHeader            *header.To
	cSeqHeader          *header.CSeq
	callIdHeader        *header.CallID
	contentLengthHeader *header.ContentLength
	maxForwardsHeader   *header.MaxForwards

	// Payload
	messageContent       string
	messageContentBytes  []byte
	messageContentObject interface{}

	// Table of headers indexed by name.
	nameTable map[string]header.ISIPHeader
}

//    /**
//     *
//     * Constructor: Initializes lists and list headers.
//     * All the headers for which there can be multiple occurances in
//     * a message are  derived from the SIPHeaderListClass. All singleton
//     * headers are derived from SIPHeader class.
//     *
//     */
func NewSIPMessage() *SIPMessage {
	this := &SIPMessage{}
	this.unrecognizedHeaders = list.New()
	this.headers = list.New()
	this.nameTable = make(map[string]header.ISIPHeader)
	//try {
	this.AttachHeader2(header.NewContentLengthFromInt(0), false)
	//} catch (Exception ex) {}
	return this
}

/** Return true if the header belongs only in a Request.
 *
 *@param sipHeader is the header to test.
 */
func (this *SIPMessage) IsRequestHeader(sipHeader header.ISIPHeader) bool {
	var ok bool
	if _, ok = sipHeader.(*header.AlertInfo); ok {
		return true
	}
	if _, ok = sipHeader.(*header.InReplyTo); ok {
		return true
	}
	if _, ok = sipHeader.(*header.Authorization); ok {
		return true
	}
	if _, ok = sipHeader.(*header.MaxForwards); ok {
		return true
	}
	if _, ok = sipHeader.(*header.Priority); ok {
		return true
	}
	if _, ok = sipHeader.(*header.ProxyAuthorization); ok {
		return true
	}
	if _, ok = sipHeader.(*header.ProxyRequire); ok {
		return true
	}
	if _, ok = sipHeader.(*header.ProxyRequireList); ok {
		return true
	}
	if _, ok = sipHeader.(*header.Route); ok {
		return true
	}
	if _, ok = sipHeader.(*header.RouteList); ok {
		return true
	}
	if _, ok = sipHeader.(*header.Subject); ok {
		return true
	}

	return false
}

//    /** Return true if the header belongs only in a response.
//     *
//     *@param sipHeader is the header to test.
//     */
func (this *SIPMessage) IsResponseHeader(sipHeader header.ISIPHeader) bool {
	var ok bool
	if _, ok = sipHeader.(*header.ErrorInfo); ok {
		return true
	}
	if _, ok = sipHeader.(*header.ProxyAuthenticate); ok {
		return true
	}
	if _, ok = sipHeader.(*header.Server); ok {
		return true
	}
	if _, ok = sipHeader.(*header.Unsupported); ok {
		return true
	}
	if _, ok = sipHeader.(*header.RetryAfter); ok {
		return true
	}
	if _, ok = sipHeader.(*header.Warning); ok {
		return true
	}
	if _, ok = sipHeader.(*header.WWWAuthenticate); ok {
		return true
	}

	return false
}

//    /** Get the headers as a linked list of encoded Strings
//     *@return a linked list with each element of the list containing a
//     * string encoded header in canonical form.
//     */
func (this *SIPMessage) GetMessageAsEncodedStrings() *list.List {
	retval := list.New()
	//synchronized (headers) {

	for li := this.headers.Front(); li != nil; li = li.Next() {
		sipHeader := li.Value.(header.ISIPHeader)
		if shl, ok := sipHeader.(header.ISIPHeaderList); ok {
			//SIPHeaderList shl = (SIPHeaderList) sipHeader;
			retval.PushBackList(shl.GetHeadersAsEncodedStrings())
		} else {
			retval.PushBack(sipHeader.String())
		}
	}
	//}
	return retval
}

//    /** Get A dialog identifier constructed from this messsage.
//     * This is an id that can be used to identify dialogs.
//     *@param isServerTransaction is a flag that indicates whether this is
//     * a server transaction.
//     */
//    public abstract String GetDialogId(boolean isServerTransaction);

//    /**
//     * Template match for SIP messages.
//     * The matchObj is a SIPMessage template to match against.
//     * This method allows you to do pattern matching with incoming
//     * SIP messages.
//     * Null matches wild card.
//     *@param other is the match template to match against.
//     *@return true if a match occured and false otherwise.
//     */
//    public boolean match(Object other) {
//        if (other == nil) return true;
//        if (! other.GetClass().equals(this.GetClass())) return false;
//        SIPMessage matchObj  = (SIPMessage) other;
//        LinkedList matchHeaders = matchObj.headers;
//        ListIterator li = matchHeaders.listIterator();
//        topOfLoop:
//            while(li.hasNext()) {
//                SIPHeader hisHeaders = (SIPHeader) li.next();
//                LinkedList  myHeaders = this.GetHeaderList
//                (hisHeaders.GetHeaderName());

//                // Could not find a header to match his header.
//                if (myHeaders == nil || myHeaders.size() == 0)
//                    return false;

//                if (hisHeaders instanceof SIPHeaderList) {
//                    ListIterator outerIterator =
//                    ((SIPHeaderList) hisHeaders).listIterator();
//                    outer:
//                        while(outerIterator.hasNext()) {
//                            SIPHeader hisHeader = (SIPHeader)
//                            outerIterator.next();
//                            ListIterator
//                            innerIterator = myHeaders.listIterator();
//                            while (innerIterator.hasNext()) {
//                                SIPHeader myHeader =
//                                (SIPHeader)innerIterator.next();
//                                if (myHeader.match(hisHeader)) 	 {
//                                    break  outer;
//                                }
//                            }
//                            return false;
//                        }
//                } else  {
//                    SIPHeader hisHeader = hisHeaders;
//                    ListIterator
//                    innerIterator = myHeaders.listIterator();
//                    while (innerIterator.hasNext()) {
//                        SIPHeader myHeader =
//                        (SIPHeader) innerIterator.next();
//                        if (myHeader.match(hisHeader))
//                            break topOfLoop;
//                    }
//                    // No match found.
//                    return false;
//                }
//            }
//            return true;

//    }

//    /**
//     * Recursively replace a portion of this object with a new Object.
//     * You cannot use this function for replacing sipheaders in
//     * a message (for that, use the remove and attach functions).
//     * Its intended use is for global find and replace of poritons of
//     * headers such as addresses.
//     * @param cText canonical representation of object that has to be
//     * 	replaced.
//     * @param newObject object that replaces the object that has the
//     * 	text cText
//     * @param matchSubstring if true then if cText is a substring of the
//     * encoded text of the Object then a match is flagged.
//     * @exception IllegalArgumentException on nil args and if
//     * replacementObject does not derive from GenericObject or
//     * GenericObjectList
//     */
//    public void replace(String cText, GenericObject newObject,
//    boolean matchSubstring )
//    throws IllegalArgumentException {
//        SIPHeader siphdr;
//        if (cText == nil || newObject == nil) {
//            throw new IllegalArgumentException("nil arguments");
//        }
//        if (SIPHeader.class.isAssignableFrom(newObject.GetClass()))  {
//            throw new IllegalArgumentException
//            ("Cannot replace object of class" + newObject.GetClass());
//        } else if (SIPHeaderList.class.
//        isAssignableFrom(newObject.GetClass()))  {
//            throw new IllegalArgumentException
//            ("Cannot replace object of class " + newObject.GetClass());
//        } else {
//            // not a sipheader or a sipheaderlist so do a find and replace.
//            synchronized (this.headers) {
//                // Concurrent modification exception noticed by Lamine Brahimi
//                ListIterator li = this.headers.listIterator();
//                while (li.hasNext()) {
//                    siphdr = (SIPHeader) li.next();
//                    siphdr.replace(cText,newObject,matchSubstring);
//                }
//            }
//        }
//    }

//    /**
//     * Recursively replace a portion of this object with a new  Object.
//     * You cannot use this function for replacing sipheaders in
//     * a message (for that, use the remove and attach functions).
//     * Its intended use is for global find and replace of poritons of
//     * headers such as addresses.
//     * @param cText canonical representation of object that has to be
//     * 	replaced.
//     * @param newObject object that replaces the object that has the
//     * 	text cText
//     * @param matchSubstring if true then flag a match if cText is a
//     * substring of the encoded text of the object.
//     * @exception IllegalArgumentException on nil args and if
//     *  replacementObject does not derive from GenericObject or
//     *  GenericObjectList
//     */
//    public void replace(String cText, GenericObjectList newObject,
//    boolean matchSubstring )
//    throws IllegalArgumentException {
//        SIPHeader siphdr;
//        if (cText == nil || newObject == nil) {
//            throw new IllegalArgumentException("nil arguments");
//        }
//        if (SIPHeaderList.class.isAssignableFrom(newObject.GetClass()))  {
//            throw new IllegalArgumentException
//            ("Cannot replace object of class " + newObject.GetClass());
//        } else if (SIPHeader.class.isAssignableFrom(newObject.GetClass()))  {
//            throw new IllegalArgumentException
//            ("Cannot replace object of class " + newObject.GetClass());
//        } else {
//            synchronized(this.headers) {
//                // not a sipheader.
//                ListIterator li = this.headers.listIterator();
//                while (li.hasNext()) {
//                    siphdr = (SIPHeader) li.next();
//                    siphdr.replace(cText,newObject,matchSubstring);
//                }
//            }
//        }
//    }

//    /**
//     * Merge a request with a template
//     */
//    public void merge(Object template) {
//        if (! template.GetClass().equals(this.GetClass()))
//            throw new IllegalArgumentException("Bad class " +
//            template.GetClass());
//        SIPMessage templateMessage = (SIPMessage) template;
//        Object[] templateHeaders = templateMessage.headers.toArray();
//        for (int i = 0; i < templateHeaders.length; i++) {
//            SIPHeader hdr = (SIPHeader) templateHeaders[i];
//            String hdrName = hdr.GetHeaderName();
//            LinkedList myHdrs = this.GetHeaderList(hdrName);
//            if (myHdrs == nil) {
//                this.attachHeader(hdr);
//            } else {
//                ListIterator it = myHdrs.listIterator();
//                while (it.hasNext()) {
//                    SIPHeader sipHdr = (SIPHeader) it.next();
//                    sipHdr.merge(hdr);
//                }
//            }
//        }

//    }

//    /**
//     * Encode this message as a string. This is more efficient when
//     * the payload is a string (rather than a binary array of bytes).
//     * If the payload cannot be encoded as a UTF-8 string then it is
//     * simply ignored (will not appear in the encoded message).
//     * @return The Canonical String representation of the message
//     * (including the canonical string representation of
//     * the SDP payload if it exists).
//     */
func (this *SIPMessage) String() string {
	var encoding bytes.Buffer
	// Synchronization added because of concurrent modification exception
	// noticed by Lamine Brahimi.
	//synchronized(this.headers) {
	for it := this.headers.Front(); it != nil; it = it.Next() {
		siphdr := it.Value.(header.ISIPHeader)
		if _, ok := siphdr.(*header.ContentLength); !ok {
			encoding.WriteString(siphdr.String())
		}
	}
	//}

	encoding.WriteString(this.contentLengthHeader.String() + core.SIPSeparatorNames_NEWLINE)

	if this.messageContentObject != nil {
		mbody := this.GetContent() //.String()

		// encoding.append(SIPHeaderNames.CONTENT_LENGTH + COLON +
		// SP + mbody.length() + NEWLINE);
		// encoding.append(NEWLINE);

		encoding.WriteString(mbody)
	} else if this.messageContent != "" || this.messageContentBytes != nil {
		var content string
		//try {
		if this.messageContent != "" {
			content = this.messageContent
		} else {
			content = string(this.messageContentBytes) + header.SIPConstants_DEFAULT_ENCODING
		}
		// } catch (UnsupportedEncodingException ex) {
		//     content = "";
		// }
		// Add the content-length header
		// encoding.append(SIPHeaderNames.CONTENT_LENGTH + COLON +
		// SP + content.length() + NEWLINE);
		// Append the content
		//encoding.append(NEWLINE);

		encoding.WriteString(content)
	} else {
		// Message content does not exist.
		// encoding.append(SIPHeaderNames.CONTENT_LENGTH + COLON +
		// SP + '0' + NEWLINE);
		//encoding.append(NEWLINE);
	}
	return encoding.String()
}

//    /**
//     * Encode the message as a byte array.
//     * Use this when the message payload is a binary byte array.
//     *
//     * @return The Canonical byte array representation of the message
//     * (including the canonical byte array representation of
//     * the SDP payload if it exists all in one contiguous byte array).
//     *
//     */
func (this *SIPMessage) EncodeAsBytes() []byte {
	//      var encoding bytes.Buffer
	//      ListIterator it = this.headers.listIterator();

	//      while (it.hasNext())  {
	//          SIPHeader siphdr = (SIPHeader) it.next();
	//          if (! (siphdr instanceof ContentLength)  )
	//              encoding.append(siphdr.encode());

	//      }
	// encoding.append(contentLengthHeader.encode()).append(NEWLINE);

	//      byte[] retval = nil;
	//      byte[] content = this.GetRawContent();
	//      if (content != nil) {
	//          // encoding.append(SIPHeaderNames.CONTENT_LENGTH +
	//          // COLON +
	//          // SP + content.length + NEWLINE);
	//          // encoding.append(NEWLINE);
	//          // Append the content

	//          byte[] msgarray = nil;
	//          try {
	//              msgarray = encoding.toString().GetBytes("UTF-8");
	//          } catch (UnsupportedEncodingException ex) {
	//              InternalErrorHandler.handleException(ex);
	//          }

	//          retval = new byte[msgarray.length + content.length];
	//          System.arraycopy(msgarray,0,retval,0,msgarray.length);
	//          System.arraycopy(content,0,retval,msgarray.
	//          length,content.length);
	//      } else  {
	//          // Message content does not exist.
	//          // encoding.append(SIPHeaderNames.CONTENT_LENGTH +
	//          // COLON + SP + '0' + NEWLINE);
	//          //encoding.append(NEWLINE);

	//          try {
	//              retval = encoding.toString().GetBytes("UTF-8");
	//          } catch (UnsupportedEncodingException ex) {
	//              InternalErrorHandler.handleException(ex);
	//          }
	//      }
	return []byte(this.String())
}

//    /**
//     * clone this message (create a new deep physical copy).
//     * All headers in the message are cloned.
//     * You can modify the cloned copy without affecting
//     * the original. The content is handled as follows:
//     * If the content is a String, or a byte array, a
//     * new copy of the content is allocated and copied over. If the
//     * content is an Object that supports the clone method, then the
//     * clone method is invoked and the cloned content is the new content.
//     * Otherwise, the content of the new message is Set equal to nil.
//     *
//     * @return A cloned copy of this object.
//     */
//    public Object clone() {
//        SIPMessage retval = nil;
//        try {
//            retval = (SIPMessage) this.GetClass().newInstance();
//        } catch ( IllegalAccessException ex) {
//            InternalErrorHandler.handleException(ex);
//        } catch (InstantiationException ex) {
//            InternalErrorHandler.handleException(ex);
//        }
//        ListIterator li = headers.listIterator();
//        while(li.hasNext()) {
//            SIPHeader sipHeader = (SIPHeader)((SIPHeader) li.next()).clone();
//            retval.attachHeader(sipHeader);
//        }
//        if (retval instanceof SIPRequest) {
//            SIPRequest thisRequest = (SIPRequest) this;
//            RequestLine rl = (RequestLine)
//            (thisRequest.GetRequestLine()).clone();
//            ((SIPRequest) retval).SetRequestLine(rl);
//        } else {
//            SIPResponse thisResponse = (SIPResponse) this;
//            StatusLine sl = (StatusLine)
//            (thisResponse.GetStatusLine()).clone();
//            ((SIPResponse) retval).SetStatusLine(sl);
//        }

//        if (this.GetContent() != nil) {
//            try {
// 	Object newContent  = nil;
// 	Object currentContent = this.GetContent();
// 	// Check the type of the returned content.
// 	if (currentContent instanceof String ) {
// 	   // If it is a string allocate a new string for the body
// 	   newContent =  new String
// 			(currentContent.toString());
// 	} else if ( currentContent instanceof byte[] ) {
// 	    // If it is raw bytes allocate a new array of bytes
// 	    // and copy over the content.
// 	    int cl = ((byte[])currentContent).length;
// 	    byte[] nc = new byte[cl];
// 	    System.arraycopy((byte[])currentContent,0,nc,0,cl);
// 	    newContent = nc;
// 	} else {
// 	    // See if the object has a clone method that is public
// 	    // If so invoke the clone method for the new content.
// 	    Class cl = currentContent.GetClass();
// 	    try {
// 	      Method meth = cl.GetMethod("clone",nil);
// 	      if (Modifier.isPublic(meth.GetModifiers())) {
// 		  newContent = meth.invoke(currentContent,nil);
// 	      } else {
// 		  newContent = currentContent;
// 	      }
// 	    } catch (Exception ex) {
// 		newContent = nil;
// 	    }
// 	}
// 	if (newContent != nil) retval.SetContent
//                 	(newContent,this.GetContentTypeHeader());
//            } catch (ParseException ex) { /** Ignore **/ }
//        }

//        return retval;
//    }

//    /**
//     * Get the string representation of this header (for pretty printing the
//     * generated structure).
//     *
//     * @return Formatted string representation of the object. Note that
//     * 	this is NOT the same as encode(). This is used mainly for
//     *	debugging purposes.
//     */

//    public String debugDump() {
//        stringRepresentation = "";
//        sprint("SIPMessage:");
//        sprint("{");
//        try {

//            Field[] fields = this.GetClass().GetDeclaredFields();
//            for (int i = 0; i < fields.length; i++) {
//                Field f = fields [i];
//                Class fieldType = f.GetType();
//                String fieldName = f.GetName();
//                if (f.Get(this) != nil  &&
//                Class.forName(SIPHEADERS_PACKAGE + ".SIPHeader").
//                isAssignableFrom(fieldType) &&
//                fieldName.compareTo("headers") != 0 ) {
//                    sprint(fieldName + "=");
//                    sprint(((SIPHeader)f.Get(this)).debugDump());
//                }
//            }
//        } catch ( Exception ex ) {
//            InternalErrorHandler.handleException(ex);
//        }

//        sprint("List of headers : ");
//        sprint(headers.toString());
//        sprint("messageContent = ");
//        sprint("{");
//        sprint(messageContent);
//        sprint("}");
//        if (this.GetContent() != nil) {
//            sprint(this.GetContent().toString());
//        }
//        sprint("}");
//        return stringRepresentation;
//    }

//    /**
//     * Attach a header and die if you Get a duplicate header exception.
//     * @param h SIPHeader to attach.
//     */
func (this *SIPMessage) AttachHeader(h header.ISIPHeader) {
	//if h == nil) throw new IllegalArgumentException("nil header!");
	//try {
	if hl, ok := h.(header.ISIPHeaderList); ok {
		//SIPHeaderList hl = (SIPHeaderList) h;
		if hl.Len() == 0 {
			return
		}
	}
	this.AttachHeader3(h, false, false)
	// } catch ( SIPDuplicateHeaderException ex) {
	//     // InternalErrorHandler.handleException(ex);
	// }
}

//    /**
//     * Attach a header (replacing the original header).
//     * @param header SIPHeader that replaces a header of the same type.
//     */
func (this *SIPMessage) SetHeader(sipHeader header.Header) {
	h, _ := sipHeader.(header.ISIPHeader)
	// if (header == nil)
	//     throw new IllegalArgumentException("nil header!");
	// try {
	if hl, ok := sipHeader.(header.ISIPHeaderList); ok {
		//SIPHeaderList hl = (SIPHeaderList) header;
		// Ignore empty lists.
		if hl.Len() == 0 {
			return
		}
	}
	this.RemoveHeader(h.GetHeaderName())
	this.AttachHeader3(h, true, false)
	// } catch ( SIPDuplicateHeaderException ex) {
	//     InternalErrorHandler.handleException(ex);
	// }
}

//    /** Set a header from a linked list of headers.
//     *
//     *@param headers -- a list of headers to Set.
//     */

func (this *SIPMessage) SetHeaders(headers *list.List) {

	for listIterator := headers.Front(); listIterator != nil; listIterator = listIterator.Next() {
		sipHeader := listIterator.Value.(header.ISIPHeader)
		//try {
		this.AttachHeader2(sipHeader, false)
		// } catch (SIPDuplicateHeaderException ex) {}
	}
}

//    /**
//     * Attach a header to the end of the existing headers in
//     * this SIPMessage structure.
//     * This is equivalent to the attachHeader(SIPHeader,replaceflag,false);
//     * which is the normal way in which headers are attached.
//     * This was added in support of JAIN-SIP.
//     *
//     * @since 1.0 (made this public)
//     * @param h header to attach.
//     * @param replaceflag if true then replace a header if it exists.
//     * @throws SIPDuplicateHeaderException If replaceFlag is false and
//     * only a singleton header is allowed (fpr example CSeq).
//     */
func (this *SIPMessage) AttachHeader2(h header.ISIPHeader, replaceflag bool) {
	//throws SIPDuplicateHeaderException {
	this.AttachHeader3(h, replaceflag, false)
}

//    /**
//     * Attach the header to the SIP Message structure at a specified
//     * position in its list of headers.
//     *
//     * @param header Header to attach.
//     * @param replaceFlag If true then replace the existing header.
//     * @param index Location in the header list to insert the header.
//     * @exception SIPDuplicateHeaderException if the header is of a type
//     * that cannot tolerate duplicates and one of this type already exists
//     * (e.g. CSeq header).
//     * @throws IndexOutOfBoundsException If the index specified is
//     * greater than the number of headers that are in this message.
//     */

func (this *SIPMessage) AttachHeader3(h header.ISIPHeader, replaceFlag, top bool) {
	//throws SIPDuplicateHeaderException {
	// if (header == nil) {
	//     throw new NullPointerException("nil header");
	// }

	//var h SIPHeader;

	// if (ListMap.hasList(sh)  &&
	// ! SIPHeaderList.class.isAssignableFrom(sh.GetClass())) {
	//     SIPHeaderList hdrList = ListMap.GetList(sh);
	//     hdrList.add(sh);
	//     h = hdrList;
	// } else {
	//     h = sh;
	// }

	if replaceFlag {
		delete(this.nameTable, strings.ToLower(h.GetName()))
	} else {
		if _, present := this.nameTable[strings.ToLower(h.GetName())]; present {
			if _, ok := h.(header.ISIPHeaderList); !ok {
				if cl, ok := h.(*header.ContentLength); ok {
					this.contentLengthHeader.SetContentLength(cl.GetContentLength())
				}
			}
			// Just ignore duplicate header.
			return
		}
	}

	originalHeader := this.GetHeader(h.GetName())

	// Delete the original sh from our list structure.
	if originalHeader != nil {
		for li := this.headers.Front(); li != nil; li = li.Next() {
			next := li.Value.(header.ISIPHeader)
			if next == originalHeader {
				this.headers.Remove(li)
			}
		}
	}

	if this.GetHeader(h.GetName()) == nil {
		this.nameTable[strings.ToLower(h.GetName())] = h
		this.headers.PushBack(h)
	} else {
		if hs, ok := h.(header.ISIPHeaderList); ok {
			hdrlist := this.nameTable[strings.ToLower(h.GetName())].(*header.SIPHeaderList)
			if hdrlist != nil {
				hdrlist.Concatenate(hs, top)
			} else {
				this.nameTable[strings.ToLower(h.GetName())] = h
			}
		} else {
			this.nameTable[strings.ToLower(h.GetName())] = h
		}
	}

	// Direct accessor fields for frequently accessed headers.
	if sh, ok := h.(*header.From); ok {
		this.fromHeader = sh
	} else if sh, ok := h.(*header.ContentLength); ok {
		this.contentLengthHeader = sh
	} else if sh, ok := h.(*header.To); ok {
		this.toHeader = sh
	} else if sh, ok := h.(*header.CSeq); ok {
		this.cSeqHeader = sh
	} else if sh, ok := h.(*header.CallID); ok {
		this.callIdHeader = sh
	} else if sh, ok := h.(*header.MaxForwards); ok {
		this.maxForwardsHeader = sh
	}

}

//    /** Remove a header given its name. If multiple headers of a given name
//     * are present then the top flag determines which end to remove headers
//     * from.
//     *
//     *@param headerName is the name of the header to remove.
//     *@param top -- flag that indicates which end of header list to process.
//     */
func (this *SIPMessage) RemoveHeader2(headerName string, top bool) {
	// System.out.println("removeHeader " + headerName);
	toRemove := this.nameTable[strings.ToLower(headerName)]

	// nothing to do then we are done.
	if toRemove == nil {
		return
	}

	if hdrList, ok := toRemove.(header.ISIPHeaderList); ok {
		if top {
			first := hdrList.Front()
			hdrList.Remove(first)
		} else {
			last := hdrList.Back()
			hdrList.Remove(last)
		}
		// Clean up empty list
		if hdrList.Len() == 0 {
			for li := this.headers.Front(); li != nil; li = li.Next() {
				sipHeader := li.Value.(header.ISIPHeader)
				if strings.ToLower(sipHeader.GetName()) == strings.ToLower(headerName) {
					this.headers.Remove(li)
				}
			}
		}
	} else {
		delete(this.nameTable, strings.ToLower(headerName))
		var ok bool
		if _, ok = toRemove.(*header.From); ok {
			this.fromHeader = nil
		} else if _, ok = toRemove.(*header.To); ok {
			this.toHeader = nil
		} else if _, ok = toRemove.(*header.CSeq); ok {
			this.cSeqHeader = nil
		} else if _, ok = toRemove.(*header.CallID); ok {
			this.callIdHeader = nil
		} else if _, ok = toRemove.(*header.MaxForwards); ok {
			this.maxForwardsHeader = nil
		} else if _, ok = toRemove.(*header.ContentLength); ok {
			this.contentLengthHeader = nil
		}

		for li := this.headers.Front(); li != nil; li = li.Next() {
			sipHeader := li.Value.(header.ISIPHeader)
			if strings.ToLower(sipHeader.GetName()) == strings.ToLower(headerName) {
				this.headers.Remove(li)
			}
		}
	}

}

//    /** Remove all headers given its name.
//     *
//     *@param headerName is the name of the header to remove.
//     */
func (this *SIPMessage) RemoveHeader(headerName string) {

	//if (headerName == nil) throw new NullPointerException("nil arg");
	toRemove := this.nameTable[strings.ToLower(headerName)]
	// nothing to do then we are done.
	if toRemove == nil {
		return
	}
	delete(this.nameTable, strings.ToLower(headerName))
	// Remove the fast accessor fields.
	var ok bool
	if _, ok = toRemove.(*header.From); ok {
		this.fromHeader = nil
	} else if _, ok = toRemove.(*header.To); ok {
		this.toHeader = nil
	} else if _, ok = toRemove.(*header.CSeq); ok {
		this.cSeqHeader = nil
	} else if _, ok = toRemove.(*header.CallID); ok {
		this.callIdHeader = nil
	} else if _, ok = toRemove.(*header.MaxForwards); ok {
		this.maxForwardsHeader = nil
	} else if _, ok = toRemove.(*header.ContentLength); ok {
		this.contentLengthHeader = nil
	}

	for li := this.headers.Front(); li != nil; li = li.Next() {
		sipHeader := li.Value.(header.ISIPHeader)
		if strings.ToLower(sipHeader.GetName()) == strings.ToLower(headerName) {
			this.headers.Remove(li)
		}
	}
}

//    /**
//     * Generate (compute) a transaction ID for this SIP message.
//     * @return A string containing the concatenation of various
//     * portions of the From,To,Via and RequestURI portions
//     * of this message as specified in RFC 2543:
//     * All responses to a request contain the same values in
//     * the Call-ID, CSeq, To, and From fields
//     * (with the possible addition of  a tag in the To field
//     * (section 10.43)). This allows responses to be matched with requests.
//     * Incorporates a bug fix  for a bug sent in by Gordon Ledgard of
//     * IPera for generating transactionIDs when no port is present in the
//     * via header.
//     * Incorporates a bug fix for a bug report sent in by Chris Mills
//     * of Nortel Networks (converts to lower case when returning the
//     * transaction identifier).
//     *
//     *@return a string that can be used as a transaction identifier
//     *  for this message. This can be used for matching responses and
//     *  requests (i.e. an outgoing request and its matching response have
//     *	the same computed transaction identifier).
//     */
// func (this *SIPMessage) GetTransactionId() string {
// 	var topVia *header.Via
// 	if this.GetViaHeaders().Len() > 0 {
// 		topVia = this.GetViaHeaders().Front().Value.(*header.Via)
// 	}
// 	// Have specified a branch Identifier so we can use it to identify
// 	// the transaction. BranchId is not case sensitive.
// 	// Branch Id prefix is not case sensitive.
// 	if topVia != nil && topVia.GetBranch() != "" &&
// 		strings.Contains(strings.ToUpper(topVia.GetBranch()),
// 			strings.ToUpper(header.SIPConstants_BRANCH_MAGIC_COOKIE)) {
// 		// Bis 09 compatible branch assignment algorithm.
// 		// implies that the branch id can be used as a transaction
// 		// identifier.
// 		return strings.ToLower(topVia.GetBranch())
// 	} else {
// 		// Old style client so construct the transaction identifier
// 		// from various fields of the request.
// 		var retval bytes.Buffer
// 		from := this.GetFrom()
// 		to := this.GetTo()
// 		hpFrom := from.GetUserAtHostPort()
// 		retval.WriteString(hpFrom + ":")
// 		if from.HasTag() {
// 			retval.WriteString(from.GetTag() + ":")
// 		}
// 		hpTo := to.GetUserAtHostPort()
// 		retval.WriteString(hpTo + ":")
// 		cid := this.callIdHeader.GetCallId()
// 		retval.WriteString(cid + ":")
// 		retval.WriteRune(rune(this.cSeqHeader.GetSequenceNumber()))
// 		retval.WriteString(":" + this.cSeqHeader.GetMethod())
// 		if topVia != nil {
// 			retval.WriteString(":" + topVia.GetSentBy().String())
// 			if !topVia.GetSentBy().HasPort() {
// 				retval.WriteString(":")
// 				retval.WriteRune(5060)
// 			}
// 		}
// 		hc := core.ToHexString([]byte(strings.ToLower(retval.String())))
// 		if len(hc) < 32 {
// 			return hc
// 		} else {
// 			return hc[len(hc)-32 : len(hc)]
// 		}
// 	}
// 	// Convert to lower case -- bug fix as a result of a bug report
// 	// from Chris Mills of Nortel Networks.
// }

//    /** Return true if this message has a body.
//     */
func (this *SIPMessage) HasContent() bool {
	return this.messageContent != "" || this.messageContentBytes != nil
}

//    /**Return an iterator for the list of headers in this message.
//     *@return an Iterator for the headers of this message.
//     */
func (this *SIPMessage) GetHeaders() header.IList {
	return this.headers
}

//    /** Get the first header of the given name.
//     *
//     *@return header -- the first header of the given name.
//     */
func (this *SIPMessage) GetHeader(headerName string) header.Header {
	// if headerName == nil) throw new NullPointerException("bad name");
	sipHeader := this.nameTable[strings.ToLower(headerName)]
	if sl, ok := sipHeader.(header.ISIPHeaderList); ok {
		return sl.Front().Value.(header.Header)
	} else {
		return sipHeader
	}
}

//    /**
//     * Get the contentType header (nil if one does not exist).
//     *@return contentType header
//     */
func (this *SIPMessage) GetContentTypeHeader() header.ContentTypeHeader {
	return this.GetHeader(core.SIPHeaderNames_CONTENT_TYPE).(header.ContentTypeHeader)
}

//    /** Get the from header.
//     *@return -- the from header.
//     */
func (this *SIPMessage) GetFrom() header.FromHeader {
	return this.fromHeader
}

//    /**
//     * Get the ErrorInfo list of headers (nil if one does not exist).
//     * @return List containing ErrorInfo headers.
//     */
func (this *SIPMessage) GetErrorInfoHeaders() *header.ErrorInfoList {
	return this.GetSIPHeaderList(core.SIPHeaderNames_ERROR_INFO).(*header.ErrorInfoList)
}

//    /**
//     * Get the Contact list of headers (nil if one does not exist).
//     * @return List containing Contact headers.
//     */
func (this *SIPMessage) GetContactHeaders() *header.ContactList {
	return this.GetSIPHeaderList(core.SIPHeaderNames_CONTACT).(*header.ContactList)
}

//    /**
//     * Get the Via list of headers (nil if one does not exist).
//     * @return List containing Via headers.
//     */
func (this *SIPMessage) GetViaHeaders() *header.ViaList {
	return this.GetSIPHeaderList(core.SIPHeaderNames_VIA).(*header.ViaList)
}

//    /** Get an iterator to the list of vial headers.
//     *@return a list iterator to the list of via headers.
//     * public ListIterator GetVia() {
//     * return this.viaHeaders.listIterator();
//     * }
//     */

//    /** Set A list of via headers.
//     *@param - a list of via headers to add.
//     */
func (this *SIPMessage) SetVia(viaList *header.ViaList) {
	vList := header.NewViaList()

	for it := viaList.Front(); it != nil; it = it.Next() {
		via := it.Value.(*header.Via)
		vList.PushBack(via)
	}
	this.SetHeader(vList)
}

//    /** Set the header given a list of headers.
//     *
//     *@param headerList a headerList to Set
//     */

func (this *SIPMessage) SetHeaderFromSIPHeaderList(sipHeaderList header.ISIPHeaderList) {
	this.SetHeader(sipHeaderList)
}

//    /** Get the topmost via header.
//     *@return the top most via header if one exists or nil if none exists.
//     */
func (this *SIPMessage) GetTopmostVia() *header.Via {
	if this.GetViaHeaders() == nil {
		return nil
	} else {
		return this.GetViaHeaders().Front().Value.(*header.Via)
	}
}

//    /**
//     * Get the CSeq list of header (nil if one does not exist).
//     * @return CSeq header
//     */
func (this *SIPMessage) GetCSeq() header.CSeqHeader {
	return this.cSeqHeader
}

//    /** Get the sequence number.
//     * @return the sequence number.
//     */
func (this *SIPMessage) GetCSeqNumber() int {
	return this.cSeqHeader.GetSequenceNumber()
}

//    /**
//     * Get the Authorization header (nil if one does not exist).
//     * @return Authorization header.
//     */
func (this *SIPMessage) GetAuthorization() *header.Authorization {
	return this.GetHeader(core.SIPHeaderNames_AUTHORIZATION).(*header.Authorization)
}

//    /**
//     * Get the MaxForwards header (nil if one does not exist).
//     * @return Max-Forwards header
//     */
func (this *SIPMessage) GetMaxForwards() header.MaxForwardsHeader {
	return this.maxForwardsHeader
}

//    /** Set the max forwards header.
//     *@param -- maxForwards is the MaxForwardsHeader to Set.
//     */
func (this *SIPMessage) SetMaxForwards(maxForwards header.MaxForwardsHeader) {
	this.SetHeader(maxForwards)
}

//    /**
//     * Get the MinExpires header.
//     * @return Min-Expires header
//     */
func (this *SIPMessage) GetMinExpires() *header.MinExpires {
	return this.GetHeader(core.SIPHeaderNames_MIN_EXPIRES).(*header.MinExpires)
}

//    /** Set the min expires header.
//     *
//     *@return the Min-Expires header.
//     */
func (this *SIPMessage) SetMinExpiresHeader(minExpires header.MinExpiresHeader) {
	this.SetHeader(minExpires)
}

//    /**
//     * Get the Organization header (nil if one does not exist).
//     * @return Orgnaization header.
//     */
func (this *SIPMessage) GetOrganizationHeader() *header.Organization {
	return this.GetHeader(core.SIPHeaderNames_ORGANIZATION).(*header.Organization)
}

//    /**
//     * Get the Priority header (nil if one does not exist).
//     * @return Priority header
//     */
func (this *SIPMessage) GetPriorityHeader() *header.Priority {
	return this.GetHeader(core.SIPHeaderNames_PRIORITY).(*header.Priority)
}

//    /**
//     * Get the ProxyAuthorization header (nil if one does not exist).
//     * @return List containing Proxy-Authorization headers.
//     */
func (this *SIPMessage) GetProxyAuthorizationHeader() *header.ProxyAuthorization {
	return this.GetHeader(core.SIPHeaderNames_PROXY_AUTHORIZATION).(*header.ProxyAuthorization)
}

//    /**
//     * Get the Route List of headers (nil if one does not exist).
//     * @return List containing Route headers
//     */
func (this *SIPMessage) GetRouteHeaders() *header.RouteList {
	return this.GetSIPHeaderList(core.SIPHeaderNames_ROUTE).(*header.RouteList)
}

//    /** Get the CallID header (nil if one does not exist)
//     *
//     * @return Call-ID header .
//     */
func (this *SIPMessage) GetCallId() header.CallIdHeader {
	return this.callIdHeader
}

//    /** Set the call id header.
//     *
//     *@param callid -- call idHeader (what else could it be?)
//     */
func (this *SIPMessage) SetCallId(callId header.CallIdHeader) {
	this.SetHeader(callId)
}

//    /** Get the CallID header (nil if one does not exist)
//     *
//     *@param callId -- the call identifier to be assigned to the call id header
//     */
func (this *SIPMessage) SetCallIdFromString(callId string) { //throws java.text.ParseException {
	if this.callIdHeader == nil {
		c, _ := header.NewCallID(callId)
		this.SetHeader(c)
	}
	this.callIdHeader.SetCallId(callId)
}

//    /**
//     * Get the call ID string.
//     * A conveniance function that returns the stuff following
//     * the header name for the call id header.
//     *
//     *@return the call identifier.
//     *
//     */
func (this *SIPMessage) GetCallIdentifier() string {
	return this.callIdHeader.GetCallId()
}

//    /**
//     * Get the RecordRoute header list (nil if one does not exist).
//     *
//     * @return Record-Route header
//     */
func (this *SIPMessage) GetRecordRouteHeaders() *header.RecordRouteList {
	return this.GetSIPHeaderList(core.SIPHeaderNames_RECORD_ROUTE).(*header.RecordRouteList)
}

//    /**
//     * Get the To header (nil if one does not exist).
//     * @return To header
//     */
func (this *SIPMessage) GetTo() header.ToHeader {
	return this.toHeader
}

func (this *SIPMessage) SetTo(to header.ToHeader) {
	this.SetHeader(to)
}

func (this *SIPMessage) SetFrom(from header.FromHeader) {
	this.SetHeader(from)
}

//    /**
//     * Get the ContentLength header (nil if one does not exist).
//     *
//     * @return content-length header.
//     */
func (this *SIPMessage) GetContentLength() header.ContentLengthHeader {
	return this.contentLengthHeader
}

//    /**
//     * Get the message body as a string.
//     *	If the message contains a content type header with a specified
//     *  charSet, and if the payload has been read as a byte array, then
//     *  it is returned encoded into this charSet.
//     *
//     * @return Message body (as a string)
//     * @throws UnsupportedEncodingException if the platform does not
//     *  support the charSet specified in the content type header.
//     *
//     */
func (this *SIPMessage) GetMessageContent() string {
	// throws UnsupportedEncodingException {
	if this.messageContent == "" && this.messageContentBytes == nil {
		return ""
	} else if this.messageContent == "" {
		//contentTypeHeader := this.nameTable[strings.ToLower(core.SIPHeaderNames_CONTENT_TYPE)].(*header.ContentType)
		//if contentTypeHeader != nil {
		// String charSet = contentTypeHeader.GetCharSet();
		// if (charSet != nil) {
		//     this.messageContent =
		//     new String(messageContentBytes,charSet);
		// } else {
		//     this.messageContent =
		//     new String(messageContentBytes,header.SIPConstants_DEFAULT_ENCODING);
		// }
		// } else this.messageContent =
		// new String(messageContentBytes,header.SIPConstants_DEFAULT_ENCODING);
		this.messageContent = string(this.messageContentBytes)
	}
	return this.messageContent
}

//    /**
//     * Get the message content as an array of bytes.
//     * If the payload has been read as a String then it is decoded using
//     * the charSet specified in the content type header if it exists.
//     * Otherwise, it is encoded using the default encoding which is
//     * UTF-8.
//     *
//     *@return an array of bytes that is the message payload.
//     *
//     */
// func (this *SIPMessage) GetRawContent() []byte {
// 	// try {
// 	if this.messageContent == "" &&
// 		this.messageContentBytes == nil &&
// 		this.messageContentObject == nil {
// 		return nil
// 	} else if this.messageContentObject != nil {
// 		messageContent := this.messageContentObject.(core.GenericObject).String()
// 		var messageContentBytes []byte
// 		contentTypeHeader := this.nameTable[strings.ToLower(core.SIPHeaderNames_CONTENT_TYPE)].(*header.ContentType)
// 		if contentTypeHeader != nil {
// 			charSet := contentTypeHeader.GetCharSet()
// 			if charSet != "" {
// 				messageContentBytes = messageContent.GetBytes(charSet)
// 			} else {
// 				messageContentBytes = messageContent.GetBytes(header.SIPConstants_DEFAULT_ENCODING)
// 			}
// 		} else {
// 			messageContentBytes = messageContent.GetBytes(header.SIPConstants_DEFAULT_ENCODING)
// 		}
// 		return messageContentBytes
// 	} else if this.messageContent != "" {
// 		var messageContentBytes []byte
// 		contentTypeHeader := this.nameTable[strings.ToLower(core.SIPHeaderNames_CONTENT_TYPE)].(*header.ContentType)
// 		if contentTypeHeader != nil {
// 			charSet := contentTypeHeader.GetCharSet()
// 			if charSet != "" {
// 				messageContentBytes = this.messageContent.GetBytes(charSet)
// 			} else {
// 				messageContentBytes = this.messageContent.GetBytes(header.SIPConstants_DEFAULT_ENCODING)
// 			}
// 		} else {
// 			messageContentBytes = this.messageContent.GetBytes(header.SIPConstants_DEFAULT_ENCODING)
// 		}
// 		return messageContentBytes
// 	} else {
// 		return this.messageContentBytes
// 	}
// 	// } catch (UnsupportedEncodingException ex) {
// 	//     InternalErrorHandler.handleException(ex);
// 	//     return nil;
// 	// }
// }

//    /** Set the message content given type and subtype.
//     *
//     *@param type is the message type (eg. application)
//     *@param subType is the message sybtype (eg. sdp)
//     *@param messageContent is the messge content as a string.
//     */

func (this *SIPMessage) SetMessageContentFromString(t string, subType string, messageContent string) {
	//if (messageContent == nil)
	//     throw new IllegalArgumentException("messgeContent is nil");
	ct := header.NewContentTypeFromString(t, subType)
	this.SetHeader(ct)
	this.messageContent = messageContent
	this.messageContentBytes = nil
	this.messageContentObject = nil
	//try {
	this.contentLengthHeader.SetContentLength(len(messageContent))
	//} catch (InvalidArgumentException ex) {}
}

//    /** Set the message content after converting the given object to a
//     * String.
//     *
//     *@param content -- content to Set.
//     *@param contentTypeHeader -- content type header corresponding to
//     *	content.
//     */
func (this *SIPMessage) SetContent(content interface{}, contentTypeHeader header.ContentTypeHeader) { //throws ParseException {
	//if content == nil) throw new NullPointerException("nil content");
	this.SetHeader(contentTypeHeader)
	length := -1
	if s, ok := content.(string); ok {
		this.messageContent = s
		length = len(s)
	} else if b, ok := content.([]byte); ok {
		this.messageContentBytes = b
		length = len(b)
	} else {
		this.messageContentObject = content
		length = len(content.(core.GenericObject).String())
	}

	//try {

	//           if (content instanceof String )
	//               length = ((String)content).length();
	//           else if (content instanceof byte[])
	//               length = ((byte[])content).length;
	//    else
	// length = content.toString().length();

	if length != -1 {
		this.contentLengthHeader.SetContentLength(length)
	}
	// } catch (InvalidArgumentException ex) {}

}

//    /** Get the content of the header.
//     *
//     *@return the content of the sip message.
//     */
func (this *SIPMessage) GetContent() string {
	if this.messageContentObject != nil {
		return this.messageContentObject.(string)
	} else if this.messageContentBytes != nil {
		return string(this.messageContentBytes)
	} else if this.messageContent != "" {
		return this.messageContent
	} else {
		return ""
	}
}

//    /** Set the message content for a given type and subtype.
//     *
//     *@param type is the messge type.
//     *@param subType is the message subType.
//     *@param messageContent is the message content as a byte array.
//     */
func (this *SIPMessage) SetMessageContent3(t string, subType string, messageContent []byte) {
	ct := header.NewContentTypeFromString(t, subType)
	this.SetHeader(ct)
	this.SetMessageContentFromByte(messageContent)
	//try {
	this.contentLengthHeader.SetContentLength(len(messageContent))
	//} catch (InvalidArgumentException ex) {}
}

//    /**
//     * Set the message content for this message.
//     *
//     * @param content Message body as a string.
//     */
func (this *SIPMessage) SetMessageContent(content string) {
	//int clength = (content == nil? 0: content.length());
	//try {
	this.contentLengthHeader.SetContentLength(len(content))
	// } catch (InvalidArgumentException ex) {}
	this.messageContent = content
	this.messageContentBytes = nil
	this.messageContentObject = nil
}

//    /** Set the message content as an array of bytes.
//     *
//     *@param content is the content of the message as an array of bytes.
//     */
func (this *SIPMessage) SetMessageContentFromByte(content []byte) {
	//try {
	this.contentLengthHeader.SetContentLength(len(content))
	//} catch (InvalidArgumentException ex) {}

	this.messageContentBytes = content
	this.messageContent = ""
	this.messageContentObject = nil
}

//    /** Remove the message content if it exists.
//     *
//     */
func (this *SIPMessage) RemoveContent() {
	this.messageContent = ""
	this.messageContentBytes = nil
	this.messageContentObject = nil
	//try {
	this.contentLengthHeader.SetContentLength(0)
	//} catch (InvalidArgumentException ex) {}
}

//    /** Get a SIP header or Header list given its name.
//     *@param headerName is the name of the header to Get.
//     *@return a header or header list that contians the retrieved header.
//     */
func (this *SIPMessage) GetHeadersFromString(headerName string) header.IList {
	// if (headerName == nil)
	//     throw new NullPointerException
	//     ("nil headerName");
	var sipHeader header.ISIPHeader
	var present bool
	if sipHeader, present = this.nameTable[strings.ToLower(headerName)].(header.ISIPHeader); !present {
		// empty iterator
		return list.New()
	}

	if shl, ok := sipHeader.(header.ISIPHeaderList); ok {
		return shl
	} else {
		l := list.New()
		l.PushBack(sipHeader)
		return l
	}
}

func (this *SIPMessage) GetSIPHeaderList(headerName string) header.ISIPHeaderList {
	return this.nameTable[strings.ToLower(headerName)].(header.ISIPHeaderList)
}

func (this *SIPMessage) GetHeaderList(headerName string) header.IList {
	sipHeader := this.nameTable[strings.ToLower(headerName)]
	if sipHeader == nil {
		return nil
	} else if shl, ok := sipHeader.(header.ISIPHeaderList); ok {
		return shl
	} else {
		ll := list.New()
		ll.PushBack(sipHeader)
		return ll
	}
}

//    /**
//     * Return true if the SIPMessage has a header of the given name.
//     *
//     *@param headerName is the header name for which we are testing.
//     *@return true if the header is present in the message
//     */

func (this *SIPMessage) HasHeader(headerName string) bool {
	_, present := this.nameTable[strings.ToLower(headerName)]
	return present
}

//    /**
//     * Return true if the message has a From header tag.
//     *
//     *@return true if the message has a from header and that header has
//     * 		a tag.
//     */
func (this *SIPMessage) HasFromTag() bool {
	return this.fromHeader != nil && this.fromHeader.GetTag() != ""
}

//    /**
//     * Return true if the message has a To header tag.
//     *
//     *@return true if the message has a to header and that header has
//     * 		a tag.
//     */
func (this *SIPMessage) HasToTag() bool {
	return this.toHeader != nil && this.toHeader.GetTag() != ""
}

//    /**
//     * Return the from tag.
//     *
//     *@return the tag from the from header.
//     *
//     */
func (this *SIPMessage) GetFromTag() string {
	if this.fromHeader == nil {
		return ""
	} else {
		return this.fromHeader.GetTag()
	}
}

//    /** Set the From Tag.
//     *
//     *@param tag -- tag to Set in the from header.
//     */
func (this *SIPMessage) SetFromTag(tag string) {
	if this.fromHeader != nil {
		this.fromHeader.SetTag(tag)
	}
	/*}
	  catch(ParseException e) {}*/
}

//    /** Set the to tag.
//     *
//     *@param tag -- tag to Set.
//     */
//     */
func (this *SIPMessage) SetToTag(tag string) {
	//try{
	if this.toHeader != nil {
		this.toHeader.SetTag(tag)
	}
	//}
	//catch(ParseException e) {}
}

//    /**
//     * Return the to tag.
//     */
func (this *SIPMessage) GetToTag() string {
	if this.toHeader == nil {
		return ""
	} else {
		return this.toHeader.GetTag()
	}
}

//    /**
//     * Return the encoded first line.
//     */
//    public abstract String GetFirstLine();

//    /** Add a SIP header.
//     *@param sipHeader -- sip header to add.
//     */
func (this *SIPMessage) AddHeader(sipHeader header.Header) {
	// Content length is never stored. Just computed.
	sh := sipHeader.(header.ISIPHeader)
	//try {
	if _, ok := sipHeader.(header.ViaHeader); ok {
		this.AttachHeader3(sh, false, true)
	} else {
		this.AttachHeader3(sh, false, false)
	}
	// } catch (SIPDuplicateHeaderException ex) {
	//try {
	//if cl, ok := sipHeader.(header.ContentLengthHeader); ok {
	//		contentLengthHeader.SetContentLength(cl.GetContentLength())
	//	}
	// } catch (InvalidArgumentException e) {}
	//}
}

//    /** Add a header to the unparsed list of headers.
//     *
//     *@param unparsed -- unparsed header to add to the list.
//     */
func (this *SIPMessage) AddUnparsed(unparsed string) {
	this.unrecognizedHeaders.PushBack(unparsed)
}

//    /** Add a SIP header.
//     *@param sipHeader -- string version of SIP header to add.
//     */

// func (this *SIPMessage) AddHeaderFromString( sipHeader string)  {
//      hdrString := strings.TrimSpace(sipHeader) + "\n";
//    // try {
//         HeaderParser parser =
//         ParserFactory.createParser(sipHeader);
//         SIPHeader sh = parser.parse();
//         this.attachHeader(sh,false);
//     // } catch (ParseException ex) {
//     //     this.unrecognizedHeaders.add(hdrString);
//     // }
// }

//    /** Get a list containing the unrecognized headers.
//     *@return a linked list containing unrecongnized headers.
//     */
func (this *SIPMessage) GetUnrecognizedHeaders() *list.List {
	return this.unrecognizedHeaders
}

//    /** Get the header names.
//     *
//     *@return a list iterator to a list of header names. These are ordered
//     * in the same order as are present in the message.
//     */
func (this *SIPMessage) GetHeaderNames() *list.List {
	return this.headers
	// ListIterator li = this.headers.listIterator();
	// LinkedList retval  = new LinkedList();
	// while (li.hasNext()) {
	//     SIPHeader sipHeader = (SIPHeader) li.next();
	//     String name = sipHeader.GetName();
	//     retval.add(name);
	// }
	// return retval.listIterator();
}

//   /** Compare for equality.
//    *
//    *@param other -- the other object to compare with.
//    *
//    */

//   public boolean equals(Object other) {
// if (!other.GetClass().equals(this.GetClass()))  {
// 	return false;
// }
// SIPMessage otherMessage = (SIPMessage) other;
// Collection values =  this.nameTable.values();
// Iterator it = values.iterator();
//        if (nameTable.size() != otherMessage.nameTable.size()) {
//  	return false;
// }

// while(it.hasNext()) {
//     SIPHeader mine = (SIPHeader) it.next();
//     SIPHeader his = (SIPHeader) (otherMessage.nameTable.Get
// 			(mine.GetName().toLowerCase()));
//     if (his == nil) {
// 	return false;
//     }
//     else if (! his.equals(mine))  {
// 	return false;
//     }
// }
// return true;
//    }

//    /** Get content disposition header or nil if no such header exists.
//     *
//     * @return the contentDisposition header
//     */
func (this *SIPMessage) GetContentDisposition() header.ContentDispositionHeader {
	return this.GetHeader(core.SIPHeaderNames_CONTENT_DISPOSITION).(header.ContentDispositionHeader)
}

//    /** Get the content encoding header.
//     *
//     *@return the contentEncoding header.
//     */
func (this *SIPMessage) GetContentEncoding() header.ContentEncodingHeader {
	return this.GetHeader(core.SIPHeaderNames_CONTENT_ENCODING).(header.ContentEncodingHeader)
}

//    /** Get the contentLanguage header.
//     *
//     *@return the content language header.
//     */
func (this *SIPMessage) GetContentLanguage() header.ContentLanguageHeader {
	return this.GetHeader(core.SIPHeaderNames_CONTENT_LANGUAGE).(header.ContentLanguageHeader)
}

//    /** Get the exipres header.
//     *
//     *@return the expires header or nil if one does not exist.
//     */
func (this *SIPMessage) GetExpires() header.ExpiresHeader {
	return this.GetHeader(core.SIPHeaderNames_EXPIRES).(header.ExpiresHeader)
}

//    /** Set the expiresHeader
//     *
//     *@param expiresHeader -- the expires header to Set.
//     */

func (this *SIPMessage) SetExpires(expiresHeader header.ExpiresHeader) {
	this.SetHeader(expiresHeader)
}

//    /** Set the content disposition header.
//     *
//     *@param contentDispositionHeader -- content disposition header.
//     */

func (this *SIPMessage) SetContentDisposition(contentDispositionHeader header.ContentDispositionHeader) {
	this.SetHeader(contentDispositionHeader)
}

func (this *SIPMessage) SetContentEncoding(contentEncodingHeader header.ContentEncodingHeader) {
	this.SetHeader(contentEncodingHeader)
}

func (this *SIPMessage) SetContentLanguage(contentLanguageHeader header.ContentLanguageHeader) {
	this.SetHeader(contentLanguageHeader)
}

//    /** Set the content length header.
//     *
//     *@param contentLength -- content length header.
//     */
func (this *SIPMessage) SetContentLength(contentLength header.ContentLengthHeader) {
	//try {
	this.contentLengthHeader.SetContentLength(contentLength.GetContentLength())
	//} catch (InvalidArgumentException ex) {}

}

//    /** Set the CSeq header.
//     *
//     *@param cseqHeader -- CSeq Header.
//     */

func (this *SIPMessage) SetCSeq(cseqHeader header.CSeqHeader) {
	this.SetHeader(cseqHeader)
}

//    //public abstract void SetSIPVersion(String sipVersion) throws ParseException;

//    //public abstract String GetSIPVersion();

//    //public abstract String toString();
