package observability

import (
	"context"
	"errors"
	"fmt"
	"github.com/arikkfir/command"
	"github.com/arikkfir/greenstar/backend/internal/util/version"
	"go.opentelemetry.io/contrib/exporters/autoexport"
	"go.opentelemetry.io/contrib/propagators/autoprop"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"os"
	"strconv"
)

type OTelHook struct {
	ServiceName string
	cleanups    []func(context.Context) error
}

func (h *OTelHook) PreRun(ctx context.Context) error {
	if disableOtelSdkEnvValue, ok := os.LookupEnv("OTEL_SDK_DISABLED"); ok {
		if disable, err := strconv.ParseBool(disableOtelSdkEnvValue); err != nil {
			return errors.New(fmt.Sprintf("failed to parse OTEL_SDK_DISABLED value: %s", disableOtelSdkEnvValue))
		} else if disable {
			return nil
		}
	}

	var cleanups []func(context.Context) error

	// shutdown calls cleanup functions registered via cleanups.
	// The errors from the calls are joined.
	// Each registered cleanup will be invoked once.
	shutdown := func(ctx context.Context) error {
		var err error
		for _, fn := range cleanups {
			err = errors.Join(err, fn(ctx))
		}
		return err
	}

	// Create resource
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName(h.ServiceName),
		semconv.ServiceVersion(version.Version),
	)

	// Set up propagator
	otel.SetTextMapPropagator(autoprop.NewTextMapPropagator())

	// Set up trace provider
	if traceExporter, err := autoexport.NewSpanExporter(context.Background()); err != nil {
		return errors.Join(err, shutdown(ctx))
	} else {
		traceProvider := trace.NewTracerProvider(trace.WithResource(res), trace.WithBatcher(traceExporter, trace.WithBlocking()))
		cleanups = append(cleanups, traceProvider.Shutdown)
		otel.SetTracerProvider(traceProvider)
	}

	// Set up meter provider
	if metricReader, err := autoexport.NewMetricReader(context.Background()); err != nil {
		return errors.Join(err, shutdown(ctx))
	} else {
		meterProvider := metric.NewMeterProvider(metric.WithResource(res), metric.WithReader(metricReader))
		cleanups = append(cleanups, meterProvider.Shutdown)
		otel.SetMeterProvider(meterProvider)
	}

	h.cleanups = cleanups
	return nil
}

func (h *OTelHook) PostRun(ctx context.Context, err error, _ command.ExitCode) error {
	var cleanupErrors error
	for _, cleanup := range h.cleanups {
		cleanupErrors = errors.Join(cleanupErrors, cleanup(ctx))
	}
	if cleanupErrors != nil {
		cleanupErrors = fmt.Errorf("failed to shutdown OTel SDK: %w", cleanupErrors)
	}
	return errors.Join(err, cleanupErrors)
}
