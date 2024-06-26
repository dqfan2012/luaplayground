package app

import (
	"context"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) StartUp(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) RunLua(script string) (string, error) {
	result, err := RunLuaScript(script)
	if err != nil {
		return "", err
	}

	return result, nil
}
