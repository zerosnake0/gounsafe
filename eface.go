package gounsafe

import "unsafe"

type Eface struct {
	RType RType
	Data  unsafe.Pointer
}
