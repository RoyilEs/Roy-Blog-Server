package main

import (
	"Goblog/core"
	"Goblog/global"
	"Goblog/utils/jwts"
	"fmt"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		Username: "Roy",
		Nickname: "Royill",
		Role:     1,
		UserID:   1,
	})
	fmt.Println(token, err)

	claims, err := jwts.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IlJveSIsIm5pY2tuYW1lIjoiUm95aWxsIiwicm9sZSI6MSwidXNlcklEIjoxLCJleHAiOjE2OTk5NjQxOTUuMDMwNTk1LCJpc3MiOiJibG9nZ2VyIn0.P-s5THfHweRdlL_rZMvEyk6hiB96ur_msDPUuCjEfvI")
	fmt.Println(claims, err)
}
