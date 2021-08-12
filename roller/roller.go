package roller

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	colorReset = "\033[0m"
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
	colorBlue  = "\033[34m"
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
	Status   bool
	Message  string
}

func (r *Roll) parseData() {
	rex := regexp.MustCompile(`[\d?]*d[\d?]+|[+->][\d]*`)
	rollTerms := rex.FindAllString(r.Request, -1)
	fmt.Println(rollTerms)
	for _, data := range rollTerms {
		if strings.Contains(data, "d") {
			r.getQntFace(data)
			continue
		}
		r.getExtras(data)
	}
	r.executeRoll()
	return
}

func (r *Roll) getQntFace(s string) {
	div := strings.Split(s, "d")
	qnt, err := strconv.Atoi(div[0])
	if err != nil {
		log.Fatal(err)
		return
	}
	r.Quantity = qnt
	r.getFace(div[1])
}

func (r *Roll) getFace(s string) {
	faces, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
		return
	}
	r.Faces = faces
}

func (r *Roll) getExtras(s string) {
	ccount := len(s)
	if s[0:1] == "+" {
		bonus, err := strconv.Atoi(s[1:ccount])
		if err != nil {
			log.Fatal(err)
			return
		}
		r.Bonus = bonus
	}
	if s[0:1] == "-" {
		penalty, err := strconv.Atoi(s[1:ccount])
		if err != nil {
			log.Fatal(err)
			return
		}
		r.Penalty = penalty
	}
	if s[0:1] == ">" {
		target, err := strconv.Atoi(s[1:ccount])
		if err != nil {
			log.Fatal(err)
			return
		}
		r.Target = target
	}
}

func (r *Roll) executeRoll() {
	separator := ","
	r.Message = "Rolando dados...\n["
	rand.Seed(time.Now().UTC().UnixNano())
	r.Status = false
	for i := 0; i < r.Quantity; i++ {
		if i+1 == r.Quantity {
			separator = "]"
		}
		value := rand.Intn(r.Faces-1) + 1
		r.Message = fmt.Sprintf("%s%s%d%s%s", r.Message, colorBlue, value, colorReset, separator)
		r.Results = append(r.Results, value)
		r.Total += value
	}
	if r.Bonus > 0 {
		r.Message = fmt.Sprintf("%s + %d", r.Message, r.Bonus)
		r.Total += r.Bonus
	}
	if r.Penalty > 0 {
		r.Message = fmt.Sprintf("%s - %d", r.Message, r.Penalty)
		r.Total -= r.Penalty
	}
	r.Message = fmt.Sprintf("%s = %d\n", r.Message, r.Total)
	if r.Total > r.Target {
		if r.Target > 0 {
			r.Message = fmt.Sprintf("%s%s%s%s\n", r.Message, colorGreen, "SUCESSO!", colorReset)
		}
		r.Status = true
	} else if r.Target > 0 {
		r.Message = fmt.Sprintf("%s%s%s%s\n", r.Message, colorRed, "FALHA!", colorReset)
	}
}

func RollDice(req string) Roll {
	roll := Roll{
		Request: req,
	}
	roll.parseData()

	return roll
}
