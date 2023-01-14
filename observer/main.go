package main

import "fmt"

type observer interface {
	update(temperature, humidity, pressure string)
}

type subject interface {
	registerObserver(o observer)
	removeObserver(o observer)
	notifyObserver()
}

type displayElement interface {
	display()
}

type WeatherData struct {
	observers   []observer
	temperature string
	humidity    string
	pressure    string
}

func newWeatherData() *WeatherData {
	return &WeatherData{
		observers: []observer{},
	}
}

func (w *WeatherData) indexOf(o observer) int {
	for k, v := range w.observers {
		if o == v {
			return k
		}
	}
	return -1
}

func (w *WeatherData) registerObserver(o observer) {
	w.observers = append(w.observers, o)
}

func (w *WeatherData) removeObserver(o observer) {
	i := w.indexOf(o)
	w.observers = append(w.observers[:i], w.observers[i+1:]...)
}

func (w *WeatherData) notifyObserver() {
	for _, val := range w.observers {
		val.update(w.temperature, w.humidity, w.pressure)
	}

}

func (w *WeatherData) Temperature() string {
	return w.temperature
}

func (w *WeatherData) SetTemperature(temperature string) {
	w.temperature = temperature
}

func (w *WeatherData) Humidity() string {
	return w.humidity
}

func (w *WeatherData) SetHumidity(humidity string) {
	w.humidity = humidity
}

func (w *WeatherData) Pressure() string {
	return w.pressure
}

func (w *WeatherData) SetPressure(pressure string) {
	w.pressure = pressure
}

func (w *WeatherData) measurementChanged() {
	w.notifyObserver()
}

func (w *WeatherData) setMeasurements(t, h, p string) {
	w.temperature = t
	w.humidity = h
	w.pressure = p
	w.measurementChanged()
}

type CurrentConditionDisplay struct {
	temperature string
	humidity    string
	data        *WeatherData
}

func newCurrentConditionDisplay(data WeatherData) *CurrentConditionDisplay {
	c := &CurrentConditionDisplay{data: &data}
	data.registerObserver(c)
	return c
}

func (c CurrentConditionDisplay) update(temperature, humidity, pressure string) {
	c.temperature = temperature
	c.humidity = humidity
	c.display()
}

func (c CurrentConditionDisplay) display() {
	fmt.Println("Current conditions: " + c.temperature + "F degrees and " + c.humidity + "% humidity")
}

func main() {
	weatherData := newWeatherData()
	c := newCurrentConditionDisplay(*weatherData)
	c.data.setMeasurements("80", "60", "30")
	c.data.setMeasurements("82", "70", "29")
}
