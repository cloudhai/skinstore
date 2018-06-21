package Handler

import (
	"net/http"
	"skinstore/common"
	"skinstore/Entity/project"
	"strconv"
	log2 "test01/log"
	"skinstore/web/router"
)

var log = log2.NewLog()

func ProjectListHander(params *router.Params,rw http.ResponseWriter) *common.WebResult{
		rowsStr := params.Get("rows")
		rows,err := strconv.Atoi(rowsStr)
		if err != nil {
			log.Error(err)
			rows = 15
		}
		pageStr := params.Get("page")
		page,err := strconv.Atoi(pageStr)
		if err != nil{
			log.Error(err)
			page = 0
		}
		index := (page-1)*rows
		list := project.GetAllProjectList(index,rows)
		return common.NewResult(1,list)
}

func ProjectLisByTypetHander(p *router.Params,rw http.ResponseWriter)*common.WebResult{
	params := make(map[string]string)
	rowsStr := p.Get("rows")
	if rowsStr == ""{
		rowsStr = strconv.Itoa(common.ROWS_SIZE)
	}
	params["rows"] = rowsStr
	indexStr := p.Get("index")
	if indexStr == ""{
		indexStr = "0"
	}
	params["index"] = indexStr
	if typep := p.Get("type");typep != ""{
		params["type"] = typep
	}
	list := project.GetProjectListByParam(params)
	return common.NewResult(1,list)
}

