## 🧩 1. Co robi przeglądarka “zwyczajnie”

Kiedy wpisujesz adres np. `https://example.com`:

1. **Przeglądarka** wysyła żądanie HTTP do serwera.

2. Serwer zwraca **statyczny plik HTML**, np.:

   ```html
   <html>
     <head><title>Example</title></head>
     <body>
       <h1>Hello world!</h1>
       <p>Strona statyczna</p>
     </body>
   </html>
   ```

3. Przeglądarka **parsuje** ten HTML, tworząc strukturę w pamięci — tzw. **DOM (Document Object Model)**.
    To drzewo obiektów, które reprezentuje każdy tag.

4. Potem:

   - ładuje CSS i stosuje style,
   - wykonuje skrypty JS (jeśli są),
   - renderuje wynikowy widok.

Ten model jest **statyczny** — jeśli chcesz zmienić coś w widoku, musisz to zrobić przez **JavaScript**, np.:

```js
document.getElementById("title").textContent = "Nowy tytuł";
```

------

## ⚙️ 2. Gdzie wchodzi JavaScript

JavaScript umożliwia **dynamiczną manipulację DOM-em** — możesz:

- dodawać elementy (`appendChild`),
- usuwać je,
- zmieniać style, teksty itp.

To jednak szybko robi się **chaotyczne** w dużych aplikacjach.
 Każda zmiana danych wymaga ręcznej aktualizacji widoku.

------

## 🚀 3. Vue.js — pomysł na “automatyzację” tego procesu

Vue.js (podobnie jak React czy Svelte) działa według idei:

> „Nie modyfikuj DOM ręcznie — opisz **jak ma wyglądać** widok dla danych, a my zrobimy resztę.”

Innymi słowy:

- Ty deklarujesz **szablon HTML** (template),
- Vue trzyma dane w obiekcie **`data()`** (stan),
- i **automatycznie aktualizuje widok**, gdy dane się zmienią.

Przykład:

```html
<div id="app">
  <p>{{ message }}</p>
  <button @click="reverseMessage">Odwróć tekst</button>
</div>

<script>
const app = Vue.createApp({
  data() {
    return { message: "Hello Vue!" }
  },
  methods: {
    reverseMessage() {
      this.message = this.message.split('').reverse().join('')
    }
  }
})
app.mount('#app')
</script>
```

Kiedy klikniesz przycisk:

- `message` zmienia się,
- Vue **automatycznie aktualizuje** `<p>{{ message }}</p>` w DOM,
   bez ręcznego `document.querySelector` itd.

------

## 🧠 4. Co Vue.js *naprawdę* wysyła do przeglądarki

Vue.js **nie wysyła plików HTML dynamicznie z serwera**.
 Zamiast tego:

1. Do przeglądarki trafia **jeden statyczny plik `index.html`**, np.:

   ```html
   <html>
   <body>
     <div id="app"></div>
     <script src="main.js"></script>
   </body>
   </html>
   ```

2. W `main.js` znajduje się zbudowana aplikacja Vue (kod JS, który framework wygenerował z twoich `.vue` plików).

3. Po załadowaniu:

   - Vue **tworzy obiekt aplikacji**,
   - “montuje” go do `<div id="app">`,
   - a następnie **buduje i aktualizuje DOM w pamięci przeglądarki** — tzn. generuje elementy HTML dynamicznie na podstawie twoich danych i szablonów.

Czyli:

> Vue.js **nie generuje plików HTML** — **buduje DOM w runtime w przeglądarce**, z pomocą JS.

------

## 🧩 5. Wnętrze Vue: Virtual DOM

Vue (od wersji 2) używa tzw. **Virtual DOM** — uproszczonej kopii drzewa DOM w pamięci.

Kiedy coś się zmienia (np. `message`):

1. Vue tworzy nową wersję Virtual DOM.
2. Porównuje ją z poprzednią (tzw. *diffing*).
3. Wprowadza **minimalne zmiany w prawdziwym DOM**, żeby odświeżyć widok.

Dzięki temu aktualizacje są szybkie i precyzyjne.

------

## 🧱 6. Vue 3 i kompilacja szablonów

W czasie “builda” (np. przez `vite` lub `vue-cli`):

- twoje pliki `.vue` (zawierające HTML, CSS i JS w jednym) są **kompilowane** do czystego JS, który tworzy struktury DOM programowo,

- czyli z:

  ```html
  <p>{{ message }}</p>
  ```

  powstaje coś w stylu:

  ```js
  function render(ctx) {
    return h("p", null, ctx.message);
  }
  ```

Ten kod trafia do przeglądarki, która tworzy prawdziwy DOM już w runtime.

------

## 🌐 7. Alternatywa: SSR (Server-Side Rendering)

Vue może też działać inaczej: **generować HTML po stronie serwera** (np. Nuxt.js).
 Wtedy serwer faktycznie wysyła gotowy HTML do przeglądarki, a Vue przejmuje kontrolę nad nim później („hydration”).

Ale w standardowym SPA (Single Page Application):

> Przeglądarka dostaje pusty HTML + duży JS,
>  a Vue „buduje” stronę po stronie klienta.

------

## 🧭 Podsumowanie

| Krok | Klasyczna strona                        | Strona z Vue.js                                      |
| ---- | --------------------------------------- | ---------------------------------------------------- |
| 1️⃣    | Serwer zwraca gotowy HTML               | Serwer zwraca pusty `index.html` + `main.js`         |
| 2️⃣    | HTML renderowany natychmiast            | Vue.js tworzy DOM w przeglądarce                     |
| 3️⃣    | Zmiany wymagają reloadu lub ręcznego JS | Vue automatycznie odświeża widok przy zmianie danych |
| 4️⃣    | Każda podstrona = nowy plik HTML        | SPA = jedna strona, różne widoki dynamicznie         |

------

Jeśli chcesz, mogę Ci teraz pokazać **diagram przepływu**: od serwera → przez przeglądarkę → do Virtual DOM i renderowania widoku w Vue.
 Chcesz, żebym go narysował (tekstowo lub graficznie)?