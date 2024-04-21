# QuadChecker

**QuadChecker** is a sophisticated tool designed to validate and process input through interactions with five other programs (quadA, quadB, quadC, quadD, quadE). It ensures data integrity and optimizes the workflow in multi-program environments by determining which of the quad programs could have produced a given output based on dimensions.  

## Prerequisites
Before you begin, ensure you have the following installed:  

**Go** (recommended version **1.18 or newer**): Go is necessary to compile and run the programs. You can download it from the [official Go website](https://go.dev/dl/).

## Installation
To get started with **QuadChecker** and its associated programs, follow these steps:

### Clone the Repository
Clone the project repository to your local machine using the following command:

```
git clone URL_OF_REPOSITORY
cd DIRECTORY_OF_THE_REPOSITORY
```

Replace  `URL_OF_REPOSITORY`  with the actual **URL of the repository** and  `DIRECTORY_OF_THE_REPOSITORY`  with the directory name where the repository is cloned.

### Compile and Prepare the Programs
Navigate to the directory containing the  **`Quads`**  subdirectory and use the builder file to compile and make the  **`quad`**  functions and the  **`quadchecker`**  utility executable:

```
go run QuidBuilder/builder.go /filepath/to/Quads
``` 

This command will compile all necessary  `.go`  files in the  **`Quads`**  directory and set the appropriate executable permissions. The builder script automates the process of making files executable, ensuring that all components are ready to use.

### Move Executables to the Parent Directory
After compiling and setting the executable permissions, move the executable files from the  **`Quads`**  directory to the parent directory to ensure they are in the same location as  `main.go`  and  `go.mod`. Use the following command:

```
mv Quads/* .
```
This command moves all files from the  **`Quads`**  directory to the current directory, where  `main.go`  and  `go.mod`  are located.

## Usage

### Running the Quad Functions
Each  **`quad`**  function can be run with dimensions as command-line arguments. For example, to run  **`quadA`**  with a width of 3 and a height of 3, use the following command:

```
./quadA 3  3
```

### Using the Quad Checker
The  **`quadchecker`**  utility can be used to identify which  **`quad`**  function(s) could have produced a given output. It accepts dimensions as command-line arguments or reads from standard input.

#### Example 1: Using Command-Line Arguments
```
./quadchecker 3  3
``` 

This command will check if the output matches any of the  **`quad`**  functions with dimensions 3x3 and print the matches.

#### Example 2: Using Standard Input
You can also pipe the output of a  **`quad`**  function into  **`quadchecker`**  to check for matches. For example:

```
./quadA 3  3  | ./quadchecker
``` 

This command pipes the output of  `./quadA 3 3`  into  `./quadchecker`, which then checks for matches and prints them.

### Using the Main Program
The  `main.go`  file can be run directly with  `go run`  to either pass command-line arguments directly to  **`quadchecker`**  or read from standard input and pass it to  **`quadchecker`**.

To pass command-line arguments:

```
go run main.go 3  3
``` 

To read from standard input:

```
./quadA 3  3  | go run main.go
``` 

This will produce the same output as Example 2, checking for matches based on the piped input.


