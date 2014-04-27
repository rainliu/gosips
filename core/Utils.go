package core

import "strings"

/**
* A few utilities that are used in various places by the stack.
* This is used to convert byte arrays to hex strings etc. Generate
* tags and branch identifiers and odds and ends.
 */
//public class Utils
//{
/**
 * to hex converter
 */
var toHex = []byte{'0', '1', '2', '3', '4', '5', '6',
	'7', '8', '9', 'a', 'b', 'c', 'd',
	'e', 'f'}

/**
 * convert an array of bytes to an hexadecimal string
 * @return a string
 * @param b bytes array to convert to a hexadecimal
 * string
 */

func ToHexString(b []byte) string {
	pos := 0
	c := make([]byte, len(b)*2)
	for i := 0; i < len(b); i++ {
		c[pos] = toHex[(b[i]>>4)&0x0F]
		pos++
		c[pos] = toHex[b[i]&0x0f]
		pos++
	}
	return string(c)
}

/**
 * Put quotes around a string and return it.
 *
 * @return a quoted string
 * @param str string to be quoted
 */
func GetQuotedString(str string) string {
	return "\"" + str + "\""
}

/**
 * Squeeze out all white space from a string and return the reduced string.
 *
 * @param input input string to sqeeze.
 * @return String a reduced string.
 */
func ReduceString(input string) string {
	newString := strings.ToLower(input)
	length := len(newString)
	var retval string
	for i := 0; i < length; i++ {
		if newString[i] == ' ' || newString[i] == '\t' {
			continue
		} else {
			retval += string(newString[i])
		}
	}
	return retval
}

/** Generate a call  identifier. This is useful when we want
 * to generate a call identifier in advance of generating a message.
 */
//        GenerateCallIdentifier( address string) string{
//        	r := rand.New(rand.NewSource(99))
//            String date = time.Date().String() + r.Float64()
//            new Double(Math.random()).toString();
//            try {
//                MessageDigest messageDigest = MessageDigest.getInstance("MD5");
//                byte cid[] = messageDigest.digest(date.getBytes());
//                String cidString = Utils.toHexString(cid);
//                return cidString + "@" + address;
//            } catch ( NoSuchAlgorithmException ex ) {
//                LogWriter.logException(ex);
//                return null;
//            }

//        }

// /** Generate a tag for a FROM header or TO header. Just return a
// * random 4 digit integer (should be enough to avoid any clashes!)
// * Tags only need to be unique within a call.
// *
// * @return a string that can be used as a tag parameter.
// */
// public static String generateTag() {
//            return new Integer((int)(Math.random() * 10000)).toString();
// }

// /** Generate a cryptographically random identifier that can be used
// * to generate a branch identifier.
// *
// *@return a cryptographically random gloablly unique string that
// *	can be used as a branch identifier.
// */
// public static String generateBranchId() {
//          String b =  new Integer((int)(Math.random() * 10000)).toString() +
// 	System.currentTimeMillis();
//          try {
//              MessageDigest messageDigest = MessageDigest.getInstance("MD5");
//              byte bid[] = messageDigest.digest(b.getBytes());
// 	// cryptographically random string.
// 	// prepend with a magic cookie to indicate we
// 	// are bis09 compatible.
//              return 	SIPConstants.BRANCH_MAGIC_COOKIE +
// 		Utils.toHexString(bid);
//           } catch ( NoSuchAlgorithmException ex ) {
//       if (LogWriter.needsLogging)
//                 LogWriter.logMessage("Algorithm not found " + ex);
//       return null;
//           }
// }
