package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const normalHrFloor = 60
const normalHrCeil = 90
const normalSPO2Floor = 95
const normalSPO2Ceil = 99
const normalBPSysFloor = 90
const normalBPSysCeil = 120
const normalBPDiaFloor = 60
const normalBPDiaCeil = 80
const normalRRFloor = 10
const normalRRCeil = 25

func randInRange(min int, max int) (randVal int) {
	return min + rand.Intn(max-min+1)
}

func pt(ptID int, interval int, event chan string) {
	ptHR := randInRange(normalHrFloor, normalHrCeil)
	ptSPO2 := randInRange(normalSPO2Floor, normalSPO2Ceil)
	ptBPSys := randInRange(normalBPSysFloor, normalBPSysCeil)
	ptBPDia := ptBPSys - int(float64(ptBPSys)*.28)
	ptRR := randInRange(normalRRFloor, normalRRCeil)

	for {
		time.Sleep(time.Duration(interval) * time.Millisecond)
		ptHR += randInRange(-1, 1)
		ptBPSys += randInRange(-2, 2)
		ptBPDia += randInRange(-2, 2)
		ptRR += randInRange(-1, 1)
		event <- fmt.Sprintf("ID: %d, HR: %d, SpO2: %d, BP: %d/%d, awRR: %d\n", ptID, ptHR, ptSPO2, ptBPSys, ptBPDia, ptRR)
	}
}

func main() {
	intervalPtr := flag.Int("interval", 1000, "Sample generation interval in MS for each simulation")
	ptCountPtr := flag.Int("pt-count", 10, "Number of patients to simulate concurrently")
	flag.Parse()

	event := make(chan string)

	for i := 1; i <= *ptCountPtr; i++ {
		go pt(i, *intervalPtr, event)
	}
	for {
		fmt.Printf(<-event)
	}
}
