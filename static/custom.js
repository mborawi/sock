// (function(){
	var ws= new WebSocket("ws://localhost:7070/ws");
	var msg = document.getElementById("messages");
	var mymap = L.map('mapid').setView([-26.0, 135.09], 5);
	L.tileLayer('https://api.tiles.mapbox.com/v4/{id}/{z}/{x}/{y}.png?access_token={accessToken}', {
	    attribution: 'Map data &copy; <a href="http://openstreetmap.org">OpenStreetMap</a> contributors, <a href="http://creativecommons.org/licenses/by-sa/2.0/">CC-BY-SA</a>, Imagery © <a href="http://mapbox.com">Mapbox</a>',
	    maxZoom: 18,
	    id: 'mapbox.streets',
	    accessToken: 'pk.eyJ1IjoibWJvcmF3aSIsImEiOiJjamJwM21hdWEzeXl5MzRxd2x6MThxd3Y5In0.1z-haSkFj0PjIqCh3WJ81g'
	}).addTo(mymap);
	var Sydney = L.circle([-33.8772695,151.2044971], {
	    color: '#e60000',
	    fillColor: '#e60000',
	    fillOpacity: 0.5,
	    radius: 10
	}).addTo(mymap);
	Sydney.bindTooltip("Syndey: 10% Occupancy").openTooltip();
	var NewCastle = L.circle([-32.9284036,151.7716398], {
	    color: '#0000e6',
	    fillColor: '#0000e6',
	    fillOpacity: 0.5,
	    radius: 10
	}).addTo(mymap);
	NewCastle.bindTooltip("NewCastle: 10% Occupancy").openTooltip();
	var Melbourne = L.circle([-37.8208437,144.947175], {
	    color: '#006600',
	    fillColor: '#006600',
	    fillOpacity: 0.5,
	    radius: 10
	}).addTo(mymap);
	Melbourne.bindTooltip("Melbourne: 10% Occupancy").openTooltip();
	var Canberra = L.circle([-35.2774059,149.1331512], {
	    color: '#ffff00',
	    fillColor: '#ffff00',
	    fillOpacity: 0.5,
	    radius: 10
	}).addTo(mymap);
	Canberra.bindTooltip("Canberra: 10% Occupancy").openTooltip();

	ws.addEventListener("message", 
		function(e){
			result = JSON.parse(e.data);

			Sydney.setRadius( result.sydney * 2000 );
			// var title = 
			Sydney.bindTooltip("Sydney: " + result.sydney.toFixed(2) + "% Occupancy").openTooltip();

			NewCastle.setRadius( result.newcastle * 1000 );
			NewCastle.bindTooltip("Newcastle: " + result.newcastle.toFixed(2) + "% Occupancy").openTooltip();

			Melbourne.setRadius( result.melbourne * 2400 );
			Melbourne.bindTooltip("Melbourne: " + result.melbourne.toFixed(2) + "% Occupancy").openTooltip();

			Canberra.setRadius( result.canberra * 2000 );
			Canberra.bindTooltip("Canberra: " + result.canberra.toFixed(2) + "% Occupancy").openTooltip();


	});
// })();
