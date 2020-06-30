package providers

import "fmt"

// enum defind
type ProviderType int

const (
	AutoMessage = 1
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
		fmt.Println("waiting..")
		break
	}
}
