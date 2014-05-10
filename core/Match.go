package core

import ()

/** Match template for pattern matching.
 */

type Match interface {
	/** Return true if a match occurs for searchString.
	 * This is used for pattern matching in the find and replace and match
	 * functions of the sipheaders and sdpfields packages. We have avoided
	 * picking a specific regexp package to avoid code dependencies.
	 * Use a package such as the jakarta regexp package to implement this.
	 */
	Match(searchString string) bool
}
