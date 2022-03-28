package main

import (
	"crypto/md5"
	"fmt"
	"strings"
	"sync"
	"time"
)

// Generator contains the parameters required to generate domains. Use New() or
// NewSeeded() to initialize.
type Generator struct {
	year, month, day, seed, i int
	tld                       string
	lock                      *sync.Mutex
}

// New initializes a new Generator and returns it.
// Year, month, and day must all be in YYYY, MM, DD format (respectively).
// Note: There is no input validation here.
func New(year, month, day int, tld string) *Generator {
	return NewSeeded(year, month, day, 0, tld)
}

// NewSeeded initializes a new Generator with a seed and returns it. See New()
// for parameter descriptions.
func NewSeeded(year, month, day, seed int, tld string) *Generator {
	if !strings.HasPrefix(tld, ".") {
		tld = "." + tld
	}

	return &Generator{
		year:  year,
		month: month,
		day:   day,
		tld:   tld,
		seed:  seed,
		lock:  new(sync.Mutex),
	}
}

// Next returns the generated domain as a string and increments the iterator
// MD5 is used to hash the generated string before adding the TLD and returning.
func (g *Generator) Next() string {
	g.lock.Lock()
	defer g.lock.Unlock()

	g.i--

	return fmt.Sprintf("%x%s", md5.Sum([]byte(
		fmt.Sprintf("%v%v%v%v%v", g.year, g.month, g.day, g.seed, g.i),
	)), g.tld)
}

func main() {
	t := time.Now()
	gen := New(t.Year(), int(t.Month()), t.Day(), "com")

	/* Print out 5 domains */
	for i := 0; i < 10; i++ {
		fmt.Println(gen.Next())
	}
}
