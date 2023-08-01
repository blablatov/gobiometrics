// Code sketch of data exchange Raspberry-ads1x15-ad8232

package main

import (
	"log"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	log.SetPrefix("I2C events: ")
	log.SetFlags(log.Llongfile)

	// Инициализатор адаптера. Inits new adaptor
	adaptor := raspi.NewAdaptor()
	log.Println("adaptor: ", adaptor)

	// Default bus/address - i2c.NewADS1115Driver(i2c.Connector, i2c.Config)
	ads := i2c.NewADS1115Driver(adaptor)
	log.Println(ads) // Тест данных. For test

	// Инициализатор датчика. Initializes the sensor.
	s := ads.Start()
	if s != nil {
		log.Fatalf("Err open sensor: ", s)
	}
	log.Println(s) // // Тест данных. For test

	// Читает данные с аналоговых выходов указанных пинов
	// Returns value from analog out pins of specified pins
	for {
		r, err := ads.AnalogRead("0-1")
		if err != nil {
			log.Fatalf("Err analog read: ", r)
		}
		log.Println(r)
		time.Sleep(2 * time.Second) // Тайм-аут опроса. Timeout
	}

	// Инициирует и стартует новый процесс робота
	// Inits the new proccess of robot
	robot := gobot.NewRobot("AnalogRead",
		[]gobot.Adaptor{adaptor},
		[]gobot.Driver{ads},
	)

	robot.Start()
}
