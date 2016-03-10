package game

type RoomType int
const (
	RoomType_0 RoomType = iota
	RoomType_1
	RoomType_2
	RoomTypeCount
)

func (rt RoomType)IsValid() bool {
	return 0 <= rt < RoomTypeCount
}

var RoomEnterCoin = [RoomTypeCount]int{200, 300, 400}