package dategroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"net/http"
)

// GroupClient is the test Infrastructure for AutoRest
type GroupClient struct {
	ManagementClient
}

// NewGroupClient creates an instance of the GroupClient client.
func NewGroupClient() GroupClient {
	return NewGroupClientWithBaseURI(DefaultBaseURI)
}

// NewGroupClientWithBaseURI creates an instance of the GroupClient client.
func NewGroupClientWithBaseURI(baseURI string) GroupClient {
	return GroupClient{NewWithBaseURI(baseURI)}
}

// GetInvalidDate get invalid date value
func (client GroupClient) GetInvalidDate() (result Date, err error) {
	req, err := client.GetInvalidDatePreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.GroupClient", "GetInvalidDate", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetInvalidDateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dategroup.GroupClient", "GetInvalidDate", resp, "Failure sending request")
		return
	}

	result, err = client.GetInvalidDateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.GroupClient", "GetInvalidDate", resp, "Failure responding to request")
	}

	return
}

// GetInvalidDatePreparer prepares the GetInvalidDate request.
func (client GroupClient) GetInvalidDatePreparer() (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/date/invaliddate"))
	return preparer.Prepare(&http.Request{})
}

// GetInvalidDateSender sends the GetInvalidDate request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) GetInvalidDateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetInvalidDateResponder handles the response to the GetInvalidDate request. The method always
// closes the http.Response Body.
func (client GroupClient) GetInvalidDateResponder(resp *http.Response) (result Date, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetMaxDate get max date value 9999-12-31
func (client GroupClient) GetMaxDate() (result Date, err error) {
	req, err := client.GetMaxDatePreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.GroupClient", "GetMaxDate", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetMaxDateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dategroup.GroupClient", "GetMaxDate", resp, "Failure sending request")
		return
	}

	result, err = client.GetMaxDateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.GroupClient", "GetMaxDate", resp, "Failure responding to request")
	}

	return
}

// GetMaxDatePreparer prepares the GetMaxDate request.
func (client GroupClient) GetMaxDatePreparer() (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/date/max"))
	return preparer.Prepare(&http.Request{})
}

// GetMaxDateSender sends the GetMaxDate request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) GetMaxDateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetMaxDateResponder handles the response to the GetMaxDate request. The method always
// closes the http.Response Body.
func (client GroupClient) GetMaxDateResponder(resp *http.Response) (result Date, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetMinDate get min date value 0000-01-01
func (client GroupClient) GetMinDate() (result Date, err error) {
	req, err := client.GetMinDatePreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.GroupClient", "GetMinDate", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetMinDateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dategroup.GroupClient", "GetMinDate", resp, "Failure sending request")
		return
	}

	result, err = client.GetMinDateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.GroupClient", "GetMinDate", resp, "Failure responding to request")
	}

	return
}

// GetMinDatePreparer prepares the GetMinDate request.
func (client GroupClient) GetMinDatePreparer() (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/date/min"))
	return preparer.Prepare(&http.Request{})
}

// GetMinDateSender sends the GetMinDate request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) GetMinDateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetMinDateResponder handles the response to the GetMinDate request. The method always
// closes the http.Response Body.
func (client GroupClient) GetMinDateResponder(resp *http.Response) (result Date, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetNull get null date value
func (client GroupClient) GetNull() (result Date, err error) {
	req, err := client.GetNullPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.GroupClient", "GetNull", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetNullSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dategroup.GroupClient", "GetNull", resp, "Failure sending request")
		return
	}

	result, err = client.GetNullResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.GroupClient", "GetNull", resp, "Failure responding to request")
	}

	return
}

// GetNullPreparer prepares the GetNull request.
func (client GroupClient) GetNullPreparer() (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/date/null"))
	return preparer.Prepare(&http.Request{})
}

// GetNullSender sends the GetNull request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) GetNullSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetNullResponder handles the response to the GetNull request. The method always
// closes the http.Response Body.
func (client GroupClient) GetNullResponder(resp *http.Response) (result Date, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetOverflowDate get overflow date value
func (client GroupClient) GetOverflowDate() (result Date, err error) {
	req, err := client.GetOverflowDatePreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.GroupClient", "GetOverflowDate", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetOverflowDateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dategroup.GroupClient", "GetOverflowDate", resp, "Failure sending request")
		return
	}

	result, err = client.GetOverflowDateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.GroupClient", "GetOverflowDate", resp, "Failure responding to request")
	}

	return
}

// GetOverflowDatePreparer prepares the GetOverflowDate request.
func (client GroupClient) GetOverflowDatePreparer() (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/date/overflowdate"))
	return preparer.Prepare(&http.Request{})
}

// GetOverflowDateSender sends the GetOverflowDate request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) GetOverflowDateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetOverflowDateResponder handles the response to the GetOverflowDate request. The method always
// closes the http.Response Body.
func (client GroupClient) GetOverflowDateResponder(resp *http.Response) (result Date, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetUnderflowDate get underflow date value
func (client GroupClient) GetUnderflowDate() (result Date, err error) {
	req, err := client.GetUnderflowDatePreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.GroupClient", "GetUnderflowDate", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetUnderflowDateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dategroup.GroupClient", "GetUnderflowDate", resp, "Failure sending request")
		return
	}

	result, err = client.GetUnderflowDateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.GroupClient", "GetUnderflowDate", resp, "Failure responding to request")
	}

	return
}

// GetUnderflowDatePreparer prepares the GetUnderflowDate request.
func (client GroupClient) GetUnderflowDatePreparer() (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/date/underflowdate"))
	return preparer.Prepare(&http.Request{})
}

// GetUnderflowDateSender sends the GetUnderflowDate request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) GetUnderflowDateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetUnderflowDateResponder handles the response to the GetUnderflowDate request. The method always
// closes the http.Response Body.
func (client GroupClient) GetUnderflowDateResponder(resp *http.Response) (result Date, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutMaxDate put max date value 9999-12-31
//
func (client GroupClient) PutMaxDate(dateBody date.Date) (result autorest.Response, err error) {
	req, err := client.PutMaxDatePreparer(dateBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.GroupClient", "PutMaxDate", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutMaxDateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "dategroup.GroupClient", "PutMaxDate", resp, "Failure sending request")
		return
	}

	result, err = client.PutMaxDateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.GroupClient", "PutMaxDate", resp, "Failure responding to request")
	}

	return
}

// PutMaxDatePreparer prepares the PutMaxDate request.
func (client GroupClient) PutMaxDatePreparer(dateBody date.Date) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/date/max"),
		autorest.WithJSON(dateBody))
	return preparer.Prepare(&http.Request{})
}

// PutMaxDateSender sends the PutMaxDate request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) PutMaxDateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// PutMaxDateResponder handles the response to the PutMaxDate request. The method always
// closes the http.Response Body.
func (client GroupClient) PutMaxDateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PutMinDate put min date value 0000-01-01
//
func (client GroupClient) PutMinDate(dateBody date.Date) (result autorest.Response, err error) {
	req, err := client.PutMinDatePreparer(dateBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.GroupClient", "PutMinDate", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutMinDateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "dategroup.GroupClient", "PutMinDate", resp, "Failure sending request")
		return
	}

	result, err = client.PutMinDateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dategroup.GroupClient", "PutMinDate", resp, "Failure responding to request")
	}

	return
}

// PutMinDatePreparer prepares the PutMinDate request.
func (client GroupClient) PutMinDatePreparer(dateBody date.Date) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/date/min"),
		autorest.WithJSON(dateBody))
	return preparer.Prepare(&http.Request{})
}

// PutMinDateSender sends the PutMinDate request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) PutMinDateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// PutMinDateResponder handles the response to the PutMinDate request. The method always
// closes the http.Response Body.
func (client GroupClient) PutMinDateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}