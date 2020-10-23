package mongoext

import (
	"github.com/gitfort/goraz/pkg/odata"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	operates = map[odata.Operate]string{
		odata.OperateEqual:        "$eq",
		odata.OperateNotEqual:     "$ne",
		odata.OperateIn:           "$in",
		odata.OperateHas:          "",
		odata.OperateLess:         "$lt",
		odata.OperateLessEqual:    "$lte",
		odata.OperateGreater:      "$gt",
		odata.OperateGreaterEqual: "$gte",
	}
	logics = map[odata.Logic]string{
		odata.LogicAnd: "$and",
		odata.LogicOr:  "$or",
	}
)

func ParseQuery(query odata.Query) (bson.M, *options.FindOptions) {
	rep := bson.M{}
	opts := options.Find()

	if filter, ok := query.Filters(); ok {
		for _, item := range filter {
			value := returnValue(item)
			if op := operates[item.Operator]; op != "" {
				value = bson.M{operates[item.Operator]: value}
			}
			rep[item.Field] = value
			if item.Logical != "" {
				rep = bson.M{logics[item.Logical]: rep}
			}
		}
	}

	if orderBy, ok := query.OrderBy(); ok {
		sort := 1
		if orderBy.Sort == odata.SortDesc {
			sort = -1
		}
		opts = opts.SetSort(bson.E{
			Key:   orderBy.Field,
			Value: sort,
		})
	}

	if selects, ok := query.Selects(); ok {
		p := bson.M{}
		for _, item := range selects {
			p[item] = 1
		}
		opts = opts.SetProjection(p)
	}

	if top, ok := query.Top(); ok {
		opts = opts.SetLimit(int64(top))
	}

	if skip, ok := query.Skip(); ok {
		opts = opts.SetSkip(int64(skip))
	}

	return rep, opts
}

func returnValue(item *odata.Filter) interface{} {
	switch value := item.Value.(type) {
	case string:
		if id, err := uuid.Parse(value); err == nil {
			return id
		}
	case []interface{}:
		var ids []uuid.UUID
		for _, i := range value {
			switch val := i.(type) {
			case string:
				if id, err := uuid.Parse(val); err == nil {
					ids = append(ids, id)
				}
			}
		}
		if len(ids) > 0 {
			return ids
		}
	}
	return item.Value
}
