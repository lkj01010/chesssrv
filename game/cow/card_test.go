package cow
import "testing"


func TestCard1(t *testing.T){
//	var cards = &[cardLen]card{card{}, card{}, card{}, card{}, card{}}
	var cards = make([]card, 5)
	DealCards(cards)
	t.Logf("cards=%+v", cards)
}