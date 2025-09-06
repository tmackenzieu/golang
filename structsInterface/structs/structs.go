package structs

type Product struct {
	ID    uint     `json:"id"`
	Name  string   `json:"name"`
	Type  Type     `json:"type"`
	Price float64  `json:"price"`
	Count float32  `json:"count"`
	Tags  []string `json:"tags"`
}

type Type struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Car struct {
	ID       uint      `json:"id"`
	UserID   uint      `json:"user_id"`
	Products []Product `json:"products"`
}

func (p Product) TotalPrice() float64 {
	return float64(p.Count) * p.Price
}

func (p *Product) SetName(name string) {
	p.Name = name
}

func (p *Product) AddTags(tags ...string) {
	p.Tags = append(p.Tags, tags...)
}

func (c *Car) AddProducts(products ...Product) {
	c.Products = append(c.Products, products...)
}

func (c Car) Total() float64 {
	var total float64
	for _, p := range c.Products {
		total += p.Price
	}
	return total
}

func NewCar(userID uint) Car {
	return Car{UserID: userID}
}
