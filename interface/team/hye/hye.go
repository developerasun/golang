package main

// Code written by Hyesu Kwon
import "fmt"

type SpoonOfJam interface {
	String() string
}

type Jam interface {
	GetOneSpoon() SpoonOfJam
}

type AddOfFriedEggs interface {
	String() string
}

type Egg interface {
	GetOneEgg() AddOfFriedEggs
}

type SliceOfTomato interface {
	String() string
}

type Tomato interface {
	GetOneTomato() SliceOfTomato
}

type Bread struct {
	val string
}

func (b *Bread) PutJam(jam Jam) {
	spoon := jam.GetOneSpoon()
	b.val += spoon.String()
}

func (b *Bread) PutEgg(egg Egg) {
	add := egg.GetOneEgg()
	b.val += add.String()
}
func (b *Bread) PutTomato(tomato Tomato) {
	slice := tomato.GetOneTomato()
	b.val += slice.String()

}
func (b *Bread) String() string {
	return "bread" + b.val
}

type StrawberryJam struct{}

func (j *StrawberryJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfStrawberryJam{}
}

type OrangeJam struct{}

func (j *OrangeJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfOrangeJam{}
}

type AppleJam struct{}

func (j *AppleJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfAppleJam{}
}

type FriedEggs struct{}

func (j *FriedEggs) GetOneEgg() AddOfFriedEggs {
	return &AddOfEggs{}
}

type AddOfEggs struct{}

func (s *AddOfEggs) String() string {
	return " + FriedEggs"
}

type SliceTomato struct{}

func (j *SliceTomato) GetOneTomato() SliceOfTomato {
	return &Slicetomato{}
}

type Slicetomato struct{}

func (s *Slicetomato) String() string {
	return " + Tomato"
}

type SpoonOfStrawberryJam struct{}

func (s *SpoonOfStrawberryJam) String() string {
	return " + strawberry jam"
}

type SpoonOfOrangeJam struct{}

func (s *SpoonOfOrangeJam) String() string {
	return " + Orange jam"
}

type SpoonOfAppleJam struct{}

func (s *SpoonOfAppleJam) String() string {
	return " + Apple jam"
}

func main() {
	bread := &Bread{}
	jam := &AppleJam{}
	egg := &FriedEggs{}
	tomato := &SliceTomato{}

	fmt.Println("동대구역 빅버거")
	bread.PutJam(jam)
	bread.PutEgg(egg)
	bread.PutTomato(tomato)
	fmt.Print(bread)
	fmt.Print(" + bread")
}
