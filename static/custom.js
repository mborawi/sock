(function(){
	var ws= new WebSocket("ws://localhost:7070/ws");
	var msg = document.getElementById("messages");

	ws.addEventListener("message", 
		function(e){
			var item = document.createElement("div");
			item.innerHTML = e.data;
			msg.appendChild(item);
	});
})();