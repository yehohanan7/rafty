package calc

import (
	"fmt"
	"io"
	"strings"
	"time"

	"strconv"

	"github.com/hashicorp/raft"
)

type Calculator struct {
	Value int
	Raft  *raft.Raft
}

func (calc *Calculator) Apply(log *raft.Log) interface{} {
	fmt.Println("command recieved: ", string(log.Data))
	c := strings.Split(string(log.Data), ":")
	v, _ := strconv.Atoi(c[1])
	switch c[0] {
	case "+":
		calc.Value = calc.Value + v
	case "-":
		calc.Value = calc.Value - v
	default:
		fmt.Printf("invalid command: %v\n", c)
	}
	return calc.Value
}

func (calc *Calculator) Add(v int) {
	calc.Raft.Apply([]byte("+:"+strconv.Itoa(v)), time.Second*2)
}

func (calc *Calculator) Subtract(v int) {
	calc.Raft.Apply([]byte("-:"+strconv.Itoa(v)), time.Second*2)
}

func (calc *Calculator) Restore(snap io.ReadCloser) error {
	return nil
}

func (clac *Calculator) Snapshot() (raft.FSMSnapshot, error) {
	return new(CalculatorSnapshot), nil
}

type CalculatorSnapshot struct {
	state string
}

func (snap *CalculatorSnapshot) Persist(sink raft.SnapshotSink) error {
	return nil
}

func (snap *CalculatorSnapshot) Release() {

}
