package rtp

import (
	"os"
	"sync"
	"time"
)

const RANDOM_2POWMIN63 = 1.08420217248550443400745280086994171142578125e-19

/** Interface for generating random numbers. */
type Random interface {
	/** Returns a random eight bit value. */
	GetRandom8() uint8

	/** Returns a random sixteen bit value. */
	GetRandom16() uint16

	/** Returns a random thirty-two bit value. */
	GetRandom32() uint32

	/** Returns a random number between $0.0$ and $1.0$. */
	GetRandomDouble() float64

	/** Can be used by subclasses to generate a seed for a random number generator. */
	PickSeed() uint32
}

/** A random number generator using the algorithm of the rand48 set of functions. */
type RandomRand48 struct {
	mutex sync.Mutex
	state uint64
}

func NewRandomRand48() *RandomRand48 {
	this := &RandomRand48{}
	this.SetSeed(this.PickSeed())
	return this
}

func NewRandomRand48FromSeed(seed uint32) *RandomRand48 {
	this := &RandomRand48{}
	this.SetSeed(seed)
	return this
}

func (this *RandomRand48) PickSeed() uint32 {
	x := uint32(os.Getpid())
	x += uint32(time.Now().Second())
	//x += (uint32_t)clock();
	//x ^= (uint32_t)((uint8_t *)this - (uint8_t *)0);

	return x
}

func (this *RandomRand48) SetSeed(seed uint32) {
	this.state = (uint64(seed) << 16) | uint64(0x330E)
}

func (this *RandomRand48) GetRandom8() uint8 {
	x := ((this.GetRandom32() >> 24) & 0xff)

	return uint8(x)
}

func (this *RandomRand48) GetRandom16() uint16 {
	x := ((this.GetRandom32() >> 16) & 0xffff)

	return uint16(x)
}

func (this *RandomRand48) GetRandom32() uint32 {
	this.mutex.Lock()

	this.state = ((uint64(0x5DEECE66D) * this.state) + uint64(0xB)) & uint64(0x0000ffffffffffff)

	x := uint32((this.state >> 16) & uint64(0xffffffff))

	this.mutex.Unlock()

	return x
}

func (this *RandomRand48) GetRandomDouble() float64 {
	this.mutex.Lock()

	this.state = ((uint64(0x5DEECE66D) * this.state) + uint64(0xB)) & uint64(0x0000ffffffffffff)

	x := int64(this.state)

	this.mutex.Unlock()

	y := 3.552713678800500929355621337890625e-15 * float64(x)
	return y
}
