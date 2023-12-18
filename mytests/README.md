```go test``` afvikler alle mine tests, i dette tilfælde fejler MakeSound() for at vise hvordan det ser ud

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
--- FAIL: TestMakeSound (0.00s)
    mytests_test.go:28: MakeSound() == ´What the fuck!´; want What the fuck!
FAIL
exit status 1
FAIL    github.com/kristiannissen/golearn/mytests       0.002s
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
