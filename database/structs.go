package database

type BlogPost struct {
	// ID     string `bson:"_id"`
	Date   string `bson:"date"`
	Title  string `bson:"title"`
	Body   string `bson:"body"`
	Author string `bson:"author"`
}
