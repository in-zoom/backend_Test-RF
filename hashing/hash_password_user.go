package hashing

import (
	"github.com/joho/godotenv"
	"encoding/hex"
	"crypto/md5"
	"os"
)

func HashPasswordUser(password string) (hashPassword string, err error) {
	err = godotenv.Load(".env") //Загрузить файл .env
	if err != nil {
		return "", err
	}
	salt := md5.Sum([]byte(os.Getenv("SALT")))
	saltHash := hex.EncodeToString(salt[:])
	hashPasswordSalt := password + saltHash

	hash := md5.Sum([]byte(hashPasswordSalt))
	return hex.EncodeToString(hash[:]), nil
}
