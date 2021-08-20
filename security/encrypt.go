package security

//#cgo LDFLAGS: -lssl -lcrypto
//#include <stdio.h>
//#include <stdlib.h>
//#include <openssl/bio.h>
//#include <openssl/evp.h>
//int aes_gcm_encrypt(const EVP_CIPHER *cipher, unsigned char *pt, int ptLen, unsigned char * ct, unsigned char * gcm_key, unsigned char * gcm_iv, unsigned char * tag, int tagLen)
//{
//	EVP_CIPHER_CTX *ctx;
//	int outlen, res = 1;
//	ctx = EVP_CIPHER_CTX_new();
//	if (ctx == NULL) return 0;
//	res = EVP_EncryptInit_ex(ctx, cipher, NULL, NULL, NULL);
//	res = EVP_CIPHER_CTX_ctrl(ctx, EVP_CTRL_AEAD_SET_IVLEN, sizeof(gcm_iv), NULL);
//	res = EVP_EncryptInit_ex(ctx, NULL, NULL, gcm_key, gcm_iv);
//	res = EVP_EncryptUpdate(ctx, ct, &outlen, pt, ptLen);
//	res = EVP_EncryptFinal_ex(ctx, ct, &outlen);
//	res = EVP_CIPHER_CTX_ctrl(ctx, EVP_CTRL_AEAD_GET_TAG, tagLen, tag);
//	EVP_CIPHER_CTX_free(ctx);
//	return res;
//}
import "C"
import (
	"errors"
)

func Encrypt(model int, encryptUnit *Unit) error {
	res := C.aes_gcm_encrypt((*C.EVP_CIPHER)(cipher[model]), (*C.uchar)(&encryptUnit.Pt[0]), C.int(len(encryptUnit.Pt)), (*C.uchar)(&encryptUnit.Ct[0]),
		(*C.uchar)(&encryptUnit.Key[0]), (*C.uchar)(&encryptUnit.Iv[0]), (*C.uchar)(&encryptUnit.Tag[0]), C.int(encryptUnit.TagLen))
	if res > 0{
		return nil
	}else{
		return errors.New("Error : cEncryption ")
	}
}
