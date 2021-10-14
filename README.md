# WASM  

## Table of contents
- WASM
- Wasmtime
- Wasmer
- (WASI)
- (WAGI)

## Notes 
`/usr/local/go/misc/wasm/go_js_wasm_exec out/cli/html2goapp-cli.js-wasm.wasm -component PF4Tabs -src example/index.html -pkg example > example/index.go
`
## WebAssembly - an Overview
- "Binary instruction format for a stack-based virtual machine. Wasm is designed as a **portable compilation target** for programming languages, enabling deployment on the web for client and server applications." - webassembly.org
- Efficient and fast, Safe, Open and debuggable, Part of the open web platform 
    - Check those for validity
- The standard is getting developed by the W3C Community Group as well as a W3C Working Group. 
- WebAssembly defines a portable binary-code format for executable programs, and a corresponding textual assembly language, as well as interfaces for facilitating interactions between such programs and their host environment.
- WebAssembly Text
    ```
    (module
  (func $multiply (param $lhs i32) (param $rhs i32) (result i32)
    get_local $lhs
    get_local $rhs
    i32.mul)
  (export "multiply" (func $multiply))
    )
    ```
    - WAT is rarely used, as WASM is a compilation target for more common languages like C, C++, Rust, C#, Go, Python and many more. Even Java has the option to be compiled to WebAssembly using TeaVM (Java, Kotlin, Scala -> Java Bytecode -> WebAssembly/JavaScript).
