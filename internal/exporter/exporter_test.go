// Copyright 2026 The NATS Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package exporter

import (
	"testing"
	"time"

	"github.com/nats-io/jsm.go/natscontext"
	"github.com/nats-io/nats.go"
)

func TestExporterConnectTimeout(t *testing.T) {
	nctx, err := natscontext.New("test", false)
	if err != nil {
		t.Fatalf("creating test context: %v", err)
	}

	resolve := func(t *testing.T, exporter *Exporter) nats.Options {
		t.Helper()

		opts, err := exporter.natsOptions(nctx)
		if err != nil {
			t.Fatalf("building connection options: %v", err)
		}

		res := nats.GetDefaultOptions()
		for _, o := range opts {
			if err := o(&res); err != nil {
				t.Fatalf("applying connection option: %v", err)
			}
		}

		return res
	}

	if to := resolve(t, &Exporter{connectTimeout: 30 * time.Second}).Timeout; to != 30*time.Second {
		t.Fatalf("expected a 30s connect timeout, got %v", to)
	}

	// unset should leave the nats.go default in place
	if to := resolve(t, &Exporter{}).Timeout; to != nats.GetDefaultOptions().Timeout {
		t.Fatalf("expected the default connect timeout of %v, got %v", nats.GetDefaultOptions().Timeout, to)
	}
}
