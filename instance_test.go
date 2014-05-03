// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"flag"
	"testing"
)

var instance = flag.Bool("instance", false, "run instance tests, should be run inside CloudSigma server instance")

func TestReadContext(t *testing.T) {
	if !*instance {
		t.SkipNow()
		return
	}

	client, err := NewClient("zrh", "user", "password", nil)
	if err != nil {
		t.Error(err)
		return
	}

	ctx, err := client.Context()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%#v", ctx)
}
