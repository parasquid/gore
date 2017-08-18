package packets

type MasterLogin struct {
	ID            uint16
	Version       uint32
	Username      [24]byte
	Password      [24]byte
	MasterVersion uint8
}
