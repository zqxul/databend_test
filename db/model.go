package db

import "time"

type T struct {
	Id1  int       `grom:"id1"`
	Id2  int       `grom:"id2"`
	Id3  int       `grom:"id3"`
	Id4  int       `grom:"id4"`
	Id5  int       `grom:"id5"`
	Id6  int       `grom:"id6"`
	Id7  int       `grom:"id7"`
	Id8  int       `grom:"id8"`
	Id9  int       `grom:"id9"`
	Id10 int       `grom:"id10"`
	Id11 int       `grom:"id11"`
	Id12 int       `grom:"id12"`
	Id13 int       `grom:"id13"`
	Id14 int       `grom:"id14"`
	Id15 int       `grom:"id15"`
	Id16 int       `grom:"id16"`
	Str1 string    `gorm:"str1"`
	Str2 string    `gorm:"str2"`
	Str3 string    `gorm:"str3"`
	Date time.Time `gorm:"dt"`
}

func (T) TableName() string {
	return "db.t"
}
