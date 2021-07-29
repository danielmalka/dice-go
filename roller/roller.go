package roller

import (
	"log"
	"strconv"
	"strings"
)

type Roll struct {
	Request  string
	Quantity int
	Faces    int
	Bonus    int
	Penalty  int
	Target   int
	Results  []int
	Total    int
	Status   []bool
}

func (r *Roll) parseData() {
	div := strings.Split(r.Request, "d")
	qnt, err := strconv.Atoi(div[0])
	if err != nil {
		log.Fatal(err)
	}
	r.Quantity = qnt

	faces, err := strconv.Atoi(div[0])
	if err != nil {
		log.Fatal(err)
	}
	r.Faces = faces
}

func RollDice(req string) Roll {
	roll := Roll{
		Request: req,
	}
	roll.parseData()

	return roll
}
