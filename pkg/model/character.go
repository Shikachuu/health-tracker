package model

type Character struct {
	TemporaryHealth int
	Health          int
	HPLog           []int
}

func (c *Character) GetCurrentHealth() int {
	if len(c.HPLog) == 0 {
		return c.Health
	} else {
		return c.HPLog[len(c.HPLog)-1]
	}
}

func (c *Character) Damage(d int) int {
	if len(c.HPLog) == 0 {
		c.HPLog = append(c.HPLog, c.Health-d)
	} else {
		c.HPLog = append(c.HPLog, c.HPLog[len(c.HPLog)-1]-d)
	}

	return c.HPLog[len(c.HPLog)-1]
}

func (c *Character) Heal(d int) int {
	if len(c.HPLog) == 0 {
		c.HPLog = append(c.HPLog, c.Health+d)
	} else {
		c.HPLog = append(c.HPLog, c.HPLog[len(c.HPLog)-1]+d)
	}

	return c.HPLog[len(c.HPLog)-1]
}
