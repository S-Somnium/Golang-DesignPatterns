# Singleton
Create a var that will hold the only instace of the struct, start it on init() and then create a function to return this single instance.

# Applicability
This is useful in cases where we only need one instance of a class, such as a database connection or a configuration file.

# Advantages
1-	Simplicity: The Singleton design pattern is very simple and easy to implement. 
2-	Ease of Access: Since only one instance of a Singleton object exists, it is very easy to access. 
3-	Reduced Memory Footprint: The Singleton design pattern helps to reduce the memory footprint of an application because only one instance of a Singleton object exists.

# Drawbacks
1-They are difficult to unit test due to the static nature of the member variables.
2-They introduce global state into an application, which can be difficult to manage and can lead to unexpected side effects.
3-They can be difficult to extend, since they often rely on static members, which cannot be overridden.
4-They can lead to code that is tightly coupled and difficult to maintain.

# Real World example
The singleton pattern is used in many places in the Go standard library. One example is the `http.DefaultClient`. The `http.DefaultClient` is created when the `http` package is initialized. This client is used by many of the functions in the `http` package.

