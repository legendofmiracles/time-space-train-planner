// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetJourneysParams creates a new GetJourneysParams object
// with the default values initialized.
func NewGetJourneysParams() *GetJourneysParams {
	var (
		bikeDefault             = bool(false)
		languageDefault         = string("en")
		polylinesDefault        = bool(false)
		remarksDefault          = bool(true)
		resultsDefault          = int64(3)
		scheduledDaysDefault    = bool(false)
		startWithWalkingDefault = bool(true)
		stopoversDefault        = bool(false)
		ticketsDefault          = bool(false)
		transferTimeDefault     = int64(0)
		walkingSpeedDefault     = string("normal")
	)
	return &GetJourneysParams{
		Bike:             &bikeDefault,
		Language:         &languageDefault,
		Polylines:        &polylinesDefault,
		Remarks:          &remarksDefault,
		Results:          &resultsDefault,
		ScheduledDays:    &scheduledDaysDefault,
		StartWithWalking: &startWithWalkingDefault,
		Stopovers:        &stopoversDefault,
		Tickets:          &ticketsDefault,
		TransferTime:     &transferTimeDefault,
		WalkingSpeed:     &walkingSpeedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetJourneysParamsWithTimeout creates a new GetJourneysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetJourneysParamsWithTimeout(timeout time.Duration) *GetJourneysParams {
	var (
		bikeDefault             = bool(false)
		languageDefault         = string("en")
		polylinesDefault        = bool(false)
		remarksDefault          = bool(true)
		resultsDefault          = int64(3)
		scheduledDaysDefault    = bool(false)
		startWithWalkingDefault = bool(true)
		stopoversDefault        = bool(false)
		ticketsDefault          = bool(false)
		transferTimeDefault     = int64(0)
		walkingSpeedDefault     = string("normal")
	)
	return &GetJourneysParams{
		Bike:             &bikeDefault,
		Language:         &languageDefault,
		Polylines:        &polylinesDefault,
		Remarks:          &remarksDefault,
		Results:          &resultsDefault,
		ScheduledDays:    &scheduledDaysDefault,
		StartWithWalking: &startWithWalkingDefault,
		Stopovers:        &stopoversDefault,
		Tickets:          &ticketsDefault,
		TransferTime:     &transferTimeDefault,
		WalkingSpeed:     &walkingSpeedDefault,

		timeout: timeout,
	}
}

// NewGetJourneysParamsWithContext creates a new GetJourneysParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetJourneysParamsWithContext(ctx context.Context) *GetJourneysParams {
	var (
		bikeDefault             = bool(false)
		languageDefault         = string("en")
		polylinesDefault        = bool(false)
		remarksDefault          = bool(true)
		resultsDefault          = int64(3)
		scheduledDaysDefault    = bool(false)
		startWithWalkingDefault = bool(true)
		stopoversDefault        = bool(false)
		ticketsDefault          = bool(false)
		transferTimeDefault     = int64(0)
		walkingSpeedDefault     = string("normal")
	)
	return &GetJourneysParams{
		Bike:             &bikeDefault,
		Language:         &languageDefault,
		Polylines:        &polylinesDefault,
		Remarks:          &remarksDefault,
		Results:          &resultsDefault,
		ScheduledDays:    &scheduledDaysDefault,
		StartWithWalking: &startWithWalkingDefault,
		Stopovers:        &stopoversDefault,
		Tickets:          &ticketsDefault,
		TransferTime:     &transferTimeDefault,
		WalkingSpeed:     &walkingSpeedDefault,

		Context: ctx,
	}
}

// NewGetJourneysParamsWithHTTPClient creates a new GetJourneysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetJourneysParamsWithHTTPClient(client *http.Client) *GetJourneysParams {
	var (
		bikeDefault             = bool(false)
		languageDefault         = string("en")
		polylinesDefault        = bool(false)
		remarksDefault          = bool(true)
		resultsDefault          = int64(3)
		scheduledDaysDefault    = bool(false)
		startWithWalkingDefault = bool(true)
		stopoversDefault        = bool(false)
		ticketsDefault          = bool(false)
		transferTimeDefault     = int64(0)
		walkingSpeedDefault     = string("normal")
	)
	return &GetJourneysParams{
		Bike:             &bikeDefault,
		Language:         &languageDefault,
		Polylines:        &polylinesDefault,
		Remarks:          &remarksDefault,
		Results:          &resultsDefault,
		ScheduledDays:    &scheduledDaysDefault,
		StartWithWalking: &startWithWalkingDefault,
		Stopovers:        &stopoversDefault,
		Tickets:          &ticketsDefault,
		TransferTime:     &transferTimeDefault,
		WalkingSpeed:     &walkingSpeedDefault,
		HTTPClient:       client,
	}
}

/*GetJourneysParams contains all the parameters to send to the API endpoint
for the get journeys operation typically these are written to a http.Request
*/
type GetJourneysParams struct {

	/*Accessibility
	  `partial` or `complete`. – Default: *not accessible*

	*/
	Accessibility *string
	/*Arrival
	  Compute journeys arriving at this date/time. Mutually exclusive with `departure`. – Default: *now*

	*/
	Arrival *strfmt.DateTime
	/*Bike
	  Compute only bike-friendly journeys?

	*/
	Bike *bool
	/*Departure
	  Compute journeys departing at this date/time. Mutually exclusive with `arrival`. – Default: *now*

	*/
	Departure *strfmt.DateTime
	/*EarlierThan
	  Compute journeys "before" an `ealierRef`.

	*/
	EarlierThan *string
	/*From*/
	From *string
	/*Language
	  Language of the results.

	*/
	Language *string
	/*LaterThan
	  Compute journeys "after" an `laterRef`.

	*/
	LaterThan *string
	/*Polylines
	  Fetch & parse a shape for each journey leg?

	*/
	Polylines *bool
	/*Remarks
	  Parse & return hints & warnings?

	*/
	Remarks *bool
	/*Results
	  Max. number of journeys.

	*/
	Results *int64
	/*ScheduledDays
	  Parse & return dates each journey is valid on?

	*/
	ScheduledDays *bool
	/*StartWithWalking
	  Consider walking to nearby stations at the beginning of a journey?

	*/
	StartWithWalking *bool
	/*Stopovers
	  Fetch & parse stopovers on the way?

	*/
	Stopovers *bool
	/*Tickets
	  Return information about available tickets?

	*/
	Tickets *bool
	/*To*/
	To *string
	/*TransferTime
	  Minimum time in minutes for a single transfer.

	*/
	TransferTime *int64
	/*Transfers
	  Maximum number of transfers. – Default: *let HAFAS decide*

	*/
	Transfers *int64
	/*WalkingSpeed
	  `slow`, `normal` or `fast`.

	*/
	WalkingSpeed *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get journeys params
func (o *GetJourneysParams) WithTimeout(timeout time.Duration) *GetJourneysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get journeys params
func (o *GetJourneysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get journeys params
func (o *GetJourneysParams) WithContext(ctx context.Context) *GetJourneysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get journeys params
func (o *GetJourneysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get journeys params
func (o *GetJourneysParams) WithHTTPClient(client *http.Client) *GetJourneysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get journeys params
func (o *GetJourneysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessibility adds the accessibility to the get journeys params
func (o *GetJourneysParams) WithAccessibility(accessibility *string) *GetJourneysParams {
	o.SetAccessibility(accessibility)
	return o
}

// SetAccessibility adds the accessibility to the get journeys params
func (o *GetJourneysParams) SetAccessibility(accessibility *string) {
	o.Accessibility = accessibility
}

// WithArrival adds the arrival to the get journeys params
func (o *GetJourneysParams) WithArrival(arrival *strfmt.DateTime) *GetJourneysParams {
	o.SetArrival(arrival)
	return o
}

// SetArrival adds the arrival to the get journeys params
func (o *GetJourneysParams) SetArrival(arrival *strfmt.DateTime) {
	o.Arrival = arrival
}

// WithBike adds the bike to the get journeys params
func (o *GetJourneysParams) WithBike(bike *bool) *GetJourneysParams {
	o.SetBike(bike)
	return o
}

// SetBike adds the bike to the get journeys params
func (o *GetJourneysParams) SetBike(bike *bool) {
	o.Bike = bike
}

// WithDeparture adds the departure to the get journeys params
func (o *GetJourneysParams) WithDeparture(departure *strfmt.DateTime) *GetJourneysParams {
	o.SetDeparture(departure)
	return o
}

// SetDeparture adds the departure to the get journeys params
func (o *GetJourneysParams) SetDeparture(departure *strfmt.DateTime) {
	o.Departure = departure
}

// WithEarlierThan adds the earlierThan to the get journeys params
func (o *GetJourneysParams) WithEarlierThan(earlierThan *string) *GetJourneysParams {
	o.SetEarlierThan(earlierThan)
	return o
}

// SetEarlierThan adds the earlierThan to the get journeys params
func (o *GetJourneysParams) SetEarlierThan(earlierThan *string) {
	o.EarlierThan = earlierThan
}

// WithFrom adds the from to the get journeys params
func (o *GetJourneysParams) WithFrom(from *string) *GetJourneysParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get journeys params
func (o *GetJourneysParams) SetFrom(from *string) {
	o.From = from
}

// WithLanguage adds the language to the get journeys params
func (o *GetJourneysParams) WithLanguage(language *string) *GetJourneysParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the get journeys params
func (o *GetJourneysParams) SetLanguage(language *string) {
	o.Language = language
}

// WithLaterThan adds the laterThan to the get journeys params
func (o *GetJourneysParams) WithLaterThan(laterThan *string) *GetJourneysParams {
	o.SetLaterThan(laterThan)
	return o
}

// SetLaterThan adds the laterThan to the get journeys params
func (o *GetJourneysParams) SetLaterThan(laterThan *string) {
	o.LaterThan = laterThan
}

// WithPolylines adds the polylines to the get journeys params
func (o *GetJourneysParams) WithPolylines(polylines *bool) *GetJourneysParams {
	o.SetPolylines(polylines)
	return o
}

// SetPolylines adds the polylines to the get journeys params
func (o *GetJourneysParams) SetPolylines(polylines *bool) {
	o.Polylines = polylines
}

// WithRemarks adds the remarks to the get journeys params
func (o *GetJourneysParams) WithRemarks(remarks *bool) *GetJourneysParams {
	o.SetRemarks(remarks)
	return o
}

// SetRemarks adds the remarks to the get journeys params
func (o *GetJourneysParams) SetRemarks(remarks *bool) {
	o.Remarks = remarks
}

// WithResults adds the results to the get journeys params
func (o *GetJourneysParams) WithResults(results *int64) *GetJourneysParams {
	o.SetResults(results)
	return o
}

// SetResults adds the results to the get journeys params
func (o *GetJourneysParams) SetResults(results *int64) {
	o.Results = results
}

// WithScheduledDays adds the scheduledDays to the get journeys params
func (o *GetJourneysParams) WithScheduledDays(scheduledDays *bool) *GetJourneysParams {
	o.SetScheduledDays(scheduledDays)
	return o
}

// SetScheduledDays adds the scheduledDays to the get journeys params
func (o *GetJourneysParams) SetScheduledDays(scheduledDays *bool) {
	o.ScheduledDays = scheduledDays
}

// WithStartWithWalking adds the startWithWalking to the get journeys params
func (o *GetJourneysParams) WithStartWithWalking(startWithWalking *bool) *GetJourneysParams {
	o.SetStartWithWalking(startWithWalking)
	return o
}

// SetStartWithWalking adds the startWithWalking to the get journeys params
func (o *GetJourneysParams) SetStartWithWalking(startWithWalking *bool) {
	o.StartWithWalking = startWithWalking
}

// WithStopovers adds the stopovers to the get journeys params
func (o *GetJourneysParams) WithStopovers(stopovers *bool) *GetJourneysParams {
	o.SetStopovers(stopovers)
	return o
}

// SetStopovers adds the stopovers to the get journeys params
func (o *GetJourneysParams) SetStopovers(stopovers *bool) {
	o.Stopovers = stopovers
}

// WithTickets adds the tickets to the get journeys params
func (o *GetJourneysParams) WithTickets(tickets *bool) *GetJourneysParams {
	o.SetTickets(tickets)
	return o
}

// SetTickets adds the tickets to the get journeys params
func (o *GetJourneysParams) SetTickets(tickets *bool) {
	o.Tickets = tickets
}

// WithTo adds the to to the get journeys params
func (o *GetJourneysParams) WithTo(to *string) *GetJourneysParams {
	o.SetTo(to)
	return o
}

// SetTo adds the to to the get journeys params
func (o *GetJourneysParams) SetTo(to *string) {
	o.To = to
}

// WithTransferTime adds the transferTime to the get journeys params
func (o *GetJourneysParams) WithTransferTime(transferTime *int64) *GetJourneysParams {
	o.SetTransferTime(transferTime)
	return o
}

// SetTransferTime adds the transferTime to the get journeys params
func (o *GetJourneysParams) SetTransferTime(transferTime *int64) {
	o.TransferTime = transferTime
}

// WithTransfers adds the transfers to the get journeys params
func (o *GetJourneysParams) WithTransfers(transfers *int64) *GetJourneysParams {
	o.SetTransfers(transfers)
	return o
}

// SetTransfers adds the transfers to the get journeys params
func (o *GetJourneysParams) SetTransfers(transfers *int64) {
	o.Transfers = transfers
}

// WithWalkingSpeed adds the walkingSpeed to the get journeys params
func (o *GetJourneysParams) WithWalkingSpeed(walkingSpeed *string) *GetJourneysParams {
	o.SetWalkingSpeed(walkingSpeed)
	return o
}

// SetWalkingSpeed adds the walkingSpeed to the get journeys params
func (o *GetJourneysParams) SetWalkingSpeed(walkingSpeed *string) {
	o.WalkingSpeed = walkingSpeed
}

// WriteToRequest writes these params to a swagger request
func (o *GetJourneysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Accessibility != nil {

		// query param accessibility
		var qrAccessibility string
		if o.Accessibility != nil {
			qrAccessibility = *o.Accessibility
		}
		qAccessibility := qrAccessibility
		if qAccessibility != "" {
			if err := r.SetQueryParam("accessibility", qAccessibility); err != nil {
				return err
			}
		}

	}

	if o.Arrival != nil {

		// query param arrival
		var qrArrival strfmt.DateTime
		if o.Arrival != nil {
			qrArrival = *o.Arrival
		}
		qArrival := qrArrival.String()
		if qArrival != "" {
			if err := r.SetQueryParam("arrival", qArrival); err != nil {
				return err
			}
		}

	}

	if o.Bike != nil {

		// query param bike
		var qrBike bool
		if o.Bike != nil {
			qrBike = *o.Bike
		}
		qBike := swag.FormatBool(qrBike)
		if qBike != "" {
			if err := r.SetQueryParam("bike", qBike); err != nil {
				return err
			}
		}

	}

	if o.Departure != nil {

		// query param departure
		var qrDeparture strfmt.DateTime
		if o.Departure != nil {
			qrDeparture = *o.Departure
		}
		qDeparture := qrDeparture.String()
		if qDeparture != "" {
			if err := r.SetQueryParam("departure", qDeparture); err != nil {
				return err
			}
		}

	}

	if o.EarlierThan != nil {

		// query param earlierThan
		var qrEarlierThan string
		if o.EarlierThan != nil {
			qrEarlierThan = *o.EarlierThan
		}
		qEarlierThan := qrEarlierThan
		if qEarlierThan != "" {
			if err := r.SetQueryParam("earlierThan", qEarlierThan); err != nil {
				return err
			}
		}

	}

	if o.From != nil {

		// query param from
		var qrFrom string
		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := qrFrom
		if qFrom != "" {
			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}

	}

	if o.Language != nil {

		// query param language
		var qrLanguage string
		if o.Language != nil {
			qrLanguage = *o.Language
		}
		qLanguage := qrLanguage
		if qLanguage != "" {
			if err := r.SetQueryParam("language", qLanguage); err != nil {
				return err
			}
		}

	}

	if o.LaterThan != nil {

		// query param laterThan
		var qrLaterThan string
		if o.LaterThan != nil {
			qrLaterThan = *o.LaterThan
		}
		qLaterThan := qrLaterThan
		if qLaterThan != "" {
			if err := r.SetQueryParam("laterThan", qLaterThan); err != nil {
				return err
			}
		}

	}

	if o.Polylines != nil {

		// query param polylines
		var qrPolylines bool
		if o.Polylines != nil {
			qrPolylines = *o.Polylines
		}
		qPolylines := swag.FormatBool(qrPolylines)
		if qPolylines != "" {
			if err := r.SetQueryParam("polylines", qPolylines); err != nil {
				return err
			}
		}

	}

	if o.Remarks != nil {

		// query param remarks
		var qrRemarks bool
		if o.Remarks != nil {
			qrRemarks = *o.Remarks
		}
		qRemarks := swag.FormatBool(qrRemarks)
		if qRemarks != "" {
			if err := r.SetQueryParam("remarks", qRemarks); err != nil {
				return err
			}
		}

	}

	if o.Results != nil {

		// query param results
		var qrResults int64
		if o.Results != nil {
			qrResults = *o.Results
		}
		qResults := swag.FormatInt64(qrResults)
		if qResults != "" {
			if err := r.SetQueryParam("results", qResults); err != nil {
				return err
			}
		}

	}

	if o.ScheduledDays != nil {

		// query param scheduledDays
		var qrScheduledDays bool
		if o.ScheduledDays != nil {
			qrScheduledDays = *o.ScheduledDays
		}
		qScheduledDays := swag.FormatBool(qrScheduledDays)
		if qScheduledDays != "" {
			if err := r.SetQueryParam("scheduledDays", qScheduledDays); err != nil {
				return err
			}
		}

	}

	if o.StartWithWalking != nil {

		// query param startWithWalking
		var qrStartWithWalking bool
		if o.StartWithWalking != nil {
			qrStartWithWalking = *o.StartWithWalking
		}
		qStartWithWalking := swag.FormatBool(qrStartWithWalking)
		if qStartWithWalking != "" {
			if err := r.SetQueryParam("startWithWalking", qStartWithWalking); err != nil {
				return err
			}
		}

	}

	if o.Stopovers != nil {

		// query param stopovers
		var qrStopovers bool
		if o.Stopovers != nil {
			qrStopovers = *o.Stopovers
		}
		qStopovers := swag.FormatBool(qrStopovers)
		if qStopovers != "" {
			if err := r.SetQueryParam("stopovers", qStopovers); err != nil {
				return err
			}
		}

	}

	if o.Tickets != nil {

		// query param tickets
		var qrTickets bool
		if o.Tickets != nil {
			qrTickets = *o.Tickets
		}
		qTickets := swag.FormatBool(qrTickets)
		if qTickets != "" {
			if err := r.SetQueryParam("tickets", qTickets); err != nil {
				return err
			}
		}

	}

	if o.To != nil {

		// query param to
		var qrTo string
		if o.To != nil {
			qrTo = *o.To
		}
		qTo := qrTo
		if qTo != "" {
			if err := r.SetQueryParam("to", qTo); err != nil {
				return err
			}
		}

	}

	if o.TransferTime != nil {

		// query param transferTime
		var qrTransferTime int64
		if o.TransferTime != nil {
			qrTransferTime = *o.TransferTime
		}
		qTransferTime := swag.FormatInt64(qrTransferTime)
		if qTransferTime != "" {
			if err := r.SetQueryParam("transferTime", qTransferTime); err != nil {
				return err
			}
		}

	}

	if o.Transfers != nil {

		// query param transfers
		var qrTransfers int64
		if o.Transfers != nil {
			qrTransfers = *o.Transfers
		}
		qTransfers := swag.FormatInt64(qrTransfers)
		if qTransfers != "" {
			if err := r.SetQueryParam("transfers", qTransfers); err != nil {
				return err
			}
		}

	}

	if o.WalkingSpeed != nil {

		// query param walkingSpeed
		var qrWalkingSpeed string
		if o.WalkingSpeed != nil {
			qrWalkingSpeed = *o.WalkingSpeed
		}
		qWalkingSpeed := qrWalkingSpeed
		if qWalkingSpeed != "" {
			if err := r.SetQueryParam("walkingSpeed", qWalkingSpeed); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
