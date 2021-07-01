package main
import (

	"fmt"
	"context"
	"time"
    //"github.com/ynsgnr/aria2go"
	//"github.com/myanimestream/arigo"
	//"github.com/auriou/godebridaria/config"
	"github.com/zyxar/argo/rpc"

)

func main(){

	rpcc, err := rpc.New(context.Background(), "https://36163-jade-hare-aip8kak5.ws-us08.gitpod.io/jsonrpc", "abcd", time.Second*5, nil)
	if err != nil {
		panic(err)
	}
	defer rpcc.Close()
	var url string =  "http://ipv4.download.thinkbroadband.com/512MB.zip"

	g, err := rpcc.AddURI([]string{url})
	if err != nil {
		panic(err)
	}
	if g != "" {

		fmt.Println("GID is", g)
	}
	fmt.Println("add download task", url)
	return 


}


