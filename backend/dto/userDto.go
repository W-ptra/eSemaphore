package dto

type UserProfile struct{
	Id			string	`json:"id"`
	Name		string	`json:"name"`
	Email		string	`json:"email"`
	ImageUrl	string	`json:"imageUrl"`
}

type UserPassword struct{
	Password	string	`json:"password"`
}