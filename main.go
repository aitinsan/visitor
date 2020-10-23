package main


import "fmt"

type Visitor1 interface{
	BuyInMainShop(s *MainShop) string
	BuyInHiddenShop(s *HiddenShop) string
	BuyInSideShop(s *SideShop) string
}
type Shop interface{
	Accept (v Visitor1) string
}
type Hero struct{

}
func (h *Hero) BuyInMainShop(m *MainShop) string {
	return m.BuyMain()
}
func (h *Hero) BuyInHiddenShop(hi *HiddenShop) string{
	return hi.BuyHidden()
}
func (h *Hero) BuyInSideShop(s *SideShop) string{
	return s.BuySide()
}


type DotaLand struct{
	shope []Shop
}
func (d *DotaLand) Add(s Shop){
	d.shope=append(d.shope,s)
}


func (d *DotaLand) Accept(v Visitor1) string {
	var result string

	for  _, p:= range d.shope {
		result += p.Accept(v)
	}
	return result
}

type MainShop struct {
}


func (m *MainShop) Accept(v Visitor1) string {
	return v.BuyInMainShop(m)
}


func (m *MainShop) BuyMain() string {
	return "buy in main shop \n"
}
type HiddenShop struct {
}


func (h *HiddenShop) Accept(v Visitor1) string {
	return v.BuyInHiddenShop(h)
}


func (h *HiddenShop) BuyHidden() string {
	return "buy in hidden shop \n"
}
type SideShop struct {
}


func (h *SideShop) Accept(v Visitor1) string {
	return v.BuyInSideShop(h)
}


func (h *SideShop) BuySide() string {
	return "buy in side shop \n"
}
func main(){
	dotaland := new(DotaLand)

	dotaland.Add(&MainShop{})
	dotaland.Add(&HiddenShop{})
	dotaland.Add(&SideShop{})
	fmt.Println(dotaland.Accept(&Hero{}))
}



