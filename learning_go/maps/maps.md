- In Go maps can only accepts types which are comparable as keys.
  These can be strings, int, bool etc..,
- Passing map[Player]int, will throw an compile time error.
- Comparable interface defined cannot be implemented, it is type parameter constraint.

## Rules to Define a Custom Type as Map Parameter

- Comparable Fields
    - All fields in the custom defined struct should be comparable types.
- Cannot use non-comparable types
    - Custom structs cannot constraints slices, maps or function as fields if they have to be used as map keys.
- Pointers & Comparable Structs
  - You can use pointers to a structs as the map keys if the struct itself is comparable.
  - Enums or type based aliases like (type goo string) are fine as map keys, as they rely on the underlying types comparability.
