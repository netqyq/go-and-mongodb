https://www.mongodb.com/blog/post/quick-start-golang--mongodb--a-quick-look-at-gridfs

This post is a continuity of getting started with Go and MongoDB series using the official Go driver. In this tutorial, we’re going to look at how to use the MongoDB GridFS feature to store files which are larger than 16 MB. Some use case of GridFS are:

- If your filesystem limits the number of files in a directory, you can use GridFS to store as many files as needed.
- When you want to access information from portions of large files without having to load files into memory, you can use GridFS to recall sections of files without reading the entire file into memory.
- hen you want to keep your files and metadata automatically synced and deployed across a number of systems and facilities, you can use GridFS. When using geographically distributed replica sets, MongoDB can distribute files and their metadata automatically to a number of mongod instances and facilities.