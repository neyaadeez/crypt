# **Everything About RSA (Rivest-Shamir-Adleman Algorithm)**  

RSA is one of the most widely used **asymmetric cryptographic algorithms** in security. It provides **encryption, decryption, and digital signatures** using a **public-private key pair**.

---

## **1. What is RSA?**
RSA is an **asymmetric encryption algorithm** that uses two keys:
- **Public Key** → Used for **encryption** (can be shared openly)
- **Private Key** → Used for **decryption** (kept secret)

It is **mathematically based on the difficulty of factoring large prime numbers**.

RSA is used in:
- **TLS/SSL (HTTPS security)**
- **PGP (Pretty Good Privacy) for secure emails**
- **Digital Signatures**
- **Secure Shell (SSH) authentication**
- **Cryptographic tokens (e.g., smart cards, 2FA hardware keys)**

---

## **2. How RSA Works**
RSA relies on **modular arithmetic and prime factorization**. The core idea is that multiplying two large prime numbers is easy, but factoring their product is computationally hard.

### **Step-by-Step Process**
1. **Key Generation**
   - Choose **two large prime numbers**: `p` and `q`
   - Compute **modulus**:  
     \[
     n = p \times q
     \]
   - Compute **Euler’s Totient Function**:
     \[
     \phi(n) = (p - 1) \times (q - 1)
     \]
   - Choose a public exponent `e` (typically **65537** for efficiency).
   - Compute **private key exponent `d`**:
     \[
     d = e^{-1} \mod \phi(n)
     \]
     (`d` is the modular inverse of `e` mod `φ(n)`)

2. **Encryption**
   - Given plaintext `M`, convert it into a number smaller than `n`.
   - Encrypt using the **public key `(e, n)`**:
     \[
     C = M^e \mod n
     \]
     (C is the ciphertext)

3. **Decryption**
   - Use the **private key `(d, n)`**:
     \[
     M = C^d \mod n
     \]
     (Retrieves the original plaintext `M`)

4. **Digital Signatures (for authentication)**
   - Sender **hashes** the message and encrypts the hash with their **private key**.
   - Receiver **decrypts** it with the sender's **public key** and verifies the hash.

---

## **3. Example: RSA Key Generation and Encryption in Python**
### **Generate RSA Key Pair**
```python
from Crypto.PublicKey import RSA

# Generate a 2048-bit key pair
key = RSA.generate(2048)

# Extract private and public keys
private_key = key.export_key()
public_key = key.publickey().export_key()

# Print the keys
print(private_key.decode())
print(public_key.decode())
```

### **Encrypt and Decrypt with RSA**
```python
from Crypto.Cipher import PKCS1_OAEP
from Crypto.PublicKey import RSA
import base64

# Load public key
public_key = RSA.import_key(open("public.pem").read())
cipher_rsa = PKCS1_OAEP.new(public_key)

# Encrypt a message
message = b"Hello, RSA Encryption!"
ciphertext = cipher_rsa.encrypt(message)

print("Encrypted:", base64.b64encode(ciphertext).decode())

# Decrypt with the private key
private_key = RSA.import_key(open("private.pem").read())
cipher_rsa = PKCS1_OAEP.new(private_key)

plaintext = cipher_rsa.decrypt(ciphertext)
print("Decrypted:", plaintext.decode())
```

---

## **4. Key Strength and Security**
### **Why RSA is Secure?**
- The security of RSA relies on the **factoring problem**: **given `n = p × q`, it's extremely hard to find `p` and `q`**.
- As **key size increases**, brute-force attacks become infeasible.
- **Common key sizes**:
  - **1024-bit** → Weak (Deprecated)
  - **2048-bit** → Secure (Recommended)
  - **4096-bit** → More secure, but slower

### **Vulnerabilities**
1. **Small Key Size** → Can be brute-forced (e.g., 512-bit keys broken by factoring attacks).
2. **Bad Randomness** → If `p` and `q` are not chosen properly, RSA can be broken.
3. **Padding Attacks**:
   - **PKCS#1 v1.5 Padding** → Vulnerable to **Bleichenbacher Attack**.
   - **Solution:** Use **OAEP (Optimal Asymmetric Encryption Padding)**.
4. **Side-Channel Attacks**:
   - **Timing Attacks** → Measure decryption time to infer private key.
   - **Countermeasure:** Implement **constant-time operations**.

---

## **5. Comparing RSA with Other Algorithms**
| Algorithm | Type | Security Level | Key Exchange | Performance |
|-----------|------|---------------|--------------|------------|
| **RSA** | Asymmetric | Strong (2048+ bit) | Public Key | **Slow for large data** |
| **AES** | Symmetric | Strong (128/256-bit) | No (Pre-shared key) | **Fast for bulk encryption** |
| **ECC (Elliptic Curve Cryptography)** | Asymmetric | Stronger at smaller key sizes | Public Key | **Faster than RSA** |

📌 **RSA is best for key exchange and digital signatures but is slower than ECC and AES for bulk encryption.**

---

## **6. Real-World Use Cases of RSA**
| Use Case | How RSA is Used |
|----------|---------------|
| **HTTPS/TLS (Web Security)** | RSA is used for key exchange in **SSL/TLS certificates** (Google, Amazon, etc.) |
| **Digital Signatures** | RSA signs **software packages, emails (PGP), and legal documents** |
| **Cryptographic Tokens** | Used in **smart cards, YubiKeys, and hardware security modules (HSMs)** |
| **Secure Email (PGP/GPG)** | Encrypting and signing emails using **PGP (Pretty Good Privacy)** |
| **Authentication (SSH Keys)** | RSA key pairs are used for **SSH login without passwords** |

---

## **7. Future of RSA**
- **ECC (Elliptic Curve Cryptography)** is replacing RSA in many applications due to:
  - **Smaller key sizes** (256-bit ECC = 3072-bit RSA).
  - **Faster computations** (better for mobile and IoT).
- **Post-Quantum Cryptography (PQC)**:
  - Quantum computers may break RSA (Shor’s Algorithm).
  - NIST is standardizing **Quantum-Safe** alternatives (e.g., Lattice-Based Crypto).

📌 **RSA is still widely used but is gradually being replaced by ECC and post-quantum cryptographic algorithms.**

---

### **RSA Calculation Example Using Small Prime Numbers**  

To make RSA calculations easier to follow, let’s use **small prime numbers** instead of large ones.

---

## **Step 1: Choose Two Prime Numbers (`p` and `q`)**
We choose two small prime numbers:

\[
p = 3, \quad q = 11
\]

📌 **Why Prime Numbers?**  
- RSA security relies on the difficulty of **factoring large numbers**.
- If `p` and `q` were **not prime**, `n` would have additional divisors, making it easier to break.
- The totient function **φ(n)** depends on `p` and `q` being prime.

---

## **Step 2: Compute `n` (Modulus)**
Multiply `p` and `q`:

\[
n = p \times q = 3 \times 11 = 33
\]

---

## **Step 3: Compute Euler’s Totient Function `φ(n)`**
Euler’s Totient Function is:

\[
\phi(n) = (p - 1) \times (q - 1)
\]

\[
\phi(33) = (3 - 1) \times (11 - 1) = 2 \times 10 = 20
\]

---

## **Step 4: Choose the Public Exponent `e`**
`e` must be:
- **1 < e < φ(n)**
- **Coprime with φ(n)** (i.e., `gcd(e, φ(n)) = 1`)

Let’s choose:

\[
e = 7
\]

Check if `gcd(7, 20) = 1`:

- Factors of `7` → `{1, 7}`
- Factors of `20` → `{1, 2, 4, 5, 10, 20}`
- **Only common factor is 1**, so **7 is a valid choice for `e`**.

---

## **Step 5: Compute the Private Exponent `d`**
`d` is the modular inverse of `e` **mod φ(n)**:

\[
d = e^{-1} \mod \phi(n)
\]

This means:

\[
d \times e \equiv 1 \mod 20
\]

We solve:

\[
7 \times d \equiv 1 \mod 20
\]

Trying values for `d`:

\[
(7 \times 3) = 21 \quad \text{(21 mod 20 = 1)}
\]

So,

\[
d = 3
\]

---

## **Step 6: RSA Key Pairs**
- **Public Key:** `(e, n) = (7, 33)`
- **Private Key:** `(d, n) = (3, 33)`

---

## **Step 7: Encryption**
Let’s encrypt a small message **M = 5**.

Encryption formula:

\[
C = M^e \mod n
\]

\[
C = 5^7 \mod 33
\]

Calculating step by step:

\[
5^7 = 5 \times 5 \times 5 \times 5 \times 5 \times 5 \times 5 = 78125
\]

\[
78125 \mod 33 = 14
\]

So, the **ciphertext (C) = 14**.

---

## **Step 8: Decryption**
To decrypt, use:

\[
M = C^d \mod n
\]

\[
M = 14^3 \mod 33
\]

Calculating step by step:

\[
14^3 = 14 \times 14 \times 14 = 2744
\]

\[
2744 \mod 33 = 5
\]

We get back the **original message `M = 5`** ✅

---

## **Why Only Prime Numbers?**
1. **Ensures `n` has only two prime factors (`p` and `q`)**  
   - If `p` and `q` were **not prime**, `n` would be easier to factor.
   
2. **Ensures Euler’s Totient Function `φ(n) = (p - 1)(q - 1)` works correctly**  
   - If `p` and `q` were not prime, `φ(n)` would be different and RSA wouldn’t work properly.

3. **Security**  
   - The difficulty of **factoring `n` into `p` and `q`** is what makes RSA secure.

---

### **Padding in Cryptography: Why and What?**

**Padding** in cryptography is **extra data added to a message** before encryption to:
1. **Prevent predictable ciphertext patterns** 🛡️  
2. **Ensure messages fit block sizes** (RSA operates on fixed-length blocks)  
3. **Add randomness for security** (some padding schemes add randomness to prevent attacks)  

---

## **Why Padding in RSA?**
RSA **encrypts data in fixed-size blocks** based on the key size.  
- **Example:** A **2048-bit RSA key** can encrypt a **256-byte message max**.  
- If the message is **too short**, we **pad** it to the correct length.  
- If it's **too long**, we must split it into chunks.

Without padding, **RSA encryption is deterministic** → meaning **the same input always produces the same ciphertext**, making it vulnerable to attacks (e.g., **chosen plaintext attacks**).

---

## **Types of Padding in RSA**
There are **three main padding schemes** used in RSA:

### **1️⃣ OAEP (Optimal Asymmetric Encryption Padding) → Modern & Secure**
✅ **Prevents deterministic encryption (i.e., same message = different ciphertext each time)**  
✅ **Protects against chosen-plaintext attacks**  
✅ **Uses a hash function (e.g., SHA-256) for randomness**  

🔹 **Used for RSA Encryption (Secure, Recommended)**
```go
ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, message, nil)
```

---

### **2️⃣ PKCS#1 v1.5 Padding → Older, Less Secure**
✅ **Still widely used**  
❌ **Vulnerable to padding oracle attacks** (e.g., Bleichenbacher’s attack)  

🔹 **Used in RSA Encryption (but less secure than OAEP)**
```go
ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, message)
```
🔹 **Used in RSA Signing**
```go
signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
```

---

### **3️⃣ PSS (Probabilistic Signature Scheme) → Secure for RSA Signing**
✅ **Recommended for RSA signatures**  
✅ **Uses randomization to prevent attacks**  
✅ **Stronger than PKCS#1 v1.5 signatures**  

🔹 **Used for Signing**
```go
signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, hashed[:], nil)
```
🔹 **Used for Verification**
```go
err := rsa.VerifyPSS(publicKey, crypto.SHA256, hashed[:], signature, nil)
```

---

## **Comparison of RSA Padding Methods**
| **Padding**  | **Used For**        | **Security** | **Attack Resistance** | **Recommended?** |
|-------------|--------------------|-------------|----------------------|----------------|
| **OAEP**    | Encryption         | ✅ Strong  | ✅ Protects against chosen-plaintext attacks | ✅ Yes |
| **PKCS#1 v1.5** | Encryption & Signing | ❌ Weaker  | ❌ Vulnerable to Bleichenbacher attacks | ❌ No (Use OAEP for encryption, PSS for signing) |
| **PSS**     | Signing            | ✅ Strong  | ✅ Protects against signature forgeries | ✅ Yes |

---
