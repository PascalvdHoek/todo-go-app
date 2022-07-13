package handlers

import (
	"context"
	"go-todo-list/database"
)

var (
	db  = database.ConnectBunDb()
	ctx = context.Background()
	err error
)
