# WASM  
## WebAssembly
- "Binary instruction format for a **stack-based virtual machine**. WebAssembly is designed as a **portable compilation target** for programming languages, enabling deployment on the web for client and server applications." - webassembly.org
- WASM's website claims that wasm is efficient and fast, safe, open and debuggable and part of the open web platform. In the following, we will check those statements for their validity
- The standard is getting developed by the W3C Community Group as well as a W3C Working Group. 
- WebAssembly defines a portable binary-code format for executable programs, and a corresponding textual assembly language, known as WebAssembly Text, as well as interfaces for facilitating interactions between such programs and their host environment.
- So how does it work? - You just need to write a program in a dedicated language; when compiling, specify `wasm` as the compilation target. In the first demo, we can look at how easy it is and what components are necessary to execute WebAssembly binaries in the browser.
-  Maybe you are currently thinking about how this thing is trying to take Javascripts place but this is actually not the case. WASM is created to peacefully coexist next to JavaScript. That's why you can call WASM functions out of Javascript and also Javascript functions from WASM.
## WAT - WebAssembly Text Format
```lisp
(module
  (func $add (param $lhs i32) (param $rhs i32) (result i32)
    local.get $lhs
    local.get $rhs
    i32.add)
  (export "add" (func $add))
) 
```

## TODO
- Wasmtime
- Wasmer
- WASI
- WAGI
- Explain the necessity of every component used to run WASM in the browser
- Stack based virtual machine 
- Java comparison (Java bytecode, JVM)
- Call javascript from wasm and the other way around