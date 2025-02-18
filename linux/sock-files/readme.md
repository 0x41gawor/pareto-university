In Linux, **`.sock` files** represent **Unix domain sockets**, which are used for **interprocess communication (IPC)** between processes running on the same machine.

---

### **Key Points about `.sock` files**:
- They act like endpoints for processes to communicate, similar to network sockets, but they operate **within the local filesystem**, not over a network.
- Faster than network sockets because they avoid the overhead of network protocols.
- Commonly used by system services and applications like:
  - **Docker** (`/var/run/docker.sock`)
  - **MySQL** or **PostgreSQL** for local database connections.
  - **Nginx** for FastCGI connections.
  
---

### **Where do you find them?**
- Typically located in `/tmp`, `/var/run`, or `/run` directories.
  
---

### **Example**:
- **`/var/run/docker.sock`** allows Docker CLI to communicate with the Docker daemon.
- **`/tmp/.X11-unix/X0`** is used by the X Window System for GUI-related IPC.

---

### **To view all `.sock` files**:
```bash
find / -type s 2>/dev/null
```

---

Would you like an example of how to use a `.sock` file, or need help with a specific one? 😊