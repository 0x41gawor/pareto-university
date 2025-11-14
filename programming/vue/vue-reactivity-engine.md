## ⚙️ 1️⃣ Reaktywność w Vue.js — fundament

Vue posiada **silnik reaktywności**, który:

- **śledzi zależności** między danymi (`data`, `ref`, `reactive`)
- i **automatycznie aktualizuje** wszystko, co od nich zależy (`template`, `computed`, `watch`, itd.)

Mechanicznie działa to przez **proxy** — Vue “owija” twoje obiekty, żeby wiedzieć, kto z nich korzysta.
Kiedy wartość się zmienia, Vue wie **co trzeba odświeżyć w DOM** lub **jakie funkcje trzeba ponownie wykonać**.

------

## 🧩 2️⃣ Cztery główne kategorie „reaktywnych konceptów”

| Koncept        | Typ                   | Służy do                                       | Reaguje automatycznie? |
| -------------- | --------------------- | ---------------------------------------------- | ---------------------- |
| **ref()**      | źródło danych         | trzymania prostych wartości (reaktywnych)      | ✅                      |
| **reactive()** | źródło danych         | trzymania złożonych obiektów (reaktywnych)     | ✅                      |
| **computed()** | pochodna (derivative) | tworzenia wartości zależnych od innych danych  | ✅                      |
| **watch()**    | obserwator (observer) | reagowania skutkami ubocznymi na zmiany danych | ✅ (wywołuje callback)  |

Zobaczmy każdy z nich po kolei 👇

------

## 🌱 3️⃣ `ref()` — najprostsza jednostka reaktywna

```ts
import { ref } from 'vue'

const count = ref(0)
```

- `ref` tworzy obiekt reaktywny z pojedynczą wartością.
- W kodzie odwołujesz się przez `.value` (bo to **opakowanie** wokół wartości).

```ts
console.log(count.value) // 0
count.value++
```

W template (`.vue`) nie musisz pisać `.value`:

```html
<p>{{ count }}</p>
<button @click="count++">+</button>
```

➡️ Każda zmiana `count` automatycznie **odświeży DOM**.

------

## 🧱 4️⃣ `reactive()` — reaktywne obiekty i tablice

```ts
import { reactive } from 'vue'

const user = reactive({
  name: 'Anna',
  age: 25
})
```

- Działa jak `ref`, ale dla złożonych struktur.

- Vue śledzi każdą właściwość:

  ```ts
  user.name = 'Ewa'  // odświeży wszystko co używa user.name
  ```

⚠️ W przeciwieństwie do `ref`, tu nie używasz `.value`.

------

## 🧮 5️⃣ `computed()` — wartości pochodne (cache + automatyka)

```ts
import { computed, ref } from 'vue'

const price = ref(10)
const quantity = ref(3)

const total = computed(() => price.value * quantity.value)
```

- `computed` to **wartość pochodna** od innych reaktywnych danych.
- Vue sam:
  - zapamiętuje (cache) wynik dopóki zależności się nie zmienią,
  - przelicza ponownie tylko gdy `price` lub `quantity` się zmieni.

➡️ To odpowiednik **“kolumny obliczeniowej” w Excelu**.

Można też mieć wersję zapisywalną:

```ts
const fullName = computed({
  get: () => `${first.value} ${last.value}`,
  set: (val) => {
    [first.value, last.value] = val.split(' ')
  }
})
```

------

## 👀 6️⃣ `watch()` — reagowanie efektami ubocznymi

`watch` to sposób na *“zrób coś, gdy dane się zmienią”*
 (np. wyślij żądanie HTTP, zapisz do localStorage, zloguj coś, itd.)

```ts
import { ref, watch } from 'vue'

const query = ref('')

watch(query, (newVal, oldVal) => {
  console.log('Nowe zapytanie:', newVal)
})
```

➡️ W przeciwieństwie do `computed`, `watch` **nie zwraca wartości** —
 to **efekt uboczny** (side effect).

------

## 🧠 7️⃣ Jak one się łączą

W uproszczeniu:

```
ref/reactive  →  źródła danych
computed      →  zależne pochodne (bez efektów ubocznych)
watch         →  efekty uboczne przy zmianie
```

### Przykład:

```ts
const count = ref(0)
const double = computed(() => count.value * 2)

watch(double, (newVal) => {
  console.log('Nowe double:', newVal)
})
```

- `count` to źródło,
- `double` automatycznie się przelicza,
- `watch` reaguje gdy `double` się zmieni.

------

## 🔁 8️⃣ Wspólne cechy wszystkich czterech

| Cechy wspólne        | Opis                                                   |
| -------------------- | ------------------------------------------------------ |
| **Reaktywne**        | Vue śledzi zmiany i reaguje automatycznie              |
| **Zależności**       | Wszystkie mogą być źródłami lub zależnościami innych   |
| **Integracja z DOM** | Template’y Vue automatycznie renderują je bez `.value` |
| **Deklaratywność**   | Opisujesz *co ma się dziać*, nie *jak to zrobić*       |

------

## 🔧 9️⃣ Dodatkowe koncepty pokrewne

Dla pełnego obrazu, Vue ma też:

- **`watchEffect()`** → jak `watch`, ale automatycznie wykrywa zależności:

  ```ts
  watchEffect(() => {
    console.log('Count changed to', count.value)
  })
  ```

- **`shallowRef()` / `shallowReactive()`** → reaktywność tylko 1. poziomu.

- **`readonly()`** → tworzy tylko-do-odczytu kopię reaktywnego obiektu.