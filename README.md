## Raspberry-ads1x15-ad8232  

### Описание. Description    
Модуль для считывания данных с аналоговых `ЭМГ` датчиков (`EMG Sensors`).  
Подключение `AD8232` (`Heart Rate Monitor`) через АЦП `ADS1115` (`Analog to Digital Converter`) к шине `I2C` (`SDA-SCL`) Raspberry Pi. 

### Схема подключения. Electrical diagram  
![Компонентная схема](https://github.com/blablatov/gobiometrics/raw/master/scheme.png)  

### Сборка. Build  
	set GOOS=linux
	set GOARCH=arm
	set GOARM=6
	go build -o hmonitor hmonitor.go  

### Использование. How use.    
	hmonitor  
:sos:
