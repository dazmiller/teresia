package models

import (
	"time"
)

type Articles struct {
	Id      int       `form:"-"`
	Title   string    `form:"title" required`
	Content string    `orm:";type(text)" form:"content,textarea"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

type Profile struct {
	Id         int    `orm:"column(id);auto"`
	UserId     int    `orm:"column(user_id);null"`
	Firstname  string `orm:"column(firstname);size(45);null"`
	Middlename string `orm:"column(middlename);size(45);null"`
	/*
		Surname    string    `orm:"column(surname);size(45);null"`
		Sex        string    `orm:"column(sex);size(45);null"`
		Birthdate  time.Time `orm:"column(birthdate);type(date);null"`
		Height     int       `orm:"column(height);null"`
		Weight     int       `orm:"column(weight);null"`
		Smoker     uint64    `orm:"column(smoker);size(1);null"`
		Avdrinks   int       `orm:"column(avdrinks);null"`
		Address1   string    `orm:"column(address1);size(90);null"`
		Address2   string    `orm:"column(address2);size(90);null"`
		Address3   string    `orm:"column(address3);size(90);null"`
		Suburb     string    `orm:"column(suburb);size(45);null"`
		State      string    `orm:"column(state);size(10);null"`
		Postcode   int       `orm:"column(postcode);null"`
		Email1     string    `orm:"column(email1);size(90);null"`
		Email2     string    `orm:"column(email2);size(90);null"`
		PhoneHome  string    `orm:"column(phone_home);size(45);null"`
		PhoneMob   string    `orm:"column(phone_mob);size(45);null"`
		Msn        string    `orm:"column(msn);size(90);null"`
		Twitter    string    `orm:"column(twitter);size(90);null"`
		Facebook   string    `orm:"column(facebook);size(90);null"`
		Wechat     string    `orm:"column(wechat);size(90);null"`
		Profession string    `orm:"column(profession);size(90);null"`
		Profilecol string    `orm:"column(profilecol);size(45);null"`
		Title      string    `orm:"column(title);size(45);null"`
		Child1Name       string    `orm:"column(child1_name);size(45);null"`
		Child1Birthdate  time.Time `orm:"column(child1_birthdate);type(date);null"`
		Child1Sex        string    `orm:"column(child1_sex);size(10);null"`
		Child2Name       string    `orm:"column(child2_name);size(45);null"`
		Child2Birthdate  time.Time `orm:"column(child2_birthdate);type(date);null"`
		Child2Sex        string    `orm:"column(child2_sex);size(10);null"`
		Child3Name       string    `orm:"column(child3_name);size(45);null"`
		Child3Birthdate  time.Time `orm:"column(child3_birthdate);type(date);null"`
		Child3Sex        string    `orm:"column(child3_sex);size(10);null"`
		Child4Name       string    `orm:"column(child4_name);size(45);null"`
		Child4Birthdatae time.Time `orm:"column(child4_birthdatae);type(date);null"`
		Child4Sex        string    `orm:"column(child4_sex);size(10);null"`
		Child5Name       string    `orm:"column(child5_name);size(45);null"`
		Child5Birthdate  time.Time `orm:"column(child5_birthdate);type(date);null"`
		Child5Sex        string    `orm:"column(child5_sex);size(10);null"`
		NoChildren       int       `orm:"column(no_children);null"`
		NoPets           int       `orm:"column(no_pets);null"`
		Pet1Name         string    `orm:"column(pet1_name);size(45);null"`
		Pet1Type         string    `orm:"column(pet1_type);size(45);null"`
		Pet1Sex          string    `orm:"column(pet1_sex);size(10);null"`
		Pet2Name         string    `orm:"column(pet2_name);size(45);null"`
		Pet2Type         string    `orm:"column(pet2_type);size(45);null"`
		Pet2Sex          string    `orm:"column(pet2_sex);size(10);null"`
		Pet3Name         string    `orm:"column(pet3_name);size(45);null"`
		Pet3Sex          string    `orm:"column(pet3_sex);size(10);null"`
		Pet3Type         string    `orm:"column(pet3_type);size(45);null"`
		LastUpdated      time.Time `orm:"column(last_updated);type(datetime);null"`
	*/
	OptionsRadios string `orm:"column(options_radios);size(191);null"`
}

func (a *Articles) TableName() string {
	return "articles"
}

func (a *Profile) TableName() string {
	return "profiles"
}
