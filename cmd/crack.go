package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"

	"github.com/aaronbbrown/caesar/pkg/cracker"
	"github.com/spf13/cobra"
)

var (
	maxKeys      *int
	percentWords *int
	parallelism  *int
)

// crackCmd represents the crack command
var crackCmd = &cobra.Command{
	Use:   "crack",
	Short: "Crack mode",
	Long:  `Crack a Caesar cipher`,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}

		dictionary, err := cracker.NewDictionary("wordlist.txt")
		if err != nil {
			log.Fatal(err)
		}

		c := cracker.NewCracker(string(data), dictionary)

		decrypted, err := c.Crack(1, *maxKeys, *percentWords, *parallelism)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Key: ", decrypted.Key.String())
		fmt.Printf("Message:\n%s", decrypted.Msg)
	},
}

func init() {
	rootCmd.AddCommand(crackCmd)
	maxKeys = crackCmd.Flags().Int("max-attempts", 1000, "maximum number of attempts to crack")
	percentWords = crackCmd.Flags().Int("percent-words", 30, "minimum percentage of strings that must be words in the dictionary")
	parallelism = crackCmd.Flags().Int("parallelism", runtime.GOMAXPROCS(0), "Number of parallel threads to run during cracking")

	crackCmd.MarkFlagRequired("max-attempts")
	crackCmd.MarkFlagRequired("percent-words")
}
