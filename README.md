# terrascout
This project provides details about flora and fauna in the wild using a preselected region from the map, then performs predictions on images taken by running two stage pipelined MobileNetV2 and EfficientNetB3 locally on Device.


Demo Map Bounding box (Sequoia National Forest Area)
SW : 35.867185, -118.612568
NE : 36.287875, -118.114577


## frontend
Displays Map with a selectable region
Allow using Webcam to click picture and display predicted output.


## backend
1. Create a type to store and fetch data in memory. (or json)
2. Run ML model on localhost
3. call model with image
4. Model detects type (insect, plant, etc) and species and returns species name.
5. Data Source: https://www.inaturalist.org/home


## Running backend

```
cd backend
go run main.go
```

##Serve ML Model
```
python3 serve.py
```





