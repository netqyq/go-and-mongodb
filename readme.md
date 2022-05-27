# Golang Using MongoDB Driver
In each of the tutorials, we made use of **bson.A, bson.D, and bson.M**, which represent **arrays**, **documents**, and **maps**. 

Working with a **bson.M** will leave you with a `map[string]interface{}` which isn't the most complicated in the world, but isn't necessarily the best for all scenarios in my opinion. Things get more challenging as you start working with **bson.D** as well.

The **omitempty** means that if there is no data in the particular field, when saved to MongoDB the field will not exist on the document rather than existing with an empty value.

