package clistate

import (
	"testing"

	"github.com/kevinklinger/open_terraform/noninternal/command/arguments"
	"github.com/kevinklinger/open_terraform/noninternal/command/views"
	"github.com/kevinklinger/open_terraform/noninternal/states/statemgr"
	"github.com/kevinklinger/open_terraform/noninternal/terminal"
)

func TestUnlock(t *testing.T) {
	streams, _ := terminal.StreamsForTesting(t)
	view := views.NewView(streams)

	l := NewLocker(0, views.NewStateLocker(arguments.ViewHuman, view))
	l.Lock(statemgr.NewUnlockErrorFull(nil, nil), "test-lock")

	diags := l.Unlock()
	if diags.HasErrors() {
		t.Log(diags.Err().Error())
	} else {
		t.Error("expected error")
	}
}
