package model

type Rating struct {
	Base

	OrganisationID string `gorm:"not null;unique"`
	A              uint64 `gorm:"column:a_star;default:0;"`
	B              uint64 `gorm:"column:b_star;default:0;"`
	C              uint64 `gorm:"column:c_star;default:0;"`
	D              uint64 `gorm:"column:d_star;default:0;"`
	E              uint64 `gorm:"column:e_star;default:0;"`
}

func (r Rating) Summation() uint64 {
	return r.A + r.B + r.C + r.D + r.E
}

func (r Rating) AverageRating() float64 {
	return float64(((1 * r.A) + (2 * r.B) + (3 * r.C) + (4 * r.D) + (5 * r.E)) / r.Summation())
}
