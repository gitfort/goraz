package gormext

import (
	"fmt"
	"strings"

	"github.com/gitfort/goraz/pkg/odata"
	"github.com/jinzhu/gorm"
)

var (
	operates = map[odata.Operate]string{
		odata.OperateEqual:        "%v = ?",
		odata.OperateNotEqual:     "%v <> ?",
		odata.OperateIn:           "%v in (?)",
		odata.OperateHas:          "? in %v",
		odata.OperateLess:         "%v < ?",
		odata.OperateLessEqual:    "%v <= ?",
		odata.OperateGreater:      "%v > ?",
		odata.OperateGreaterEqual: "%v >= ?",
	}
)

func ApplyQuery(db *gorm.DB, query odata.Query) *gorm.DB {
	res := db

	if filter, ok := query.Filters(); ok {
		var stmt string
		var args []interface{}
		for _, item := range filter {
			stmt = fmt.Sprintf("%v %v "+operates[item.Operator], stmt, item.Logical, item.Field)
			args = append(args, item.Value)
		}
		stmt = strings.TrimSpace(stmt)
		res = res.Where(stmt, args...)
	}

	if orderBy, ok := query.OrderBy(); ok {
		res = res.Order(fmt.Sprintf("%v %v", orderBy.Field, orderBy.Sort))
	}

	if selects, ok := query.Selects(); ok {
		res = res.Select(selects)
	}

	if top, ok := query.Top(); ok {
		res = res.Limit(top)
	}

	if skip, ok := query.Skip(); ok {
		res = res.Offset(skip)
	}

	if expands, ok := query.Expand(); ok {
		for _, expand := range expands {
			res = res.Preload(expand)
		}
	}

	return res
}
