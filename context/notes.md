# context chapter notes

I think i slowly start to understand what context is how and why to use it.
Standard library http and httptest are very powerful and it would be beneficial to learn how they work.
Go has great support on concurrent and distributed programming from standard library,
the tools are simple and cohesive but there composition allows building complex systems! (I guess I'm a fanboy now)

```
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}
```
This code snippet looks no cool i will print it and hand on wall.

OK after whole chapter it seem context may be a bit more complicated 😅
