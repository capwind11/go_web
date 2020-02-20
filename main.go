package main

import (
	"awesomeProject/model"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main(){
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/",indexView)
	http.HandleFunc("/list",ListView)
	http.HandleFunc("/api/index/data",indexData)
	http.HandleFunc("/api/list/data",ListData)
	http.HandleFunc("/api/list/del",ListDel)
	http.ListenAndServe(":8080",nil)


}

func indexView(w http.ResponseWriter, r *http.Request){
	f,_ := os.Open("./view/index.html")
	io.Copy(w,f)
	f.Close()
}

func ListView(w http.ResponseWriter, r *http.Request)  {
	f,_ := os.Open("./view/list.html")
	io.Copy(w,f)
	f.Close()
}

func indexData(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	w.Header().Set("Content-Type","application/json")
	idStr := r.Form.Get("id")
	id,_:=strconv.ParseInt(idStr,10,64)
	mod,err:=model.ArticleGet(id)
	if err!=nil{
		Fail(w,err.Error())
		return
	}
	buf,_ := json.Marshal(mod)
	w.Write(buf)
}

func ListData(w http.ResponseWriter, r *http.Request)  {
	mods,err:=model.ArticleList()
	if err!=nil{
		Fail(w,err.Error())
	}
	buf,_ := json.Marshal(mods)
	w.Write(buf)
}

func ListDel(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	idStr := r.Form.Get("id")
	id,_:=strconv.ParseInt(idStr,10,64)
	if model.ArticleDel(id){
		Succ(w,"删除成功")
		return
	}
	Fail(w,"删除失败")
	return

}

func Fail(w http.ResponseWriter, msg string){
	w.Write([]byte(msg))
}
func Succ(w http.ResponseWriter, msg string){
	w.Write([]byte(msg))
}