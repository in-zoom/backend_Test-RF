package password

import(
	"crypto/md5"
	"encoding/hex"
	"github.com/joho/godotenv"
	"os"
)
func PreparationPasswordUser(password string) (hashPassword string, err error) {
	err = godotenv.Load(".env") //Загрузить файл .env
	if err != nil {
		//fmt.Print(e)
		return "", err
	}
	salt := md5.Sum([]byte(os.Getenv("salt")))
	saltHash := hex.EncodeToString(salt[:])
	hashPasswordSalt := password + saltHash
	
	hash := md5.Sum([]byte(hashPasswordSalt))
	return hex.EncodeToString(hash[:]), nil
}