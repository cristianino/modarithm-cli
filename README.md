# Modarithm CLI

A **Command-Line Interface (CLI)** written in Go for exploring **modular arithmetic** as a foundation for cryptography.  
The project is both an educational journey and a practical toolkit: learn, experiment, and extend toward cryptographic applications.

## 📥 Quick Download & Install

[![Latest Release](https://img.shields.io/github/v/release/cristianino/modarithm-cli?style=for-the-badge&logo=github&color=blue)](https://github.com/cristianino/modarithm-cli/releases/latest)
[![Downloads](https://img.shields.io/github/downloads/cristianino/modarithm-cli/total?style=for-the-badge&logo=download&color=green)](https://github.com/cristianino/modarithm-cli/releases)

> **🚀 Ready to use! No compilation needed.**

### Choose your platform:

<table>
<tr>
<td align="center">
<img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/linux/linux-original.svg" width="48" height="48"><br>
<strong>Linux x64</strong><br>
<a href="https://github.com/cristianino/modarithm-cli/releases/latest/download/modarithm-v1.0.0-linux-amd64.tar.gz">
� Download
</a>
</td>
<td align="center">
<img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/linux/linux-original.svg" width="48" height="48"><br>
<strong>Linux ARM64</strong><br>
<a href="https://github.com/cristianino/modarithm-cli/releases/latest/download/modarithm-v1.0.0-linux-arm64.tar.gz">
📥 Download
</a>
</td>
<td align="center">
<img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/apple/apple-original.svg" width="48" height="48"><br>
<strong>macOS Intel</strong><br>
<a href="https://github.com/cristianino/modarithm-cli/releases/latest/download/modarithm-v1.0.0-darwin-amd64.tar.gz">
📥 Download
</a>
</td>
<td align="center">
<img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/apple/apple-original.svg" width="48" height="48"><br>
<strong>macOS Apple Silicon</strong><br>
<a href="https://github.com/cristianino/modarithm-cli/releases/latest/download/modarithm-v1.0.0-darwin-arm64.tar.gz">
📥 Download
</a>
</td>
<td align="center">
<img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/windows8/windows8-original.svg" width="48" height="48"><br>
<strong>Windows x64</strong><br>
<a href="https://github.com/cristianino/modarithm-cli/releases/latest/download/modarithm-v1.0.0-windows-amd64.zip">
📥 Download
</a>
</td>
</tr>
</table>

### 🔧 Installation Steps:

#### 🐧 **Linux / macOS:**
```bash
# 1. Download and extract (replace URL with your platform)
wget https://github.com/cristianino/modarithm-cli/releases/latest/download/modarithm-v1.0.0-linux-amd64.tar.gz
tar -xzf modarithm-v1.0.0-linux-amd64.tar.gz

# 2. Install system-wide (optional)
sudo mv modarithm-v1.0.0-linux-amd64 /usr/local/bin/modarithm

# 3. Test installation
modarithm --help
```

#### 🪟 **Windows:**
1. **Download** the `.zip` file for Windows
2. **Extract** the archive 
3. **Add to PATH** or move `modarithm.exe` to a directory in your PATH
4. **Test** by opening Command Prompt and running `modarithm --help`

### ✅ **Verify Download (Optional)**
```bash
# Download checksums
wget https://github.com/cristianino/modarithm-cli/releases/latest/download/checksums.txt

# Verify integrity
sha256sum -c checksums.txt --ignore-missing
```

---

## 🚀 Quick Start Examples

Get started immediately with these common modular arithmetic operations:

### 🔢 **Greatest Common Divisor (GCD)**
```bash
# Find GCD of two numbers
modarithm gcd --numbers "48,18"
# Output: 6

# GCD of multiple numbers  
modarithm gcd --numbers "12,18,24"
# Output: 6
```

### 📏 **Least Common Multiple (LCM)**
```bash
# Find LCM of multiple numbers
modarithm lcm --numbers "4,6,8"
# Output: 24

# LCM with larger numbers
modarithm lcm --numbers "15,25,35"
# Output: 525
```

### ⚡ **Modular Exponentiation**
```bash
# Compute (base^exponent) mod modulus efficiently
modarithm modexp --base 3 --exponent 4 --modulus 5
# Output: 1 (3^4 = 81, 81 mod 5 = 1)

# Large numbers handled efficiently
modarithm modexp --base 123 --exponent 456 --modulus 789
```

### 🔄 **Modular Inverse**
```bash
# Find modular inverse (if exists)
modarithm inverse --number 3 --modulus 7
# Output: 5 (because 3 * 5 ≡ 1 (mod 7))

# Check if inverse exists
modarithm inverse --number 6 --modulus 9
# Output: No inverse exists (gcd(6,9) ≠ 1)
```

> 💡 **Tip**: Run `modarithm --help` or `modarithm <command> --help` to see all available options!

---

## ✨ Vision

- Provide **modular arithmetic utilities** (multiples, divisibility, modular addition/multiplication, inverses, congruences).
- Serve as a **learning companion** for studying number theory concepts relevant to cryptography.
- Act as a **playground** before building production-grade crypto tools.

---

## 📦 Installation

### Option 1: Download Pre-built Binary (Recommended)

Download the latest binary from the links above, then:

```bash
# Linux/macOS
tar -xzf modarithm-*.tar.gz
chmod +x modarithm-*
./modarithm-* --help

# Windows
# Extract the .zip file
# Run modarithm-windows-amd64.exe --help
```

### Option 2: Build from Source

```bash
git clone https://github.com/cristianino/modarithm-cli.git
cd modarithm-cli
go mod tidy
make build
./modarithm --help
```

### Option 3: Install with Go

```bash
go install github.com/cristianino/modarithm-cli@latest
modarithm-cli --help
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

## �️ Build from Source

For developers who want to modify or contribute:

```bash
# Clone repository
git clone https://github.com/cristianino/modarithm-cli.git
cd modarithm-cli

# Install dependencies
go mod tidy

# Build
make build

# Run tests
make test

# Build for all platforms
make build-all
```

---

## 📖 Command Reference

Run `modarithm --help` to see all available commands:

• `gcd` - Calculate Greatest Common Divisor using Euclidean algorithm
• `lcm` - Calculate Least Common Multiple efficiently  
• `modexp` - Perform modular exponentiation with large numbers
• `inverse` - Find modular multiplicative inverse
• `congruence` - Solve and check linear congruences
• `multiples` - Generate and analyze multiples of numbers

> 💡 **Tip**: Use `modarithm <command> --help` for detailed options on each command.

---

## 🎯 Educational Purpose

This tool is designed for learning modular arithmetic and number theory. Each command demonstrates real-world mathematical concepts:

• Understand how fundamental algorithms work (Euclidean, Extended Euclidean)
• Practice with efficient implementations
• Experiment with parameters and see mathematical relationships
• Learn foundations essential for cryptography

---

## 🛡️ Security Notes

• Uses standard mathematical algorithms (Euclidean algorithm, binary exponentiation)
• Proper handling of edge cases and large numbers
• Educational focus with mathematically sound implementations
• No cryptographic secrets or random number generation

---

## �📚 Motivation

This project was born as a **personal study plan** in modular arithmetic, with the ultimate goal of reinforcing cryptography knowledge. It is designed for students, enthusiasts, and developers who want a **hands-on math lab** in the terminal.

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

---

<div align="center">

**⭐ If you find this tool useful, please give it a star on GitHub! ⭐**

[Report Bug](https://github.com/cristianino/modarithm-cli/issues) • [Request Feature](https://github.com/cristianino/modarithm-cli/issues) • [Documentation](https://github.com/cristianino/modarithm-cli)

</div>

