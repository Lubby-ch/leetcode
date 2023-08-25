package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func getPermutation(n int, k int) string {
	var ans []byte
	for i := 1; i <= n; i++ {
		ans = append(ans, byte(i))
	}
	var dfs func(idx int)
	fmt.Println(ans)
	dfs = func(idx int) {
		if idx == n-1 {
			return
		}
		for i := idx; i < n; i++ {
			if idx == i {
				dfs(idx+1)
				continue
			}
			ans[idx], ans[i] = ans[i], ans[idx]
			fmt.Println(ans)
			dfs(idx + 1)
			ans[idx], ans[i] = ans[i], ans[idx]
		}
	}
	dfs(0)
	return string(ans)
}

func decryptAES(encrypted string, key string) (string, error) {
	// 将加密后的数据解码为字节数组
	encryptedData, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", fmt.Errorf("failed to decode encrypted data: %v", err)
	}

	// 创建 AES 解密器
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("failed to create AES cipher: %v", err)
	}

	// 从加密数据中分离出 IV 和密文
	ivSize := block.BlockSize()
	if len(encryptedData) < ivSize {
		return "", fmt.Errorf("encrypted data is too short")
	}
	iv := encryptedData[:ivSize]
	ciphertext := encryptedData[ivSize:]

	// 创建 CBC 解密器
	mode := cipher.NewCBCDecrypter(block, iv)

	// 解密数据
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	// 返回解密后的数据
	return string(plaintext), nil
}

func main() {
	encrypted := "04Fo0+9L8t6oDnD6oOWctd5br/7/DFrnFjPLWaSj7eg="
	key := "27g8exBhY3WTYi7W"
	plaintext := CBCDecrypt(encrypted, key)
	fmt.Println(plaintext)

}

//解密
func CBCDecrypt(ciphertext string, key string) string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("cbc decrypt err:", err)
		}
	}()

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return ""
	}

	ciphercode, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return ""
	}

	iv := ciphercode[:aes.BlockSize]        // 密文的前 16 个字节为 iv
	ciphercode = ciphercode[aes.BlockSize:] // 正式密文

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphercode, ciphercode)

	plaintext := string(ciphercode) // ↓ 减去 padding
	return string(pkcs7UnPadding([]byte(plaintext)))
}

func pkcs7UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
