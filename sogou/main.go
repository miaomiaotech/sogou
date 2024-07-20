package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/miaomiaotech/sogou"
)

func main() {
	fromLang := flag.String("from", sogou.English, "source language")
	toLang := flag.String("to", sogou.Chinese, "target language")
	both := flag.Bool("both", false, "output both language")
	flag.Parse()

	text := flag.Arg(0)

	if text == "" {
		bs, err := os.ReadFile("/dev/stdin")
		if err != nil {
			fmt.Printf("read stdin err: %v\n", err)
			os.Exit(1)
		}

		text = string(bs)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res := sogou.Translate(ctx, &sogou.Request{
		FromLang: *fromLang,
		ToLang:   *toLang,
		Text:     text,
	})
	if res.Err != nil {
		fmt.Printf("translate err: %v\n", res.Err)
		os.Exit(1)
	}

	if *both {
		fmt.Println("<-", text)
		fmt.Println("->", res.Result)
	} else {
		fmt.Println(res.Result)
	}
}
