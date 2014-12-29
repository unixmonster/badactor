package badactor

import (
	"strconv"
	"testing"
	"time"
)

func BenchmarkIsJailed(b *testing.B) {
	var err error
	d := NewDirector()
	d.Run()
	an := "an_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	rn := "rn_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	rm := "rm_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	r := &Rule{
		Name:        rn,
		Message:     rm,
		StrikeLimit: 3,
		ExpireBase:  time.Second * 60,
		Sentence:    time.Minute * 5,
	}

	err = d.AddRule(r)
	if err != nil {
		b.Errorf("AddRule [%s] should not fail", rn)
	}

	for i := 0; i < 4; i++ {
		d.Infraction(an, rn)
	}

	for i := 0; i < b.N; i++ {
		d.IsJailed(an)
	}
}

func BenchmarkIsJailedFor(b *testing.B) {
	var err error
	d := NewDirector()
	d.Run()
	an := "an_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	rn := "rn_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	rm := "rm_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	r := &Rule{
		Name:        rn,
		Message:     rm,
		StrikeLimit: 3,
		ExpireBase:  time.Second * 60,
		Sentence:    time.Minute * 5,
	}

	err = d.AddRule(r)
	if err != nil {
		b.Errorf("AddRule [%s] should not fail", rn)
	}

	for i := 0; i < 4; i++ {
		d.Infraction(an, rn)
	}

	for i := 0; i < b.N; i++ {
		d.IsJailedFor(an, rn)
	}
}

func BenchmarkIsJailedInfraction(b *testing.B) {
	var err error
	d := NewDirector()
	d.Run()
	rn := "rn_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	rm := "rm_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	r := &Rule{
		Name:        rn,
		Message:     rm,
		StrikeLimit: 3,
		ExpireBase:  time.Second * 60,
		Sentence:    time.Minute * 5,
	}

	err = d.AddRule(r)
	if err != nil {
		b.Errorf("AddRule [%s] should not fail", rn)
	}

	for i := 0; i < b.N; i++ {
		an := "an_" + strconv.FormatInt(time.Now().UnixNano(), 10)
		for i := 0; i < 3; i++ {
			d.Infraction(an, rn)
		}
		d.IsJailed(an)
	}
}

func BenchmarkIsJailedForInfraction(b *testing.B) {
	var err error
	d := NewDirector()
	d.Run()
	rn := "rn_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	rm := "rm_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	r := &Rule{
		Name:        rn,
		Message:     rm,
		StrikeLimit: 3,
		ExpireBase:  time.Second * 60,
		Sentence:    time.Minute * 5,
	}

	err = d.AddRule(r)
	if err != nil {
		b.Errorf("AddRule [%s] should not fail", rn)
	}

	for i := 0; i < b.N; i++ {
		an := "an_" + strconv.FormatInt(time.Now().UnixNano(), 10)
		for i := 0; i < 3; i++ {
			d.Infraction(an, rn)
		}
		d.IsJailedFor(an, rn)
	}
}

func BenchmarkInfraction(b *testing.B) {
	var err error
	d := NewDirector()
	d.Run()
	an := "an_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	rn := "rn_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	rm := "rm_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	r := &Rule{
		Name:        rn,
		Message:     rm,
		StrikeLimit: 3,
		ExpireBase:  time.Second * 60,
		Sentence:    time.Minute * 5,
	}

	err = d.AddRule(r)
	if err != nil {
		b.Errorf("AddRule for [%v] should not fail", rn)
	}

	for i := 0; i < b.N; i++ {
		d.Infraction(an, rn)
	}
}

func BenchmarkInfractionMostCostly(b *testing.B) {
	var err error
	d := NewDirector()
	d.Run()
	an := "an_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	rn := "rn_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	rm := "rm_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	r := &Rule{
		Name:        rn,
		Message:     rm,
		StrikeLimit: 3,
		ExpireBase:  time.Second * 60,
		Sentence:    time.Minute * 5,
	}

	err = d.AddRule(r)
	if err != nil {
		b.Errorf("AddRule for Actor [%s] should not fail", an)
	}

	// create initial infraction
	d.Infraction(an, rn)

	// bench the Least Costly way
	for i := 0; i < b.N; i++ {
		d.MostCostlyInfraction(an, rn)
	}
}

func Benchmark10000Actors(b *testing.B) {
	var err error
	d := NewDirector()
	d.Run()
	rn := "rn_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	rm := "rm_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	r := &Rule{
		Name:        rn,
		Message:     rm,
		StrikeLimit: 3,
		ExpireBase:  time.Second * 60,
		Sentence:    time.Minute * 5,
	}

	err = d.AddRule(r)
	if err != nil {
		b.Errorf("AddRule for [%v] should not fail", rn)
	}

	aN := 10000

	for i := 0; i < b.N; i++ {
		for a := 0; a < aN; a++ {
			an := "an_" + strconv.FormatInt(time.Now().UnixNano(), 10)
			d.Infraction(an, rn)
		}
	}
}

func Benchmark100000Actors(b *testing.B) {
	var err error
	d := NewDirector()
	d.Run()
	rn := "rn_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	rm := "rm_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	r := &Rule{
		Name:        rn,
		Message:     rm,
		StrikeLimit: 3,
		ExpireBase:  time.Second * 60,
		Sentence:    time.Minute * 5,
	}

	err = d.AddRule(r)
	if err != nil {
		b.Errorf("AddRule for [%v] should not fail", rn)
	}

	aN := 100000

	for i := 0; i < b.N; i++ {
		for a := 0; a < aN; a++ {
			an := "an_" + strconv.FormatInt(time.Now().UnixNano(), 10)
			d.Infraction(an, rn)
		}
	}
}

func Benchmark1000000Actors(b *testing.B) {
	var err error
	d := NewDirector()
	d.Run()
	rn := "rn_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	rm := "rm_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	r := &Rule{
		Name:        rn,
		Message:     rm,
		StrikeLimit: 3,
		ExpireBase:  time.Second * 60,
		Sentence:    time.Minute * 5,
	}

	err = d.AddRule(r)
	if err != nil {
		b.Errorf("AddRule for [%v] should not fail", rn)
	}

	aN := 1000000

	for i := 0; i < b.N; i++ {
		for a := 0; a < aN; a++ {
			an := "an_" + strconv.FormatInt(time.Now().UnixNano(), 10)
			d.Infraction(an, rn)
		}
	}
}

func Benchmark10000Actors4Infractions(b *testing.B) {
	var err error
	d := NewDirector()
	d.Run()
	rn := "rn_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	rm := "rm_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	r := &Rule{
		Name:        rn,
		Message:     rm,
		StrikeLimit: 3,
		ExpireBase:  time.Second * 60,
		Sentence:    time.Minute * 5,
	}

	err = d.AddRule(r)
	if err != nil {
		b.Errorf("AddRule for [%v] should not fail", rn)
	}

	aN := 10000

	for i := 0; i < b.N; i++ {
		for a := 0; a < aN; a++ {
			an := "an_" + strconv.FormatInt(time.Now().UnixNano(), 10)
			for inf := 0; inf < 4; inf++ {
				d.Infraction(an, rn)
			}
		}
	}
}
