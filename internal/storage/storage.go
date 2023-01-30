package storage

type Storage struct {
	dB DB
}

type DB interface {
	Init()
	// Close()
	// Ping() error
	// GetBalance(user string)
	// PostBalance(user string, balance float32)
	// PostOrder(order string) error
	// GetOrdersInformation() ([]string, error)
	// GetUser(uid string)
}

const dbTimeout = 5

func NewStore(db DB) *Storage {
	return &Storage{dB: db}
}

func (s *Storage) SaveUser(username string, password string) (string, error) {
	return "", nil
}

func (s *Storage) GetUser(username string, password string) (string, error) { return "", nil }

func (s *Storage) GetBalance() {}
func (s *Storage) SetBalance() {}
