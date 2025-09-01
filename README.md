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

## 🔧 Example Usage

### Multiples Commands

```bash
# Generate symmetric multiples of 5 (4 on each side) - now defaults to 'all'
modarithm multiples --n 5 --limit 4
# Output: [-20 -15 -10 -5 0 5 10 15 20]

# Generate asymmetric multiples of 3 (2 negative, 4 positive)  
modarithm multiples --n 3 --limitNeg 2 --limitPos 4
# Output: [-6 -3 0 3 6 9 12]

# JSON output for multiples (structured + flat array)
modarithm multiples --n 7 --limit 3 --json
# Output:
# {
#   "all": [-21, -14, -7, 0, 7, 14, 21],
#   "negatives": [-21, -14, -7],
#   "zero": 0,
#   "positives": [7, 14, 21]
# }

# Show only specific groups (consistent flag names: neg|zero|pos|all)
modarithm multiples --n 5 --limit 4 --group neg
# Output: [-20 -15 -10 -5]

modarithm multiples --n 5 --limit 4 --group zero
# Output: 0

modarithm multiples --n 5 --limit 4 --group pos
# Output: [5 10 15 20]

# Get all multiples in a single ordered array (default behavior)
modarithm multiples --n 5 --limit 4 --group all
# Output: [-20 -15 -10 -5 0 5 10 15 20]

# Works with asymmetric limits too
modarithm multiples --n 3 --limitNeg 2 --limitPos 3 --group all
# Output: [-6 -3 0 3 6 9]

# Combine specific groups with JSON for clean arrays
modarithm multiples --n 3 --limit 3 --group pos --json
# Output: [3, 6, 9]
```

**Notes:**
- Default `--group` is `all` for convenience
- Use consistent flag names: `neg`, `zero`, `pos`, `all`  
- When `n=0`, all multiples are zero (edge case handled)
- Limits must be >= 0 (validation included)
- JSON with `--group all` provides both structured object and flat array
```

### Other Commands

```bash
# Compute gcd(84, 30)
modarithm gcd --a 84 --b 30

# Compute lcm(12, 18)
modarithm lcm --a 12 --b 18

# Compute modular inverse of 7 mod 26
modarithm inverse --a 7 --mod 26

# Fast modular exponentiation
modarithm modexp --base 5 --exp 117 --mod 19

# Check congruence
modarithm congruence --a 45 --b 9 --mod 6
```

---

## 📚 Motivation

This project was born as a **personal study plan** in modular arithmetic, with the ultimate goal of reinforcing cryptography knowledge. It is designed for students, enthusiasts, and developers who want a **hands-on math lab** in the terminal.

---

## ⚠️ Disclaimer

This tool is **educational**. Do not use it for production cryptography. For real-world security, rely on audited libraries and standards.

