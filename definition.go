package sogou

const (
	Chinese  = "zh-CN"
	English  = "en"
	Russian  = "ru"
	Japanese = "ja"
	German   = "de"
	French   = "fr"
	Korean   = "ko"
	Spanish  = "es"

	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36"
)

// Request translation request
type Request struct {
	FromLang string
	ToLang   string
	Text     string
}

type Response struct {
	Result string
	Err    error
}
