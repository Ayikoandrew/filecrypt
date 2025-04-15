# FileCrypt

A command-line tool for file encryption and decryption using AES-256.

## Usage

### Generate an encryption key
```bash
./bin/filecrypt -mode=genkey
```

### Encrypt a file
```bash
./bin/filecrypt -mode=encrypt -in=inputfile -out=encrypted.enc -key=YOUR_HEX_KEY
```

### Decrypt a file
```bash
./bin/filecrypt -mode=decrypt -in=encrypted.enc -out=decrypted.file -key=YOUR_HEX_KEY
```

## Example Usage

1. Generate a random 32-byte hex key:
```bash
openssl rand -hex 32
8a0da5163e913b317cd9b695e98a4322e2cda658e1bfc82bc18ef9121e484ddd
```

2. Encrypt a file:
```bash
./bin/filecrypt -mode=encrypt -in=password.txt -out=password.enc -key=8a0da5163e913b317cd9b695e98a4322e2cda658e1bfc82bc18ef9121e484ddd
```

3. Decrypt a file:
```bash
./bin/filecrypt -mode=decrypt -in=password.enc -out=decrypt.doc -key=8a0da5163e913b317cd9b695e98a4322e2cda658e1bfc82bc18ef9121e484ddd
```