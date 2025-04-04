// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Swagger Petstore - OpenAPI 3.0
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.11
 * Contact: apiteam@swagger.io
 */

package openapi

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	
	
)

//var _ = &reflect.Value{}

// TaskAPIServicer defines the api actions for the TaskAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type TaskAPIServicer interface { 
	GetTask(context.Context, string) (*GetTaskResult,  int,error)
	UpdateTask(context.Context, Task) (*UpdateTaskResult,  int,error)
	CreateTask(context.Context, Task) (*CreateTaskResult,  int,error)
	DeleteTask(context.Context, string) (*DeleteTaskResult,  int,error)
	GetTaskList(context.Context, int32, int32) (*GetTaskListResult,  int,error)
}

type ApiService interface { 

	// TaskAPIServicer defines the api actions for the TaskAPI service
	// This interface intended to stay up to date with the openapi yaml used to generate it,
	// while the service implementation can be ignored with the .openapi-generator-ignore file
	// and updated with the logic required for the API.
	TaskAPIServicer


	InstallMiddleware(engine *gin.Engine)
}


type Server struct {
	router *gin.Engine
	apiService ApiService
}

func NewServer(apiService ApiService) *Server {
	server := &Server{
		apiService: apiService,
	}

	fTaskAPIController := NewTaskAPIController(apiService)
	
	server.router = NewRouter(apiService, fTaskAPIController)
	return server
}


func (s *Server) Run(addr string) error {
    return s.router.Run(addr)
}


var Module = fx.Module("server",
	fx.Provide(NewServer),
	fx.Provide(NewApiService),
)
