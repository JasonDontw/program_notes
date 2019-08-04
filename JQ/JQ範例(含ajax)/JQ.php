<!DOCTYPE html>
<html>

<head>
    <meta charset="UTF-8">
    <script src="https://code.jquery.com/jquery-3.4.1.min.js" integrity="sha256-CSXorXvZcTkaix6Yvo6HppcZGetbYMGWSFlBw8HfCJo=" crossorigin="anonymous">
    </script> <!--必須放於任何JQ語法之前-->
</head>

<body>
<!--------元素------------->
<h1 id="title">fuck</h1>
<h1 id="title">you</h1>

<h1 class="title">you</h1>
<h1 class="title">you</h1>
<h1 class="title">you</h1>
<h1 class="title">you</h1>
<form>
<div id="myDiv"><h2>使用 Ajax 修改内容</h2></div>
<input type="text" id="in">
<button type="button" id="btn1">修改内容</button>
<button type="button" id="btn2">修改内容2</button>
</form>
<script>
//-------------JQ修改方法-------------
    $('#title').html("id修改");//js叫innerHTML在JQ只需hmtl，若有2個相同ID只會改第一個
    
    $(document).ready(function(){ //網頁載完後才載入，精簡寫法$(function(){ });
        $('.title').html("class修改");//相同class會全部改
    });

$(document).ready(function(){
  $("#btn1").click(function(){
    $('#myDiv').load('JQ.txt');
  })
})


$("#btn2").click(function(){
  $.ajax({
        url: "JQ2.php",
        type:'get', 
        
        data:{'y':$("#in").val()},
        success: function(x){ $("#myDiv").html(x);
        },
        error:function(y){
          $("#myDiv").html("no");
        }
        }
      );
    }
)

</script>


</body>
</html>