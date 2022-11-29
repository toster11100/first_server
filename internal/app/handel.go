package handel

type Handel struct {
	db int
}

func New(db int) Handel {
	return Handel{
		db,
	}
}
