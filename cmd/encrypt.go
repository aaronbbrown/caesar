package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aaronbbrown/caesar/pkg/cipher"
	"github.com/spf13/cobra"
)

var encryptKey *int

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt mode",
	Long:  `Encrypt a string using the Caesar cipher.`,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}

		caesar := cipher.NewCaesar(string(data), cipher.NewCaesarKey(*encryptKey))
		encrypted := caesar.Encrypt()
		fmt.Println(encrypted)
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)
	encryptKey = encryptCmd.Flags().Int("key", 0, "The encryption key (ignored in crack mode)")
	encryptCmd.MarkFlagRequired("key")
}
