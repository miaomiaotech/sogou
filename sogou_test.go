package sogou

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSogouTranslate(t *testing.T) {
	sogou := &Sogou{}
	{
		resp := sogou.Translate(context.Background(), Request{
			ToLang: Chinese,
			Text:   "test",
		})
		assert.Nil(t, resp.Err)
		expectedResp := &Response{Result: "试验"}
		assert.Equal(t, expectedResp, resp)
	}

	{
		resp := sogou.Translate(context.Background(), Request{
			ToLang: English,
			Text:   "测试",
		})
		assert.Nil(t, resp.Err)
		expectedResp := &Response{Result: "test"}
		assert.Equal(t, expectedResp, resp)
	}

	{
		assert.Equal(t, "试验", ToChinese("test"))
	}
}
