package core

import (
	"strings"
)

type Token struct {
      tokenValue	string;
      tokenType	int;
}
      func (this *Token) GetTokenValue() string { 
      	return this.tokenValue; 
      }
      func (this *Token) GetTokenType() int{ 
      	return this.tokenType; 
      }
      func (this *Token) ToString() string {
		return "tokenValue = " + tokenValue + 
			   "/tokenType = " + tokenType;
      }

/** Base string token splitter.
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
*/

type StringTokenizer string {
	
	buffer	string;
	ptr	int;
	savedPtr	int;
}
        
        
        
        func NewStringTokenizer (buffer string) *StringTokenizer {
			this := &StringTokenizer{}
			this.buffer = buffer;
			this.ptr = 0;
			
			return this;
		}
        
        func (this *StringTokenizer) NextToken() string {
            var retval string;
          
            for this.ptr < len(this.buffer)  {
                if buffer[ptr] == '\n'  {
                    retval += buffer[this.ptr:this.ptr+1];
                    ptr++;
                    break;
                }else{
                    retval += buffer[this.ptr:this.ptr+1];
                    ptr++;
                }
            }
          
            return retval;
        }
        
        
	func (this *StringTokenizer) HasMoreChars() bool {
		return this.ptr < len(this.buffer);
	}

        func IsHexDigit(ch byte) bool{
            if IsDigit(ch){ 
            	return true;
            }else {
                ch1 := strings.ToUpper(string(ch))[0];
                return  ch1 == 'A' || ch1 == 'B' || ch1 == 'C' ||
                		ch1 == 'D' || ch1 == 'E' || ch1 == 'F';
            }
        }
            
                

	func IsAlpha(ch byte) bool{
            //boolean retval = strings..isUpperCase(ch) ||
            //    Character.isLowerCase(ch);
	   // Debug.println("isAlpha is returning " + retval  + " for " + ch);
	    ch1 := strings.ToUpper(string(ch))[0];
	    return ch1 >= 'A' && ch1 <= 'Z';
	}

	func IsDigit(ch byte) bool{
            //boolean retval =  Character.isDigit(ch);
	    // Debug.println("isDigit is returning " + retval + " for " + ch);
	    return 	ch == '0' || ch == '1' || ch == '2' || ch == '3' || ch == '4' || 
	    		ch == '5' || ch == '6' || ch == '7' || ch == '8' || ch == '9';
	}

	

	func (this *StringTokenizer) GetLine() string {
		var retval string;//= new StringBuffer();
		for this.ptr < len(this.buffer) && buffer[this.ptr] != '\n' {
		    retval += this.buffer[this.ptr:this.ptr+1];
		    this.ptr++;
		}
        if this.ptr < len(this.buffer) && buffer[this.ptr]s == '\n' {
            retval += "\n";
            this.ptr++;
        }  
		return retval;
	}

	func (this *StringTokenizer) PeekLine() string {
	    curPos := this.ptr;
	    retval := this.GetLine();
		this.ptr = curPos;
		return retval;
	}

	func (this *StringTokenizer) LookAhead() byte {//throws ParseException {
		return this.LookAheadK(0);
	}
	
	func (this *StringTokenizer) LookAheadK( k int) byte{//throws ParseException  {
		// Debug.out.println("ptr = " + ptr);
		if this.ptr+k < len(this.buffer) {
			return buffer[this.ptr + k];
		}
		return '\0';
	}

	func (this *StringTokenizer) GetNextChar()  byte{//throws ParseException {
	   if (ptr >= buffer.length()) 
	       throw new ParseException
		(buffer + " getNextChar: End of buffer",ptr);
	    else return buffer.charAt(ptr++);
	}

	func (this *StringTokenizer) Consume() {
		ptr = savedPtr;
	}

	func (this *StringTokenizer) ConsumeK(int k) {
		ptr += k;
	}

        /** Get a Vector of the buffer tokenized by lines
         */ 
        func (this *StringTokenizer) GetLines() map[int]string{
            //Vector result=new Vector();
            result :=make(map[int]string);
            for this.HasMoreChars() {
                line:=this.GetLine();
                result[len(result)]=line;
            }
            return result;
        }
        

	/** Get the next token from the buffer.
	*/
    func (this *StringTokenizer) GetNextToken(delim byte) string{//throws ParseException {
        var retval string;// new StringBuffer();
        for {
	    	la := this.LookAheadK(0);
	    	// System.out.println("la = " + la);
	    	if la == delim {
	    		break;
	    	}else if la == '\0'{
	    	 	return "";//throw new ParseException( "EOL reached", 0);
            }
            retval += this.buffer[this.ptr:this.ptr+1];
            this.ConsumeK(1);
        }
        return retval;
    }

        /** get the SDP field name of the line
         *  @return String
         */
        func GetSDPFieldName(line string) string{
           if line=="" {
           	return "";
           }
           
           //var fieldName string;
           //try{
                begin:=strings.Index(line, "=");
                if begin!=-1 {
                	return line[0,begin];
                }
           //}
           //catch(IndexOutOfBoundsException e) {
           //    return null;
           //}
           return "";
        }

