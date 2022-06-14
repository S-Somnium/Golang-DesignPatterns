# Abstract Factory
We need to create an interface that is common to all factories, then create a specific factory that implements that interface. And in this interface, we need to create a method to create products.
Basically it's a factory of factories.

# Applicability
Can be used in any situation where you need to create objects that belong to a certain family of objects. This might be objects that are all related to each other in some way, or that all serve a similar purpose.

# Benefits
Abstract factories can be used to provide a layer of abstraction between a client and a concrete implementation. This can allow for more flexibility when changing implementations.

# Advantages
1-Provides a way to encapsulate a group of individual factories that have a common theme.
2-Ensures that all products returned by the factory are compatible.
3-Produces products that are suitable for use together. 
4-Makes it easier to change the implementation of products. 
5-Makes it easier to extend the products.

# Drawbacks
1-Abstract factories can be difficult to maintain if the number of products and product variants starts to grow.
2-Abstract factories can be inflexible if the products they create need to be changed frequently.
3-Abstract factories can be complex to implement if the products they create are very different from each other.

# Real World example
The `database/sql` package uses the Abstract Factory pattern to abstract over the different database drivers. The `sql.DB` type is the abstract factory and the `driver.Driver` interface is the abstract product.

