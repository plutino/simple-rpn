package calc

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	MODE_BIN = 2
	MODE_OCT = 8
	MODE_DEC = 10
	MODE_HEX = 16
)

type Calculator struct {
	mode   int
	fmtStr string
	stk    *stack
}

var operators = map[string]operator{
	"+": add{},
	"-": sub{},
	"*": mul{},
	"/": div{},
	"^": pow{},
}

func isOperator(token string) bool {
	_, ok := operators[token]
	return ok
}

func isCmd(token string) bool {
	return strings.HasPrefix(token, "\\")
}

func NewCalculator(mode int, stack_depth int) (*Calculator, error) {
	c := &Calculator{
		stk: newStack(stack_depth),
	}
	c.SetMode(mode)
	return c, nil
}

func (c *Calculator) Insert(token string) error {
	if isOperator(token) {
		return c.do(operators[token])
	}
	if isCmd(token) {
		return c.exec(token)
	}
	val, err := strconv.ParseInt(token, c.mode, 64)
	if err != nil {
		return fmt.Errorf("unparsable token: %s", token)
	}
	return c.stk.Push(int(val))
}

func (c *Calculator) PrintStack() (string, error) {
	values := c.stk.Snapshot()
	if len(values) == 0 {
		return "", nil
	}
	res := fmt.Sprintf(c.fmtStr, values[0])
	for _, v := range values[1:] {
		res += fmt.Sprintf(" "+c.fmtStr, v)
	}
	return res, nil
}

func (c *Calculator) SetMode(mode int) {
	c.mode = mode
	switch mode {
	case MODE_BIN:
		c.fmtStr = "%b"
	case MODE_OCT:
		c.fmtStr = "%o"
	case MODE_DEC:
		c.fmtStr = "%d"
	case MODE_HEX:
		c.fmtStr = "%x"
	}
}

func (c *Calculator) do(op operator) error {
	b, a, err := c.stk.Pop2()
	if err != nil {
		return err
	}
	return c.stk.Push(op.Do(a, b))
}

func (c *Calculator) exec(cmd string) error {
	switch cmd {
	case "\\d":
		c.SetMode(MODE_DEC)
	case "\\b":
		c.SetMode(MODE_BIN)
	case "\\o":
		c.SetMode(MODE_OCT)
	case "\\x":
		c.SetMode(MODE_HEX)
	case "\\r":
		c.stk.Pop()
	case "\\c":
		c.stk.Clear()
	default:
		return fmt.Errorf("unknown command: %s", cmd)
	}
	return nil
}
