package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aaronbbrown/caesar/pkg/cipher"
	"github.com/spf13/cobra"
)

var decryptKey *int

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt mode",
	Long:  `Decrypt a string encrypted with the Caesar cipher`,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}

		caesar := cipher.NewCaesar(string(data), cipher.NewCaesarKey(*decryptKey))
		decrypted := caesar.Decrypt()
		fmt.Println(decrypted)
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)
	decryptKey = decryptCmd.Flags().Int("key", 0, "The encryption key (ignored in crack mode)")
	decryptCmd.MarkFlagRequired("key")
}
