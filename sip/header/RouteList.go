package header

import "gosips/core"

/**
* A list of Route Headers.
 */
type RouteList struct {
	SIPHeaderList

	//private HashSet routeSet;
}

/** default constructor
 */
func NewRouteList() *RouteList {
	this := &RouteList{}
	this.SIPHeaderList.super(core.SIPHeaderNames_ROUTE)
	//this.routeSet = new HashSet();
	return this
}

// public boolean add(Object rh) {
//    if (! routeSet.contains(rh) ) {
// 	this.routeSet.add(rh);
// 	return super.add(rh);
//    } else return false;
// }

/** Constructor
 * @param sip SIPObjectList to set
 */
// public RouteList (SIPObjectList sip) {
// 	super(sip, RouteHeader.NAME);
// }

/**
* Order is important when comparing route lists.
 */
// public boolean equals(Object other) {
// 	if (!(other instanceof RouteList)) return false;
// 	RouteList that  = (RouteList) other;
// 	if (this.size() != that.size()) return false;
// 	ListIterator it = this.listIterator();
// 	ListIterator it1 = that.listIterator();
// 	while (it.hasNext()) {
// 	    Route route = (Route) it.next();
// 	    Route route1 = (Route) it1.next();
// 	    if (! route.equals(route1)) return false;
// 	}
// 	return true;
// }
