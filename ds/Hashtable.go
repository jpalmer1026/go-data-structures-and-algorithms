package ds

type HashTable struct {
	data []*int
}

func NewHashTable(size int) *HashTable {
	h := HashTable{}
	h.data = make([]*int, size)

	return &h
}

func (h *HashTable) hash(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = (hash + int(key[i])*i) % len(h.data)
	}

	return hash
}

func (h *HashTable) Set(key string, value int) {
	i := h.hash(key)
	if h.data[i] == nil {
		h.data[i] = &value
	}
}

func (h *HashTable) Get(key string) int {
	i := h.hash(key)

	return *h.data[i]
}
