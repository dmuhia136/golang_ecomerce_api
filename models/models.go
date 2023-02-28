package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type User struct {
	ID				primitive.ObjectID
	First_Name		*string
	Last_Name		*string
	Password		*string	
	Email			*string
	Phone			*string
	Token			*string
	Refresh_Token	*string
	Created_At		time.Time
	Updated_At		time.Time
	User_ID			*string
	UserCart		[]ProductUser
	Address_Details	[]Address
	Order_Status	[]Order
}

type Product struct {
	Product_ID 		primitive.ObjectID
	Product_Name	*string
	Price			*unint64
	Rating			*unint8
	Image			*string
}

type ProductUser struct {
	Product_ID		primitive.ObjectID
	Product_Name	*string
	Price			*unint64
	RatingImage		*unint8
}

type Address struct {
	Address_ID		primitive.ObjectID
	House			*string
	Street			*string
	City			*string
	Pincode			*string
}
type Order struct {
	Order_ID		primitive.ObjectID
	Order_Cart		*string
	Ordered_At		time.Time
	Price			*unint64
	Discount		*unint64
	Payment_Method	
}

type Payment struct {
	Digital bool
	COD     bool
}
