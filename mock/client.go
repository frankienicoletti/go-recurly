package mock

import (
	"net/http"

	"github.com/blacklightcms/recurly"
)

// NewClient sets the unexported fields on the struct and returns a Client.
func NewClient(httpClient *http.Client) *recurly.Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := recurly.NewClient("a", "b", httpClient)

	// Services not implemented in mock package are nil so that they panic when used.
	client.Accounts = &AccountsService{}
	client.Adjustments = &AdjustmentsService{}
	client.Billing = &BillingService{}
	client.Coupons = &CouponsService{}
	client.Redemptions = &RedemptionsService{}
	client.Invoices = &InvoicesService{}
	client.Plans = &PlansService{}
	client.AddOns = &AddOnsService{}
	client.Subscriptions = &SubscriptionsService{}
	client.Transactions = &TransactionsService{}

	return client
}
