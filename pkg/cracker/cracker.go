package cracker

import (
	"fmt"
	"sync"

	"github.com/aaronbbrown/caesar/pkg/cipher"
)

type Cracker struct {
	msg        string
	dictionary Dictionary
}

func NewCracker(msg string, dictionary Dictionary) *Cracker {
	return &Cracker{
		msg:        msg,
		dictionary: dictionary,
	}
}

func (c *Cracker) Crack(minKey, maxKey, percentWords, parallelism int) (*cipher.Caesar, error) {
	wg := sync.WaitGroup{}
	keyC := make(chan int)
	decryptedC := make(chan *cipher.Caesar, 1)

	foundC := make(chan struct{})

	for i := 0; i < parallelism; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			for key := range keyC {
				caesar := cipher.NewCaesar(c.msg, cipher.NewCaesarKey(key))
				decrypted := caesar.Decrypt()
				percent := c.dictionary.PercentMatches(decrypted)
				if percent >= percentWords {
					close(foundC)
					// This sends to a buffered channel so that the go routine
					// will exit before the value has been read, allowing us
					// to signal that an answer was found and wait for all the
					// go routines to exit before checking the actual value.
					decryptedC <- cipher.NewCaesar(decrypted, cipher.NewCaesarKey(key))
					return
				}
			}
		}()
	}

	for i := minKey; i < maxKey; i++ {
		select {
		case keyC <- i:
			continue
		case <-foundC:
			break
		}
	}
	close(keyC)
	wg.Wait()

	select {
	case <-foundC:
		decrypted := <-decryptedC
		return decrypted, nil
	default:
		break
	}

	return nil, fmt.Errorf("Unable to decrypt")
}
