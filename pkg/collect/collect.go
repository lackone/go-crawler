package collect

type IFetch interface {
	Get(url string) ([]byte, error)
}
