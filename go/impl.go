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
	"github.com/sherlockhua/koala/middleware/logid"
	"github.com/gin-gonic/gin"
	
)

//var _ = &reflect.Value{}

// TaskAPIServicer defines the api actions for the TaskAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.

type ApiServiceImp struct { 
}


func NewApiService()ApiService {
	return &ApiServiceImp{}
}


func (p *ApiServiceImp) InstallMiddleware(engine *gin.Engine) {
	// Add middleware here
	//engine.Use(gin.Logger())
    //engine.Use(gin.Recovery())
    //engine.Use(cors.Default())
    //engine.Use(gzip.Gzip(gzip.DefaultCompression))
    //engine.Use(jwt.AuthMiddleware(jwt.New(jwt.SigningMethodHS256, []byte("your_secret_key"))))
    //engine.Use(httprate.Limit(httprate.NewLimiter(httprate.Every(time.Minute), 100)))
	engine.Use(logid.LogIDMiddleware())
}


//TaskAPIServicer
func (p *ApiServiceImp) GetTask(ctx context.Context, taskId string) (result *GetTaskResult, code int, err error) {
	return 
}

func (p *ApiServiceImp) UpdateTask(ctx context.Context, task Task) (result *UpdateTaskResult, code int, err error) {
	return 
}

func (p *ApiServiceImp) CreateTask(ctx context.Context, task Task) (result *CreateTaskResult, code int, err error) {
	return 
}

func (p *ApiServiceImp) DeleteTask(ctx context.Context, taskId string) (result *DeleteTaskResult, code int, err error) {
	return 
}

func (p *ApiServiceImp) GetTaskList(ctx context.Context, offset int32, size int32) (result *GetTaskListResult, code int, err error) {
	return 
}




