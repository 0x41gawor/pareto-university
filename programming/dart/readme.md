- [ ] const constructor
- [ ] named params, optional params
- [ ] super-parameter


Świetne pytanie! Rozbijmy to na trzy rzeczy: `const` konstruktor, klamry w parametrach i `super.key`.

---

## 1) `const MyApp({super.key});` – o co chodzi z `const`?

To **nie** „stały konstruktor”, tylko **konstruktor stałej**.
`const` przed konstruktorem oznacza, że **obiekt może być utworzony w czasie kompilacji** (compile-time), a nie dopiero w runtime. Warunek: klasa musi być *niemutowalna* (wszystkie pola `final`) i konstruktor nie może robić nic „dynamicznego”.

Dzięki temu możesz pisać:

```dart
runApp(const MyApp()); // instancja kanonizowana, zero alokacji w runtime
```

Korzyści w Flutterze:

* mniejszy churn GC (nie tworzy się nowa instancja przy każdym rebuildzie),
* framework łatwiej porównuje drzewo widgetów (stałe są „te same” między rebuildami).

> Gdyby konstruktor nie był `const`, musiałbyś pisać `runApp(MyApp())` i instancja powstawałaby w runtime.

---

## 2) Skąd te **klamry** w parametrach?

W Dart:

* **`{ ... }`** oznacza **parametry nazwane** (opcjonalne, chyba że dasz `required`).
* **`[ ... ]`** oznacza **parametry pozycyjne opcjonalne**.

Tu masz parametr **nazwany** `key`, więc wywołanie wyglądałoby np. tak:

```dart
MyApp(key: someKey)
```

---

## 3) Co to jest `super.key`?

To skrót (tzw. **super-parameter**) wprowadzony w Dart 2.17.
Zamiast pisać dłużej:

```dart
class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);
}
```

piszesz krócej:

```dart
class MyApp extends StatelessWidget {
  const MyApp({super.key});
}
```

To:

* **deklaruje** nazwany parametr `key` w konstruktorze `MyApp`,
* **automatycznie przekazuje** go do konstruktora klasy bazowej (`StatelessWidget`), która ma `const StatelessWidget({Key? key})`.

`Key` służy Flutterowi do zachowania tożsamości widgetów między przebudowami (ważne przy listach, animacjach, itp.).

---

## Bonus: kiedy używać `const` przy widgetach?

Zawsze, gdy **wszystkie argumenty są stałe w compile-time**:

```dart
const Text('Hello');         // OK
const SizedBox(height: 20);  // OK
// Ale:
final title = getTitle();     // runtime
Text(title);                  // bez const
```
