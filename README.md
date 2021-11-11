# WASM  
## WebAssembly
- "Binary instruction format for a **stack-based virtual machine**. WebAssembly is designed as a **portable compilation target** for programming languages, enabling deployment on the web for client and server applications." - webassembly.org
- Add "portable compilation target" graphic (Left side langs, mid wasm, right side target architectures)
- WASM's website claims that wasm is efficient and fast, safe, open and debuggable and part of the open web platform. In the following, we will check those statements for their validity
- The standard is getting developed by the W3C Community Group as well as a W3C Working Group and is supported by all major browser vendors. 
- WebAssembly defines a portable binary-code format for executable programs, and a corresponding textual assembly language, known as WebAssembly Text, as well as interfaces for facilitating interactions between such programs and their host environment.
- So how does it work? - You just need to write a program in a dedicated language; when compiling, specify `wasm` as the compilation target. In the first demo, we can look at how easy it is and what components are necessary to execute WebAssembly binaries in the browser.
-  Maybe you are currently thinking about how this thing is trying to take JavaScripts place but this is actually not the case. WASM is created to peacefully coexist next to JavaScript. That's why you can call WASM functions out of JavaScript and also JavaScript functions from WASM using the WebAssembly JavaScript APIs.
- MDN Web Docs summarizes WASM as follows: "WebAssembly is a new type of code that can be run in modern web browsers - it is a low-level assembly-like language with a compact binary format that runs with near-native performace and provides languages such as C/C++, C# and Rust a compilation target so that they can run on the web. It is also designed to run alongside JavaScript, allowing both to work together.
- Run code on the web near native speed, run client apps on the web that previously couldn't have done so. 
- For example can Doom be run in the browser using wasm. That would have probably surprised you a frew years ago before Doom ran on a pregnancy test or an IKEA lamp.

## Stack-based Virtual Machines 
- Abstraction of a computer that emulates a real machine
- Generally built as an interpreter of a special bytecode, that translates it in real time for execution on the CPU.
- Famous stack-based VM's are the JVM and the CLR (Common language runtime, base for .NET)
- To add two numbers in a stack VM, the program will push the first number to the stack, push the second and then execute some form of the special instruction add, that will pop the first two elements of the stack and replace them with their sum
- two main data structures for a stack-based VM are the code listing, with the instruction pointer, and the stack data, thich is accessed only via the stack pointer. By adding some external memory (heap), real languages like Java or **WebAssembly** (is this the right choice of word here?) can be created. 

## WAT - WebAssembly Text Format
- Low level, assembly-like language
- TODO: Add an example in go code and in WebAssembly
```lisp
(module
  (func $add (param $lhs i32) (param $rhs i32) (result i32)
    local.get $lhs
    local.get $rhs
    i32.add)
  (export "add" (func $add))
) 
```
- Code section, Function incl. Code section, Export section
- WASM is not primarily intented to be written by hand
- WASM can be importet into a web (or Node.js) app, exposing WebAssembly functions for use via JavaScript 

## The Web Platform
- The web platform can be thought of as having two parts:
  - A virtual machine (VM) that runs the Web app's code, e.g. the JavaScript code that powers your apps.
  - A set of WebAPIs that the Web app can call to control web browser/device functionality and make things happen (DOM, CSSOM, WebGL, IndexedDB, WebAudio API, etc.)
- Historically, the VM has been able to load only JavaScript. This became a bottleneck when more intensive use cases like 3D games, Virtual and Augmented Reality, computer vision etc. entered the space

## WebAssembly key concepts
- Several key concepts needed to understand how WebAssembly runs in the browser at near native speed. See those concepts in the WebAssembly JavaScript API (https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/WebAssembly)
  - Module: a WASM binary that has been compiled by the browser into executable machine code.
  - Memory: a resizable ArrayBuffer that contains the linear array of bytes read and written by the WASM's low-level memory access instructions.
  - Table: a resizable typed array of references (e.g. to functions) that could not otherwhise be stored as raw bytes in Memory (for safety and portability reasons)
  - Instance: a module paired wit hall the state it uses at runtime including a memory, table and set of importet values. (An instance is like an ES2015 module that has been loaded into a particular global with a particular set of imports)
- The JS API provides developers with the ability to create modules, memories, tables and instances. Given a WebAssembly instance, JavaScript code can synchronously call its exports, which are exposed as normal JavaScript functions. Arbitrary JavaScript functions can also be synchronously called by WebAssembly code by passing in those JavaScript functions as the import to a WebAssembly instance.
- Since JavaScript has complete control over how WebAssembly code is downloaded, compiled and run, JavaScript developers could even think of WebAssembly as just a JavaScript feature for efficiently generating high-performance functions.

## Workflow
- Emscripten is able to take any C/C++ source code and compile it into a .wasm module, plus the necessary JavaScript "glue" code for loading and running the module, and an HTML document to display the results of the code.

![alt text](/assets/img/emscripten-diagram.png "Logo Title Text 1")

1. Emscripten feeds the C/C++ into clang + LLVM. 
2. Emscripten transforms the compiled result of clang + LLVM into a .wasm binary.
3. By itself, WebAssembly cannot currently directly access the DOM; it can only call JavaScript, passing in integer and floating point primitive data types. Thus, to access any Web API, WebAssembly needs to call out to JavaScript, which then makes the Web API call. Emscripten therefore creates the HTML and JavaScript glue code needed to achieve this.

**There are plans to allow WASM to call Web APIs directly**

### Javascript glue code
- Emscripten implements popular C/C++ libraries like OpenGL, SDL, OpenAL, and parts of POSIX. Thse libraries are implemented in terms of Web APIs and thus each one requires some JavaScript glue code to connect WebAssembly to the underlaying Web API. 
- Part of the glue code is implementing the functionality of each respective library used by the C/C++ code. The glue code also contains the logic for calling the above-mentioned WebAssembly JavaScript APIs to fetch, load and run the .wasm file.
- The generated HTML document loads the JavaScript glue file and writes stdout to a `<textarea>`. If the application uses OpenGL, the HTML also contains a `<canvas>` element that is used as the rendering target. It's very easy to modify the Emscripten output and turn it into whatever web app you require.

## Wasmer
- fast and secure WebAssembly runtime that enables super lightweight containers to run anywhere: from Desktop to the Cloud, Edge and IoT devices.
- Secure by default: No file, network, or environment access, unless explicitly enabled. 
- Execute *.wasm on your computer
- installation commands
- `wapm run cowsay Hello World`
- E.g. wasmer-go
  - wasmer-go embeds the Wasmer runtime (written in Rust) inside Go. It provides most of the Wasmer features. If you are using Go, then wasmer-go is a good solution for you. If you are using Rust, then Wasmer itself is the best solution.
  - It is performant, and it is production ready. The Wasmer runtime (in Rust) provides better runtime performance than wasmer-go; the latter goes through cgo to hit the wasmer-c-api; there is a penalty for each call due to Go/cgo overall design, compared to using the Rust API directly.
- Run some of the wasmer examples in rust and explain the workflow there

## Wasmtime
- better user experience IMHO 
- Standalone JIT-style runtime for WebAssembly, using Cranelift
- Demo

## WASI 
- The WebAssembly System Interface
- Modular system interface for WebAssembly
- "Developers start to push WebAssembly beyond the browser, because it provides a fast, scalable and secure way to run the same code across all machines." - Lin Clark
- WebAssembly is a language for a conceptual machine, not a physical one. This is why it can be run across a variety of different machine architectures. 
- This implies that WASM needs a system interface for a conceptual operating system, not any single one. We want to target a variety of plattforms again.

### System interface
- Let's first talk about what a system interface is in the first place. 
- Imagine the kernel and the operating system. If we want to open a file we need to make a syscall. The kernel then has time to check if the user has the necessary permissions. The operating system makes these system calls available. We also don't need multiple implementations of the code for every operating system as the standard library in your programming language most likely handles this. Otherwhise you would need to know which system you are targeting while programming and therefore implement your functions differently if the interface is a different one. 

### What was the state of the art solution before WASI
- First, there was Emscripten
- It emulates POSIX on the web (a programmer can use functions from the C standard library)
- To do this, Emscripten created its own version of libc. Partly this implementation is compiled into the WASM module and partly, it is contained in the JS glue code.
- The JS glue code would then call into the browser, which would then talk to the OS.
- (Show picture to improve understanding)
- The runtimes also implemented all the JS glue code functions so WASM compiled with Emscripten could be executed without the browser.
- (Show picture of emulation of emulation)
- All we have is an emulation of an emulation when trying to run WASM outside of the browser and want to interact with the system.

### What does WASI do now?
- WASI wants to uphold the most important goals of WASM:
  - portability
  - security
- As we saw the old approach couldn't totally do that for us with the emulation in an emulation

#### Portability
- Posix provides source code portability because you can compile the same code with different versions of libc to target different machines.
- One more step, compile once and run on different platforms. The goal is a portable binary. (Use images from mozilla)

#### Security 
- When you ask the operating system for some input or output, it needs to determine if you are allowed the information.
- This is handled with access control based on ownership and groups.
- That protects users from each other and was useful back in time when systems often were multi-user systems
- But nowadays most systems are single user systems. Nowadays it's actually way more important to protect the user from itself.
- We are pulling in so much third pary code of unknown trustworthyness that we need to have control over what it does.
- The solution is simple and probably known to everyone in this room. Sandboxing.
- That basically means that code can't directly talk to the OS. But instead we create a sandbox, we lock the program in and expose some functions to it with just the functionality we allow. 
- This isn't automatically a secure solution because the host still can possibly grant full permissions to the sandbox in which case we are no better than in the first solution. But at least we have the option of being secure. 
- Files
  - In the normal OS, if code needs to open a file, it calls the open function containing a string with the path to the file. This succeeds if we have the necessary user permissions. As we now know this is not secure enough.
  - Instead, if you call a function interacting with files in WASI, you have to pass a file descriptor, which has permissions attached to it. This way you can't randomly open a file. Instead the code can only operate in the directories it was given. 
  - So the runtime passes in the file descriptors that an app can use at the top level code. The code itself can then propagate those file descriptors through the rest of the modules as needed. (Insert cute picture here) 
- TODO Demo

## WAGI 
- WebAssembly Gateway Interface
- Write microservices and webapps 
- Experimental 
- "Write HTTP handlers in WebAssembly with minimal amount of work"

## asm.js
- "asm.js is a subset of JavaScript designed to allow computer software written in languages such as C to be run as web applications" - https://en.wikipedia.org/wiki/Asm.js

## Major Goal: Efficient and Fast
- Fast is relative
- Fast than Javascript: Statically typed and simple to optimize 

### Startup time
- It's hard to improve the size of a minified gzipped JavaScript file. WASM is designed with size in mind which results in usually 10%-20% smaller gzip files.
- WASM is parsed way faster than JavaScript, which results out of the fact that binary formats are usually faster to parse. WASM also makes it easy to parse functions in parallel
- There are of course other factors playing into the startup time but file size and parsing time are the two factors which can't be avoided or improved in some other way

### CPU features
- WASM is able to use many CPU features which e.g. ASM did not: 
  - 64-Bit integers -> Operations up to 4x faster.
  - Load and store offsets
  - CPU instructions like copysign, popcount which can be useful in special cases
- These CPU features lead to a speed improvement of ~5% compared to ASM
- Could be even faster in the future when WASM uses SIMD (Single-Instruction, Multiple-Data) which could then perform the same instruction on multiple data points at the same time

### Toolchain improvements
- Improving the compiler 
  - Improving the Relooper Algorithm 
    - Translate basic blocks into higher level structures using loops and ifs
    - Is needed to compile C++ into JS
    - Can't use goto untyped in a stack based virtual machine, which can be included in basic blocks
    - Example of a basic block to set a 10*10 matrix to an identity matrix by https://www.geeksforgeeks.org/basic-blocks-in-compiler-design/:
    ```
    1)  i=1        //Leader 1 (First statement)
    2)  j=1             //Leader 2 (Target of 11th statement)
    3)  t1 = 10 * i     //Leader 3 (Target of 9th statement) 
    4)  t2 = t1 + j
    5)  t3 = 8 * t2
    6)  t4 = t3 - 88
    7)  a[t4] = 0.0
    8)  j = j + 1
    9)  if j <= goto (3)       
    10) i = i + 1                    //Leader 4 (Immediately following Conditional goto statement)
    11) if i <= 10 goto (2)
    12) i = 1                        //Leader 5 (Immediately following Conditional goto statement)
    13) t5 = i - 1                   //Leader 6 (Target of 17th statement) 
    14) t6 = 88 * t5
    15) a[t6] = 1.0
    16) i = i + 1
    17) if i <= 10 goto (13)
    ```
- Another ~5% speedup compared to ASM

### Conclusion
- Faster than ASM
- ASM was not an actual standard and was not jointly designed by all major browsers 

## Resources
- https://app.element.io/#/room/!zfXkSajYpjFUicXtCA:matrix.org
- https://webassembly.org/
- https://en.wikipedia.org/wiki/WebAssembly
- https://developer.mozilla.org/en-US/docs/WebAssembly/Understanding_the_text_format
- https://developer.mozilla.org/en-US/docs/WebAssembly
- https://doom.fandom.com/wiki/Wasm-doom
- https://www.popularmechanics.com/science/a33957256/this-programmer-figured-out-how-to-play-doom-on-a-pregnancy-test/
- https://developer.mozilla.org/en-US/docs/WebAssembly/Concepts
- https://emscripten.org/
- https://github.com/wasmerio/wasmer/tree/master/examples
- https://github.com/wasmerio/wasmer-go
- https://spectrum.chat/wasmer/runtime/wasmer-go-vs-wasmer-rust~4f8d36bd-fb3d-4c6b-9bc2-4e04559ab038
- https://github.com/bytecodealliance/wasmtime
- https://hacks.mozilla.org/2019/03/standardizing-wasi-a-webassembly-system-interface/
- https://events.mi.hdm-stuttgart.de/2021-04-09-5-developers-day/DEATH%202%20JAVASCRIPT
- https://hacks.mozilla.org/2017/03/why-webassembly-is-faster-than-asm-js/
- http://mozakai.blogspot.com/2012/05/reloop-all-blocks.html
- https://en.wikipedia.org/wiki/Basic_block
- https://www.geeksforgeeks.org/basic-blocks-in-compiler-design/

## Interesting Links
- https://wasdk.github.io/WasmFiddle/
- https://anonyco.github.io/WasmFiddlePlusPlus/
- https://mbebenita.github.io/WasmExplorer/
- https://www.assemblyscript.org/ (If you are a web developer wjp wants to try wasm without needing to learn languages like C, Rust etc. AssemblyScript compiles a strict variant of TypeScript to WebAssembly, allowing to keep using TS compatible tooling such as Prettier, ESLint, intellisense etc.)

## TODO
- Theory on how wasmer and wasmtime work
- WAGI
- Explain the necessity of every component used to run WASM in the browser
- Stack based virtual machine 
- Java comparison (Java bytecode, JVM)
- Call javascript from wasm and the other way around
- actually benchmark myself
- a wasm binary that has been compiled **by the browser**
- "native speed"
- Emscripten
- Why can we only call exported functions and not run the code itself
- Use the wasmtime wasm binary in the browser as well
- Actually build a tool using nearly all the possible tools simultaneously
- Create a overview / Table of Contents