
type DogAdapter struct {
	*Dog
}

func (adapter *DogAdapter) Reaction() {
	adapter.Woof()
}

func NewDogAdapter(dog *Dog) AnimalAdapter {
	return &DogAdapter{dog}
}

type CatAdapter struct {
	*Cat
}

func (adapter *CatAdapter) Reaction() {
	adapter.Meow(true)
}

func NewCatAdapter(cat *Cat) AnimalAdapter {
	return &CatAdapter{cat}
}