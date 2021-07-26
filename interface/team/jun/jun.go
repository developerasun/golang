package main

// Code written by Junseok Oh
import (
	"fmt"
	"time"
)

type MakeBugger interface {
	String() string
}

type Meterial interface {
	GetOnMeterial() MakeBugger
}

type Bread struct {
	val string
}

//////////////////////////////////
type Patty struct {
}

func (j *Patty) GetOnMeterial() MakeBugger {
	return &PutOnPatty{}
}

type PutOnPatty struct {
}

func (s *PutOnPatty) String() string {
	return " + 쇠고기패티"
}

//////////////////////////////////

type MakeMenu interface {
	Setting() string
}

type Menus interface {
	PulsMenu() MakeMenu
}

type Set struct {
	me string
}

func (b *Bread) PutMeterial(meterial Meterial) {
	shief := meterial.GetOnMeterial()
	b.val += shief.String()
}

func (b *Bread) String() string {
	return "참깨빵" + b.val + " + 참깨빵"
}

func (m *Set) PulsMenus(menu Menus) {
	setmenu := menu.PulsMenu()
	m.me += setmenu.Setting()
}

func (m *Set) Setting() string {
	return "따뜻한 아메리카노" + m.me
}

type Poteto struct {
}

func (j *Poteto) PulsMenu() MakeMenu {
	return &MakePoteto{}
}

type Cheese struct {
}

func (j *Cheese) GetOnMeterial() MakeBugger {
	return &PutOnCheese{}
}

type Sauce struct {
}

func (j *Sauce) GetOnMeterial() MakeBugger {
	return &PutOnSauce{}
}

type Lettuce struct {
}

func (j *Lettuce) GetOnMeterial() MakeBugger {
	return &PutOnLettuce{}
}

type Onion struct {
}

func (j *Onion) GetOnMeterial() MakeBugger {
	return &PutOnOnion{}
}

type Pickle struct {
}

func (j *Pickle) GetOnMeterial() MakeBugger {
	return &PutOnPickle{}
}

type PutOnCheese struct {
}

func (s *PutOnCheese) String() string {
	return " + 치즈"
}

type PutOnSauce struct {
}

func (s *PutOnSauce) String() string {
	return " + 특별한소스"
}

type PutOnLettuce struct {
}

func (s *PutOnLettuce) String() string {
	return " + 양상추"
}

type PutOnPickle struct {
}

func (s *PutOnPickle) String() string {
	return " + 피클"
}

type PutOnOnion struct {
}

func (s *PutOnOnion) String() string {
	return " + 양파"
}

type MakePoteto struct {
}

func (s *MakePoteto) Setting() string {
	return " + 감자튀김"
}

func main() {
	// 햄버거 재료
	bread := &Bread{}
	cheese := &Cheese{}
	patty := &Patty{}
	sauce := &Sauce{}
	lettuce := &Lettuce{}
	onion := &Onion{}
	pickle := &Pickle{}

	// 셋트 메뉴
	set := &Set{}
	poteto := &Poteto{}

	// 주문 시간 표현 1
	maketime := time.Now()

	print("[ 동대구 맥도날드 ]")

	fmt.Println("주문 시간 :", maketime)

	time.Sleep(3 * time.Second)

	bread.PutMeterial(patty)
	bread.PutMeterial(patty)
	bread.PutMeterial(sauce)
	bread.PutMeterial(lettuce)
	bread.PutMeterial(cheese)
	bread.PutMeterial(pickle)
	bread.PutMeterial(onion)

	fmt.Println("햄버거 :", bread)

	set.PulsMenus(poteto)

	fmt.Println("세트 메뉴 :", set)

	// 주문 시간 표현 2
	endtime := time.Since(maketime)

	fmt.Printf("약 %v 걸렸습니다.", endtime)
}
