package cow

const cardLen = 5	// 一副牌

type cardColor int
const (
	cc_diamond    cardColor = 1 + iota
	cc_club
	cc_heart
	cc_spade
)

type cardNumber int
const (
	cn_1 cardNumber = iota + 1
	cn_2
	cn_3
	cn_4
	cn_5
	cn_6
	cn_7
	cn_8
	cn_9
	cn_10
	cn_j
	cn_q
	cn_k
	cn_blackJk    // black joker
	cn_redJk    // red joker
)

var cardPool = [...]card{
	card{cn_1,cc_spade},
	card{cn_1,cc_heart},
	card{cn_1,cc_club},
	card{cn_1,cc_diamond},

	card{cn_2,cc_spade},
	card{cn_2,cc_heart},
	card{cn_2,cc_club},
	card{cn_2,cc_diamond},

	card{cn_3,cc_spade},
	card{cn_3,cc_heart},
	card{cn_3,cc_club},
	card{cn_3,cc_diamond},

	card{cn_4,cc_spade},
	card{cn_4,cc_heart},
	card{cn_4,cc_club},
	card{cn_4,cc_diamond},

	card{cn_5,cc_spade},
	card{cn_5,cc_heart},
	card{cn_5,cc_club},
	card{cn_5,cc_diamond},

	card{cn_6,cc_spade},
	card{cn_6,cc_heart},
	card{cn_6,cc_club},
	card{cn_6,cc_diamond},

	card{cn_7,cc_spade},
	card{cn_7,cc_heart},
	card{cn_7,cc_club},
	card{cn_7,cc_diamond},

	card{cn_8,cc_spade},
	card{cn_8,cc_heart},
	card{cn_8,cc_club},
	card{cn_8,cc_diamond},

	card{cn_9,cc_spade},
	card{cn_9,cc_heart},
	card{cn_9,cc_club},
	card{cn_9,cc_diamond},

	card{cn_10,cc_spade},
	card{cn_10,cc_heart},
	card{cn_10,cc_club},
	card{cn_10,cc_diamond},

	card{cn_j,cc_spade},
	card{cn_j,cc_heart},
	card{cn_j,cc_club},
	card{cn_j,cc_diamond},

	card{cn_q,cc_spade},
	card{cn_q,cc_heart},
	card{cn_q,cc_club},
	card{cn_q,cc_diamond},

	card{cn_k,cc_spade},
	card{cn_k,cc_heart},
	card{cn_k,cc_club},
	card{cn_k,cc_diamond},

	card{cn_blackJk,cc_spade},
	card{cn_redJk,cc_heart},
}


type compareRet int
const (
	cp_small compareRet = iota - 1
	cp_equal
	cp_big
)

type cardType int
const (
	ct_cow0 cardType = iota
	ct_cow1
	ct_cow2
	ct_cow3
	ct_cow4
	ct_cow5
	ct_cow6
	ct_cow7
	ct_cow8
	ct_cow9
	ct_cow10
)

// max player in a game
const maxPlayer = 5

// 发牌后到结算之间的时间
const timeout_settle = 15
//打扑克牌用到的英语
//card games 打牌
//cards 纸牌
//pack (of cards),deck 一副牌
//suit 一组
//joker 百搭
//ace A牌
//king 国王,K
//queen 王后,Q
//Jack 王子,J
//face cards,court cards 花牌(J,Q,K)
//clubs 梅花,三叶草
//diamonds 方块,红方,钻石
//hearts 红桃,红心
//spades 黑桃,剑
//trumps 胜
//to shuffle 洗牌
//to cut 倒牌
//to deal 分牌
//banker 庄家
//hand 手,家
//to lead 居首
//to lay 下赌
//to follow suit 出同花牌
//to trump 出王牌
//to overtrump 以较大的王牌胜另一张王牌
//to win a trick 赢一圈
//to pick up,to draw 偷
//stake 赌注
//to stake 下赌注
//to raise 加赌注
//to see 下同样赌注
//bid 叫牌
//to bid 叫牌
//to bluff 虚张声势
//royal flush 同花大顺
//straight flush 同花顺
//straight 顺子
//four of a kind 四张相同的牌
//full house 三张相同和二张相同的牌
//three of a kind 三张相同的牌
//two pairs 双对子
//one pair 一对,对子
//大王(red Joker) 小王( black Joker)