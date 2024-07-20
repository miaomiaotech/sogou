package sogou

import "context"

var sogou = &Sogou{}

func ToChinese(text string) string {
	resp := sogou.Translate(context.Background(), Request{
		ToLang: Chinese,
		Text:   text,
	})
	return resp.Result
}
