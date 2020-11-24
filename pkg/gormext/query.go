package gormext

import (
	"fmt"
	"github.com/gitfort/goraz/pkg/odata"
	"github.com/jinzhu/gorm"
	"strings"
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

func ApplyQuery(db *gorm.DB, query odata.Query, withOpts bool) *gorm.DB {
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

	if selects, ok := query.Selects(); withOpts && ok {
		res = res.Select(selects)
	}

	if expands, ok := query.Expand(); withOpts && ok {
		for _, expand := range expands {
			res = res.Preload(expand)
		}
	}

	if orderBy, ok := query.OrderBy(); withOpts && ok {
		res = res.Order(fmt.Sprintf("%v %v", orderBy.Field, orderBy.Sort))
	}

	if top, ok := query.Top(); withOpts && ok {
		res = res.Limit(top)
	}

	if skip, ok := query.Skip(); withOpts && ok {
		res = res.Offset(skip)
	}

	return res
}
