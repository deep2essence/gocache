package cache

type Data []byte

type Node struct {
	key  int
	data *Data
	next *Node
	prev *Node
}

type NodeList struct {
	header   *Node
	footer   *Node
	capacity int
}

func NewNodeList() *NodeList {
	return &NodeList{}
}

func (c *NodeList) Get(key int) *Data {
	if c.header == nil {
		return nil
	}
	for node := c.header; node != nil; node = node.next {
		if node.key == key {
			return node.data
		}
	}
	return nil
}

func (c *NodeList) Remove(key int) {
	if c.header == nil {
		return
	}
	for node := c.header; ; node = node.next {
		if node.key == key {
			prev := node.prev
			next := node.next
			prev.next = next
			next.prev = prev
			c.capacity--
			return
		}

	}
}

func (c *NodeList) Put(key int, value *Data) {
	node := Node{key: key, data: value}
	c.capacity++
	if c.header == nil {
		c.header = &node
		c.footer = &node
		return
	}
	node.prev = c.footer
	c.footer.next = &node
	c.footer = &node
}

type Cache struct {
	capacity int
	pool     []NodeList
}

func NewCache(capacity int) *Cache {
	pool := make([]NodeList, capacity)
	return &Cache{capacity: capacity, pool: pool}
}

func (c *Cache) id(key int) int {
	return key % c.capacity
}

func (c *Cache) Get(key int) *Data {
	return c.pool[c.id(key)].Get(key)
}

func (c *Cache) Remove(key int) {
	c.pool[c.id(key)].Remove(key)
}
func (c *Cache) Put(key int, value *Data) {
	c.pool[c.id(key)].Put(key, value)
}
