package models 


type PageData struct {
	Page 	string
	Title	string
	Message string
}

type Address struct {
	Address_ID	int
	Street_1	string
	Street_2	string
	City		string
	State		string
	Zip			string
}


type ContactInfo struct {
	Name	  string
	Email	  string
	Number    string
	Password  string
}


type Booking struct {
	Contact_Info	[]ContactInfo

}