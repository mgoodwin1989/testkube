package validator

import (
	"errors"

	"github.com/spf13/cobra"
)

func ExecutorName(cmd *cobra.Command, args []string) error {
	// TODO validate executor name as valid kubernetes resource name

	if len(args) < 1 {
		return errors.New("please pass valid executor-name")
	}
	return nil
}
