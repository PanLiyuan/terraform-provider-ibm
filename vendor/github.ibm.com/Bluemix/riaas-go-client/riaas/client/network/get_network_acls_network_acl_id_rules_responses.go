// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// GetNetworkAclsNetworkACLIDRulesReader is a Reader for the GetNetworkAclsNetworkACLIDRules structure.
type GetNetworkAclsNetworkACLIDRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkAclsNetworkACLIDRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNetworkAclsNetworkACLIDRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetNetworkAclsNetworkACLIDRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworkAclsNetworkACLIDRulesOK creates a GetNetworkAclsNetworkACLIDRulesOK with default headers values
func NewGetNetworkAclsNetworkACLIDRulesOK() *GetNetworkAclsNetworkACLIDRulesOK {
	return &GetNetworkAclsNetworkACLIDRulesOK{}
}

/*GetNetworkAclsNetworkACLIDRulesOK handles this case with default header values.

dummy
*/
type GetNetworkAclsNetworkACLIDRulesOK struct {
	Payload *models.GetNetworkAclsNetworkACLIDRulesOKBody
}

func (o *GetNetworkAclsNetworkACLIDRulesOK) Error() string {
	return fmt.Sprintf("[GET /network_acls/{network_acl_id}/rules][%d] getNetworkAclsNetworkAclIdRulesOK  %+v", 200, o.Payload)
}

func (o *GetNetworkAclsNetworkACLIDRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetNetworkAclsNetworkACLIDRulesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkAclsNetworkACLIDRulesInternalServerError creates a GetNetworkAclsNetworkACLIDRulesInternalServerError with default headers values
func NewGetNetworkAclsNetworkACLIDRulesInternalServerError() *GetNetworkAclsNetworkACLIDRulesInternalServerError {
	return &GetNetworkAclsNetworkACLIDRulesInternalServerError{}
}

/*GetNetworkAclsNetworkACLIDRulesInternalServerError handles this case with default header values.

error
*/
type GetNetworkAclsNetworkACLIDRulesInternalServerError struct {
	Payload *models.Riaaserror
}

func (o *GetNetworkAclsNetworkACLIDRulesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /network_acls/{network_acl_id}/rules][%d] getNetworkAclsNetworkAclIdRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNetworkAclsNetworkACLIDRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}