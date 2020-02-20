package model

import "time"

type Article struct {
	Id      int64     `json:"id,omitempty"`
	Title   string    `json:"title,omitempty"`
	Author  string    `json:"author,omitempty"`
	Content string    `json:"content,omitempty"`
	Hits    string    `json:"hits,omitempty"`
	Utime   time.Time `json:"utime,omitempty"`
}
// 查询一条数据
func ArticleGet(id int64) (Article,error) {
	mod := Article{}
	err := DB.Unsafe().Get(&mod,"select * from Article where id =? limit 1",id)
	return mod,err
}
// 查询10条数据
func ArticleList()([]Article,error)  {
	mods:=make([]Article,0,10)
	err:=DB.Unsafe().Select(&mods,"select * from Article limit 10")
	return mods,err
}

func ArticleDel(id int64) bool {
	res,_:=DB.Exec("delete from Article where id = ?",id)
	if res==nil {
		return false
	}
	rows,_:=res.RowsAffected()
	if rows>=1{
		return true
	}
	return false
}