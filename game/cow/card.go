package cow
import (
	"github.com/lkj01010/log"
	"strconv"
	"chess/fw"
)

type card struct {
	n cardNumber
	c cardColor
}

func (c *card)score() int {
	if c.n < cn_1 {
		log.Error("c.n is invalid(<1)")
	}
	if c.n < cn_10 {
		return int(c.n)
	} else {
		return 10
	}
}

//func (c *card)compare(c1 *card) compareRet {
//	if c.n < c1.n {
//		return cp_small
//	} else if c.n > c1.n {
//		return cp_big
//	} else {
//		if c.t < c1.t {
//			return cp_small
//		} else if c.t > c1.t {
//			return cp_big
//		} else {
//			return cp_equal
//		}
//	}
//}

func (c *card)string() string {
	return strconv.Itoa(int(c.n)) + "#" + strconv.Itoa(int(c.c))
}

// utils
func CompareCard(c0, c1 *card) compareRet {
	if c0.n < c1.n {
		return cp_small
	} else if c0.n > c1.n {
		return cp_big
	} else {
		if c0.c < c1.c {
			return cp_small
		} else if c0.c > c1.c {
			return cp_big
		} else {
			return cp_equal
		}
	}
}

func JudgeCardType(cards *[]card) cardType {
	total := 0
	for _, card := range (*cards) {
		total += card.score()
	}
	left := total / 10
	return cardType(left)
}

func DealCards(cards []card) {
	for i, _ := range (cards) {
		cards[i] = cardPool[fw.Randn(54)]
	}
	log.Debugf("cards=%+v", cards)
}