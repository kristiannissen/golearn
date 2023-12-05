# Pointers

Reference materiale https://google.github.io/styleguide/go/decisions#receiver-type
https://go.dev/tour/methods/8

> Correctness wins over speed or simplicity

...Ok!

## Hvornår skal du bruge pointer receiver
Der er ingen entydig logik, anbefalingen er, at for ikke at kopiere store data sæt, bruger du en pointer.

Hvis en metode kan modificere data som dens receiver peger på, skal du bruge en pointer.

```golang
var list []string
// l modificeres
func addItem(s string, l *[]string) {
	*l = append(*l, s)
}
```

Hvis en metode ikke modificerer data er der ingen grund til at bruge en pointer.

```
var list []string
// l modificeres ikke
func listLen(l []string) int {
	return len(l)
}
```
