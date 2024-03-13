# TMPLZ

Recursive, Non-intrusive templates.

## Rules

* A Template can be defined on its own or in a group with others.
  * Templates defined together can reference each other.
* Unused Variables are assumed to be literals.
* A Variable clause begins with a `Prefix`(e.g. `@`).
  * A Variable clause may terminate with a `Suffix`(e.g. `_`).
  * Additional `Prefix`es increase the "depth" of the Variable.
  * `Suffix`es within a Variable can be used to escape the "depth".
* Use `Prefix`es and `Suffix`es to create a Variable Graph. 

### Examples

All examples explicitly include the suffix even when it can be ignored. (Optional suffixes denoted `‗`)

0. Zero Variables:
   * `@` by itself is not a variable.
   * `@_` evaluates to `_`.
1. One Variable:
   * A standard template variable: `@solo‗ -> {@solo_}`.
2. Two Variables:
   * Two standard variables one after the other: `@one_@two‗ -> {@one_}{@two_}`.
     * Assigning one does not affect the other.
   * A variable with a sub-variable: `@one@two‗‗ -> {@one@two__{@two_}}`.
     * Neither `@one` nor `@one@two` can be assigned to.
     * You must first assign `@two` and that will set the actual name of `@one@two`.
       * e.g. ["two" -> "123"] then `@one@two__` becomes `@one123_`.
