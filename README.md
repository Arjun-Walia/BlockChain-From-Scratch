---

This project is a simple implementation of a blockchain in Golang. It includes core blockchain features such as blocks, transactions, wallets, Proof of Work (PoW), and a REST API for interaction. The project is designed to help you understand distributed systems, cryptography, and consensus mechanisms.

---

## Table of Contents
1. [Features](#features)
2. [Project Structure](#project-structure)
3. [Installation](#installation)
4. [Usage](#usage)
5. [API Endpoints](#api-endpoints)
6. [Consensus Mechanisms](#consensus-mechanisms)
7. [Contributing](#contributing)
8. [License](#license)

---

## Features
- **Blockchain Core**:
  - Block creation and chaining using cryptographic hashes.
  - Proof of Work (PoW) consensus mechanism.
- **Transactions**:
  - Secure transactions using ECDSA (Elliptic Curve Digital Signature Algorithm).
  - Merkle tree for efficient transaction validation.
- **Wallet System**:
  - Public/private key generation.
  - Digital signatures for transaction security.
- **P2P Network**:
  - Decentralized peer-to-peer communication.
- **REST API**:
  - Interact with the blockchain (view blocks, create transactions, mine blocks).

---

## Project Structure
```
blockchain-from-scratch/
├── go.mod
├── go.sum
├── main.go
├── block/
│   ├── block.go
│   └── block_test.go
├── blockchain/
│   ├── blockchain.go
│   └── blockchain_test.go
├── transaction/
│   ├── transaction.go
│   └── transaction_test.go
├── wallet/
│   ├── wallet.go
│   └── wallet_test.go
├── network/
│   ├── p2p.go
│   └── p2p_test.go
├── api/
│   ├── api.go
│   └── api_test.go
├── consensus/
│   ├── pow.go
│   └── pos.go
└── utils/
    ├── merkle.go
    └── crypto.go
```

---

## Installation

### Prerequisites
- Go 1.21 or higher.
- Git (optional, for cloning the repository).

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/Arjun-Walia/blockchain-from-scratch.git
   cd blockchain-from-scratch
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Build the project:
   ```bash
   go build -o blockchain .
   ```
4. Run the application:
   ```bash
   ./blockchain
   ```

---

## Usage

### Running the Blockchain
1. Start the blockchain node:
   ```bash
   ./blockchain
   ```
2. Use the REST API to interact with the blockchain (see [API Endpoints](#api-endpoints)).

### Creating a Wallet
1. Generate a new wallet:
   ```bash
   curl -X POST http://localhost:8080/wallet/create
   ```
2. Use the wallet address and private key to send transactions.

### Mining a Block
1. Add transactions to the pending pool:
   ```bash
   curl -X POST -H "Content-Type: application/json" -d '{
     "sender": "sender_address",
     "receiver": "receiver_address",
     "amount": 1.0
   }' http://localhost:8080/transaction
   ```
2. Mine a new block:
   ```bash
   curl -X POST http://localhost:8080/mine
   ```

---

## API Endpoints
| Endpoint              | Method | Description                          |
|-----------------------|--------|--------------------------------------|
| `/blocks`             | GET    | Get the entire blockchain.           |
| `/block/{index}`      | GET    | Get a specific block by index.       |
| `/transaction`        | POST   | Create a new transaction.            |
| `/mine`               | POST   | Mine a new block.                    |
| `/wallet/create`      | POST   | Generate a new wallet.               |
| `/wallet/balance`     | GET    | Check the balance of a wallet.       |

---

## Consensus Mechanisms

### Proof of Work (PoW)
- Miners solve a cryptographic puzzle to add a new block.
- Implemented in `consensus/pow.go`.

### Proof of Stake (PoS) (Optional)
- Validators are chosen based on the amount of cryptocurrency they hold.
- Implemented in `consensus/pos.go`.

---

## Contributing
Contributions are welcome! Follow these steps:
1. Fork the repository.
2. Create a new branch:
   ```bash
   git checkout -b feature/your-feature
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add your feature"
   ```
4. Push to the branch:
   ```bash
   git push origin feature/your-feature
   ```
5. Open a pull request.

---

## Acknowledgments
- Inspired by the original Bitcoin whitepaper by Satoshi Nakamoto.
- Built with ❤️ using Golang.

---