// Code generated by "go.opentelemetry.io/collector/cmd/builder". DO NOT EDIT.

package main

import (
	"go.opentelemetry.io/collector/component"
	lokiexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/lokiexporter"
	groupbyattrsprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbyattrsprocessor"
	filterprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor"
	memorylimiterprocessor "go.opentelemetry.io/collector/processor/memorylimiterprocessor"
	batchprocessor "go.opentelemetry.io/collector/processor/batchprocessor"
	fluentforwardreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/fluentforwardreceiver"
)

func components() (component.Factories, error) {
	var err error
	factories := component.Factories{}

	factories.Extensions, err = component.MakeExtensionFactoryMap(
	)
	if err != nil {
		return component.Factories{}, err
	}

	factories.Receivers, err = component.MakeReceiverFactoryMap(
		fluentforwardreceiver.NewFactory(),
	)
	if err != nil {
		return component.Factories{}, err
	}

	factories.Exporters, err = component.MakeExporterFactoryMap(
		lokiexporter.NewFactory(),
	)
	if err != nil {
		return component.Factories{}, err
	}

	factories.Processors, err = component.MakeProcessorFactoryMap(
		groupbyattrsprocessor.NewFactory(),
		filterprocessor.NewFactory(),
		memorylimiterprocessor.NewFactory(),
		batchprocessor.NewFactory(),
	)
	if err != nil {
		return component.Factories{}, err
	}

	return factories, nil
}
