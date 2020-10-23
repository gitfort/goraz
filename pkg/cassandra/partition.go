package cassandra

type Partition struct {
	From int
	To   int
}

func ToPartitions(from, to int) []*Partition {
	if from > to {
		return nil
	}
	var list []*Partition
	start := from
	finish := from/100*100 + 99
	for {
		if finish > to {
			finish = to
		}
		list = append(list, &Partition{
			From: start,
			To:   finish,
		})
		if finish == to {
			break
		}
		start = finish + 1
		finish += 100
	}
	return list
}
