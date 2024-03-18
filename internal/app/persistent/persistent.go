package persistent

/*
# Method sets (Rer: https://go.dev/ref/spec#Method_sets)
The method set of a type determines the methods that can be called on an operand of that type. Every type has a (possibly empty) method set associated with it:

- The method set of a defined type T consists of all methods declared with receiver type T.
- The method set of a pointer to a defined type T (where T is neither a pointer nor an interface) is the set of all methods declared with receiver *T or T.
- The method set of an interface type is the intersection of the method sets of each type in the interface's type set (the resulting method set is usually just the set of declared methods in the interface).

......

# Explanation (Ref: https://stackoverflow.com/questions/33587227/method-sets-pointer-vs-value-receiver/33591156#33591156);
- If you have a *T you can call methods that have a receiver type of *T as well as methods that have a receiver type of T.
- If you have a T and it is addressable you can call methods that have a receiver type of *T as well as methods that have a receiver type of T, because the method call t.Meth() will be equivalent to (&t).Meth().
- If you have a T and it isn't addressable (for instance, the result of a function call, or the result of indexing into a map), Go can't get a pointer to it, so you can only call methods that have a receiver type of T, not *T.
- If you have an interface I, and some or all of the methods in I's method set are provided by methods with a receiver of *T (with the remainder being provided by methods with a receiver of T),
  then *T satisfies the interface I, but T doesn't. That is because *T's method set includes T's, but not the other way around (back to the first point again).
```S
Why do T and *T have different method sets? (Ref: https://go.dev/doc/faq#different_method_sets)
This distinction arises because if an interface value contains a pointer *T, a method call can obtain a value by dereferencing the pointer, but if an interface
value contains a value T, there is no safe way for a method call to obtain a pointer. (Doing so would allow a method to modify the contents of the value inside
the interface, which is not permitted by the language specification.)
```


In short, you can mix and match methods with value receivers and methods with pointer receivers, and use them with variables containing values and pointers, without worrying about which is which. Both will work, and the syntax is the same.
However, if methods with pointer receivers are needed to satisfy an interface, then only a pointer will be assignable to the interface — a value won't be valid.
*/

type NoSqlOperator interface {
	FetchAll()
	FetchById(id string)
	UpdateById(id string)
	DeleteById(id string)
}

type CacheOperator interface {
	Connect(uri string)
}
