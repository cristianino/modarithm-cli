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
```

## 📖 Documentation

- **[MATH_FORMULAS.md](MATH_FORMULAS.md)** - Detailed mathematical formulas, algorithms, and implementation analysis`

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

# Matrix output: multiple n values (comma-separated)
modarithm multiples --n 4,6 --limit 6 --group pos
# Output:
# [4 8 12 16 20 24]
# [6 12 18 24 30 36]

# Matrix with more values
modarithm multiples --n 3,5,7 --limit 4 --group all
# Output:
# [-12 -9 -6 -3 0 3 6 9 12]
# [-20 -15 -10 -5 0 5 10 15 20]
# [-28 -21 -14 -7 0 7 14 21 28]

# JSON output for multiples (structured + flat array)
modarithm multiples --n 7 --limit 3 --json
# Output:
# {
#   "all": [-21, -14, -7, 0, 7, 14, 21],
#   "negatives": [-21, -14, -7],
#   "zero": 0,
#   "positives": [7, 14, 21]
# }

# JSON matrix output
modarithm multiples --n 4,6 --limit 3 --group pos --json
# Output:
# {
#   "group": "pos",
#   "matrix": [[4, 8, 12], [6, 12, 18]],
#   "n_values": [4, 6]
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
- **Matrix support**: Use comma-separated n values (e.g., `--n 4,6,8`) for matrix output
- When `n=0`, all multiples are zero (edge case handled)
- Limits must be >= 0 (validation included)
- JSON with `--group all` (single n) provides both structured object and flat array
- JSON with matrix provides `n_values`, `group`, and `matrix` fields
```

### GCD (Greatest Common Divisor)

```bash
# Basic GCD computation
modarithm gcd --a 84 --b 30
# Output: GCD(84, 30) = 6

# GCD with smaller numbers
modarithm gcd --a 12 --b 18
# Output: GCD(12, 18) = 6

# Coprime numbers (GCD = 1)
modarithm gcd --a 17 --b 19
# Output: GCD(17, 19) = 1

# Works with negative numbers
modarithm gcd --a -15 --b 25
# Output: GCD(-15, 25) = 5
```

### LCM (Least Common Multiple)

The LCM command computes the least common multiple using the formula: `LCM(a,b) = |a*b| / GCD(a,b)`

```bash
# Basic LCM computation
modarithm lcm --a 12 --b 18
# Output: LCM(12, 18) = 36

# LCM relationship: 12 = 2² × 3, 18 = 2 × 3²
# LCM takes highest powers: 2² × 3² = 36

# LCM with different examples
modarithm lcm --a 15 --b 25
# Output: LCM(15, 25) = 75
# (15 = 3 × 5, 25 = 5², LCM = 3 × 5² = 75)

# Coprime numbers: LCM equals their product
modarithm lcm --a 7 --b 13
# Output: LCM(7, 13) = 91
# (Since GCD(7,13) = 1, LCM = 7 × 13 = 91)

# Edge case: zero input
modarithm lcm --a 0 --b 15
# Output: LCM(0, 15) = 0

# Works with negative numbers (result is always positive)
modarithm lcm --a -12 --b 18
# Output: LCM(-12, 18) = 36
```

**LCM Applications:**
- **Finding common periods:** When two cyclic processes repeat every `a` and `b` units
- **Fraction arithmetic:** Finding common denominators
- **Scheduling problems:** When events occur at different intervals
- **Cryptography:** Key generation and period calculations

### GCD ↔ LCM Relationship

The fundamental relationship between GCD and LCM: **GCD(a,b) × LCM(a,b) = |a × b|**

```bash
# Demonstrate the relationship with a = 12, b = 18:
modarithm gcd --a 12 --b 18  # Result: 6
modarithm lcm --a 12 --b 18  # Result: 36

# Verify: GCD × LCM = 6 × 36 = 216 = |12 × 18|
# This relationship holds for any integers a, b

# Another example with a = 15, b = 25:
modarithm gcd --a 15 --b 25  # Result: 5  
modarithm lcm --a 15 --b 25  # Result: 75
# Verify: 5 × 75 = 375 = |15 × 25|
```

### Other Mathematical Commands

```bash
# Modular inverse (when it exists)
modarithm inverse --a 7 --mod 26
# Output: Inverse of 7 mod 26 = 15
# Verification: (7 × 15) mod 26 = 105 mod 26 = 1

# Fast modular exponentiation
modarithm modexp --base 5 --exp 117 --mod 19
# Output: 5^117 mod 19 = 1
# Uses exponentiation by squaring for efficiency

# Check congruence relationships
modarithm congruence --a 45 --b 9 --mod 6
# Output: 45 ≡ 9 (mod 6)? true
# Output: Canonical residues: 3, 3
# (Both 45 and 9 have remainder 3 when divided by 6)
```

### Practical Use Cases

**Combining GCD & LCM for problem solving:**

```bash
# Problem: Two gears with 12 and 18 teeth - when do they align?
modarithm lcm --a 12 --b 18
# Answer: Every 36 teeth rotations they align completely

# Problem: Simplify fraction 84/30
modarithm gcd --a 84 --b 30  # GCD = 6
# Simplified: (84÷6)/(30÷6) = 14/5

# Problem: Find when two events coincide
# Event A every 15 minutes, Event B every 25 minutes
modarithm lcm --a 15 --b 25
# Answer: They coincide every 75 minutes (1h 15min)
```

---

## 📚 Motivation

This project was born as a **personal study plan** in modular arithmetic, with the ultimate goal of reinforcing cryptography knowledge. It is designed for students, enthusiasts, and developers who want a **hands-on math lab** in the terminal.

---

## ⚠️ Disclaimer

This tool is **educational**. Do not use it for production cryptography. For real-world security, rely on audited libraries and standards.

