package days

import (
	"fmt"
	"github.com/mkideal/cli"
	"os"
	"strconv"
	"strings"
)

func computeFuel(masse int) int {
	return masse/3 - 2
}
