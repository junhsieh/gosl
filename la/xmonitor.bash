#!/bin/bash

FILE="*.go"

while true; do
    inotifywait -q -e modify $FILE
    echo
    echo
    echo
    echo
    #go test -test.run="SpSolver06M"
    #go test
    #mpirun -np 3 go run t_sp_mpi_main.go
    #mpirun -np 2 go run t_mumpssol01b_main.go
    #mpirun -np 2 go run t_mumpssol01a_main.go
    #mpirun -np 5 go run t_mumpssol02_main.go
    #mpirun -np 2 go run t_mumpssol03_main.go
    #mpirun -np 2 go run t_mumpssol04_main.go
    #mpirun -np 3 go run t_mumpssol05_main.go
    #go test -test.run="Vector03"
    #go test -test.run="Triplet02"
    #go test -test.run="DenSolve01"
    #go test -test.run="Matrix06"
    #go test -test.run="Eigen05"
    go test -test.run="Eqs04"
done
