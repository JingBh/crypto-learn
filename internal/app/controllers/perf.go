package controllers

import (
	"crypto/rand"
	"encoding/json"
	"github.com/JingBh/crypto-learn/pkg/aes"
	"github.com/JingBh/crypto-learn/pkg/des"
	"net/http"
	"sync"
	"time"
)

type perfResponse struct {
	Sizes      []string      `json:"sizes"`
	Algorithms []string      `json:"algorithms"`
	Data       [][][]float64 `json:"data"` // [size][algorithm][encrypt, decrypt] (avg time)
}

type perfSize struct {
	name string
	size int
}

type perfAlgorithm struct {
	name        string
	generateKey func() []byte
	encipher    func([]byte, []byte) []byte
	decipher    func([]byte, []byte) []byte
}

func PostPerf(w http.ResponseWriter, r *http.Request) {
	sizes := []perfSize{
		{"1KB", 1024},
		{"16KB", 16384},
		{"256KB", 262144},
	}

	algorithms := []perfAlgorithm{
		{
			"DES",
			des.GenerateKey,
			func(input []byte, key []byte) []byte {
				return des.ECB(key).Encipher(input)
			},
			func(input []byte, key []byte) []byte {
				return des.ECB(key).Decipher(input)
			},
		},
		{
			"AES-128",
			func() []byte {
				return aes.GenerateKey(128)
			},
			func(input []byte, key []byte) []byte {
				return aes.ECB(key).Encipher(input)
			},
			func(input []byte, key []byte) []byte {
				return aes.ECB(key).Decipher(input)
			},
		},
		{
			"AES-192",
			func() []byte {
				return aes.GenerateKey(192)
			},
			func(input []byte, key []byte) []byte {
				return aes.ECB(key).Encipher(input)
			},
			func(input []byte, key []byte) []byte {
				return aes.ECB(key).Decipher(input)
			},
		},
		{
			"AES-256",
			func() []byte {
				return aes.GenerateKey(256)
			},
			func(input []byte, key []byte) []byte {
				return aes.ECB(key).Encipher(input)
			},
			func(input []byte, key []byte) []byte {
				return aes.ECB(key).Decipher(input)
			},
		},
	}

	perf := func(f func()) float64 {
		const N = 10
		var sum float64
		for i := 0; i < N; i++ {
			start := time.Now()
			f()
			sum += time.Since(start).Seconds()
		}
		return sum / N
	}

	result := perfResponse{
		make([]string, len(sizes)),
		make([]string, len(algorithms)),
		make([][][]float64, len(sizes)),
	}

	for i, algorithm := range algorithms {
		result.Algorithms[i] = algorithm.name
	}

	var wg sync.WaitGroup
	mu := &sync.Mutex{}

	for i, size := range sizes {
		result.Sizes[i] = size.name

		data := make([][]float64, len(result.Algorithms))
		result.Data[i] = data

		for j, algorithm := range algorithms {
			wg.Add(1)
			go func(i int, j int, algorithm perfAlgorithm) {
				defer wg.Done()

				key := algorithm.generateKey()
				text := make([]byte, size.size)
				_, err := rand.Read(text)
				if err != nil {
					panic(err)
				}
				cipher := algorithm.encipher(text, key)

				res := []float64{
					perf(func() {
						algorithm.encipher(text, key)
					}),
					perf(func() {
						algorithm.decipher(cipher, key)
					}),
				}

				mu.Lock()
				result.Data[i][j] = res
				mu.Unlock()
			}(i, j, algorithm)
		}
	}
	wg.Wait()

	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
