package main

import (
	"io/ioutil"
	"testing"

	"github.com/aaronbbrown/caesar/pkg/cipher"
	"github.com/aaronbbrown/caesar/pkg/cracker"
)

func BenchmarkCrack1SM(b *testing.B)   { benchmarkCrack(1, 1, b) }
func BenchmarkCrack10SM(b *testing.B)  { benchmarkCrack(10, 1, b) }
func BenchmarkCrack100SM(b *testing.B) { benchmarkCrack(100, 1, b) }

func BenchmarkCrack1KP1(b *testing.B)   { benchmarkCrack(1000, 1, b) }
func BenchmarkCrack10KP1(b *testing.B)  { benchmarkCrack(10000, 1, b) }
func BenchmarkCrack100KP1(b *testing.B) { benchmarkCrack(100000, 1, b) }
func BenchmarkCrack1MP1(b *testing.B)   { benchmarkCrack(1000000, 1, b) }

func BenchmarkCrack1KP2(b *testing.B)   { benchmarkCrack(1000, 2, b) }
func BenchmarkCrack10KP2(b *testing.B)  { benchmarkCrack(10000, 2, b) }
func BenchmarkCrack100KP2(b *testing.B) { benchmarkCrack(100000, 2, b) }
func BenchmarkCrack1MP2(b *testing.B)   { benchmarkCrack(1000000, 2, b) }

func BenchmarkCrack1KP4(b *testing.B)   { benchmarkCrack(1000, 4, b) }
func BenchmarkCrack10KP4(b *testing.B)  { benchmarkCrack(10000, 4, b) }
func BenchmarkCrack100KP4(b *testing.B) { benchmarkCrack(100000, 4, b) }
func BenchmarkCrack1MP4(b *testing.B)   { benchmarkCrack(1000000, 4, b) }

func BenchmarkCrack1KP8(b *testing.B)   { benchmarkCrack(1000, 8, b) }
func BenchmarkCrack10KP8(b *testing.B)  { benchmarkCrack(10000, 8, b) }
func BenchmarkCrack100KP8(b *testing.B) { benchmarkCrack(100000, 8, b) }
func BenchmarkCrack1MP8(b *testing.B)   { benchmarkCrack(1000000, 8, b) }

func benchmarkCrack(key int, threads int, b *testing.B) {
	data, err := ioutil.ReadFile("text.txt")
	if err != nil {
		b.Fatal("Could not load text.txt")
	}

	dictionary, err := cracker.NewDictionary("wordlist.txt")
	if err != nil {
		b.Fatal("Could not load wordlist.txt")
	}

	caesar := cipher.NewCaesar(string(data), cipher.NewCaesarKey(key))
	encrypted := caesar.Encrypt()
	c := cracker.NewCracker(encrypted, dictionary)

	for n := 0; n < b.N; n++ {
		_, err := c.Crack(1, key+1, 30, threads)
		if err != nil {
			b.Fatal("Unable to decrypt")
		}
	}
}
