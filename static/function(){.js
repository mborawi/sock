(function(){
	var ws= new WebSocket("ws://localhost:7070/ws");
	ws.addEventListener("message", function(e){console.log(e)});
})();