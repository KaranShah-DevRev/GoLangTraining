package constants

func CreateMongoUrl(mongoUser string, mongoPass string) string {
	mongoUrl := "mongodb+srv://" + mongoUser + ":" + mongoPass + "@" + "crudcluster.g3ktvdf.mongodb.net/?retryWrites=true&w=majority"
	return mongoUrl
}
