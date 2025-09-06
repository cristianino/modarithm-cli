# Release Notes - v1.0.0

## 🎉 First Stable Release!

This is the first stable release of **modarithm-cli**, a comprehensive command-line tool for learning and practicing modular arithmetic as a foundation for cryptography. This Go implementation provides essential mathematical operations with support for multiple platforms.

## ✨ Features

### Core Mathematical Operations
- **🔢 GCD (Greatest Common Divisor)**: Euclidean algorithm for finding the greatest common divisor
- **📏 LCM (Least Common Multiple)**: Efficient computation of least common multiples
- **⚡ Modular Exponentiation**: Fast modular exponentiation using binary method
- **🔄 Modular Inverse**: Find modular multiplicative inverses when they exist
- **🔍 Congruence Relations**: Check and solve linear congruences
- **📊 Multiples Analysis**: Generate and analyze multiples of numbers

### Platform Support
- **Linux**: x86_64 and ARM64
- **macOS**: Intel (x86_64) and Apple Silicon (ARM64)  
- **Windows**: x86_64

### User Experience
- **Intuitive CLI**: Easy-to-use command structure with comprehensive help
- **Educational Focus**: Perfect for learning number theory and cryptographic foundations
- **Performance**: Optimized algorithms for handling large numbers
- **Cross-platform**: Works seamlessly across all major operating systems

## 📦 Installation Options

### Quick Install (Precompiled Binaries)
Download ready-to-use binaries from the [releases page](https://github.com/cristianino/modarithm-cli/releases/v1.0.0):

- `modarithm-v1.0.0-linux-amd64.tar.gz` - Linux x64
- `modarithm-v1.0.0-linux-arm64.tar.gz` - Linux ARM64
- `modarithm-v1.0.0-darwin-amd64.tar.gz` - macOS Intel
- `modarithm-v1.0.0-darwin-arm64.tar.gz` - macOS Apple Silicon
- `modarithm-v1.0.0-windows-amd64.zip` - Windows x64

### Build from Source
```bash
git clone https://github.com/cristianino/modarithm-cli.git
cd modarithm-cli
make build
```

## 🚀 Quick Start

```bash
# Find GCD of two numbers
modarithm gcd --numbers "48,18"

# Calculate LCM of multiple numbers
modarithm lcm --numbers "4,6,8"

# Compute modular exponentiation
modarithm modexp --base 3 --exponent 4 --modulus 5

# Find modular inverse
modarithm inverse --number 3 --modulus 7

# Check congruence relations
modarithm congruence --number1 17 --number2 5 --modulus 6
```

## 🔒 Security

All mathematical operations use well-tested algorithms and implementations:
- Euclidean algorithm for GCD computation
- Binary exponentiation for modular exponentiation
- Extended Euclidean algorithm for modular inverses
- Secure random number generation when applicable

## 📖 Educational Purpose

This tool was specifically designed for educational purposes to help students and developers understand:
- Fundamental concepts in number theory
- Modular arithmetic operations essential for cryptography
- Practical implementation of mathematical algorithms
- The relationship between different mathematical operations

## 🛡️ Verification

All release binaries include SHA256 checksums for integrity verification. Download `checksums.txt` and verify your download:

```bash
sha256sum -c checksums.txt --ignore-missing
```

## 🤝 Contributing

We welcome contributions! Please see our [GitHub repository](https://github.com/cristianino/modarithm-cli) for more information.

## 📄 License

This project is licensed under the MIT License. See the LICENSE file for details.

---

**Full Changelog**: https://github.com/cristianino/modarithm-cli/commits/v1.0.0
