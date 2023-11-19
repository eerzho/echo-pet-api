package service_i

type JWTServiceI interface {
	GenerateToken(id uint) (string, error)
	GenerateHash(word string) (string, error)
	IsEqual(hash, word string) bool
}
