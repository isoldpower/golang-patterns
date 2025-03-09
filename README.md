# Golang Pipeline
This is a simple project to practice with golang concurrency features and implement 
the pipeline concurrency pattern with starting a list of goroutines to optimize
slow stages of pipeline

## Usage
Just run the following command to execute the app:
```bash
go run main.go
```

## App Flow
The app is generating big random numbers (in range from _10.000.000_ to _99.999.999_) 
until **10 prime numbers** are generated. When the program is run, it runs 2 processes one-by-one:
1) Parallel process. It runs goroutines on CPU cores separately to identify prime numbers
2) Default process. It uses only one CPU core to identify prime numbers
