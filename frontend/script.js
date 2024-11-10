// Initialize the map
var map = L.map('map', {
    zoomControl: false,           // Disable zoom controls
    scrollWheelZoom: false,       // Disable zooming via scroll wheel
    doubleClickZoom: false,       // Disable zooming via double-click
    touchZoom: false
}).setView([36.07753, -118.36357], 11);
// Add OpenStreetMap tiles
L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
}).addTo(map);

// Initialize the FeatureGroup to store editable layers
var drawnItems = new L.FeatureGroup();
map.addLayer(drawnItems);
var selectAreaBtn = document.getElementById('select-area-btn');

// Configure Leaflet Draw
var drawControl = new L.Control.Draw({

    draw: {
        polygon: false,
        polyline: false,
        circle: false,
        marker: false,
        circlemarker: false,
        rectangle: true
    }
});
map.addControl(drawControl);

// Event listener for when a rectangle is created
map.on(L.Draw.Event.CREATED, function (event) {
    var layer = event.layer;
    drawnItems.addLayer(layer);

    // Get the bounds of the selected rectangle
    var bounds = layer.getBounds();
    // alert("You selected a region with bounds:\n" +
    //     "SouthWest: " + bounds.getSouthWest().toString() + "\n" +
    //     "NorthEast: " + bounds.getNorthEast().toString());

    selectAreaBtn.disabled = false;
    selectAreaBtn.classList.add('active');
});
