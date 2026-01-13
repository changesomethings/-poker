package service
import "fmt"
import "poker/model"
type Texas struct {
   
}
func (t*Texas)Check(play[]*model.Player){
	fmt.Println("德州judge")
     for _,v:=range play{
       fmt.Println(v)
	 }
}