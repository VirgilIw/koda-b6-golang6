package internal

type Data struct {
	Name string
	Wait int
}

func Queque(ch chan Data) {
	data := []Data{
		{Name: "vir", Wait: 4},
		{Name: "ver", Wait: 3},
		{Name: "vur", Wait: 6},
		{Name: "vor", Wait: 2},
		{Name: "vmr", Wait: 1},
	}

	for _, x := range data {
		ch <- x
	}
	close(ch)
}
