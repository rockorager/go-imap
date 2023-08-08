package imap

// StoreOptions contains options for the STORE command.
type StoreOptions struct {
	UnchangedSince uint64 // requires CONDSTORE
}

type StoreData interface {
	dataName() string
}

// StoreFlagsOp is a flag operation: set, add or delete.
type StoreFlagsOp int

const (
	StoreFlagsSet StoreFlagsOp = iota
	StoreFlagsAdd
	StoreFlagsDel
)

// StoreFlags alters message flags.
type StoreFlags struct {
	Op     StoreFlagsOp
	Silent bool
	Flags  []Flag
}

func (s StoreFlags) dataName() string {
	return "FLAGS"
}

type StoreXGmLabels struct {
	Op        StoreFlagsOp
	Silent    bool
	XGmLabels []string
}

func (s StoreXGmLabels) dataName() string {
	return "X-GM-LABELS"
}
