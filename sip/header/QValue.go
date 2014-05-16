package header

type QValue interface {

	/**

	 * Gets q-value of the media-range in AcceptLanguageHeader. A value of

	 * <code>-1</code> indicates the<code>q-value</code> is not set.

	 *

	 * @return q-value of media-range, -1 if q-value is not set.

	 */

	GetQValue() float32

	/**

	 * Sets q-value for media-range in AcceptLanguageHeader. Q-values allow the

	 * user to indicate the relative degree of preference for that media-range,

	 * using the qvalue scale from 0 to 1. If no q-value is present, the

	 * media-range should be treated as having a q-value of 1.

	 *

	 * @param qValue - the new float value of the q-value

	 * @throws InvalidArgumentException if the q parameter value is not between <code>0 and 1</code>.

	 */

	SetQValue(qValue float32) (InvalidArgumentException error)

	RemoveQValue()
	
	HasQValue() bool
}
