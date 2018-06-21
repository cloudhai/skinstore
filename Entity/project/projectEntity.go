package project

import (
	"skinstore/utils/SqliteUtil"
	"skinstore/common"
	"fmt"
	"strings"
)

type projectEntity struct{
	Id int
	Name string
	Description string
	Type string
	ImgUrl string
	OriginalPrice int
	CurPrice int
	Status byte
}

func GetAllProjectList(index,rows int) []projectEntity{
	db := SqliteUtil.NewSqlDb()
	stmt,err := db.Db.Prepare(`select id as Id,name as Name, description as Description,
				type as Type,img_url as ImgUrl,original_price as OriginalPrice,cur_price as CurPirce,
				status as Status from project limit ?,?`)
	common.CheckErr(err)
	defer stmt.Close()
	res,err := stmt.Query(index,rows)
	common.CheckErr(err)
	var list []projectEntity
	for res.Next(){
		var project projectEntity
		err := res.Scan(&project.Id,&project.Name,&project.Description,&project.Type,&project.ImgUrl,&project.OriginalPrice,&project.CurPrice,&project.Status)
		common.CheckErr(err)
		list = append(list,project)
	}
	return list
}

func GetProjectListByParam(params map[string]string) []projectEntity{
	db := SqliteUtil.NewSqlDb()
	sql := []string{"select id as Id,name as Name, description as Description,type as Type,img_url as ImgUrl,original_price as OriginalPrice,cur_price as CurPirce, status as Status from project where status=1"}
	if paramType,ok := params["type"];ok{
		sql = append(sql,fmt.Sprintf(" and type='%s'",paramType))
	}
	index,_ := params["index"]
	rows,_ := params["rows"]
	sql = append(sql,fmt.Sprintf(" order by id desc limit %s,%s",index,rows))
	stmt,err := db.Db.Prepare(strings.Join(sql,""))
	fmt.Println(strings.Join(sql,""))
	common.CheckErr(err)
	defer stmt.Close()
	res,err := stmt.Query()
	common.CheckErr(err)
	var list []projectEntity
	for res.Next(){
		var project projectEntity
		err := res.Scan(&project.Id,&project.Name,&project.Description,&project.Type,&project.ImgUrl,&project.OriginalPrice,&project.CurPrice,&project.Status)
		common.CheckErr(err)
		list = append(list,project)
	}
	return list

}
