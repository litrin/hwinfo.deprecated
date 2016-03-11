package routes

type Route interface {
	SetTTL(int)
	Get() error
}

type route struct {
	Destination string `json:"destination"`
	Gateway     string `json:"gateway"`
	Genmask     string `json:"genmask"`
	Flags       string `json:"flags"`
	MSS         int    `json:"mss"` // Maximum segment size
	Window      int    `json:"window"`
	IRTT        int    `json:"irtt"` // Initial round trip time
	Interface   string `json:"interface"`
}

// New memory constructor.
func New() *route {
	return &route{
		TTL: 5 * 60,
	}
}

// Get memory info.
func (m *memory) Get() error {
	if m.Last.IsZero() {
		if err := m.get(); err != nil {
			return err
		}
		m.Last = time.Now()
	} else {
		expire := m.Last.Add(time.Duration(m.TTL) * time.Second)
		if expire.Before(time.Now()) {
			if err := m.get(); err != nil {
				return err
			}
		}
	}

	return nil
}
