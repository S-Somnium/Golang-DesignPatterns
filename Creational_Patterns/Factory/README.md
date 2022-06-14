# Simple Factory
We need to create an interface that is common to all objects of the factory. This interface will specify all methods that should be implemented. Then, we create a struct that implements the methods that are common to all objects, like getName(). Next, we create the struct of the object itself and implement any methods that are unique to each object. Finally, we create a function that will serve all of the objects that implemented the interface.

# Applicability
Encapsulate the logic for creating objects. This is especially useful in Go because it is a statically typed language. This means that the type of an object needs to be known at compile time. The Simple Factory pattern helps to abstract away this type information so that objects can be created without knowing their type ahead of time.

# Benefits
It improves the flexibility of the application and makes it easier to manage the object creation. It also makes the code more readable and maintainable.

# Advantages
 1-  You avoid tight coupling between the creator and the concrete products.
 2-  You can produce products without specifying their concrete classes.
 3-  The factory method decouples the client code from the concrete classes of the products.
 4-  The factory method allows you to change the concrete classes of the products without changing the client code.

# Drawbacks
1-  The client code must be aware of the types of products that the factory can create.
2-  The client code must specify the type of product that it wants to create.
3-  The client code is coupled to the concrete classes of the products.
4-  The client code is responsible for creating and managing the lifecycle of the products.

# Real World example
The Simple Factory pattern is used in many places in the Go standard library. One example is the bytes.NewBuffer function. This function creates a new byte buffer and returns it. The type of the byte buffer is not known until runtime, but the bytes.NewBuffer function knows how to create it.


