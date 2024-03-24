### Getting Started 
* Create a main.go file 
* Do `go mod init github.com/krahul2024/some-app` -> but this is only for the projects yk, for simple cases, just `go mod init app-name` is enough for local stuff, this creates a go.mod file

```go
import "fmt"
func main() {
    fmt.Println("This is some string")
}
```
* The important thing in here is all the package names start with Capital letters, also we can't use single quoutes inside the Println. 
* To run the file use `go run filename.go or main.go`, this just compiles and runs the file and not build it. 
* Look at the documentation with `go --help` for various commands. 
* Look at various path commands like gopath etc, also build and compile commands. 
* Learn some things about the lexer and where not to place the semi-colons and where to put the semi-colons. 
* lexer is to check the syntax of the code for a given language. 

### Types 
```go
    fmt.Printf("Variable is of type : %T\n", username)
```
* This prints the type of the username variable.
* Integer types : uint8, uint16, uint32, uint64, int8, int16, int32, int64, uint(either 32 or 64 bytes), same goes for the int. 
* Float types : float32 or float64
* complex types : complex64, complex128
```go
var some_val int
fmt.Println(some_val)
// This prints 0 as default value, no garbage stuff. 

// Undeclared types
var some_string = "this is some string"

//  without the var mention
some_variable := 3200

// some more variable declaration ways
var (
		name    string
		age     string
		address int
	)

var a, b int = 0, 342 // this requires all the variables have same type
x, y, z := 4, 6, "this"

```
* using the `:=` syntax we can only declare variables inside a method only, we can't declare global variables using this way, rather we have to resort to the var key declaration. 

* Global variable declarations 
```go 
const Variable = "This is some constant" // this is available to all the files since it starts with capital letter, it is equivalent to Public keyword in other langs and can be use anywhere, basically exporting the variable. 
const variable = "this is another constant" // this is available only to the current file 

// Variable is exported and visible outside the package
const Variable = "This is some constant"

// variable is not exported and visible only within the package
const variable = "this is another constant"

var ExportedVar string

func init() {
    ExportedVar = "This is exported via init function"
}

```
**List of types(not all)**
1. %d: Integer
2. %f: Float (floating point number)
3. %s: String
4. %c: Character (rune)
5. %t: Boolean
5. %v: Value (can parse any type)
7. %b: Base 2 (binary)
8. %o: Base 8 (octal)
9. %x: Base 16 (hexadecimal)
10. %X: Base 16 (hexadecimal), uppercase
11. %p: Pointer


### I/O
1. Using the fmt.Scan/Scanln/Scanf
```go
// to take the input
var name string 
fmt.Scan(&name) 
fmt.Scanln(&name)
fmt.Scanf("%s", &name) 

// to print
fmt.Println()
fmt.Printf()
fmt.Fprintf()
// usage of the Fprintf
file, err := os.Create("output.txt") // here when doing os.Create, we have to pass two variables since it returns two variables. 
	if err != nil {
		fmt.Println("There was an error!")
	}
	defer file.Close()

	fmt.Fprintf(file, "This is some text, %s", "from somewhere!")

// this can also be followed 
file, _ := os.Create("output.txt")
fmt.Fprintf(file, "This is the content of the file : %d", 6) 

```

2. File I/O (os package):
 Opening and reading from files using `os.Open`, `os.OpenFile`, `os.ReadFile`, or `os.ReadAll`.
Writing to files using `os.Create`, `os.OpenFile` (with appropriate flags), `os.WriteFile`, or `File.Write`.

3. Reading/Writing Specific File Formats:
For JSON files: Use encoding/json package to encode/decode JSON data.
For CSV files: Use packages like encoding/csv for parsing and generating CSV files.
For XML files: Use encoding/xml package for parsing and generating XML files.
For YAML files: Use third-party packages like gopkg.in/yaml.v2 for YAML support.

4. Reading Directories and File Metadata:
Using os.ReadDir to read the contents of a directory.
Getting file metadata using os.Stat or os.Lstat.

5. Buffered I/O (bufio package):
Reading/writing data using buffered I/O for better performance.
Wrapping os.File with bufio.Reader and bufio.Writer for buffered reading and writing.
*Examples*
> Note : In golang, there is no try/catch paradigm, so in case there is some error we might be in trouble and to check that we have this `comma ok / err ok -> result, error := some operation`, if we don't care about the error or even result we can simply replace them with the `result, _ := some_operation`. 

```go 
// reading till newline is encountered
reader := bufio.NewReader(os.Stdin)
fmt.Print("Enter some input :")

input, _ := reader.ReadString('\n')
fmt.Println(input) // the type of input is string currently

//--------Type conversion
reader := bufio.NewReader(os.Stdin)
fmt.Print("Enter some input :")

input, _ := reader.ReadString('\n')
result, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // here if we try to just pass the input like this strconv.ParseFloat(input, 64), this would give error because the input string also contains \n at the end of it and thus the parse error
if err != nil {
    fmt.Println("There was an error", err)
}
fmt.Printf("result is of the type %T : %f", result, result)


```


6. Seeking in Files:
Using File.Seek to move the file pointer to a specific location.
Reading or writing at a specific offset within a file.

7. Working with Paths (path/filepath package):
Building platform-independent paths using filepath.Join.
Cleaning paths using filepath.Clean.

8. Command-Line Arguments (os.Args):
Accessing command-line arguments passed to the Go program via os.Args.

9. Environment Variables (os package):
Getting and setting environment variables using os.Getenv and os.Setenv.

### Working with time
```go
    present_time := time.Now()
	/*
		01-02-2006 -> this prints the time in mm-dd-yyyy format, these values are fixed,
		so to print in the dd-mm-yyyy format we have to use 02-01-2006
		To get the weekday we need to pass the value as monday(always)
		To get the time : 15:04:05(fixed values)
	*/
	epoch_now, epoch_now_nano := time.Now().Unix(), time.Now().UnixNano()

	fmt.Println(present_time)
	fmt.Println(present_time.Format("02-01-2006 15:04:05 Monday"))

	date_from_epoch := time.Unix(epoch_now, 0)

	fmt.Println(epoch_now, epoch_now_nano, date_from_epoch, date_from_epoch.Format("Monday"))
```

### Building for different OSs
* Using the `go env` command we can get to know list of all the environment variables. 
* To build for different operating systems, we can use this `go env` command to select for different operating systems. 
* Build for windows : `GOOS="windows" go build` -> this builds for windows, for macos -> it is darwin and for linux it is linux. 
* Explore more aobut the go-envs. 


### Memory mangagement
* In go memory management is done by the language itself and it does have its garbage collector automatic(out of scope or nil).
* new() -> this allocates memory but no initilization, not used much, zeroed storage(we can't put data initially), here we get memory address. 
* make() -> this allocates memory and iniitialization, we get memory address and non-zeroed storage(we can put date on initilization). 
* https://pkg.go.dev/runtime -> package for low level information about the sytem like cpu usage and stuff.
**About the garbage collector**
  *The GOGC variable sets the initial garbage collection target percentage. A collection is triggered when the ratio of freshly allocated data to live data remaining after the previous collection reaches this percentage. The default is GOGC=100. Setting GOGC=off disables the garbage collector entirely. runtime/debug.SetGCPercent allows changing this percentage at run time.*

 *The GOMEMLIMIT variable sets a soft memory limit for the runtime. This memory limit includes the Go heap and all other memory managed by the runtime, and excludes external memory sources such as mappings of the binary itself, memory managed in other languages, and memory held by the operating system on behalf of the Go program. GOMEMLIMIT is a numeric value in bytes with an optional unit suffix. The supported suffixes include B, KiB, MiB, GiB, and TiB. These suffixes represent quantities of bytes as defined by the IEC 80000-13 standard. That is, they are based on powers of two: KiB means 2^10 bytes, MiB means 2^20 bytes, and so on. The default setting is math.MaxInt64, which effectively disables the memory limit. runtime/debug.SetMemoryLimit allows changing this limit at run time.*

### Pointers 
*Below example should help with basic pointer usage*
```go
    var one_ptr *int // value is <nil>
	value := 435
	// var val_ptr *int = &value
	val_ptr := &value // both above and this are valid approaches
	// address := 0xc00012a006 /* this can be done but can't be used to get the value stored at this address */
	// fmt.Println(address) /* it's decimal value is printed. */
	fmt.Println(one_ptr, val_ptr, *val_ptr, reflect.TypeOf(val_ptr))
    fmt.Println(*(val_ptr + 8)) // this is also not allowed in go
    *val_ptr = *val_ptr + 89 // operation performed on actual value
	fmt.Println(val_ptr, *val_ptr)
    value += 423// operation performed on the actual value and val_ptr and &value are same while *val_ptr an d value are same . 
	fmt.Println(value, &value, val_ptr, *val_ptr)
```
### Basics 
**Arrays** 
* In go, arrays are static and fixed size and can hold elements of single type only. 
```go 
    var arr [5]int = [5]int{4, 34, 2342}
    var arr_3 = [6]string{"fasd", "fasdf", "fd"}
	arr_1 := [...]int{3, 43, 42, 42323, 90}
	arr_2 := [5]int{53, 323, 3123}
	fmt.Println(arr, arr_1, len(arr), len(arr_2))
    // if we don't have elements then it puts 0s or blank spaces or some default value instead of garbage values. 
```
*Not much to do with the arrays in go* 

**Slices** 
```go
   	// var arr = []string{}
	arr_2 := []int{5, 312, 32, 322}
	// fmt.Println(arr, reflect.TypeOf(arr), arr_2, len(arr_2))
	arr_2 = append(arr_2, 80980, 90)
	fmt.Print(arr_2, len(arr_2), "\n")

	arr_2 = append(arr_2[1:], 313, 9013) // this removes the 1st element as expected
	fmt.Print(arr_2, "\n")

	arr_2 = append(arr_2[1:4], 813, 113) // this removes the 1st element and all the elements from 4th index as expected
	fmt.Print(arr_2, "\n")

	// using the make and new keywords
	scores := make([]int, 5)
	fmt.Println(scores)

	scores[0] = 53
	scores[4] = 190
	scores[2] = 17
	sorted := sort.IntsAreSorted(scores)
	fmt.Println(sorted)
	fmt.Println(scores)
	// scores[5] = 89, this would give us an error

	scores = append(scores, 52, 32, 901, 3132) // doing this let's the array resize and accomodate all the values instead of giving an error
	fmt.Println(scores)

	sort.Ints(scores)
	fmt.Println(scores)
	sorted = sort.IntsAreSorted(scores)
	fmt.Println(sorted)

	// remove values from slice(index-based), we need to do this using append
	index := 3
	scores = append(scores[:index], scores[index+1:]...) // when appending 2 slices we need to spread the elements last slice, we can't concatenate more than two slices
	fmt.Println(scores)

	fmt.Println(scores[3:4])
```