# Pointers and Errors Chapter

- in Go you don't need to dereference pointer to struct to call its methods or read field values

- t.Helper is useful when you want to get good error messages

- t.Fatal is good when there is no sense in continuing a test

- it can be error if it implements "Error() string" method
