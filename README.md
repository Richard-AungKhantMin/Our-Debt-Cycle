# Our Debt Cycle

## Overview
**Our Debt Cycle** is a Go-based tool that helps track and calculate shared expenses among friends when buying groceries together. Since different individuals buy items at different times, keeping track of debts manually can be cumbersome. This program automates the process, calculating outstanding balances and updating debt history within seconds.

## Features
- **Automated Debt Calculation**: Processes multiple transactions and determines the final debt balances.
- **History Tracking**: Maintains a stored history of debts and updates it whenever a new calculation is made.
- **Dual Usage Modes**: Run it as a terminal command or a web server for convenience.
- **Simple Text Input Format**: Transactions are recorded in a text file, making it easy to track purchases.

## Usage

### Running in Terminal
You can run the program in the terminal by specifying a text file containing transactions:
```sh
go run . fileName.txt
```

### Running as a Web Server
Alternatively, you can start a web server to interact with the tool via a browser:
```sh
go run .
```

## Input Format
Each line in the input file should follow this format:
```
[payer-name]-[payee-name] [amount]
```
Example:
```
All-Chan 15.50
All-Richie 10.00
Chan-Milli 8.75
Richie-Milli 20
```
This means:
- Everyone owes Chan €15.50
- Everyone owes Richie €10.00
- Chan owes Milli €8.75
- Richie owes Milli €20.00

Results will be printed on the terminal and stored in the `historyFiles/` repository as well, under the name of the latest time.

## How It Works
1. The program reads the transaction history from the given text file.
2. It calculates who owes whom and how much.
3. It updates and stores the latest debt history.
4. When run as a web server, it provides a user-friendly interface to input and view transactions.

## Installation
Ensure you have Go installed on your system. Then, clone the repository:
```sh
git clone https://github.com/Richard-AungKhantMin/Our-Debt-Cycle
```

## Contributions
Feel free to contribute! Open an issue or submit a pull request if you have improvements or bug fixes.

---
**Our Debt Cycle** simplifies expense tracking so you can focus on sharing meals without worrying about the numbers!

