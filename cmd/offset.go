package cmd

import (
	"fmt"

	"github.com/aaronbbrown/caesar/pkg/cipher"
	"github.com/spf13/cobra"
)

var debugKey *int

// offsetCmd represents the key command
var offsetCmd = &cobra.Command{
	Use:   "offset",
	Short: "debug keys and offsets",
	Run: func(cmd *cobra.Command, args []string) {
		key := cipher.NewCaesarKey(*debugKey)
		fmt.Println("Key: ", key.String())
	},
}

func init() {
	rootCmd.AddCommand(offsetCmd)
	debugKey = offsetCmd.Flags().Int("key", 0, "The key to debug")
	offsetCmd.MarkFlagRequired("key")
}
