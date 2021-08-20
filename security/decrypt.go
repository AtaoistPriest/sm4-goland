package security

//#cgo LDFLAGS: -lssl -lcrypto
//#include <stdio.h>
//#include <stdlib.h>
//#include <openssl/bio.h>
//#include <openssl/evp.h>
//int aes_gcm_decrypt(const EVP_CIPHER *cipher, unsigned char * pt, unsigned char *ct, int ctLen, unsigned char * gcm_key, unsigned char * gcm_iv, unsigned char *gcm_tag, int tagLen)
//{
//	EVP_CIPHER_CTX *ctx;
//	int outlen, res = 1;
//	ctx = EVP_CIPHER_CTX_new();
//	if (ctx == NULL) return 0;
//	res = EVP_DecryptInit_ex(ctx, cipher, NULL, NULL, NULL);
//	res = EVP_CIPHER_CTX_ctrl(ctx, EVP_CTRL_AEAD_SET_IVLEN, sizeof(gcm_iv), NULL);
//	res = EVP_DecryptInit_ex(ctx, NULL, NULL, gcm_key, gcm_iv);
//	res = EVP_DecryptUpdate(ctx, pt, &outlen, ct, ctLen);
//	res = EVP_CIPHER_CTX_ctrl(ctx, EVP_CTRL_AEAD_SET_TAG, tagLen, (void *)gcm_tag);
//	res = EVP_DecryptFinal_ex(ctx, pt, &outlen);
//	EVP_CIPHER_CTX_free(ctx);
//	return res;
//}
import "C"
import (
	"errors"
)
/*
	加密模式：
	0 ： sm4（国密）
	1 ： aes—128
*/
var cipher = []*C.EVP_CIPHER{(*C.EVP_CIPHER)(C.EVP_sm4_gcm()), (*C.EVP_CIPHER)(C.EVP_aes_128_gcm())}

func Decrypt(model int, decryptUnit *Unit) error {
	res := C.aes_gcm_decrypt((*C.EVP_CIPHER)(cipher[model]), (*C.uchar)(&decryptUnit.Pt[0]), (*C.uchar)(&decryptUnit.Ct[0]), C.int(len(decryptUnit.Ct)), (*C.uchar)(&decryptUnit.Key[0]),
		(*C.uchar)(&decryptUnit.Iv[0]), (*C.uchar)(&decryptUnit.Tag[0]), C.int(decryptUnit.TagLen))
	if res > 0{
		return nil
	}else{
		return errors.New("Error : cDecryption ")
	}
}
