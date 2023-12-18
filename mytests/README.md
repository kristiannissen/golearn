# Go test, Examples & benchmarks

```go test -v``` afvikler alle mine tests, i dette tilfælde fejler MakeSound() for at vise hvordan det ser ud

```golang
func TestMakeSound(t *testing.T) {
	// Simpel test
	got := MakeSound()
	//
	if got != "Hello Kitty" {
		t.Errorf("MakeSound() == ´%s´; want What the fuck!", got)
	}
}
```

Følgende vises i terminalen

```
=== RUN   TestMakeSound
    mytests_test.go:28: MakeSound() == ´What the fuck!´; want What the fuck!
--- FAIL: TestMakeSound (0.00s)
=== RUN   TestFeed
--- PASS: TestFeed (0.00s)
=== RUN   TestFart
--- PASS: TestFart (0.00s)
FAIL
exit status 1
FAIL    github.com/kristiannissen/golearn/mytests       0.008s
```

## setup og teardown

Hvis du skal bruge setup og teardown kan du gøre det på følgende måde

```golang
func setup() {
	log.Println("Hello from setup()")
}

func teardown() {
	log.Println("Goodbye from teardown()")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
```

Du kan se beskederne fra setup() og teardown()

```
2023/12/18 09:26:57 Hello from setup()
--- FAIL: TestMakeSound (0.00s)
    mytests_test.go:29: MakeSound() == ´What the fuck!´; want What the fuck!
FAIL
coverage: 100.0% of statements
2023/12/18 09:26:57 Goodbye from teardown()
exit status 1
FAIL    github.com/kristiannissen/golearn/mytests       0.005s
```

## Test Coverage

```go test -cover``` fortæller mig at 87.5% af min kode er med i mine tests da jeg endnu ikke har skrevet en test til Fart() metoden

```
2023/12/18 08:57:13 Hello from setup()
--- FAIL: TestMakeSound (0.00s)
    mytests_test.go:28: MakeSound() == ´What the fuck!´; want What the fuck!
FAIL
coverage: 87.5% of statements
2023/12/18 08:57:13 Goodbye from teardown()
exit status 1
FAIL	github.com/kristiannissen/golearn/mytests	0.004s
```

## Test coverage rapport

Du kan få vist en coverage rapport med følgende kommando ```go test . -coverprofile=cover.out``` efterfulgt af ```go tool cover -html=cover.out``` som giver dig en HTML rapport

## Benmark

For at du kan benchmarke koden skal alle tests kunne gennemføres, jeg rettede i MakeSound() for at kunne se min benchmark

```go test -bench=.``` afvikler benchmarks

```golang
func BenchmarkFart(b *testing.B) {
        for i := 0; i < b.N; i++ {
                Fart()
        }
}
```

b.N er et er antallet af gange koden afvikles for at få et retvisende resultat. I terminalen kan du så se at koden blev afviklet 1000000000 gange for at nå et resultat

```
go test -bench=. -count=3
goos: linux
goarch: amd64
pkg: github.com/kristiannissen/golearn/mytests
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
BenchmarkFart-8         1000000000               0.2496 ns/op
BenchmarkFart-8         1000000000               0.2487 ns/op
BenchmarkFart-8         1000000000               0.2527 ns/op
PASS
ok      github.com/kristiannissen/golearn/mytests       0.844s
```

I eksemplet her krydrer jeg det med at fortælle go test at min benchmark skal afvikles 3 gange ```-count=3```

## Examples

En anden måde at skrive tests på er ved at bruge ExampleXXX. Her fanger ```go test``` alt der skrives til standard output ved f.eks. at anvende ```fmt.Println()``` og sammenligne det med en ```Output:``` strengen i ```ExampleXXX()``` funktionen

```golang
func ExampleFart() {
        fmt.Println(Fart())
        // Output: Fart all day, every day!
}
```

Du kan bruge _test.go filerne til dine ExampleXXX som vist herover. Når ```go test -v``` afvikles kaldes ```ExampleXXX()``` på samme måde som ```TestXXX(t *testing.T)``` og output strengen sammenlignes med hvad der skrives til standard output

```
=== RUN   TestMakeSound
--- PASS: TestMakeSound (0.00s)
=== RUN   TestFeed
--- PASS: TestFeed (0.00s)
=== RUN   TestFart
--- PASS: TestFart (0.00s)
=== RUN   ExampleFart
--- PASS: ExampleFart (0.00s)
PASS
ok      github.com/kristiannissen/golearn/mytests       0.004s
```

Hvis jeg ændrer i ```Output:``` strengen så testen fejler får jeg en fejl besked som den der vises her

```
=== RUN   TestMakeSound
--- PASS: TestMakeSound (0.00s)
=== RUN   TestFeed
--- PASS: TestFeed (0.00s)
=== RUN   TestFart
--- PASS: TestFart (0.00s)
=== RUN   ExampleFart
--- FAIL: ExampleFart (0.00s)
got:
Fart all day, every day!
want:
Fart all day, every day
FAIL
exit status 1
FAIL    github.com/kristiannissen/golearn/mytests       0.002s
```

Og jeg kan se at ExampleFart fejlede med besked som hvad go test fangede via ```fmt.Println()``` vist her i ```got:``` og værdien af min ```Output``` streng vises i ```want:```
