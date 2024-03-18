package main

/*
1. Using lowercase to define Constant as other usual variables
2. Grouping variables as possible as you can
3. There is nothing wrong with Panic functions. Remember to prefix function name with "Must" e.g., MustCreateConnection…
4. Specifying field name when initialising struct
5. Grouping Mutex with protected field
   type Server struct {
      ID string

      Mu	sync.RWMutex
      peers map[string]string
   }
6. Interface defines behaviours. Using ...er as interface name and preferring small to complex interface.
7. Putting important & external variables and functions at the top of the file for highlighting
8. Prefixing http handler names with "handler"
*/
