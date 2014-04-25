package parser

import (
	"testing"
)

func TestSubscriptionStateParser(t *testing.T) {
	var tvs = []string{
		"Subscription-State: active \n",
		"Subscription-State: terminated;reason=rejected \n",
		"Subscription-State: pending;reason=probation;expires=36\n",
		"Subscription-State: pending;retry-after=10;expires=36\n",
		"Subscription-State: pending;generic=void\n",
	}
	var tvs_o = []string{
		"Subscription-State: active \n",
		"Subscription-State: terminated;reason=rejected \n",
		"Subscription-State: pending;reason=probation;expires=36\n",
		"Subscription-State: pending;retry-after=10;expires=36\n",
		"Subscription-State: pending;generic=void\n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewSubscriptionStateParser(tvs[i])
		testHeaderParser(t, shp, tvs_o[i])
	}
}

/** Test program
  public static void main(String args[]) throws ParseException {
      String subscriptionState[] = {
          "Subscription-State: active \n",
          "Subscription-State: terminated;reason=rejected \n",
          "Subscription-State: pending;reason=probation;expires=36\n",
          "Subscription-State: pending;retry-after=10;expires=36\n",
          "Subscription-State: pending;generic=void\n"
      };

      for (int i = 0; i < subscriptionState.length; i++ ) {
          SubscriptionStateParser parser =
          new SubscriptionStateParser(subscriptionState[i]);
          SubscriptionState ss= (SubscriptionState) parser.parse();
          System.out.println("encoded = " + ss.encode());
      }

  }
*/
