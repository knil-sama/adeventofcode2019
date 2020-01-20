package adventofcode2019

import (
	"fmt"
	"knilsama/adventofcode2019/days"
)

type argOne struct {
	cli.Helper
	Masses string `cli:"m,masses" usage:"raw of masses comma separated"`
}

func main() {
	switch {
	}

	os.Exit(cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		fmt.Println(argv.Masses)
		masses := strings.Split(argv.Masses, ",")
		fuelRequirement := 0
		for _, masse := range masses {
			m, _ := strconv.Atoi(masse)
			fuelRequirement += computeFuel(m)
		}
		TotalFuel := 0
		for fuelRequirement > 0 {
			TotalFuel += fuelRequirement
			fuelRequirement = computeFuel(fuelRequirement)
		}
		fmt.Println(TotalFuel)
		return nil
	}))
}
