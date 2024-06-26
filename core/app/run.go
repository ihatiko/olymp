package app

import (
	"context"
	"os"
	"strconv"

	"github.com/ihatiko/olymp/core/iface"
	"github.com/ihatiko/olymp/core/store"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

type Option func(*App)

type App struct {
	context    context.Context
	Components []iface.IComponent
}
type SharedComponents map[string][]iface.IComponent

func Modules(components ...iface.IComponent) {
	app := new(App)
	app.context = context.Background()
	buffer := map[string]struct{}{}
	for _, component := range components {
		if _, ok := buffer[component.Name()]; !ok {
			buffer[component.Name()] = struct{}{}
			app.Components = append(app.Components, component)
		}
	}
	for _, pkg := range store.PackageStore.Get() {
		level := zap.FatalLevel
		if env := os.Getenv("TECH_SERVICE_DEBUG"); env != "" {
			if state, err := strconv.ParseBool(env); err == nil {
				if state {
					level = zap.ErrorLevel
				}
			}
		}
		if pkg.HasError() {
			otelzap.S().Logf(level, "init package %s %v", pkg.Name(), pkg.Error())
		}
	}
	for _, component := range app.Components {
		store.LivenessStore.Load(component)
		if component == nil {
			otelzap.L().Fatal("empty struct [func Deployment(components ...iface.IComponent)]")
			return
		}
		go func(component iface.IComponent) {
			defer func() {
				if r := recover(); r != nil {
					otelzap.L().Error("recovered from panic", zap.Any("recover", r))
				}
			}()
			component.Run()
		}(component)
	}
	app.Graceful(app.Components)
}
