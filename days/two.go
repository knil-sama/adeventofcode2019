package days

import (
	"fmt"
	"github.com/mkideal/cli"
	"os"
	"strconv"
	"strings"
)

func computeIntCode(source []int, position int) int {
	optCode := source[position]
	if optCode == 99 {
		return -1
	}
	noun := source[position+1]
	verb := source[position+2]
	storage := source[position+3]
	if optCode == 1 {
		source[storage] = source[noun] + source[verb]
	} else {
		source[storage] = source[noun] * source[verb]
	}
	print(fmt.Sprintf("%d %d %d\n", noun, verb, source[storage]))
	return position + 4
}

type argT struct {
	cli.Helper
	Input string `cli:"i,input usage:"comma separated input"`
}

func main() {
	os.Exit(cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		fmt.Println(argv.Input)
		var inputList []int
		for _, x := range strings.Split(argv.Input, ",") {
			v, _ := strconv.Atoi(x)
			inputList = append(inputList, v)
		}
		var operationIndex int = 0
		for operationIndex >= 0 {
			operationIndex = computeIntCode(inputList[:], operationIndex)
		}
		for _, x := range inputList {
			print(x, ",")
		}
		return nil
	}))
}
