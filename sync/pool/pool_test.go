package pool

import (
	"encoding/json"
	"sync"
	"testing"
)

type Student struct {
	Name string
	Age  int
	A    string
	B    string
	C    string
	D    string
}

var buf, _ = json.Marshal(Student{
	Name: "YST",
	Age:  20,
	A:    "1",
	B:    "2",
	C:    "2",
	D:    "1",
})

var stuPool = sync.Pool{
	New: func() any {
		return new(Student)
	},
}

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := &Student{}
		json.Unmarshal(buf, stu)
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := stuPool.Get().(*Student)
		json.Unmarshal(buf, stu)
		stuPool.Put(stu)
	}
}
