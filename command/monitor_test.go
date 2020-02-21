package command

import (
	"strings"
	"testing"

	"github.com/mitchellh/cli"
)

func testMonitorCommand(tb testing.TB) (*cli.MockUi, *MonitorCommand) {
	tb.Helper()

	ui := cli.NewMockUi()
	return ui, &MonitorCommand{
		BaseCommand: &BaseCommand{
			UI: ui,
		},
	}
}

func TestMonitorCommand_Run(t *testing.T) {
	t.Parallel()

	t.Run("output", func(t *testing.T) {
		t.Parallel()

		client, closer := testVaultServer(t)
		defer closer()

		ui, cmd := testMonitorCommand(t)
		cmd.client = client

		code := cmd.Run(nil)
		if exp := 0; code != exp {
			t.Errorf("expected %d to be %d", code, exp)
		}

		// some stuff here about what we expect in the UI after logs are streaming in
	})

	t.Run("no_tabs", func(t *testing.T) {
		t.Parallel()

		_, cmd := testMonitorCommand(t)
		assertNoTabs(t, cmd)
	})
}
