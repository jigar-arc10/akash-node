package ipoperator

import (
	provider_flags "github.com/ovrclk/akash/provider/cmd/flags"
	"github.com/ovrclk/akash/provider/operator/operatorcommon"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "ip-operator",
		Short:        "kubernetes operator interfacing with Metal LB",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doIPOperator(cmd)
		},
	}
	operatorcommon.AddOperatorFlags(cmd, "0.0.0.0:8086")
	operatorcommon.AddIgnoreListFlags(cmd)

	if err := provider_flags.AddServiceEndpointFlag(cmd, serviceProvider); err != nil {
		return nil
	}

	if err := provider_flags.AddServiceEndpointFlag(cmd, serviceMetalLb); err != nil {
		return nil
	}
	err := provider_flags.AddKubeConfigPathFlag(cmd)
	if err != nil {
		panic(err)
	}

	return cmd
}
