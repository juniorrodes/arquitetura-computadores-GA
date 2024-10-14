package pipeline

import (
	"strings"
)

type Queue [5]*ExecInstructions

func (q *Queue) Push(item *ExecInstructions) {
    helper := q[0]
    q[0] = item

    for i := 1; i < 5; i++ {
        temp := q[i]
        q[i] = helper
        helper = temp 
    }
}

func (q *Queue) String() string {
    var stringBuilder strings.Builder
    
    stringBuilder.WriteString("[ ")

    for _, item := range q {
        stringBuilder.WriteString(item.String())
        stringBuilder.WriteString(" ")
    }
    
    stringBuilder.WriteByte(']')

    return stringBuilder.String()
}
