# Modarithm CLI

A **Command-Line Interface (CLI)** written in Go for exploring **modular arithmetic** as a foundation for cryptography.  
The project is both an educational journey and a practical toolkit: learn, experiment, and extend toward cryptographic applications.

---

## ✨ Vision

- Provide **modular arithmetic utilities** (multiples, divisibility, modular addition/multiplication, inverses, congruences).
- Serve as a **learning companion** for studying number theory concepts relevant to cryptography.
- Act as a **playground** before building production-grade crypto tools.

---

## 📦 Installation

```bash
git clone https://github.com/cristianino/modarithm-cli.git
cd modarithm-cli
go mod tidy
go build -o modarithm
./modarithm --help
````

---

## 🚀 Roadmap (Planned Features)

* **Stage 1 – Fundamentals**

  * Multiples & divisibility
  * Prime numbers & factorization
  * Modular addition, subtraction, multiplication
  * Congruence relations

* **Stage 2 – Intermediate**

  * Modular inverses
  * Greatest Common Divisor (GCD) & Extended Euclidean Algorithm
  * Least Common Multiple (LCM)
  * Modular exponentiation

* **Stage 3 – Cryptographic Bridges**

  * Euler’s totient function
  * Fermat’s Little Theorem
  * Chinese Remainder Theorem
  * Random number checks (primality tests)

* **Stage 4 – Cryptography Applications**

  * Key generation building blocks
  * Simple cryptosystems based on modular arithmetic
  * Linking with existing `crypto-cli` project

---

## 🔧 Example Usage (planned)

```bash
# List first 10 multiples of 7
modarithm multiples --n 7 --limit 10

# Check if 45 is a multiple of 9
modarithm is-multiple --n 45 --a 9

# Compute gcd(84, 30)
modarithm gcd --a 84 --b 30

# Compute modular inverse of 7 mod 26
modarithm inverse --a 7 --mod 26
```

---

## 📚 Motivation

This project was born as a **personal study plan** in modular arithmetic, with the ultimate goal of reinforcing cryptography knowledge. It is designed for students, enthusiasts, and developers who want a **hands-on math lab** in the terminal.

---

## ⚠️ Disclaimer

This tool is **educational**. Do not use it for production cryptography. For real-world security, rely on audited libraries and standards.

