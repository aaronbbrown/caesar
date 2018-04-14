## Caesar Cipher Demo

This project encrypts & decrypts a modified version of the Caesar cipher. It also
implements a rudimentary, naive cracking algorithm to crack Caesar ciphers using
an English dictionary.

## Encryption Algorithm

For keys under 26, the encryption algorithm is a typical Caesar cipher shift
algorithm, otherwise known as ROT _n_. Each letter in each word is shifted by
_key_ positions, so if the key is 1, `a` becomes `b`, `y` becomes `z` and `z`
becomes `a`.

For Example:
```
$ echo "abc xyz"| bin/caesar encrypt --key 1
bcd yza
```

For keys over 25, the key is converted into base 25 and we get an array of shift
values.  For example, a key of 27 converts to offsets (1,2). This would offset
the first char by 1, second char by 2, third char by 3, etc.  Unknown characters
are skipped.

For example:
```
$ bin/caesar offset --key 27
Key:  27 [1 2]


$ echo "abc xyz" | bin/caesar encrypt --key 27
bdd zzb
```

## Cracking

The project includes a rudimentary cracking algorithm that iterates through keys
and if a certain percentage of the words in the decrypted text are included in
the dictionary (`wordlist.txt`), the text is considered decrypted.

```
$ bin/caesar encrypt --key 1234 < text.txt | bin/caesar crack --max-attempts 10000 --parallelism 4 --percent-words 30
Key:  1234 [1 24 9]
Message:
the action of a caesar cipher is to replace each plaintext letter with a different one a fixed number of places down the alphabet. the cipher illustrated here uses a left shift of three, so that (for example) each occurrence of e in the plaintext becomes b in the ciphertext.
in cryptography, a caesar cipher, also known as caesar's cipher, the shift cipher, caesar's code or caesar shift, is one of the simplest and most widely known encryption techniques. it is a type of substitution cipher in which each letter in the plaintext is replaced by a letter some fixed number of positions down the alphabet. for example, with a left shift of 3, d would be replaced by a, e would become b, and so on. the method is named after julius caesar, who used it in his private correspondence.[1]

the encryption step performed by a caesar cipher is often incorporated as part of more complex schemes, such as the vigenère cipher, and still has modern application in the rot13 system. as with all single-alphabet substitution ciphers, the caesar cipher is easily broken and in modern practice offers essentially no communication security.

```

## Building

Run `make build`.  This will create `bin/caesar`
