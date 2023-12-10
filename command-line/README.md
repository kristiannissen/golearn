Flag er navngivne variabler der kan bruges når et program afvikles, f.eks. go run main.go -name=Kitty hvor -name er et flag, en navngiven variable vi kan bruge i programmet.

https://pkg.go.dev/flag@go1.21.5#pkg-overview

## Erklæring og definering af variabler
```golang
	var name string
	flag.StringVar(&name, "name", "Pussy", "The name")

	num := flag.Int("number", 1, "")
```
Når variablerne er erklæret og defineret kaldes
```golang
flag.Parse()
```
for at parse command-line ind i de definerede variabler.

## os.Args
Vi kan også bruge
```golang
os.Args
```
som giver os []string retur hvor alle kommandolinje værdierne er at finde i den indtastede rækkefølge.

**Eksempel**

go run main.go Hello Kitty
```golang
// Prints main, Hello, Kitty
fmt.Println(strings.Join(os.Args, ", "))
```
Bruges os.Args er programmets navn først værdi, i dette tilfælde "main".
