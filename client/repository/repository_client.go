// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new repository API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for repository API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteRepository(params *DeleteRepositoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteRepositoryOK, error)

	GetRepository(params *GetRepositoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRepositoryOK, error)

	ListAllRepositories(params *ListAllRepositoriesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAllRepositoriesOK, error)

	ListRepositories(params *ListRepositoriesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListRepositoriesOK, error)

	UpdateRepository(params *UpdateRepositoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateRepositoryOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteRepository deletes repository

  Delete the repository specified by name
*/
func (a *Client) DeleteRepository(params *DeleteRepositoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteRepositoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRepositoryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteRepository",
		Method:             "DELETE",
		PathPattern:        "/projects/{project_name}/repositories/{repository_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteRepositoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteRepositoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteRepository: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRepository gets repository

  Get the repository specified by name
*/
func (a *Client) GetRepository(params *GetRepositoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRepositoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRepositoryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRepository",
		Method:             "GET",
		PathPattern:        "/projects/{project_name}/repositories/{repository_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRepositoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRepositoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRepository: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAllRepositories lists all authorized repositories

  List all authorized repositories
*/
func (a *Client) ListAllRepositories(params *ListAllRepositoriesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAllRepositoriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAllRepositoriesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAllRepositories",
		Method:             "GET",
		PathPattern:        "/repositories",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListAllRepositoriesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAllRepositoriesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listAllRepositories: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListRepositories lists repositories

  List repositories of the specified project
*/
func (a *Client) ListRepositories(params *ListRepositoriesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListRepositoriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRepositoriesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listRepositories",
		Method:             "GET",
		PathPattern:        "/projects/{project_name}/repositories",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListRepositoriesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListRepositoriesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listRepositories: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateRepository updates repository

  Update the repository specified by name
*/
func (a *Client) UpdateRepository(params *UpdateRepositoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateRepositoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRepositoryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateRepository",
		Method:             "PUT",
		PathPattern:        "/projects/{project_name}/repositories/{repository_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateRepositoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateRepositoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateRepository: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
