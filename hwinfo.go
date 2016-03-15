package hwinfo

type HWInfo interface {
	Update() error
	GetData() data
	GetCache() cache
}

type hwInfo struct {
	data  *data  `json:"data"`
	cache *cache `json:"cache"`
}

func New() HWInfo {
	return &hwInfo{
		data:  &data{},
		cache: &cache{},
	}
}

func (h *hwInfo) GetData() data {
	return *h.data
}

func (h *hwInfo) GetCache() cache {
	return *h.cache
}
