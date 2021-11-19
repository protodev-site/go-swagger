// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/protodev-site/runtime"
	"github.com/protodev-site/runtime/middleware"
	"github.com/protodev-site/runtime/security"
	"github.com/protodev-site/spec"
	"github.com/protodev-site/loads"

	"github.com/protodev-site/go-swagger/examples/task-tracker/restapi/operations/tasks"
)

// NewTaskTrackerAPI creates a new TaskTracker instance
func NewTaskTrackerAPI(spec *loads.Document) *TaskTrackerAPI {
	return &TaskTrackerAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer:          runtime.JSONConsumer(),
		MultipartformConsumer: runtime.DiscardConsumer,

		JSONProducer: runtime.JSONProducer(),

		TasksAddCommentToTaskHandler: tasks.AddCommentToTaskHandlerFunc(func(params tasks.AddCommentToTaskParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation tasks.AddCommentToTask has not yet been implemented")
		}),
		TasksCreateTaskHandler: tasks.CreateTaskHandlerFunc(func(params tasks.CreateTaskParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation tasks.CreateTask has not yet been implemented")
		}),
		TasksDeleteTaskHandler: tasks.DeleteTaskHandlerFunc(func(params tasks.DeleteTaskParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation tasks.DeleteTask has not yet been implemented")
		}),
		TasksGetTaskCommentsHandler: tasks.GetTaskCommentsHandlerFunc(func(params tasks.GetTaskCommentsParams) middleware.Responder {
			return middleware.NotImplemented("operation tasks.GetTaskComments has not yet been implemented")
		}),
		TasksGetTaskDetailsHandler: tasks.GetTaskDetailsHandlerFunc(func(params tasks.GetTaskDetailsParams) middleware.Responder {
			return middleware.NotImplemented("operation tasks.GetTaskDetails has not yet been implemented")
		}),
		TasksListTasksHandler: tasks.ListTasksHandlerFunc(func(params tasks.ListTasksParams) middleware.Responder {
			return middleware.NotImplemented("operation tasks.ListTasks has not yet been implemented")
		}),
		TasksUpdateTaskHandler: tasks.UpdateTaskHandlerFunc(func(params tasks.UpdateTaskParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation tasks.UpdateTask has not yet been implemented")
		}),
		TasksUploadTaskFileHandler: tasks.UploadTaskFileHandlerFunc(func(params tasks.UploadTaskFileParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation tasks.UploadTaskFile has not yet been implemented")
		}),

		// Applies when the "token" query is set
		APIKeyAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (api_key) token from query param [token] has not yet been implemented")
		},
		// Applies when the "X-Token" header is set
		TokenHeaderAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (token_header) X-Token from header param [X-Token] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*TaskTrackerAPI This application implements a very simple issue tracker.
It's implemented as an API which is described by this swagger spec document.

The go-swagger project uses this specification to test the code generation.
This document contains all possible values for a swagger definition.
This means that it exercises the framework relatively well.
*/
type TaskTrackerAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/vnd.goswagger.examples.task-tracker.v1+json
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for the following mime types:
	//   - multipart/form-data
	MultipartformConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/vnd.goswagger.examples.task-tracker.v1+json
	JSONProducer runtime.Producer

	// APIKeyAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key token provided in the query
	APIKeyAuth func(string) (interface{}, error)

	// TokenHeaderAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key X-Token provided in the header
	TokenHeaderAuth func(string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// TasksAddCommentToTaskHandler sets the operation handler for the add comment to task operation
	TasksAddCommentToTaskHandler tasks.AddCommentToTaskHandler
	// TasksCreateTaskHandler sets the operation handler for the create task operation
	TasksCreateTaskHandler tasks.CreateTaskHandler
	// TasksDeleteTaskHandler sets the operation handler for the delete task operation
	TasksDeleteTaskHandler tasks.DeleteTaskHandler
	// TasksGetTaskCommentsHandler sets the operation handler for the get task comments operation
	TasksGetTaskCommentsHandler tasks.GetTaskCommentsHandler
	// TasksGetTaskDetailsHandler sets the operation handler for the get task details operation
	TasksGetTaskDetailsHandler tasks.GetTaskDetailsHandler
	// TasksListTasksHandler sets the operation handler for the list tasks operation
	TasksListTasksHandler tasks.ListTasksHandler
	// TasksUpdateTaskHandler sets the operation handler for the update task operation
	TasksUpdateTaskHandler tasks.UpdateTaskHandler
	// TasksUploadTaskFileHandler sets the operation handler for the upload task file operation
	TasksUploadTaskFileHandler tasks.UploadTaskFileHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *TaskTrackerAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *TaskTrackerAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *TaskTrackerAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *TaskTrackerAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *TaskTrackerAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *TaskTrackerAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *TaskTrackerAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *TaskTrackerAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *TaskTrackerAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the TaskTrackerAPI
func (o *TaskTrackerAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}
	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.APIKeyAuth == nil {
		unregistered = append(unregistered, "TokenAuth")
	}
	if o.TokenHeaderAuth == nil {
		unregistered = append(unregistered, "XTokenAuth")
	}

	if o.TasksAddCommentToTaskHandler == nil {
		unregistered = append(unregistered, "tasks.AddCommentToTaskHandler")
	}
	if o.TasksCreateTaskHandler == nil {
		unregistered = append(unregistered, "tasks.CreateTaskHandler")
	}
	if o.TasksDeleteTaskHandler == nil {
		unregistered = append(unregistered, "tasks.DeleteTaskHandler")
	}
	if o.TasksGetTaskCommentsHandler == nil {
		unregistered = append(unregistered, "tasks.GetTaskCommentsHandler")
	}
	if o.TasksGetTaskDetailsHandler == nil {
		unregistered = append(unregistered, "tasks.GetTaskDetailsHandler")
	}
	if o.TasksListTasksHandler == nil {
		unregistered = append(unregistered, "tasks.ListTasksHandler")
	}
	if o.TasksUpdateTaskHandler == nil {
		unregistered = append(unregistered, "tasks.UpdateTaskHandler")
	}
	if o.TasksUploadTaskFileHandler == nil {
		unregistered = append(unregistered, "tasks.UploadTaskFileHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *TaskTrackerAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *TaskTrackerAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "api_key":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.APIKeyAuth)

		case "token_header":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.TokenHeaderAuth)

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *TaskTrackerAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *TaskTrackerAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/vnd.goswagger.examples.task-tracker.v1+json":
			result["application/vnd.goswagger.examples.task-tracker.v1+json"] = o.JSONConsumer
		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *TaskTrackerAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/vnd.goswagger.examples.task-tracker.v1+json":
			result["application/vnd.goswagger.examples.task-tracker.v1+json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *TaskTrackerAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the task tracker API
func (o *TaskTrackerAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *TaskTrackerAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/tasks/{id}/comments"] = tasks.NewAddCommentToTask(o.context, o.TasksAddCommentToTaskHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/tasks"] = tasks.NewCreateTask(o.context, o.TasksCreateTaskHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/tasks/{id}"] = tasks.NewDeleteTask(o.context, o.TasksDeleteTaskHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tasks/{id}/comments"] = tasks.NewGetTaskComments(o.context, o.TasksGetTaskCommentsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tasks/{id}"] = tasks.NewGetTaskDetails(o.context, o.TasksGetTaskDetailsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tasks"] = tasks.NewListTasks(o.context, o.TasksListTasksHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/tasks/{id}"] = tasks.NewUpdateTask(o.context, o.TasksUpdateTaskHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/tasks/{id}/files"] = tasks.NewUploadTaskFile(o.context, o.TasksUploadTaskFileHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *TaskTrackerAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *TaskTrackerAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *TaskTrackerAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *TaskTrackerAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *TaskTrackerAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
