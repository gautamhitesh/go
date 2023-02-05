# Concepts

1. Type assertion
   - Usage: get the underlying concrete value of an interface
   - Example: typeassertion.go
2. Variable type check during runtime:
   - Usage: check the type of the interface
   - Example: runtime.go
3. Empty interfaces
   - Usage: handle the values of unknown type
   - Example: emptyinterface.go
4. Empty struct: to save memory, to show developer no value is stored
   - Usage: 
        1. While implementing a data set, if only keys need to be saved then each key can point to empty struct instead of setting dummy value
        2. In graph traversals to keep track of visited nodes
        3. to send a signal of an event via channel without data
   - Example: emptystruct.go
5. Slice copy: copy(to, from) for reference | = for values
6. Map copy: iterate for values | = for description
7. Error Handling:
    - How to: implement Error method
    - Example: errorHandling.go
8. Channels: ?? concurrent data access
9. Sorting:
    Example: customSort.go
10. Shadowing: variable overriding for a specific scope
11. Variadic Functions: function with variable number of arguments
    - Example: variadic.go
    - Usage: func vfun(arg, arge2,...type)type{}
12. Method overloading alternative
13. Inheritance alternative
14. Constants in go
15. 

