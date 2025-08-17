
**FuzzIt: When programs say ‘whatever’, chaos begins.** 

## About

FuzzIt is a **directory/file fuzzer** written in Go. It sends HTTP requests to a target URL using a custom wordlist and reports the status codes. Useful for discovering hidden files, directories, and endpoints.  

- Written in Go (fast & portable)  
- Configurable speed & wordlist  
- Helps find hidden paths in web apps  

---

## Installation

```bash
git clone https://github.com/sudebozkurt/FuzzIt.git
cd fuzzit
go build -o fuzzit main.go
go run main.go -u https://target.com -txt wordlist.txt -s 50
