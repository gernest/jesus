package jesus

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type TextPipeline struct {
	In          []string
	Pos         int
	Out         map[string]string
	FilterFuncs map[string]interface{}
	Salio       string
}

func (s *TextPipeline) Next() string {
	if len(s.In) < s.Pos+1 {
		return "end of line"
	}
	pos := s.Pos
	elem := s.In[pos]
	s.Pos++
	return elem
}

func (s *TextPipeline) Process() error {
	last := len(s.In) - 1
	s.Salio = s.In[last]
	s.In[last] = ""
	fmt.Printf("Salio ni %s \n", s.Salio)
	s.RegisterFilterFunc()
	for {
		node := s.Next()
		if node == "end of line" {
			return nil
		}
		err := s.Filter(node)
		if err != nil {
			return err
		}
	}
}

func (s *TextPipeline) Filter(node string) error {
	for k, _ := range s.FilterFuncs {
		val, err := CallFilter(s.FilterFuncs, k, node)
		if err != nil {
			return err
		}
		a := val[0].String()
		b := val[1].String()
		if a != "" {
			s.Out[a] = b
		}

	}
	return nil
}

func (s *TextPipeline) CheckOutput() map[string]string {
	fmt.Println()
	return s.Out
}
func NewTextPipeline(s string) *TextPipeline {
	text := strings.Split(s, " ")
	pipeline := &TextPipeline{In: text, Pos: 0, Out: make(map[string]string)}
	return pipeline
}
func (s *TextPipeline) RegisterFilterFunc() {
	s.FilterFuncs = map[string]interface{}{
		"Thibitsha": Imethibitishwa,
		"pokea":     Umepokea,
		"Tarehe":    Tarehe,
		"Kiasi":     Kiasi,
	}
}

func CallFilter(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}
