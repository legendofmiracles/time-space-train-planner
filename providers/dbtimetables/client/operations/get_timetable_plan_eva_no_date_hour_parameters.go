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
)

// NewGetTimetablePlanEvaNoDateHourParams creates a new GetTimetablePlanEvaNoDateHourParams object
// with the default values initialized.
func NewGetTimetablePlanEvaNoDateHourParams() *GetTimetablePlanEvaNoDateHourParams {
	var ()
	return &GetTimetablePlanEvaNoDateHourParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTimetablePlanEvaNoDateHourParamsWithTimeout creates a new GetTimetablePlanEvaNoDateHourParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTimetablePlanEvaNoDateHourParamsWithTimeout(timeout time.Duration) *GetTimetablePlanEvaNoDateHourParams {
	var ()
	return &GetTimetablePlanEvaNoDateHourParams{

		timeout: timeout,
	}
}

// NewGetTimetablePlanEvaNoDateHourParamsWithContext creates a new GetTimetablePlanEvaNoDateHourParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTimetablePlanEvaNoDateHourParamsWithContext(ctx context.Context) *GetTimetablePlanEvaNoDateHourParams {
	var ()
	return &GetTimetablePlanEvaNoDateHourParams{

		Context: ctx,
	}
}

// NewGetTimetablePlanEvaNoDateHourParamsWithHTTPClient creates a new GetTimetablePlanEvaNoDateHourParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTimetablePlanEvaNoDateHourParamsWithHTTPClient(client *http.Client) *GetTimetablePlanEvaNoDateHourParams {
	var ()
	return &GetTimetablePlanEvaNoDateHourParams{
		HTTPClient: client,
	}
}

/*GetTimetablePlanEvaNoDateHourParams contains all the parameters to send to the API endpoint
for the get timetable plan eva no date hour operation typically these are written to a http.Request
*/
type GetTimetablePlanEvaNoDateHourParams struct {

	/*Date
	  Date in format YYMMDD.

	*/
	Date string
	/*EvaNo
	  Station EVA-number.

	*/
	EvaNo string
	/*Hour
	  Hour in format HH.

	*/
	Hour string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get timetable plan eva no date hour params
func (o *GetTimetablePlanEvaNoDateHourParams) WithTimeout(timeout time.Duration) *GetTimetablePlanEvaNoDateHourParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get timetable plan eva no date hour params
func (o *GetTimetablePlanEvaNoDateHourParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get timetable plan eva no date hour params
func (o *GetTimetablePlanEvaNoDateHourParams) WithContext(ctx context.Context) *GetTimetablePlanEvaNoDateHourParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get timetable plan eva no date hour params
func (o *GetTimetablePlanEvaNoDateHourParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get timetable plan eva no date hour params
func (o *GetTimetablePlanEvaNoDateHourParams) WithHTTPClient(client *http.Client) *GetTimetablePlanEvaNoDateHourParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get timetable plan eva no date hour params
func (o *GetTimetablePlanEvaNoDateHourParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDate adds the date to the get timetable plan eva no date hour params
func (o *GetTimetablePlanEvaNoDateHourParams) WithDate(date string) *GetTimetablePlanEvaNoDateHourParams {
	o.SetDate(date)
	return o
}

// SetDate adds the date to the get timetable plan eva no date hour params
func (o *GetTimetablePlanEvaNoDateHourParams) SetDate(date string) {
	o.Date = date
}

// WithEvaNo adds the evaNo to the get timetable plan eva no date hour params
func (o *GetTimetablePlanEvaNoDateHourParams) WithEvaNo(evaNo string) *GetTimetablePlanEvaNoDateHourParams {
	o.SetEvaNo(evaNo)
	return o
}

// SetEvaNo adds the evaNo to the get timetable plan eva no date hour params
func (o *GetTimetablePlanEvaNoDateHourParams) SetEvaNo(evaNo string) {
	o.EvaNo = evaNo
}

// WithHour adds the hour to the get timetable plan eva no date hour params
func (o *GetTimetablePlanEvaNoDateHourParams) WithHour(hour string) *GetTimetablePlanEvaNoDateHourParams {
	o.SetHour(hour)
	return o
}

// SetHour adds the hour to the get timetable plan eva no date hour params
func (o *GetTimetablePlanEvaNoDateHourParams) SetHour(hour string) {
	o.Hour = hour
}

// WriteToRequest writes these params to a swagger request
func (o *GetTimetablePlanEvaNoDateHourParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param date
	if err := r.SetPathParam("date", o.Date); err != nil {
		return err
	}

	// path param evaNo
	if err := r.SetPathParam("evaNo", o.EvaNo); err != nil {
		return err
	}

	// path param hour
	if err := r.SetPathParam("hour", o.Hour); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}