<?php 
namespace App\Http\Controllers;
use App\Http\Controllers\dealCardController;
class cardController extends Controller{
    public function cardget(){
     $color=["S","H","D","C"];
     $number=[2,3,4,5,6,7,8,9,10,"J","Q","K","A"];
    $cards=[];
    
    foreach($number as $n){
        foreach($color as $c){
            $cards[]=$c.$n;
        }
    }
      //$bitwise=dealCardController::bitwise();
      
     return view('home',compact('cards'));
    }   
}