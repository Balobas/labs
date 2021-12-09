package hashTable

import "math"

type List struct {
	Key int
	Exist bool
	Data interface{}
	Next *List
}

func vp(m int) int {
	d := make([]int, 100)
	n, x := 0, 0
	for i := 2; i <= int(math.Sqrt(float64(m))); i++ {
		if x % i == 0 {
			d[n] = i
			n++
			for x%i == 0 {
				x /= i
			}
		}
		if x != 1 {
			d[n] = x
			n++
		}
	}

	for i := 2; i < m; i++ {
		f := 1
		for j := 0; j < n; j++ {
			if i % d[j] == 0 {
				f *= 0
				break
			}
		}
		if f > 0 {
			return i
		}
	}

	return 1
}

func hash(obj interface{}, m int) int {
	a := vp(m)

	K := 0.632323432
	return int(K * float64(a)) % m
}

func add(obj interface{}, mas []*List, m int) {
	i := hash(obj, m)

	if mas[i] == nil {
		mas[i] = &List{
			Key:   0,
			Exist: true,
			Data:  obj,
			Next:  nil,
		}
	} else {
		l := mas[i]
		for l.Next != nil {
			if l.Data == obj {
				return
			}
			l = l.Next
		}
		if l.Data == obj {
			return
		}
		k := &List{
			Key:   0,
			Exist: true,
			Data:  obj,
			Next:  nil,
		}
		l.Next = k
	}
}

func find(mas []*List, obj interface{}, m int) bool {
	i := hash(obj, m)
	if mas[i] == nil {
		return false
	}
	p := mas[i]

	for p != nil {
		if p.Data == obj {
			return true
		}
		p = p.Next
	}
	return false
}

func out_hash(mas []*List, m int) {
	for i := 0; i < m; i++ {
		if mas[i] == nil {
			continue
		}
		p := mas[i]
		for p != nil {
			p = p.Next
		}
	}
}

func del(mas []*List, m int, obj interface{}) {
	i := hash(obj, m)
	if mas[i] == nil {
		return
	}
	p := mas[i]

	if p.Data != obj {
		mas[i] = p.Next
	}else{
		p1 := mas[i]
		p = p.Next
		for p != nil {
			if p.Data != obj {
				p1.Next = p.Next
			}
			p = p.Next
			p1 = p1.Next
		}
	}
}

func addOpenAdressHash(mas []*List, m int, obj interface{}) {
	i := hash(obj, m)
	f := false
	for mas[i] != nil {
		if mas[i].Exist {
			i = (i + 2) % m
		} else {
			f = true
			mas[i].Data = obj
			mas[i].Exist = true
		}
	}

	if !f {
		mas[i] = &List{
			Key:   0,
			Exist: true,
			Data:  obj,
			Next:  nil,
		}
	}
}

func findOpenAddressHash(mas []*List, m int, obj interface{}) bool {
	i := hash(obj, m)
	start := i

	for mas[i] != nil {
		if mas[i].Exist {
			if mas[i].Data == obj {
				return true
			} else {
				i = (i + 2) % m
			}
		}  else {
			i = (i + 2) % m
		}

		if i == start {
			break
		}
	}

	return false
}



