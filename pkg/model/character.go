package model

type Character struct {
	Health          int
	TemporaryHealth int
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
	c.TemporaryHealth -= d

	if c.TemporaryHealth < 0 {
		newHP := c.GetCurrentHealth() + c.TemporaryHealth

		if newHP < 0 {
			c.HPLog = append(c.HPLog, 0)
		} else {
			c.HPLog = append(c.HPLog, newHP)
		}

		c.TemporaryHealth = 0
	}

	return c.GetCurrentHealth()
}

func (c *Character) Heal(h int) int {
	newHP := c.GetCurrentHealth() + h

	if newHP > c.Health {
		c.HPLog = append(c.HPLog, c.Health)
	} else {
		c.HPLog = append(c.HPLog, c.GetCurrentHealth()+h)
	}

	return c.GetCurrentHealth()
}

func (c *Character) TemporaryHeal(thp int) int {
	c.TemporaryHealth += thp

	return c.TemporaryHealth
}
