package poker

//
//type InMemoryPlayerStore struct {
//	store map[string]int
//}
//
//func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
//	return i.store[name]
//}
//
//func (i *InMemoryPlayerStore) RecordWin(name string) {
//	i.store[name]++
//}
//
//func (i *InMemoryPlayerStore) GetLeague() League {
//	var league League
//	for name, wins := range i.store {
//		league = append(league, Player{name, wins})
//	}
//	return league
//}
//
//func NewInMemoryPlayerStore() *InMemoryPlayerStore {
//	return &InMemoryPlayerStore{store: make(map[string]int)}
//}
