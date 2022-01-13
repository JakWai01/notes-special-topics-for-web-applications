# Script 

## Motivation

- **1995**: Javascript taucht auf
- **2008**: Erster Wendepunkt in der Performanceentwicklung - Just in Time Compiler: Kompilierung und Optimierung während der Runtime 
- Die Steigung hat dann mit NodeJS abgenommen, da jetzt auch mehr Performance für die komplexer werdenden Applikationen nötig werden
- **2017** bis heute: Haben wir einen weiteren Wendepunkt der Performance im Web durch WebAssembly erreicht? Diese Frage können wir hoffentlich im Laufe der Präsentation klären

## Definition

- WebAssembly ist ein Offener Standard, welcher ermöglicht, Software in irgendeiner Sprache zu schreiben und diese dann über den Browser, mit near-native-speed auszuführen. 
- Die Spezifikation enthält eine Low-Level language ähnlich zu Assembly, daher auch der Name, welche zu einer Binary kompiliert werden kann, welche dann auf allen modernen Browsern läuft.
- Man wird diesen Code allerdings nicht mit der mitgelieferten Assemblyartigen Sprache schreiben, sondern das Binary Instruction Format als Compilationtarget für andere Sprachen verwenden. 
- E.g. ein Spiel, welches mit Unity und C# erstellt wurde im Browser ausführen. Das wurde übrigens auch mit dem Spiel AngryBots beim ursprünglichen WebAssembly-Pitch gemacht.

- Um uns mit den Schlüsselkonzepten vertraut zu machen hangeln wir uns an der Definition von `webassembly.org` entlang

## Binary Instruction Format

- Binaries können grundlegend zweier Natur sein: Maschinencode oder Bytecode, welcher dann von einer virtuellen Maschine ausgeführt werden kann
- Im Falle WebAssembly redet man von einem Bytecode
- Binaries können verschiedene Instruction Set Architectures als Ziel haben. Euch ist da wahrscheinlich x86, ARM oder RISC-V bekannt
- WASM verwendet allerdings virtuelle instruktionen für eine konzeptionelle Maschine, um am Ende mit möglichst vielen der Instruction Set Architectures kompatibel zu sein
- Man muss sich das WebAssembly Binary Instruction Format also eher als überschneidung mehrerer Instruction Formate vorstellen, welches nicht einfach so auf ein einzelnes gemapped werden kann, da unsere Binary ja überall laufen soll

## Stack-Based Virtual Machine

- Eine Virtuelle Maschine, welches einen Stack anstatt Registers für Operationen verwendet
- Bekannte SVM sind die JVM (Java Virtual Machine) und die CLR (Common Language Runtime, .NET)
- Will man zwei Zahlen in einer Stack-Based-VM addieren, pusht man die beiden Zahlen mit der add Instruktion auf den Stack, die beiden Zahlen werden gepopped und mit ihrer Summe replaced
- Wenn der Browser den WebAssembly Code zum jeweiligen Maschinencode übersetzt, werden allerdings wieder Register verwendet. WASM spezifiziert keine Register, was dem Browser maximale Flexibilität die Register für diese Maschine möglichst Effizient zu allokieren

## Portable Compilation Target

- Beim kompilieren kann als target WASM ausgewählt werden, was dann erlaubt auf verschiedenen Betriebssystemen und Architekturen ausgeführt zu werden. Im Browser oder auch Nativ.
- Geringe Anforderungen an WASM execution environments: 
    - 8-bit Bytes
    - Adressierbare Bytes 
    - Little Endian 
    - ...

## Deployment on the Web Platform

- Web Platform kann in zwei Teilen beschrieben werden
    - VM/Engine, welche den Code ausführt. Beispiele dafür sind V8 in Chrome oder SpiderMonkey in Firefox
    - Web APIs welche weitere Browser- oder Gerätefunktionen freischalten (DOM, WebGL, Web Audio API etc.)
- Stand von Dezember 2021 war, dass 95% aller installierten Browser WebAssembly supporten. 
- Schließen sich nicht gegenseitig aus

## Definition

- Wenn wir uns jetzt nochmal die Definition anschauen merken wir, dass wir nun mit den Schlüsselkonzepten vertraut sind. Aber was können wir jetzt mit WASM machen? 

## Use Cases 

- Im Browser: 
    - VR and AR, da diese sehr viel Performance benötigen
    - Simulationen/Emulationen (DOSBox, QEMU, ...)
    - Remote Desktop 
    - Spiele 
    - Cloud IDEs
- Außerhalb des Browsers:
    - Serverseitige Anwendungen
    - Spieleplattformen 
- Projekte, welche gerade WASM verwenden: 
    - Blazor
    - AutoCAD
    - Figma
    - Jitsi für deren virtuelle Hintergründe
    - Zoom für ihr video encoding
    - D3wasm (Doom 3 im Browser, WebAssembly Demo)
    - go-app (Frontend mit Hilfe von WASM in Go schreiben)

In der WebAssembly spezifikation ist nicht nur das Binary Instruction Format definiert, sondern auch eine Assemblyartige Sprache, mit dem man WebAssembly Module schreiben kann. Und dabei handelt es sich um WebAssemblyText.

## WebAssembly Text

- Eine Menschenlesbare Text-Repräsentation von WASM-Binaries
- Konzipiert um in Editoren oder im Browser-Inspektor betrachtet zu werden
- .wat Dateien können mit Hilfe des WebAssembly Binary Toolkits (WABT) zu .wasm kompiliert werden
- Wer sich da noch mehr Informieren möchte, oder auch allgemein zu WebAssembly, die MDN Webdocs sind dafür meiner Ansicht nach die beste Quelle
- Im folgenden ein kleines Beispiel: 
    - Wir sehen hier eine simples WASM module mit einer add function, welche zwei Parameter nimmt und diese addiert
    - Module: ist einfach eine stateless WASM Binary mit einer gewissen Menge an Code. Angelehnt an ES2015 Module.
    - WebAssembly Text verwendet S-Expressions, welche ein altes Format ist um Baumstrukturen darzustellen
    - Function signature
    - Function body:
        - Stack Machine - Push zwei Values auf den Stack und addiere mit der i32 add Instruktion
        - i32 add: Poppe die zwei Values und Pushe das Ergebnis 
    - Export block: "add" ist der identifier, welcher von JavaScript verwendet wird und $add ist intern, um die zu exportierende Funktion festzustellen
- Diesen nennen wir jetzt ein exportiertes WASM Module mit einer add function
- Unsere exportierte Funktion könnte jetzt so aus JavaScript aufgerufen werden. Diese Funktion kompiliert und instanziiert ein WebAssemembly Module direkt von einer gestreamten Quelle. Das ist der Effizienteste Weg, ein WASM-Modul zu laden, da hier eben streaming verwendet wird und so eben nicht alles auf einmal heruntergeladen werden muss. Will man das machen gibt es eine seperate compile und instantiate methode. Dieses Code snippet, werden wir in den Demos noch öfter sehen

Dann würde ich sagen haben wir erstmal über die grundlegende Spezifikation und Definition von WebAssembly geredet und schauen uns jetzt mal ein paar Demos an, wie wir konkret mit WebAssembly umgehen können.

## Demo - Prequisitories

- Kleiner Disclaimer. Die Demo wird primär in Go stattfinden. Teile sind auch in Rust. Der Syntax ist aber sehr trivial, weswegen wir uns dazu entschieden haben Go zu verwenden.

## Unterstützte Sprachen

- Es gibt allerdings noch viel mehr Sprachen mit WASM support. Hier eine kleine Aufzählung:
    - C
    - C++
    - Go
    - Rust
    - Zig
    - C# 
    - AssemblyScript, ein Typescript-Dialekt ist, der mit WebAssembly im Hinterkopf designed wurde.
- Es gibt noch viel mehr, aber die sind entweder nicht so interessant oder noch nicht bereit für production usage

Dann würde ich sagen, ab zur Demo. Wir beginnen mal mit einem simplen Hello World!

## Demo - File Structure

- Unsere Demos werden die Folgende struktur haben. Das ist so ziemlich das minimalste finale Setup, das wir brauchen.
 - Wir haben einen assets Ordner mit einer `index.html`, der kompilierten `main.wasm` und der `wasm_exec.js` mit dem JavaScript glue code
- Dann haben wir einen `cmd` ordner mit einem `server` und einen `wasm` Ordner mit unserer zu kompilierenden Applikation

## Demo - cmd/wasm/main.go

- Die erste Demo wir ein simples Hello World programm. Es soll uns Hello World in der Browser-Konsole printen
- Die `main.go` Datei im wasm folder ist, wie man sehen kann ein simples Hello World in Go. Keine Besonderheiten

## Demo - Compilation

- Diese datei kompilieren wir jetzt zu WebAssembly. Dafür spezifizieren wir bei der kompilierung einfach das target WASM mit den Optionen `GOOS=js` und `GOARCH=wasm`
- Jetzt sollten wir die kompilierte `main.wasm` im assets Ordner finden

## Demo - Javascript Glue Code

- Es ist eine gewisse Menge an JavaScript Code nötig, um das WebAssembly Modul, welches wir gerade erstellt haben im Browser zu laden. 
- Dieser Code ist nötig, da WASM nicht auf Web APIs zugreifen kann. Daher muss WASM JavaScript callen, welche dann den API call machen. Dafür ist der Glue Code da. In e.g. Emscripten implementiert der glue code Zugriff auf SDL, OpenGL, OpenAL und teile von POSIX.
- Dieser Code ist in der Regel immer bereitgestellt
- In Go erhalten wir den Glue Code aus der Go installation (`wasm_exec.js`)
- Ganz wichtig hierbei ist, dass sich dieser mit den Go Versionen ändert. Wenn also das Program so nicht bei euch läuft, versucht mal die `wasm_exec.js` durch eure eigene auszutauschen.

# Demo - assets/index.html

- Die `index.html` included die `wasm_exec.js` Datei, welches der glue code ist
- Im script laden wir das WebAssembly Module wie auch schon im WebAssembly Text Beispiel
- Das importObject enthält dabei die Werte, die in die neue Instanz einfließen sollen, also Funktionen oder `WebAssembly.Memory` Objekte

# Demo - cmd/server/main.go

- Hier schreiben wir uns einen kleinen WebServer, welcher unseren `assets` Ordner served

# Demo - Execection

- Diesen WebServer können wir ausführen und das Ergebnis im Browser betrachten
- In der Console können wir unser Hello World betrachten, welches aus unserem WASM Module entspringt
- Im Inspektor können wir auch unsere WASM-Datei entdecken und auch den Code in WebAssembly Text-Repräsentation

Die nächsten Demos werden jetzt auf die interoperability abzielen, die ich schon vorher erwähnt hattte. Wir können WebAssembly von JavaScript callen und JavaScript von WebAssembly callen. 
In der foldenden Demo werde ich nun die Exports präsentieren, die es erlauben WebAssembly-Module in JavaScript zu callen.

## Demo Exports 

- In dieser zweiten Demo wollen wir eine WASM funktion exportieren und diese dann in JavaScript nutzen.
- Hierfür verändern wir die main.go
- Wir definieren eine eine add function
- Jetzt erstellen wir einen Wrapper mit Hilfe des `syscall/js` packages. Das Package erlaubt die Interaktion mit JavaScript. Dieses Package ist nur legitim, wenn das Compilation-Target WASM ist. Deshalb zeigt das Go LSP dies auch als Fehler an. Func ist eine gewrappete Go function, welche von JavaScript gecalled werden kann
- FuncOf returned eine function, welche von JS gecalled werden kann
- Mit Global erhalten wir ein JavaScript Global Object
- Mit Set setzen wir die value zu einem property string

- Außerdem verändern wir die `index.html` um ein wenig Interaktion mit dem DOM zu haben. Sonst sieht der Code wieder gleich aus

- Jetzt können wir wieder den Server starten und das Ergebnis betrachten
- Wir haben nun ein simples Interface, mit welchem wir unsere im WASM Module definierte `add` function callen können.
- Genauso können wir jetzt add in der Konsole aufrufen


## Demo Imports

- Die letzte Demo zeigt uns noch, wie wir auf JavaScript aus unserem WASM Module zugreifen können.
- Hierfür erstellen wir uns eine kleine JavaScript library
    - add function
    - hello function
    - env variable
    - config object

- Die `main.go` greift auf JavaScript mit dem `syscall/js` package zu. 
    - Callen add function mit 2 Parametern 
    - Printen das Ergebnis. Müssen aber den Typ mitgeben angeben, da JavaScript dynamische Typen hat und es daher theoretisch alles sein könnte. 
    - Callen hello
    - Greifen auf die `name` Variable zu
    - Ändern die `name` Variable und schauen den veränderten Wert an 
    - Setzen einen key und eine value im `config` Object 
    - Betrachten die value zum `key`

- In der `index.html` müssen wir zusätzlich nur die `lib.js` Datei includen

Und schon können wir das Ergebnis im Browser betrachten.

## Additional Information 

- WASM-Funktionen unterstützen nur Integer und Floats. Mit anderen Datentypen kann über den Linear Memory gearbeitet werden
- Linear Memory - Ein kontinuierlicher Buffer von Bytes, auf welche von JavaScript und WebAssembly zugegriffen werden kann 
- Keine DOM Manipulation - Kann aber über JS erreicht werden `syscall/js` in Go oder `wasm-bindgen` in Rust. AssemblyScript provides those features out of the box

## Terminology 

- Module und Memory sollten aus den Demos halbwegs klar geworden sein
- Table ist einfach ein typisierter array von Referenzen zu e.g. Funktionen oder anderen Informationen, welche aus Sicherheitsgründen nicht einfach so im Linear Memory gespeichert werden sollten
- Instance: Module inklusive Memory, Table und Imports

## WebAssembly Runtimes

- WebAssembly überall ausführen, also auch außerhalb des Browsers
- WASM Runtimes sind einfach low level virtual stack machines
- Bekannte Beispiele sind Wasmer, Wasmtime und die WebAssembly Micro Runtime
- Vielleicht kennen manche von euch auch noch Lucit, aber Lucit wurde für `wasmtime` deprecated

## wasmtime demo

Wenn wir nun schon bei dem lokalen ausführen von WASM Binaries sind, brauchen wir eigentlich noch einen Weg, um mit dem System zu interargieren.

## WASI 

- Zitat von Lin Clark
- Mit dem System interagieren, wenn wir WASM außerhalb des Browsers ausführen
- Normalerweise übernimmt die Standard library einer programmiersprache diese Aufgabe und stellt das system interface für die jeweilige Sprache bereit
- Zum Beispiel könnte für die `printf` Function auf einer Windows Maschine die Windows API genutzt werden. Auf einer Linux oder Mac Maschine stattdessen könnte POSIX verwendet werden
- WASM weiß nicht, auf welchem System die Binary ausgeführt wird, daher kann kein vorgefertigtes OS interface verwendet werden
- Realisiert wird das ganze wie bei CloudABI's capability-orientiert, weswegen WASI auch gut in das Sanboxing Modell von WASM passt. WASI hat keine Möglichkeit nach außen zu kommunizieren ohne von außen mitgegebene capabilities e.g. File descriptors. 
- E.g. Statt open, openat, hier wird dann ein FileDescriptor benötigt. 
    - Es gibt aber trotzdem auf lipreopen basierte open functions
- Es gibt eine system call wrapper layer, welche calls zur eigentlicheen WASI Implementation macht, welche diese calls dann auf die Umgebung mapped
- Ziel damit ist eben portabilität und sandboxing auf system interface level
    - Geplant ist eine full-featured libc implementation 

## WAGI 

- HTTP handlers mit Hilfe von STDIN, STDOUT und Umgebungsvariablen
- Damit können WASM microservices und web apps geschrieben werden
- Grundlegende problem welches gelöst werden soll ist, dass eine WASM binary kein Server ist. Sie kann nicht aktiv listenen, noch läuft sie die ganze Zeit. Also wie bekommen wir HTTP requests
- WAGI-Server verwenden. Dies leitet die HTTP request weiter, indem es die Binary mit STDIN und Umgebungsvariablen mit den nötigen Information versorgt. Die HTTP Response kommt dann über STDOUT und kann wieder von WAGI-Server gehandled werden

## Evaluation

- Nachdem wir wissen was WebAssembly ist, wofür wir es brauchen und wie wir damit arbeiten können, können wir mal schauen, was WASM macht, um die von ihnen selbstgesteckten Ziele, welche auf ihrer Website stehen zu erreichen
- Wir fokussieren in dieser Präsentation mal auf "Efficient and fast" und "Safe", beziehungsweise auf die Security
- Zuerst betrachten wir aber, was WebAssembly so schnell macht, denn WebAssembly verspricht uns near native speed beim ausführen von Code

## Efficient and fast 

- Zwei Grafiken die die Zeitverteilung von Prozessen bei der ausführung von JavaScript und WASM im Browser betrachten
- Oben JavaScript 
- Unten WASM
- Am Signifikantestens sind Compilation und Execution. Diese Schritte werden wir mal näher beleuchten.

## Efficient and fast - Compilation + Optimization

- JavaScript wird während der Ausführung mit einem JIT Compiler kompiliert. Da Javascript dynamic typing besitzt kann es sein, dass einzelnen Schritte mehrmals durchgeführt werden müssen, bis die richtigen Typen gefunden wurden
- WebAssembly spart in diesem Schritt Zeit, da wir die Typen bereits kennen und der Code bereits mit e.g. LLVM optimiert wurde, als die WebAssembly Binary erstellt wurde

## Efficient and fast - Execution

- Die Meisten Entwickler wissen nicht über JIT internals bescheid, daher ist es auch schwer speziell für ihn optimalen Code zu schreiben. 
- Im Gegensatz dazu ist WASM als compiler target designed und daher auch dementsprechend performant
- WASM kann in nahezu native speed ausgeführt werden (native - wenn man den Code so bei sich laufen lassen würde). Im Browser können die Programme trotzdem mit 60-70% der Geschwindigkeit ausgeführt werden (aus einem C++ Beispiel)
    - Schlechte register allokierung (Erinerrung: WASM keine Vorgaben)
    - Reservierte Register
    - Schlechte Wahl der Instruktionen
    - Stack overflow checks
- Meistens ist trotzdem ein speedup von bis zu 800% im vergleich zu JavaScript möglich

## Evaluation

- Jetzt haben wir also die Performance von JavaScript und WASM verglichen
- Jetzt sollten wir uns anschauen, was WASM für Security unternimmt 

## Security 

- Es gelten also die gleichen Sicherheitsbeschränkungen wie auch schon für JavaScript Code
    - Same-Origin Policy
    - Auf Dateisystem oder Hardware kann auch nur mit Permission zugegriffen werden
- Um nicht beliebigen Speicher e.g. Passwörter oder andere Tokens auslesen zu können werden WASM-Moudlen sperate Speicherbereiche zugewiesen
    - Untermenge des JS-Heaps und wird durch einen ArrayBuffer realisiert.
    - JS-Umgebung weiß daher zu jedem Zeitpunkt die größe und Inhalt des Buffers -> Unerlaubte Speicherzugriffe unterbinden, da Zugriffe nur innerhalb des Buffers erlaubt sind
    - Overhead, aber notwendig
- Execution Stack wird außerhalb des WebAssembly-Speichers gespeichert und WASM hat nur Lesezugriff
- Bei Function calls wird mit vorher erwähnten Tabellen gearbeitet 
    - Statt Zieladresse im call-Befehl verwendet der call Befehl zwei Parameter. Einen Index und eine Funktionssignatur. Der Index zeigt in die Tabelle mit Funktionspointern, die auch außerhalb des WebAssembly-Speichers gespeichert wird, sodass das Überschreiben nicht möglich ist.
    - Prüfung ob Funktionssignatur mit der an Index X übereinstimmt
        - Nur wenn ja, wird die Funktion an der Adresse aufgerufen
        - Wenn nein, wird das Modul sofort gestoppt
- Traps, JavaScript Exceptions um abnormales Verhalten an die Laufzeitumgebung zu melden. Module wird sofort gestoppt

- Trotz dessen will ich euch nicht vorgaukeln, dass WASM die perfekt sichere Lösung ist. Es gab natürlich auch mal Sicherheitslücken. Crypto Mining, Spectre Angriffe und ausbrüche aus der Sandbox sind schon vorgekommen. Es gibt auch interessante side-channel-attacks, aber ich glaube das ist eher was für eine Master oder Doktorarbeit, so wie ich das aus dem Arikel herauslesen konnte. 

## Conclusion

- Don't use it for your website
- Use it for performance intensive usecases
- Wenn wirklich performance benötigt wird und nicht aller Gain den Overhead verschwindet, dann lohnt sich WebAssembly allerdings extrem, weswegen auch schon so viele Firmen auf WebAssembly in gewissen Teilen ihrer Anwendungen schwören
- Portability and interoperability
- Here to stay 
    - Offener Standard, der Unterstützt und Gefördert wird e.g. von der ByteCode Allegiance, die sich gegründet, um WASM weiterzuentwickeln (ARM, Google, Intel, Microsoft und Mozilla)
    - Krustlet, was die Entwicklung weitertreibt  und auch nochmal einen kleine Verbindung zu Felix seinem Teil darstellt, da Krustlet im Endeeffekt erlaubt WASM-Modules statt container images auf Kubernetes Clustern laufen zu lassen

## Conclusion

- Um abzuschließen nochmal kurz zur Grafik vom Anfang 
- Ich glaube sehr wohl, dass wir einen weiteren Wendepunkt in der performance im Web erreicht haben
- Und ich glaube auch, dass es in Zukunft nicht mehr so sein wird, dass WASM JS benötigt, aber JS WASM nicht unbedingt benötigt, sondern dass auch JavaScript WASM benötigen wird, da wie auch schon bei NodeJS einfach die Anforderungen an Web Applikationen steigen werden und daher nicht mehr auf den Performancezuwachs durch WASM verzichtet werden kann.

## Webshop
 
- Zum Abschluss noch mein Web Shop
- Web Shop mit go-app erstellt, einem Go package um web apps zu erstellen
- HTML kann deklarativ in Go geschrieben werden