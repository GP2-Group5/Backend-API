package classes

type ClassCore struct {
	ID           uint
	Name         string `validate:"required"`
	User_id      uint   `validate:"required"`
	Users        Users
	Start_date   string `validate:"required"`
	Graduated_at string `validate:"required"`
}

type Users struct {
	ID        int
	Full_Name string
}

type IClassRepository interface {
	Create(input ClassCore) (row int, err error)
}

type IClassService interface {
	Create(input ClassCore) (err error)
}
