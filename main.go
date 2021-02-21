package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func main(){
	fasthttp.ListenAndServe(":8080",requestHandler)
}

func requestHandler(ctx *fasthttp.RequestCtx){
	fmt.Fprintf(ctx,"m:%q,user:%q\n",ctx.Method(),ctx.UserAgent())
}
