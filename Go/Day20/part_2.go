package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run() {
	m := make(map[string]struct {
		t string
		d []string
	})
	f := make(map[string]string)
	c := make(map[string]map[string]string)
	l := make(map[string]int) // Initialize map l
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	var r string

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, "->", "")
		splitLine := strings.Fields(line)

		s, d := splitLine[0], splitLine[1:]
		var t string
		if s != "broadcaster" {
			t = string(s[0])
			s = s[1:]
		}

		m[s] = struct {
			t string
			d []string
		}{t, d}
		f[s] = "off"

		if _, exists := c[s]; !exists {
			c[s] = make(map[string]string)
		}

		for _, o := range d {
			c[o][s] = "low"
		}

		if len(d) == 1 && d[0] == "rx" {
			r = s
		}
	}

	b := 0

	for k := range c[r] {
		l[k] = 0
	}

	for {
		b++
		q := []struct{ i, n, p string }{{"button", "broadcaster", "low"}}

		for len(q) > 0 {
			item := q[0]
			q = q[1:]

			i, n, p := item.i, item.n, item.p
			if _, exists := m[n]; !exists {
				continue
			}

			t, d := m[n].t, m[n].d

			if t == "" {
				for _, o := range d {
					q = append(q, struct{ i, n, p string }{n, o, p})
				}
			} else if t == "%" {
				if p == "low" {
					s := f[n]
					if s == "off" {
						f[n] = "on"
						for _, o := range d {
							q = append(q, struct{ i, n, p string }{n, o, "high"})
						}
					} else {
						f[n] = "off"
						for _, o := range d {
							q = append(q, struct{ i, n, p string }{n, o, "low"})
						}
					}
				}
			} else if t == "&" {
				c[n][i] = p
				allHigh := true
				for _, v := range c[n] {
					if v != "high" {
						allHigh = false
						break
					}
				}
				if allHigh {
					for _, o := range d {
						q = append(q, struct{ i, n, p string }{n, o, "low"})
					}
				} else {
					for _, o := range d {
						q = append(q, struct{ i, n, p string }{n, o, "high"})
					}
				}
				if n == r {
					for k, v := range c[n] {
						if v == "high" && l[k] == 0 {
							l[k] = b
						}
					}
				}
			}
		}

		allPositive := true
		for _, v := range l {
			if v <= 0 {
				allPositive = false
				break
			}
		}

		if allPositive {
			product := 1
			for _, v := range l {
				product *= v
			}
			fmt.Println(product)
			break
		}
	}
}
