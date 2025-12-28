<?php
namespace App\Http\Controllers;
use Illuminate\Support\Facades\Http;
class sendCard extends Controller{

public function sendmessage($card1,$card2){
   $data =[
    "hand1"=>$card1,
    "hand2"=>$card2
   ];
   error_log("收到牌组数据: " . print_r($data, true));
   //给go发数据 JSON
   $send=Http::post("http://localhost:8080/message",$data);
   
}
}