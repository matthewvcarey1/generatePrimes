# generatePrimes
Utility to generate list of prime numbers
By default it prints a comma seperated list of 1000 prime numbers on each line

The simplest way to run this is
```
./generatePrimes 10000 > output.txt
```
This will generate all the prime numbers less than 10000.

The possible parameters are ./generatePrimes number reverse|forward primes_per_line

Number can be in Octal (leading 0), Decimal or Hexadecimal (leading 0x).

reverse or forward specifies the sort order of the output.

primes_per_line is a positive number that specifies the number of primes on each line of input.

The reason for these possible options is that I am uncertain of the best way of rolling the data into the very slow roll-in program that fills the db.

To build this program you need to have golang installed at least version 18

In the base folder of the repository type to following
```
go build -o generatePrimes cmd/generatePrimes/main.go
```
If you are running windows you will need to type
```
go build -o generatePrimes.exe cmd/generatePrimes/main.go
```
You may need to run at least once to get it to compile
```
go mod tidy
```
On my 32GB Linux box the highest number that I can start with is 0x1000000000 or 2³⁶ or 68719476736 any larger and the program dies. It just uses a sieve of Eratosthenes that fills a bitarray with 1s or 0s depending if a number is prime. It walks the bitarray at the end to list just the prime numbers. It is reasonably quick.
