/*
 * AirGateway NDC JSON API
 *
 * This API allows shopping and booking with IATA's New Distribution Capabilities (NDC) standard. It provides aggregated shopping capabilities (AirShopping), detailed offer description (OfferPrice), flight seat selection (SeatAvailability) and booking flight reservations (OrderCreate). Some fields in our API (when noticed) use the PADIS Standard v16.1. Find more information <a href=http://www.iata.org/whatwedo/workgroups/Pages/padis.aspx>here</a>
 *
 * API version: 1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Price struct {

	// 670 or 670.41
	BaseAmount float64 `json:"baseAmount,omitempty"`

	Currency string `json:"currency,omitempty"`

	// The price for each flight offer. There is an offer for each type of passenger (ADT, CHD, INF).
	OfferPrices []OfferPrice `json:"offerPrices,omitempty"`

	// 570 or 570.41
	ProviderBaseAmount float64 `json:"providerBaseAmount,omitempty"`

	ProviderCurrency string `json:"providerCurrency,omitempty"`

	// 617 or 617.21
	ProviderTotalPrice float64 `json:"providerTotalPrice,omitempty"`

	// 136 or 136.80
	ProviderTotalTaxes float64 `json:"providerTotalTaxes,omitempty"`

	// The value price, currency, code and description for each tax for each AirlineOffer
	Taxes []Tax `json:"taxes,omitempty"`

	// 720 or 720.62
	TotalPrice float64 `json:"totalPrice,omitempty"`

	// 120 or 120.63
	TotalTaxes float64 `json:"totalTaxes,omitempty"`
}