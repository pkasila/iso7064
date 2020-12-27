# iso7064
ISO7064 implementation to compute and verify check digits (+ IBAN is implemented too)

## Installation
To install `iso7064` you need to, as usual, run this command:
```sh
go get github.com/pkosilo/iso7064
```

## How to use?
First of all you need to create a new `Calculator` to compute and verify ISO7064 check digits.
`BaseCalculator` and `IBANCalculator` are interfaces of `Calculator`, they implement methods
like `Verify(input string) (bool, error)`, `Compute(input string) (string, error)` and
`ComputeChars(input string) (string, error)`. So, you just need to create their instances using
one of the methods below

### IBAN
To create IBAN compatible `Calculator` you need to do something like this:
```golang
calc := iso7064.NewIBANCalculator()
```
It will create an instance of `IBANCalculator` which will have `BaseCalculator` to compute
check digits according to ISO7064 MOD97-10, but also will automatically rearrange input string,
replace letters with digits and etc. **Remember**, that you need to add a placeholder for check
digits when you compute with `IBANCalculator`, it can be any digit-only placeholder (preferred one
is `00`)

### MOD11-2
To create ISO7064 MOD11-2 compatible `Calculator` you need to do:
```golang
calc := iso7064.NewMod112Calculator()
```
It will return `BaseCalculator` instance with `Modulus` equal to `11`, `Radix` equal to `2`,
`Charset` equal to `0123456789X` and `IsDouble` equal to `false`.

### MOD37-2
To create ISO7064 MOD37-2 compatible `Calculator` you need to do:
```golang
calc := iso7064.NewMod372Calculator()
```
It will return `BaseCalculator` instance with `Modulus` equal to `37`, `Radix` equal to `2`,
`Charset` equal to `0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ*` and `IsDouble` equal to `false`.

### MOD97-10
To create ISO7064 MOD97-10 compatible `Calculator` you need to do:
```golang
calc := iso7064.NewMod9710Calculator()
```
It will return `BaseCalculator` instance with `Modulus` equal to `97`, `Radix` equal to `10`,
`Charset` equal to `0123456789` and `IsDouble` equal to `true`.

### MOD661-26
To create ISO7064 MOD661-26 compatible `Calculator` you need to do:
```golang
calc := iso7064.NewMod66126Calculator()
```
It will return `BaseCalculator` instance with `Modulus` equal to `661`, `Radix` equal to `26`,
`Charset` equal to `ABCDEFGHIJKLMNOPQRSTUVWXYZ` and `IsDouble` equal to `true`.

### MOD1271-36
To create ISO7064 MOD1271-36 compatible `Calculator` you need to do:
```golang
calc := iso7064.NewMod127136Calculator()
```
It will return `BaseCalculator` instance with `Modulus` equal to `1271`, `Radix` equal to `36`,
`Charset` equal to `0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ` and `IsDouble` equal to `true`.

### BaseCalculator
As you may spot in examples before, functions `NewMod****Calculator` always return `BaseCalculator`
with different `Modulus`, `Radix`, `Charset` and `IsDouble`. So, it means that `BaseController` can
take any `Modulus`, `Radix`, `Charset` and `IsDouble`, and you can create your own custom `BaseCalculator`:
```golang
calc := &iso7064.BaseCalculator {
  Modulus:  97,
  Radix:    10,
  Charset:  "0123456789",
  IsDouble: true,
}
```

### Compute and verify
#### Verify(input string) (bool, error)
To verify existing string that it has correct check digits, you need to do:
```golang
input := "1234123435"

calc := iso7064.NewMod9710Calculator()
correct, err := calc.Verify(input)
fmt.Println(correct, err) // Will return `true <nil>`
```
In case of an error it will return `false` and error itself. Error may happen if
the `input` is too short (shorter or equal to check digits length), the `input` has
`InvalidRune`, which is not present in the `Charset`.

#### Compute(input string) (string, error)
This function will return `input` string the check digits added to it. If it is `BaseCalculator`,
check digits will be located at the end of the string. It it is `IBANCalculator`, they will be
located at position 2-3 (couting from zero, remember) in the output string. **Remember**, you
need to add a placeholder for check digits if you use `IBANCalculator` (**again** preferred one
is `00`).
```golang
input := "12341234"

calc := iso7064.NewMod9710Calculator()
computed, err := calc.Compute(input)
fmt.Println(computed, err) // Will return `1234123435 <nil>`
```

#### ComputeChars(input string) (string, error)
This function will return `input` string the check digits added to it. If it is `BaseCalculator`,
check digits will be located at the end of the string. It it is `IBANCalculator`, they will be
located at position 2-3 (couting from zero, remember) in the output string. **Remember**, you
need to add a placeholder for check digits if you use `IBANCalculator` (**again** preferred one
is `00`).
```golang
input := "12341234"

calc := iso7064.NewMod9710Calculator()
computedChars, err := calc.ComputeChars(input)
fmt.Println(computedChars, err) // Will return `35 <nil>`
```
