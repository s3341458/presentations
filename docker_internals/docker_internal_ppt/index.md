Docker Internals

* Why it is just a illusion?

---

Main Points

* Namespace
* Overlay file system

---

Uncovered but related

* Cgroups
* Linux virtual networks

---

Namespace 1 (Question)

* How many programs are running on my machine?

---

Namespace 1 (Answer)

* Only one! Operating System!

---

Namespace 2 (Process)

* Programs are binary code (eg 1)
* Programs load and run by OS as process
* Programs are commands and systemcalls

---

Namespace 3 (Virtual machine system calls)

* Virtual machine process can load and execute OS
* system calls -> VM OS -> VM process -> OS

---

Namespace 4 (kernel magic)

* OS can directly tricks system calls
* System calls -> OS (much simpler)

---

Namespace 5 (examples)

* Let's see examples (eg 2, eg 3)

---

Overlay filesystem 1 (Makes things portable)

* It is wrapper beyond native filesystem
* Can merge folders into one folder

---

Overlay filesystem 2 (examples)
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
