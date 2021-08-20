# sm4-goland
SM4 encryption and decryption demo by Goland base on open source “https://github.com/BabaSSL/BabaSSL”。 

The code supports switching between SM4 and AES by a parameter。You can expand it as long as OpenSSL supports.

Although the code is written in go, in fact, it still calls the C function in terms of performance。

Normal encryption and decryption takes about 20 microseconds.
