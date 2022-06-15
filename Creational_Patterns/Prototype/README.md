# Prototype
The object's interface must have a clone method that returns a clone of itself.

# Applicability
The Prototype design pattern is only useful when we have complex objects to create and we need to copy them without doing all the steps one by one.

# Advantages
1- It provides a skeleton for the new object, and the new object can inherit the attributes and behavior of the prototype.
2- It is easy to change the attributes of the prototype without affecting the existing objects.
3- It is easy to create new objects by copying the prototype.
4- Easy to maintain: Prototype design pattern is easy to maintain and modify.

# Drawbacks
1- Cloning complex objects that have circular references might be very tricky

# Real World example

