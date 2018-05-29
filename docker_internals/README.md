# What is the purpose of this talk.
Let you understand the basic ideas about how did the docker engine(dockerd)
was written, so every time when you use docker you do not feel like you working
on mysterious black magic.
I will not cover too much the details about code implementaions, if you guys want
to know more details please follow https://github.com/s3341458/mydocker

In my talk, I will emphasize idea > mechanism > implementaion details

# what will be definitely describe in this talk
core container isolation technology namespace
core container mounting overlayfs

basically explains:
how docker build works
how docker run works
how docker stop works
how docker restart works
how docker rm works
how docker exec works
how docker export works

# what will be just just briefly mentioned in this talk
core container resource control technology cgroups
core container network technology linux bridge and iptables

the material mostly comes from a 195 pages of a chinese book
"Write Docker by yourself" https://github.com/xianlubird/mydocker

if you visit the repo, you can see I have done tiny little contribution
to its codebase.

I will use about 25 minutes to describe the core ideas of this book, and use
5 minutes to summarize the my opinion. And leave about 10 min for Q&A

I want to make the presentation more interactive, so please feel free to
interrupt me.


# Namespace [15min]
In order to understand namespace, firstly we should first start to get a basic
understanding virtualization technology. I will do a over simplified discussion
about this topic.

[check if every one has used virtual machine]

[section] start with a tricky question, how many programs are running on my computer.

answer is one! that is my operating system, for this computer this is my linux
kernel.
[eg1, show a binary example]
You can see a program is actually just a bunch of binary code. Those binary codes
are actually just commands that can be only understanded by CPU.

Operating system is also a chunk of excutable binary code, but it is special.
it can read other programs(in binary format) and make those programs running as part of it(a process). In order to be read
and understanded by OS, the binary codes file should formatted undered the OS
covention, that is why a excutable program can run on linux can not run on windows.

In order manage those running programs, OS create a concept called "process",
a process is a program that is excuting by the OS. From OS point of view, it is a
a data structure(bunch of data) which is recorded the excuting code of the program,
what is its state and what resource is using.

A program that can run on the OS are formatted by two kinds commands, one is normal command
which OS can directly taken without any conversion and can be throw to CPU directly.
(mathmatic operaters, if else, for loop stuff are all belongs to this kind)
The other kind of commands are system calls, which will trigger operating system to do some special works.  (eg, file open function, sleep function is just a wrapper for those commands).
If you have not heard system calls you might heard of ABI, yes actually they are the same thing.

Those system calls are the only way that process can communicate with operating
system. By system calls, process can ask OS do something for them, eg, "give me a file",
"open up the camera" or "even put me in sleep". The thing need to notice is that
from cpu point of view dealing with system calls takes much longer time than normal
command (Context switching involved). All virtualization technologies are actually play
tricks on those system calls.

From the process point of view, it can only know which world it is in from the response
of those system calls.

Normally, a virtual machine as a program is also a bunch of binary code which will
be loaded to OS and executed as a "process". This process is a little bit special that
It can actually load binary code of a operating system and make it run as part of it,
the OS runs on virtual machine let us call it virtual OS. This virtual OS like a real OS can load normal
program and execute it as a part of it. So when process that runs on virtual OS, It makes a system call
will be made towards virtual OS in virtual machine process, then the underly virtual machine process
may need to make several system calls to its host OS. As you can see, things become nasty
and expensive. But from this way, process makes system calls to the virtual operating system,
instead of native operating system, so an isolation has been achieved.
[graph needs to be shown]

So the OS developers firstly come out with a idea, why can not we enhance the native
OS so that it can respond the system calls of the process more cleverly. It can put processes
in different groups so the process within the same group can feels like they are in a separate
OS.

Yeah this the idea of namespace, internally it is just a label on process data structure,
used to separate processes into groups. So when operating system respond to a system call,
before directly respond it from the current context. It check the context of process namespace
and respond accordingly.

The way to create a container is just when you create a process, you make a fork system call and indicate the system
that the process is in another namespace
lets see the code exampe.
[eg2, code example of namespace]

there are actually six namespace modules available for linux.
UTS namespace: used to isolate hostnames
IPC namespace: used to isolate POSIX message queues, or in a simpler this isolate the communication for processes.
PID namespace: used to isolate the ids of processes
mount namespace: used to isolated the mount point each process can see, in a simpler words this isolate the file hierachies that each process can see.
network namespace: used to isolated networks devices each process can use.
user namespace: used to isolate users that each process recongized. For example, user can be as a root user in namespace 2 but a normal user in namespace 1.

Actually you can see namespace in linux pesudo file system.
[eg3, check namespace]

# Overlay filesystem https://wiki.archlinux.org/index.php/Overlay_filesystem
Once we understand namespace understand overlay filesystem become much more straightforward.
Like process, folder hierachies in kernel is just another data structure(tree liked), when a process want
to find a file by give a file path, kernel will search the data structure whether the path folder exists and contains file with the
same name and give file reference to the process. What overlay filesystem is doing is that it wrap up normal file systems data structure.
So when process want to find a file by give a file path, operating system will search some candidates paths and try to find the file.
This sounds a little too abstract, let me describe it in shorter word, overlay filesystem is a wrapper on normal file system.
so it can combine multiple folders and showed it in one folder in a certain manner. In a even more simpler words,
it makes you feel you are visiting one folder but actually you are visit multiple folders. Still feel weird, let me show you
some examples.
[eg4, overlay file system demo]{8min}
see images folders, containers folders and volume folder, and merged folder.
how images folder has been build. (docker build will be covered here)

Now you are understand overlay fs.
things become easy.

what docker run has done?
start a new process with new namespace, and use merged folder as root folder the mount namespace of new process.
Keep a record about what has happened.

what docker stop has done?
stop the process, but still keep the record and container folder.

what docker restart has done?
Restart the process based on the previous start record, use undeleted container folders.

what docker rm has done?
stop the process, and delete the record and container folder.

what docker exec has done?
starting a new process and put it in the namespace of previous container and then execute the command

what docker export has done?
copy the merged folder, and export it.

# Q&A time.
