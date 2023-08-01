package main

import (
	"log"
	"testing"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

func TestAdaptor(t *testing.T) {
	// Инициализатор адаптера. Inits new adaptor
	Adaptor := raspi.NewAdaptor()
	if Adaptor != nil {
		log.Fatalf("err test adaptor")
	}
	log.Println("adaptor: ", adaptor)
}

func TestNewADS(t *testing.T) {
	// Default bus/address - i2c.NewADS1115Driver(i2c.Connector, i2c.Config)
	Ads := i2c.NewADS1115Driver(adaptor)
	if Ads != nil {
		log.Fatalf("err test newADSdriver")
	}
	log.Println(ads) // Тест данных. For test
}

func TestStart(t *testing.T) {
	// Инициализатор датчика. Initializes the sensor.
	s := ads.Start()
	if s != nil {
		log.Fatalf("Err open sensor: ", s)
	}
	log.Println(s) // // Тест данных. For test
}

func TestAnalogRead(t *testing.T) {
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
}

func TestInitRobot(t *testing.T) {
	// Инициирует и стартует новый процесс робота
	// Inits the new proccess of robot
	robot := gobot.NewRobot("AnalogRead",
		[]gobot.Adaptor{Adaptor},
		[]gobot.Driver{Ads},
	)

	robot.Start()
}
