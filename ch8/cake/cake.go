package cake

import (
	"fmt"
	"math/rand"
	"time"
)

type Shop struct {
	Verbose        bool
	Cakes          int           // number of cakes to bake
	BakeTime       time.Duration // time to bake a cake
	BakeStdDev     time.Duration // standard deviaton of baking time
	BakeBuf        int           // buffer slots between baking and icing
	NumIcers       int           // number of cooks icing
	IceTime        time.Duration // time to ice a cake
	IceStdDev      time.Duration // standard deviaton of icing time
	IceBuf         int           // buffer slots between icing and inscribing
	InscribeTime   time.Duration // time to inscribe a cake
	InscribeStdDev time.Duration // standard devaition of inscribing time
}

type cake int

func (s *Shop) baker(baked chan<- cake) {
	for i := 0; i < s.Cakes; i++ {
		c := cake(i)
		if s.Verbose {
			fmt.Println("baking", c)
		}
		work(s.BakeTime, s.BakeStdDev)
		baked <- c
	}
	close(baked)
}

func (s *Shop) icer(baked <-chan cake, iced chan<- cake) {
	for c := range baked {
		if s.Verbose {
			fmt.Println("icing", c)
		}
		work(s.IceTime, s.IceStdDev)
		iced <- c
	}
}

func (s *Shop) inscriber(iced <-chan cake) {
	for i := 0; i < s.Cakes; i++ {
		c := <-iced
		if s.Verbose {
			fmt.Println("inscribing", c)
		}
		work(s.InscribeTime, s.InscribeStdDev)
		if s.Verbose {
			fmt.Println("finished", c)
		}
	}
}

func work(t, stdDev time.Duration) {
	delay := t + time.Duration(rand.NormFloat64()*float64(stdDev))
	time.Sleep(delay)
}

func (s *Shop) Work(runs int) {
	for run := 0; run < runs; run++ {
		baked := make(chan cake, s.BakeBuf)
		iced := make(chan cake, s.IceBuf)
		go s.baker(baked)
		for i := 0; i < s.NumIcers; i++ {
			go s.icer(baked, iced)
		}
		s.inscriber(iced)
	}
}
