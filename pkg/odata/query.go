package odata

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Query map[string]string

func (q Query) String() string {
	values := make(url.Values)
	for key, value := range q {
		values[key] = []string{value}
	}
	return values.Encode()
}

func (q Query) Count() (bool, bool) {
	val, ok := q[keyCount]
	if !ok {
		return false, false
	}
	return strings.ToLower(strings.TrimSpace(val)) == "true", true
}

func (q Query) Expand() ([]string, bool) {
	val, ok := q[keyExpand]
	if !ok {
		return nil, false
	}
	var items []string
	for _, item := range strings.Split(val, ",") {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		items = append(items, item)
	}
	return items, items != nil
}

func (q Query) OrderBy() (*OrderBy, bool) {
	val, ok := q[keyOrderBy]
	if !ok {
		return nil, false
	}
	raws := orderByRegex.FindAllStringSubmatch(strings.TrimSpace(val), 1)
	if len(raws) == 1 {
		orderBy := &OrderBy{
			Field: raws[0][1],
			Sort:  SortAsc,
		}
		if raws[0][2] != "" {
			orderBy.Sort = Sort(raws[0][2])
		}
		return orderBy, true
	}
	return nil, false
}

func (q Query) Search() (string, bool) {
	val, ok := q[keySearch]
	if !ok {
		return "", false
	}
	val = strings.TrimSpace(val)
	return val, val != ""
}

func (q Query) Selects() ([]string, bool) {
	val, ok := q[keySelect]
	if !ok {
		return nil, false
	}
	var items []string
	for _, item := range strings.Split(val, ",") {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		items = append(items, item)
	}
	return items, items != nil
}

func (q Query) Select(fields ...string) Query {
	selects, ok := q[keySelect]
	if ok {
		selects = fmt.Sprint(selects, ",", strings.Join(fields, ","))
	} else {
		selects = strings.Join(fields, ",")
	}
	q[keySelect] = selects
	return q
}

func (q Query) Skip() (int, bool) {
	val, ok := q[keySkip]
	if !ok {
		return 0, false
	}
	i, _ := strconv.Atoi(strings.TrimSpace(val))
	if i <= 0 {
		return 0, false
	}
	return i, true
}

func (q Query) Top() (int, bool) {
	val, ok := q[keyTop]
	if !ok {
		return 0, false
	}
	i, _ := strconv.Atoi(strings.TrimSpace(val))
	if i <= 0 {
		return 0, false
	}
	return i, true
}

func (q Query) Filters() ([]*Filter, bool) {
	val, ok := q[keyFilter]
	if !ok {
		return nil, false
	}
	var filters []*Filter
	for _, item := range filterRegex.FindAllStringSubmatch(strings.TrimSpace(val), -1) {
		filter := &Filter{
			Logical:  Logic(item[1]),
			Field:    item[2],
			Operator: Operate(item[3]),
		}
		_ = json.Unmarshal([]byte(strings.ReplaceAll(item[4], "'", "\"")), &filter.Value)
		filters = append(filters, filter)
	}
	return filters, filters != nil
}

func (q Query) AndFilter(field string, operate Operate, value interface{}) Query {
	return q.filter(LogicAnd, field, operate, value)
}

func (q Query) OrFilter(field string, operate Operate, value interface{}) Query {
	return q.filter(LogicOr, field, operate, value)
}

func (q Query) filter(logic Logic, field string, operate Operate, value interface{}) Query {
	bin, _ := json.Marshal(value)
	val := strings.ReplaceAll(string(bin), "\"", "'")
	filters, ok := q[keyFilter]
	if ok {
		filters = fmt.Sprintf("%v %v %v %v %v", filters, logic, field, operate, val)
	} else {
		filters = fmt.Sprintf("%v %v %v", field, operate, val)
	}
	q[keyFilter] = filters
	return q
}
