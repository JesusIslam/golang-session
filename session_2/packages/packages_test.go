package packages

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrintingMachine(t *testing.T) {
	t.Run("Test printing data", func(t *testing.T) {
		p := &PrintingMachine{}
		p.PublicData = "data"
		p.Print()

		require.True(t, p.Printed)
	})
}
