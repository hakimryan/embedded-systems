package main

/*test with curl
curl -X POST -H "Content-Type: application/json" \
-d '{"method":"HelloService.Say","params":[{"Who":"Test"}], "id":"1"}' \
http://localhost:10000/rpc
*/

import (
	"fmt"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"io/ioutil"
	"net/http"
	_ "os"
	_ "reflect"
)

type TRHArgs struct {
	Where string
	Time  string
	Temp  string
	RH    string
}

type TRHReply struct {
	Message string
}

type TRHService struct{}

func (trh *TRHService) TRH(r *http.Request, args *TRHArgs, reply *TRHReply) error {
	// f, _ := os.Open("dat2")

	reply.Message = "Hello, " + args.Where + "!"

	TRHData := *args

	lokasi := TRHData.Where
	waktu := TRHData.Time
	suhu := TRHData.Temp
	rh := TRHData.RH
	fmt.Println(lokasi + " " + waktu + " " + suhu + " " + rh)
	mystring := lokasi + " " + waktu + " " + suhu + " " + rh + "\n"
	data := []byte(mystring)
	_ = ioutil.WriteFile("dat2", data, 0644)

	// f.WriteString(lokasi + " " + waktu + " " + suhu + " " + rh + "\n")
	// f.Sync()

	// defer f.Close()

	return nil
}

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(new(TRHService), "")
	http.Handle("/rpc", s)
	http.Handle("/", http.FileServer(http.Dir(".")))
	fmt.Println("Web server will be run...")
	http.ListenAndServe(":10000", nil)

}
