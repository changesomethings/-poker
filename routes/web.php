<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\cardController;
use App\Http\Controllers\dealCardController;
Route::get('/', function () {
    return view('welcome');
});
Route::get('/home',[cardController::class,'cardget']);
Route::post('/set',[dealCardController::class,'deal']);