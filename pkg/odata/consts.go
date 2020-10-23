package odata

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	keyCount   = "$count"
	keyExpand  = "$expand"
	keyOrderBy = "$orderby"
	keySearch  = "$search"
	keySelect  = "$select"
	keyTop     = "$top"
	keySkip    = "$skip"
	keyFilter  = "$filter"
)

type Sort string

const (
	SortAsc  Sort = "asc"
	SortDesc Sort = "desc"
)

type Operate string

const (
	OperateEqual        Operate = "eq"
	OperateNotEqual     Operate = "ne"
	OperateGreater      Operate = "gt"
	OperateGreaterEqual Operate = "ge"
	OperateLess         Operate = "lt"
	OperateLessEqual    Operate = "le"
	OperateIn           Operate = "in"
	OperateHas          Operate = "has"
)

type Logic string

const (
	LogicAnd Logic = "and"
	LogicOr  Logic = "or"
)

var (
	keys = []string{
		keyCount,
		keyExpand,
		keyOrderBy,
		keySearch,
		keySelect,
		keyTop,
		keySkip,
		keyFilter,
	}
	sorts = []string{
		string(SortAsc),
		string(SortDesc),
	}
	logics = []string{
		string(LogicAnd),
		string(LogicOr),
	}
	operates = []string{
		string(OperateIn),
		string(OperateLessEqual),
		string(OperateLess),
		string(OperateGreaterEqual),
		string(OperateGreater),
		string(OperateNotEqual),
		string(OperateEqual),
		string(OperateHas),
	}
	filterRegex = regexp.MustCompile(fmt.Sprintf(`(%v|)\s?([^\s]*)\s(%v)\s(\[[^\]]*\]|'[^']*'|true|false|(\d+(\.\d+)?))`,
		strings.Join(logics, "|"),
		strings.Join(operates, "|"),
	))
	orderByRegex = regexp.MustCompile(fmt.Sprintf(`(\w+)\s?(%v|)`,
		strings.Join(sorts, "|"),
	))
)
