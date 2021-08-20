package security

type Unit struct {
	Key []byte
	Iv []byte
	AAD []byte
	Ct []byte
	Pt []byte
	Tag []byte
	TagLen int
}
