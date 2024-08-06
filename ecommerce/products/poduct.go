package products

type Product struct {
	Name string
	Qty, Id int32
	Price float64
}

type Products struct {
	Products []Product
}

func (p *Products) getProduct(Id int32) Product {
	var pd Product;

	for _, ele := range p.Products {
		if ele.Id == Id {
			pd = ele
		}
	}

	return pd
}

func (p *Products) getProductPrice(Id int32) float64 {
	var pr float64

	for _, ele := range p.Products {
		if ele.Id == Id {
			pr = ele.Price
		}
	}

	return pr
}

func (p *Products) getAllProducts() []Product {
	var allP []Product

	allP = append(allP, p.Products...)

	return allP
}

func (p *Products) isProductAvailable(Id int32) bool {
	pp := p.getProduct(Id)

	return pp.Qty != 0
}