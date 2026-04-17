# 🧮 Go FizzBuzz: A Data-Driven Refactoring Exercise

An idiomatic and scalable Go implementation of the classic FizzBuzz algorithm.

Although the mathematical problem itself is fundamentally simple, this repository practically demonstrates Software Engineering principles. It showcases a properly structured, clean, and `data-driven` codebase, highlighting the Separation of Concerns and the Open/Closed Principle (SOLID).

## 🧠 Architecture & Design Decisions

This codebase implements the following practices:

* **Data-Driven Design (Separation of Concerns):** The "business logic" (which numbers result in which words) was completely decoupled from the execution engine. I used a slice of structs (`combinationType`) to store the divisors, making the `main` function blind to the specific rules.

* **Open/Closed Principle:** The system is open to extending the number of rules, but closed to logic modification. If requirements change and we need to add a new rule (e.g., multiples of 7 return "Bazz" or multiples of 11 return "Jazz"), we just need to add a single line to the data structure, without altering the main loop's logic.

* **Efficiency and Optimization:** Instead of redundantly checking the modulo (`%`) and trying to predict all possible logical combinations within nested `if/else` statements, the result is built sequentially. This ensures the mathematical operation only needs to be executed once per rule, scaling predictably.

* **Idiomatic Go:** Correct use of the language's semantics and safe iteration using the `range` loop.

# 🚀 Running the Project

The project was built using only the Go standard library, with no external dependencies required.

**Prerequisites:** Go 1.18+ installed.

1. Clone the repository:
```bash
git clone [https://github.com/CasimiroDev/fizzbuzz-challenge](https://github.com/CasimiroDev/fizzbuzz-challenge)
```

2. Navigate to the project folder:
```bash
cd fizzbuzz-challenge
```

3. Run the application:
```bash
go run main.go
```

# 🔮 Next Steps (To-Do)

- [ ] Unit Tests (TDD): Implement a test suite using the native `testing` package with Table-Driven Tests, validating different combinations of divisors.

- [ ] Dynamic Configuration: Allow the use of a `.json` file to inject business rules. An alternative would be using CLI (Command Line Interface) arguments.