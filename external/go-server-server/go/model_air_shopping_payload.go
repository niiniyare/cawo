/*
 * AirGateway NDC JSON API
 *
 * This API allows shopping and booking with IATA's New Distribution Capabilities (NDC) standard. It provides aggregated shopping capabilities (AirShopping), detailed offer description (OfferPrice), flight seat selection (SeatAvailability) and booking flight reservations (OrderCreate). Some fields in our API (when noticed) use the PADIS Standard v16.1. Find more information <a href=http://www.iata.org/whatwedo/workgroups/Pages/padis.aspx>here</a>
 *
 * API version: 1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// AirShoppingPayload is the type for AirShopping endpoint request body
type AirShoppingPayload struct {

	// Corporate Discount Codes Map
	CorporateDiscountCodes interface{} `json:"corporateDiscountCodes,omitempty"`

	Metadata *Metadata `json:"metadata,omitempty"`

	// Origin/Destination pair in a shopping search scenario
	OriginDestinations []OriginDestinationAsrq `json:"originDestinations"`

	Preferences *Preferences `json:"preferences,omitempty"`

	Travelers *Travelers `json:"travelers"`
}
