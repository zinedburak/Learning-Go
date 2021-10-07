 # Learning-Go
This is a project that is based on the teachings of the learning go class from linekdin e-learning

# Essential Syntax Rules
- Go is case sensitive
- Variable and package  names are lowercase and mixed case
- Names of public fields have initial upperCase characters (Meaning that if a field starts with upper case it is equivalent of public variable_name in java )
- Initial uppercase character means the symbol is exported
- Line feed ends a statement; no semicolon required
- Semicolons are part of the formal language spec
- The lexer (Software component that parses your code and analyzes your code ) adds the semicolons as needed
- If a statement determined to be complete and the lexer encounters a line feed that means that it is end of the statement.
- Numeric types don't implicitly convert. Example - you can not add float and integer


# Memory Management
- The Go runtime is statically linked into application
- Memory is allocated and deallocated automatically
- Use make() or new() to initialize maps, slices, and channels
- The new() function
    - Allocates but does not initialize memory
    - Results in zeroed storage but returns a memory address
    - (Meaning that it will create an map object but you can not add any element in it)
- The make() function
    - Allocates and initializes memory
    - Allocates non-zeroed storage and returns a memory address
- Memory is deallocated by garbage collector
- Objects out of scope or set to nil are eligible
- GC was rebuilt in Go 1.5 for very low latency