package main
import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)
//Class db class
type Class struct {
	ID   int
	Name string
	Desc string
}

func view(w http.ResponseWriter, r *http.Request){
	f,_ := os.Open("./view/index.html")
	io.Copy(w,f)
	f.Close()
}

func data(w http.ResponseWriter, r *http.Request){
	db, err := sqlx.Open(`mysql`, `root:112112@tcp(127.0.0.1:3306)/community?charset=utf8&parseTime=true`)
	log.Println(db, err)
	mod := Class{}
	r.ParseForm()
	idstr := r.Form.Get("id")
	log.Println("id===>",idstr)
	id,_ := strconv.ParseInt(idstr,10,64)
	//查询多条数据
	//错误信息 = db.Select(要保存数据的变量的指针,`sql语句`,参数...
	err = db.Get(&mod, `select * from class where id = ?`,id)
	log.Println("------",mod)
	js,_ := json.Marshal(mod)
	w.Write(js)
}

func main() {
	http.HandleFunc("/",view)
	http.HandleFunc("/data",data)
	http.ListenAndServe(":8087",nil)

}