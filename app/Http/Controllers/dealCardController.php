<?php 
namespace App\Http\Controllers;
use Illuminate\Http\Request;
use App\Http\Controllers\cardController;
use App\Http\Controllers\sendCard;
class dealCardController extends Controller{
 public   static function   bitwise(){
    $color=[0,1,2,3];
    $number=range(2,14);
    $bitcards=[];
    foreach($number as$n){
        foreach ($color as $a){
            $bitcards[]=$n<<2|$a;
        }     
    }
      return $bitcards;
 }
 public  static function dealCard($card){
    $colorMap=["S"=>0,"H"=>1,"D"=>2,"C"=>3];
    $numberMap=["2"=>2,"3"=>3,"4"=>4,"5"=>5,"6"=>6,"7"=>7,"8"=>8,
    "9"=>9,"10"=>10,"J"=>11,"Q"=>12,"K"=>13,"A"=>14];
    $c=substr($card,0,1);
    $n=substr($card,1);
    $number=$numberMap[$n];
    $color=$colorMap[$c];
    return $number<<2|$color;
 }
 public function deal(Request $request){
    $card1=$request->input('card1');
    $card2=$request->input('card2');
    $bitcard1=dealCardController::dealCard($card1);
    $bitcard2=dealCardController::dealCard($card2);
    
    //$bitcards=dealCardController::bitwise();//显
     $sender=new sendCard();

    $request=$sender->sendmessage($bitcard1,$bitcard2);//把数据打包JSON发给go去算
    return view('welcome');
 }
}