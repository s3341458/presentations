Docker Internals

* Why it is just a illusion.

---

Main Points

* namespace
* overlay file system

---

Uncovered but related

* cgroups
* linux virtual networks

---

Namespace 1 (Question)

* How many programs are running on my machine?

---

Namespace 2 (Process)

* program are binary code (eg 1)
* program load and run by OS as process
* program are commands and systemcalls

---

Namespace 3 (Virtual machine system call)

* virtual machine execute OS
* system call -> VM OS -> VM -> OS

---

Namespace 4 (Namespace wrap systemcalls)

* OS can directly tricks systemcalls
* systemcalls -> OS (much simpler)

---

Namespace 5 (examples)

* Let's see examples (eg 2, eg 3)

---

Overlay filesystem 1

* It is wrapper beyond native filesystem
* can merge folders in to fold

---

Overlay filesystem 2
* let's see examples (eg 4)


---
describe examples
* docker build
* docker run
* docker stop
* docker restart
* docker rm
* docker volume
* docker export


---
Q&A
