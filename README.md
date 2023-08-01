### Описание. Description  
Пакет для считывания данных с аналовоговых `ЭМГ` датчиков (`EMG Sensors`).  
Подключение `AD8232` (`Heart Rate Monitor`) через АЦП `ADS1115` (`Analog to Digital Converter`) к шине `I2C` (`SDA-SCL`) Raspberry Pi. 

### Схема подключения. Electrical diagram  
![Компонентная схема](https://github.com/blablatov/raspberry-ads1x15-ad8232/raw/master/scheme.png)  

### Сборка. Build  
	set GOOS=linux
	set GOARCH=arm
	set GOARM=6
	go build -o hmonitor hmonitor.go  

### Использование. How use.    
	hmonitor  
:sos: