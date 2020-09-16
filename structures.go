package main

type Test struct {
	On    bool
	Ammo  int
	Power int
}

func (ts *Test) Shoot() bool {
	if ts.On == false {
		return false
	}
	if ts.Ammo > 0 {
		ts.Ammo--
		return true
	} else {
		return false
	}

}
func (ts *Test) RideBike() bool {
	if ts.On == false {
		return false
	}
	if ts.Power > 0 {
		ts.Power--
		return true
	} else {
		return false
	}

}
