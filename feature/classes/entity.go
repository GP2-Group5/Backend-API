package classes

type ClassCore struct {
	ID           uint
	Name         string `validate:"required"`
	User_id      uint   `validate:"required"`
	Users        string
	Start_date   string `validate:"required"`
	Graduated_at string `validate:"required"`
}

type Users struct {
	ID        int
	Full_Name string
}

type IClassRepository interface {
	GetAll() (data []ClassCore, err error)
	Create(input ClassCore) (row int, err error)
	GetByID(id int) (data ClassCore, err error)
	Delete(data ClassCore, id int) (row int, err error)
	Update(data ClassCore, id int) (row int, err error)
}

type IClassService interface {
	GetAll() (data []ClassCore, err error)
	Create(input ClassCore) (err error)
	GetByID(id int) (data ClassCore, err error)
	Delete(data ClassCore, id int) (err error)
	Update(data ClassCore, id int) (err error)
}
