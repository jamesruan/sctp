# sctp
SCTP (Stream Control Transmission Protocol) in golang

This is a userland implementation of SCTP in pure golang.

Status: work in progress

## Features

- [RFC4960 Stream Control Transmission Protocol](https://tools.ietf.org/html/rfc4960)
  - [ ] 3. Packet Format
    - [X] 3.1 SCTP Common Header Format
    - [X] 3.2 Chunk Field Format
      - [X] 3.2.1 Optional/Variable-Length Parameter Format
      - [X] 3.2.2. Reporting of Unrecognized Parameters
    - [ ] 3.3 Chunk Definition
      - [ ] 3.3.1. Payload Data (DATA) (0)
      - [ ] 3.3.2. Initiation (INIT) (1)
  - [X] Appendix B: CRC32c Checksum Calculation

- [RFC1982 Serial Number Arithmetic](https://tools.ietf.org/html/rfc1982)
  - [X] Addition
  - [X] Comparison
