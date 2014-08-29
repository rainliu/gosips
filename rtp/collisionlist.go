package rtp

import (
	"container/list"
	"errors"
)

type AddressTime struct {
	addr     Address
	recvtime RTPTime
}

/** This class represents a list of addresses from which SSRC collisions were detected. */
type CollisionList struct {
	addresslist *list.List
}

/** Constructs an instance, optionally installing a memory manager. */
func NewCollisionList() *CollisionList {
	this := &CollisionList{}
	this.addresslist = list.New()
	return this
}

/** Clears the list of addresses. */

func (this *CollisionList) Clear() {
	this.addresslist.Init()
}

/** Updates the entry for address \c addr to indicate that a collision was detected at time \c receivetime.
 *  If the entry did not exist yet, the flag \c created is set to \c true, otherwise it is set to \c false.
 */
func (this *CollisionList) UpdateAddress(addr Address, receivetime RTPTime) (bool, error) {
	if addr == nil {
		return false, errors.New("ERR_RTP_COLLISIONLIST_BADADDRESS")
	}

	for it := this.addresslist.Front(); it != nil; it = it.Next() {
		v := it.Value.(*AddressTime)
		if v.addr.IsSameAddress(addr) {
			v.recvtime = receivetime
			return false, nil
		}
	}

	newaddr := addr.Clone()
	this.addresslist.PushBack(&AddressTime{newaddr, receivetime})
	return true, nil
}

/** Returns \c true} if the address \c addr appears in the list. */
func (this *CollisionList) HasAddress(addr Address) bool {
	for it := this.addresslist.Front(); it != nil; it = it.Next() {
		v := it.Value.(*AddressTime)
		if v.addr.IsSameAddress(addr) {
			return true
		}
	}

	return false
}

/** Assuming that the current time is given by \c currenttime, this function times out entries which
 *  haven't been updated in the previous time interval specified by \c timeoutdelay.
 */

func (this *CollisionList) Timeout(currenttime *RTPTime, timeoutdelay *RTPTime) {
	checktime := currenttime.Clone()
	checktime.Sub(timeoutdelay)

	it := this.addresslist.Front()
	for it != nil {
		v := it.Value.(*AddressTime)
		if v.recvtime.LT(checktime) { // timeout
			d := it
			it = it.Next()
			this.addresslist.Remove(d)
		} else {
			it = it.Next()
		}
	}
}
