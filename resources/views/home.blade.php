<!DOCTYPE html>
<html>
<head>
    
    <title>Home</title>
    <meta name="viewport" content="width=device-width,initial-scale=1,maximum-scale=1">
</head>
<body>
    <div class="cards">
     @foreach($cards as $c)
   <span>{{$c}}</span>
   @endforeach
    </div>
    <form action="/set" method="post">
        @csrf
       <input type="text" name="card1">
       <input type="text" name="card2">
       <button>计算<button>
    </form>
</body>
</html>