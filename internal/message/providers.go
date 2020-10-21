package message

import "fmt"

type ProviderType int
const (
	AutoMessage ProviderType = iota
)
func (p ProviderType) String() string {
	return [...]string{"Message"}[p]
}

type Payload struct {
	KeyWords []string
}

func DispatchToProvider(pt ProviderType, p Payload) {
	switch pt {
	case AutoMessage:
		fmt.Println("AutoMessage", p.KeyWords)
		break
	default:
		fmt.Println("do nothing...")
		break
	}
}
