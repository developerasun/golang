package main // interface 방식으로 변경한 햄버거 만들기~
// Code written by Taewoong Kim
import "fmt"

type OneFilling interface { //외부 메서드 = 관계
	String() string // OneFilling은 String()이라는 외부공개메서드이다.
}
type Filling interface { //오로지 관계만 선언한다.
	GetOneFilling() OneFilling // Filling(속재료)은 GetOneFilling() 이라는 외부 공개 메서드이다.
}
type Bread struct { // bread 객체를 만든다
	val string
}

func (b *Bread) AddFilling(filling Filling) { // filling(속)은 종류가 무관하다.
	onefilled := filling.GetOneFilling()
	b.val += onefilled.String()
}

func (b *Bread) String() string { //Bread 매서드2 string
	return "[ bread " + b.val + " + bread ]"
}

////////////////////////////////////////
type MeetPatty struct{}

func (i *MeetPatty) GetOneFilling() OneFilling {
	return &FilledMeetPatty{}
}

type FilledMeetPatty struct{}

func (l *FilledMeetPatty) String() string {
	return " + MeetPatty"
}

//////////////////////////////////////////

type SlicedCheese struct{}

func (j *SlicedCheese) GetOneFilling() OneFilling {
	return &FilledSlicedCheese{}
}

type FriedEggs struct{}

func (k *FriedEggs) GetOneFilling() OneFilling {
	return &FilledFriedEggs{}
}

type FilledSlicedCheese struct{}

func (m *FilledSlicedCheese) String() string {
	return " + SlicedCheese"
}

type FilledFriedEggs struct{}

func (n *FilledFriedEggs) String() string {
	return " + FriedEggs"
}

func main() {
	bread := &Bread{}
	filling := &MeetPatty{}
	filling2 := &FriedEggs{}
	filling3 := &SlicedCheese{}

	bread.AddFilling(filling)
	bread.AddFilling(filling2)
	bread.AddFilling(filling3)

	fmt.Println("[버거킹]")
	fmt.Println(bread, "\n[MeetPattybuger complete]")
}
