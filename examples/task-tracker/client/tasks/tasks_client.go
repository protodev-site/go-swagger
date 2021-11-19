// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/protodev-site/runtime"
)

// New creates a new tasks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tasks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddCommentToTask(params *AddCommentToTaskParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddCommentToTaskCreated, error)

	CreateTask(params *CreateTaskParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateTaskCreated, error)

	DeleteTask(params *DeleteTaskParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteTaskNoContent, error)

	GetTaskComments(params *GetTaskCommentsParams, opts ...ClientOption) (*GetTaskCommentsOK, error)

	GetTaskDetails(params *GetTaskDetailsParams, opts ...ClientOption) (*GetTaskDetailsOK, error)

	ListTasks(params *ListTasksParams, opts ...ClientOption) (*ListTasksOK, error)

	UpdateTask(params *UpdateTaskParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateTaskOK, error)

	UploadTaskFile(params *UploadTaskFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UploadTaskFileCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddCommentToTask adds a comment to a task

  The comment can contain ___github markdown___ syntax.
Fenced codeblocks etc are supported through pygments.

*/
func (a *Client) AddCommentToTask(params *AddCommentToTaskParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddCommentToTaskCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddCommentToTaskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addCommentToTask",
		Method:             "POST",
		PathPattern:        "/tasks/{id}/comments",
		ProducesMediaTypes: []string{"application/vnd.goswagger.examples.task-tracker.v1+json"},
		ConsumesMediaTypes: []string{"application/vnd.goswagger.examples.task-tracker.v1+json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddCommentToTaskReader{formats: a.formats},
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
	success, ok := result.(*AddCommentToTaskCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddCommentToTaskDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CreateTask creates a task object

  Allows for creating a task.
This operation requires authentication so that we know which user
created the task.

*/
func (a *Client) CreateTask(params *CreateTaskParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateTaskCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTaskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createTask",
		Method:             "POST",
		PathPattern:        "/tasks",
		ProducesMediaTypes: []string{"application/vnd.goswagger.examples.task-tracker.v1+json"},
		ConsumesMediaTypes: []string{"application/vnd.goswagger.examples.task-tracker.v1+json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateTaskReader{formats: a.formats},
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
	success, ok := result.(*CreateTaskCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateTaskDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteTask deletes a task

  This is a soft delete and changes the task status to ignored.

*/
func (a *Client) DeleteTask(params *DeleteTaskParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteTaskNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTaskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteTask",
		Method:             "DELETE",
		PathPattern:        "/tasks/{id}",
		ProducesMediaTypes: []string{"application/vnd.goswagger.examples.task-tracker.v1+json"},
		ConsumesMediaTypes: []string{"application/vnd.goswagger.examples.task-tracker.v1+json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteTaskReader{formats: a.formats},
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
	success, ok := result.(*DeleteTaskNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteTaskDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetTaskComments gets the comments for a task

  The comments require a size parameter.

*/
func (a *Client) GetTaskComments(params *GetTaskCommentsParams, opts ...ClientOption) (*GetTaskCommentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTaskCommentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTaskComments",
		Method:             "GET",
		PathPattern:        "/tasks/{id}/comments",
		ProducesMediaTypes: []string{"application/vnd.goswagger.examples.task-tracker.v1+json"},
		ConsumesMediaTypes: []string{"application/vnd.goswagger.examples.task-tracker.v1+json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTaskCommentsReader{formats: a.formats},
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
	success, ok := result.(*GetTaskCommentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetTaskCommentsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetTaskDetails gets the details for a task

  The details view has more information than the card view.
You can see who reported the issue and who last updated it when.

There are also comments for each issue.

*/
func (a *Client) GetTaskDetails(params *GetTaskDetailsParams, opts ...ClientOption) (*GetTaskDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTaskDetailsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTaskDetails",
		Method:             "GET",
		PathPattern:        "/tasks/{id}",
		ProducesMediaTypes: []string{"application/vnd.goswagger.examples.task-tracker.v1+json"},
		ConsumesMediaTypes: []string{"application/vnd.goswagger.examples.task-tracker.v1+json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTaskDetailsReader{formats: a.formats},
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
	success, ok := result.(*GetTaskDetailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetTaskDetailsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListTasks lists the tasks

  Allows for specifying a number of filter parameters to
narrow down the results.
Also allows for specifying a **sinceId** and **pageSize** parameter
to page through large result sets.

*/
func (a *Client) ListTasks(params *ListTasksParams, opts ...ClientOption) (*ListTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListTasksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listTasks",
		Method:             "GET",
		PathPattern:        "/tasks",
		ProducesMediaTypes: []string{"application/vnd.goswagger.examples.task-tracker.v1+json"},
		ConsumesMediaTypes: []string{"application/vnd.goswagger.examples.task-tracker.v1+json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListTasksReader{formats: a.formats},
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
	success, ok := result.(*ListTasksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListTasksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateTask updates the details for a task

  Allows for updating a task.
This operation requires authentication so that we know which user
last updated the task.

*/
func (a *Client) UpdateTask(params *UpdateTaskParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateTaskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateTask",
		Method:             "PUT",
		PathPattern:        "/tasks/{id}",
		ProducesMediaTypes: []string{"application/vnd.goswagger.examples.task-tracker.v1+json"},
		ConsumesMediaTypes: []string{"application/vnd.goswagger.examples.task-tracker.v1+json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateTaskReader{formats: a.formats},
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
	success, ok := result.(*UpdateTaskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateTaskDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UploadTaskFile adds a file to a task

  The file can't be larger than **5MB**
*/
func (a *Client) UploadTaskFile(params *UploadTaskFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UploadTaskFileCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadTaskFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "uploadTaskFile",
		Method:             "POST",
		PathPattern:        "/tasks/{id}/files",
		ProducesMediaTypes: []string{"application/vnd.goswagger.examples.task-tracker.v1+json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UploadTaskFileReader{formats: a.formats},
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
	success, ok := result.(*UploadTaskFileCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UploadTaskFileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
