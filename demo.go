package main
import (
	"log"
	"sm4-goland/security"
	"time"
)

func main()  {

	var plaintext = []byte{84,50,48,48,48,50,0,0,0,0,0,0,0,0,222,38,30,110,78,125,132,94,171,232,177,163,50,156,147,35,0,46,96,88,93,242,66,52,41,134,87,175,166,250,48,44}
	var key = []byte{87,240,113,33,159,120,209,57,94,98,2,111,66,72,127,109}
	var iv = []byte{153,170,62,104,237,129,115,160,238,208,102,132}
	ciphertext := make([]byte, 4096)
	tag := make([]byte, 1024)
  // encrypt
	task := security.Unit{
		Key:    key,
		Iv:     iv,
		Ct:     ciphertext[:46],
		Pt:     plaintext[:46],
		Tag:    tag[:16],
		TagLen: 16,
	}
	startTime := time.Now().UnixNano()
	err := security.Encrypt(0, &task)
	if err != nil {
		log.Print("encrypt : ", err)
	}
	endTime := time.Now().UnixNano()
	log.Print("encrypt : ", float64((endTime - startTime) / 1e3), " us")
  // decrypt
	task1 := security.Unit{
		Key: key,
		Iv: iv,
		Ct: task.Ct,
		Pt: task.Pt,
		Tag: tag,
		TagLen: 16,
	}
	startTime = time.Now().UnixNano()
	err = security.Decrypt(0, &task1)
	if err != nil {
		log.Print("decrypt : ", err)
	}
	endTime = time.Now().UnixNano()
	log.Print("decrypt : ", float64((endTime - startTime) / 1e3), " us")
}
