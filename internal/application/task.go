package application

import (
	"context"
)

type App struct {
	ctx         context.Context
	taskService TaskService
}
