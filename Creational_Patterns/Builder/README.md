# Builder
The builder pattern is used to create instances of complex objects step by step. First, a struct is created with all of the possible fields that our object can have. Next, methods are created that always return the struct itself while implementing the fields.

# Applicability
when is need to create complex object step by step and the object creation process can be generalized.
The created object should be independent of the creation process. This means that the same construction process can create different representations of the object. The construction process must allow different representations for the object that is being created.

# Advantages
1- It separates the construction of a complex object from its representation so that the same construction process can create different representations.
2- It allows for better control over the construction process, making it possible to produce different types of products from the same set of construction instructions.
3- It makes it easier to change the internal representation of a product without affecting its external interface.
4- It makes it possible to reuse existing objects in the construction of new objects, thereby reducing the overall complexity of the system.
5- It allows for the construction of objects that are not possible to create using other creational patterns.

# Drawbacks
1- The builder design pattern can be difficult to implement correctly, and it can be easy to introduce errors into the resulting code.
2- The builder design pattern can make code difficult to read and understand.
3- The builder design pattern can make code difficult to maintain and change.

# Real World example
The Builder pattern is used in many places in the Go standard library. One example is the `bufio.Reader` type that allows reading from a buffer as if it was an io.Reader implementation. For example, when parsing a JSON document, we might want to first parse the document into a generic map[string]interface{} type, and then create a more specific type based on the contents of the map.



