package main

type Calcu struct {
	Num int
}

func (c Calcu) A() {
	c.Num++
}

func (c *Calcu) B() {
	c.Num++
}
