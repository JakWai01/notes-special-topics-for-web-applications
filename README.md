# WASM  
## WebAssembly
- "Binary instruction format for a **stack-based virtual machine**. WebAssembly is designed as a **portable compilation target** for programming languages, enabling deployment on the web for client and server applications." - webassembly.org
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
```lisp
(module
  (func $add (param $lhs i32) (param $rhs i32) (result i32)
    local.get $lhs
    local.get $rhs
    i32.add)
  (export "add" (func $add))
) 
```
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


https://developer.mozilla.org/en-US/docs/WebAssembly/Concepts

## Resources
- https://app.element.io/#/room/!zfXkSajYpjFUicXtCA:matrix.org
- https://webassembly.org/
- https://en.wikipedia.org/wiki/WebAssembly
- https://developer.mozilla.org/en-US/docs/WebAssembly/Understanding_the_text_format
- https://developer.mozilla.org/en-US/docs/WebAssembly
- https://doom.fandom.com/wiki/Wasm-doom
- https://www.popularmechanics.com/science/a33957256/this-programmer-figured-out-how-to-play-doom-on-a-pregnancy-test/
- https://developer.mozilla.org/en-US/docs/WebAssembly/Concepts

## Interesting Links
- https://wasdk.github.io/WasmFiddle/
- https://anonyco.github.io/WasmFiddlePlusPlus/
- https://mbebenita.github.io/WasmExplorer/
## TODO
- Wasmtime
- Wasmer
- WASI
- WAGI
- Explain the necessity of every component used to run WASM in the browser
- Stack based virtual machine 
- Java comparison (Java bytecode, JVM)
- Call javascript from wasm and the other way around
- actually benchmark myself
- a wasm binary that has been compiled **by the browser**
- "native speed"