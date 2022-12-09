// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build windows
// +build windows

package sqlserverreceiver

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/comparetest"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/comparetest/golden"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/winperfcounters"
)

// MockPerfCounterWatcher is an autogenerated mock type for the PerfCounterWatcher type
type MockPerfCounterWatcher struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m MockPerfCounterWatcher) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Path provides a mock function with given fields:
func (_m MockPerfCounterWatcher) Path() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ScrapeData provides a mock function with given fields:
func (_m MockPerfCounterWatcher) ScrapeData() ([]winperfcounters.CounterValue, error) {
	ret := _m.Called()

	var r0 []winperfcounters.CounterValue
	if rf, ok := ret.Get(0).(func() []winperfcounters.CounterValue); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]winperfcounters.CounterValue)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func TestSqlServerScraper(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig().(*Config)
	logger, obsLogs := observer.New(zap.WarnLevel)
	settings := componenttest.NewNopReceiverCreateSettings()
	settings.Logger = zap.New(logger)
	s := newSqlServerScraper(settings, cfg)

	s.start(context.Background(), nil)
	assert.Equal(t, 0, len(s.watcherRecorders))
	assert.Equal(t, 21, obsLogs.Len())
	assert.Equal(t, 21, obsLogs.FilterMessageSnippet("failed to create perf counter with path \\SQLServer:").Len())
	assert.Equal(t, 21, obsLogs.FilterMessageSnippet("The specified object was not found on the computer.").Len())
	assert.Equal(t, 1, obsLogs.FilterMessageSnippet("\\SQLServer:General Statistics\\").Len())
	assert.Equal(t, 3, obsLogs.FilterMessageSnippet("\\SQLServer:SQL Statistics\\").Len())
	assert.Equal(t, 2, obsLogs.FilterMessageSnippet("\\SQLServer:Locks(_Total)\\").Len())
	assert.Equal(t, 6, obsLogs.FilterMessageSnippet("\\SQLServer:Buffer Manager\\").Len())
	assert.Equal(t, 1, obsLogs.FilterMessageSnippet("\\SQLServer:Access Methods(_Total)\\").Len())
	assert.Equal(t, 8, obsLogs.FilterMessageSnippet("\\SQLServer:Databases(*)\\").Len())

	metrics, err := s.scrape(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 0, metrics.ResourceMetrics().Len())

	err = s.shutdown(context.Background())
	require.NoError(t, err)
}

var goldenScrapePath = filepath.Join("testdata", "golden_scrape.json")
var dbInstance = "db-instance"

func TestScrape(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig().(*Config)
	settings := componenttest.NewNopReceiverCreateSettings()
	scraper := newSqlServerScraper(settings, cfg)

	for i, rec := range perfCounterRecorders {
		perfCounterWatcher := &MockPerfCounterWatcher{}
		perfCounterWatcher.On("ScrapeData").Return([]winperfcounters.CounterValue{{InstanceName: dbInstance, Value: float64(i)}}, nil)
		i++
		for _, recorder := range rec.recorders {
			scraper.watcherRecorders = append(scraper.watcherRecorders, watcherRecorder{
				watcher:  *perfCounterWatcher,
				recorder: recorder,
			})
		}
	}

	scrapeData, err := scraper.scrape(context.Background())
	require.NoError(t, err)

	expectedMetrics, err := golden.ReadMetrics(goldenScrapePath)
	require.NoError(t, err)

	err = comparetest.CompareMetrics(expectedMetrics, scrapeData)
	require.NoError(t, err)
}
