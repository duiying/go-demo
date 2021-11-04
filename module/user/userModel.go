package user

type User struct {
	ID    		int   	`json:"id"`
	Name  		string	`json:"name"`
	Email   	string	`json:"email"`
	Root 		int   	`json:"root"`
	Mtime    	string 	`json:"mtime"`
	Ctime    	string 	`json:"ctime"`
}