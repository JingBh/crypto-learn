package controllers

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"github.com/JingBh/crypto-learn/pkg/aes"
	"github.com/JingBh/crypto-learn/pkg/des"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type cipherRequest struct {
	Type           string `json:"type"`
	Mode           string `json:"mode"`
	Key            string `json:"key"`
	KeyType        string `json:"key_type"`
	Iv             string `json:"iv"`
	IvType         string `json:"iv_type"`
	PlainText      string `json:"text"`
	PlainTextType  string `json:"text_type"`
	CipherText     string `json:"cipher"`
	CipherTextType string `json:"cipher_type"`
}

type cipherResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Data    map[string]string `json:"data"`
}

func decodeValue(value, encoding string) ([]byte, error) {
	if len(value) == 0 {
		return []byte{}, nil
	}
	switch encoding {
	case "base64":
		return base64.StdEncoding.DecodeString(value)
	case "hex":
		return hex.DecodeString(value)
	}
	return []byte(value), nil
}

func encodeValue(value []byte, encoding string) string {
	switch encoding {
	case "base64":
		return base64.StdEncoding.EncodeToString(value)
	case "hex":
		return hex.EncodeToString(value)
	}
	return string(value)
}

func sendError(w http.ResponseWriter, msg string) {
	resp := cipherResponse{
		Success: false,
		Message: msg,
		Data:    nil,
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func PostCipher(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		sendError(w, "Invalid request body")
		return
	}

	data := new(cipherRequest)
	err = json.Unmarshal(body, data)
	if err != nil {
		sendError(w, "Invalid request body")
		return
	}

	data.Type = strings.ToLower(data.Type)
	if data.Type != "encrypt" && data.Type != "decrypt" {
		sendError(w, "Invalid operation type")
		return
	}

	data.Mode = strings.ToUpper(data.Mode)
	if data.Mode != "ECB" && data.Mode != "CBC" {
		sendError(w, "Invalid mode of operation")
		return
	}

	key, err := decodeValue(data.Key, data.KeyType)
	if err != nil {
		sendError(w, "Invalid key encoding")
		return
	}
	if len(key) != 8 && len(key) != 16 && len(key) != 24 && len(key) != 32 {
		sendError(w, "For DES, the key must be 64-bit in length; for AES, the key must be 128, 192, or 256-bit in length")
		return
	}

	iv, err := decodeValue(data.Iv, data.IvType)
	if data.Mode == "CBC" {
		if err != nil {
			sendError(w, "Invalid iv encoding")
			return
		}
		if len(key) == 8 && len(iv) != 8 {
			sendError(w, "The DES iv must be 64-bit in length")
			return
		}
		if len(key) != 8 && len(iv) != 16 {
			sendError(w, "The AES iv must be 128-bit in length")
			return
		}
	}

	resData := make(map[string]string)
	if data.Type == "encrypt" {
		plainText, err := decodeValue(data.PlainText, data.PlainTextType)
		if err != nil {
			sendError(w, "Invalid plaintext encoding")
			return
		}

		var cipherText []byte
		if data.Mode == "ECB" {
			if len(key) == 8 {
				cipherText = des.ECB(key).Encipher(plainText)
			} else {
				cipherText = aes.ECB(key).Encipher(plainText)
			}
		} else if data.Mode == "CBC" {
			if len(key) == 8 {
				cipherText = des.CBC(key, iv).Encipher(plainText)
			} else {
				cipherText = aes.CBC(key, iv).Encipher(plainText)
			}
		} else {
			sendError(w, "Invalid mode of operation")
			return
		}
		resData["cipher"] = encodeValue(cipherText, data.CipherTextType)
	} else if data.Type == "decrypt" {
		cipherText, err := decodeValue(data.CipherText, data.CipherTextType)
		if err != nil {
			sendError(w, "Invalid ciphertext encoding")
			return
		}

		var plainText []byte
		if data.Mode == "ECB" {
			if len(key) == 8 {
				plainText = des.ECB(key).Decipher(cipherText)
			} else {
				plainText = aes.ECB(key).Decipher(cipherText)
			}
		} else if data.Mode == "CBC" {
			if len(key) == 8 {
				plainText = des.CBC(key, iv).Decipher(cipherText)
			} else {
				plainText = aes.CBC(key, iv).Decipher(cipherText)
			}
		} else {
			sendError(w, "Invalid mode of operation")
			return
		}
		resData["text"] = encodeValue(plainText, data.PlainTextType)
	} else {
		sendError(w, "Invalid operation type")
		return
	}

	resp := cipherResponse{
		Success: true,
		Data:    resData,
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetCipherKey(w http.ResponseWriter, r *http.Request) {
	length, err := strconv.Atoi(r.URL.Query().Get("l"))
	if err != nil {
		sendError(w, "Invalid key length")
		return
	}

	var key []byte
	if length == 64 {
		key = des.GenerateKey()
	} else if length == 128 || length == 192 || length == 256 {
		key = aes.GenerateKey(length)
	} else {
		sendError(w, "Invalid key length")
		return
	}

	resp := cipherResponse{
		Success: true,
		Data: map[string]string{
			"key":      encodeValue(key, "hex"),
			"key_type": "hex",
		},
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
