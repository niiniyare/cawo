/*
 * AirGateway NDC JSON API
 *
 * This API allows shopping and booking with IATA's New Distribution Capabilities (NDC) standard. It provides aggregated shopping capabilities (AirShopping), detailed offer description (OfferPrice), flight seat selection (SeatAvailability) and booking flight reservations (OrderCreate). Some fields in our API (when noticed) use the PADIS Standard v16.1. Find more information <a href=http://www.iata.org/whatwedo/workgroups/Pages/padis.aspx>here</a>
 *
 * API version: 1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type OrderReshopMediaType struct {

	// List of disclosures corresponding to offer.
	Disclosures []Disclosure `json:"disclosures,omitempty"`

	// Flight Segment information. Including departure, arrival and airline information.
	Flights []OriginDestinationRs `json:"flights,omitempty"`

	// Offer identifier
	OfferID string `json:"offerID,omitempty"`

	// Offer type
	OfferType string `json:"offerType,omitempty"`

	// 2 letter abbreviation of airline
	Owner string `json:"owner,omitempty"`

	Price *Price `json:"price,omitempty"`

	PriceBreakdown *PriceBreakdown `json:"priceBreakdown,omitempty"`

	// Indicates if the price has change from the original in the hold order.
	PriceChange bool `json:"priceChange,omitempty"`

	Warnings []OrderLog `json:"warnings,omitempty"`
}
