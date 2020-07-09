package main

import (
	//"encoding/json"
	"fmt"
	"io/ioutil"

	//"io/ioutil"
	"net/http"
	"github.com/strive-after/gojenkins/api"
)






var jenkins  *api.Jenkins






func main() {
	var req *http.Request

	ar := api.NewJsonAPIRequest("get","/")

	req, err := http.NewRequest(ar.Method, "http://39.105.114.198:8080/job/test/api/json", ar.Payload)
	jenkins = api.NewJenkins("admin","123","http://39.105.114.198:8080","")
	req.SetBasicAuth(jenkins.User,jenkins.Password)
	responce,err := jenkins.Client.Do(req)
	if err != nil {
		fmt.Println(err,"err1")
		return
	}
	data ,err := ioutil.ReadAll(responce.Body)
	if err != nil {
		fmt.Println(err,"err2")
		return
	}
	//var data []byte = []byte{}
	//json.Unmarshal(data,responce.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	fmt.Printf("%#v\n",string(data))
}
