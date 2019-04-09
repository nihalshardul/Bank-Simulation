# Bank-Simulation

In this assignment we are going to build a simulation model of a bank branch.

We have to model the bank cashiers and the customers.

A cashier can be in two states, Occupied & Not Occupied. 
If a cashier is not occupied, the cashier takes the next customer from a queue. 
If the cashier is processing a customer the status is Occupied.
It takes 'n' seconds for a bank cashier to process one customer.
Customers are labelled in increasing numberical order 1,2,3.
All cashiers are operating in parallel independent of each other.

The input will be in the following format: 

./runprogram --numCashiers=3 --numCustomers=100 --timePerCustomer=3


The output will be the following format:

2019-04-05 13:20:22 --> Bank Simulation Started

2019-04-05 13:20:22 --> Cashier 1: Customer 1 Started

2019-04-05 13:20:22 --> Cashier 2: Customer 2 Started

2019-04-05 13:20:22 --> Cashier 3: Customer 3 Started

2019-04-05 13:20:22 --> Cashier 2: Customer 2 Completed

2019-04-05 13:20:22 --> Cashier 1: Customer 1 Completed

2019-04-05 13:20:22 --> Cashier 2: Customer 4 Started

.......

.......

.......

2019-04-05 13:20:22 --> Bank Simulated Ended


Notes:

One cashier model will carry partial marks

You may have to use mutex, go routines, wait groups etc to solve this problem.
