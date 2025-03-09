# Golang Patterns
This is a simple project to practice with golang concurrency features and implement 
the concurrency pattern with starting a list of goroutines to optimize processes. <br>
For now, the project has the following patterns implemented:
1) [Pipeline Pattern](#pipeline)

## Usage
Just run the following command to execute the app:
```bash
go run main.go
```

## App Flow
When the user runs the app, he will first meet the selection menu 
where it will be asked to choose the pattern it wants to run. After
choosing the pattern, it will execute the specific pattern runtime. See
below for more insights.

### Pipeline
The app is generating big random numbers (in range from _10.000.000_ to _99.999.999_) 
until **10 prime numbers** are generated. When the program is run, it runs 2 processes one-by-one:
1) Parallel process. It runs goroutines on CPU cores separately to identify prime numbers
2) Default process. It uses only one CPU core to identify prime numbers
