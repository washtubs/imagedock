package imagedock

type FileAddr interface {
	String() string
}

type FileStore interface {
	save(addr FileAddr)
}
