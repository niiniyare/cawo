/*
 * AirGateway NDC JSON API
 *
 * This API allows shopping and booking with IATA's New Distribution Capabilities (NDC) standard. It provides aggregated shopping capabilities (AirShopping), detailed offer description (OfferPrice), flight seat selection (SeatAvailability) and booking flight reservations (OrderCreate). Some fields in our API (when noticed) use the PADIS Standard v16.1. Find more information <a href=http://www.iata.org/whatwedo/workgroups/Pages/padis.aspx>here</a>
 *
 * API version: 1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// ServiceListMedia media type (default view)
type ServiceListMedia struct {

	// List of disclosures
	Disclosures []Disclosure `json:"disclosures,omitempty"`

	FlightSegments []FlightSegment `json:"flightSegments,omitempty"`

	// List of referenced travelers
	Passengers []OfferPassenger `json:"passengers,omitempty"`

	// List of services
	Services []ServiceResponse `json:"services,omitempty"`

	Warnings []OrderLog `json:"warnings,omitempty"`
}
