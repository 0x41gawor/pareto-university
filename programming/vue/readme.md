Roadmap:
- [ ] ETAP 0 — Infrastruktura projektu (1 dzień)
- [ ] ETAP 1 — Fundamenty Composition API (2–3 dni)
- [ ] ETAP 2 — Router + Przepływ danych (3 dni)
- [ ] ETAP 3 — Praca z HTTP i architektura usług (3–5 dni)
- [ ] ETAP 4 — Formy i walidacje (2–3 dni)
- [ ] ETAP 5 — Stan globalny: Pinia (3–4 dni)
- [ ] ETAP 6 — Reużywalność i architektura dużych aplikacji (5–6 dni)
- [ ] ETAP 7 — Optymalizacja i Performance (5–7 dni)
- [ ] ETAP 8 — Integracja z TypeScript (3–5 dni)
- [ ] ETAP 9 — Testy (2–4 dni)
- [ ] ETAP 10 — Architektura profesjonalnej aplikacji CRUD (pełny projekt)
- [ ] (Opcjonalnie) ETAP 11 — Zaawansowane koncepcje

---

# 🚀 ROADMAPA NAUKI VUE.JS 

Całość zakłada **Vue 3 + Composition API + TypeScript + Vite** jako standard.

## ⭐ ETAP 0 — Infrastruktura projektu (1 dzień)

**Cel:** od razu pracujesz na realnym środowisku, bez CLI-magii.

### Wiedza

* Struktura projektu Vite + Vue.
* `main.ts`, `createApp()`, montowanie aplikacji.
* `App.vue` jako root-component.
* SFC (Single File Component): `<script setup>`, `<template>`, `<style>`.
* Różnice pomiędzy:

  * Composition API vs Options API (wybierasz Composition).
  * reactive vs ref.
  * DOM rendering pipeline Vue.

### Mini-projekt

**Projekt [00] Minimal SPA**
Tworzysz:

* `<Navbar/>`
* `<Counter/>` z `ref`, `reactive`
* `<AppLayout/>`

Czysty komponentyzm + propsy.

---

# ⭐ ETAP 1 — Fundamenty Composition API (2–3 dni)

### Wiedza

* `ref`, `reactive`, `computed`, `watch`, `watchEffect`
* Lifecycle hooks: `onMounted`, `onUpdated`, `onUnmounted`, `onBeforeRouteLeave` (router).
* Emity: `defineEmits`, `defineProps`.
* Różnice między reactive → shallowReactive, markRaw, toRaw.
* Jak działa reactivity tracking (effect graph, dependency tracking).
* Jak działa virtual DOM i diffing w Vue (wysoko abstrakcyjnie, inżyniersko).

### Mini-projekt

**Projekt [01] Advanced Counter**

* Counter z timerem (interval).
* Oddzielny plik `useCounter.ts` jako pierwszy composable.
* Watch różnych źródeł.
* Computed z zależnościami.

---

# ⭐ ETAP 2 — Router + Przepływ danych (3 dni)

### Wiedza

* Vue Router (4.x):

  * historia: hash vs history API
  * lazy loading route-level code splitting
  * navigation guards
  * dynamiczne paramy
  * route meta → np. permissions
* Przekazywanie danych:

  * props down / emits up
  * dependency injection: `provide` / `inject`
  * composables jako data/logic sharing

### Mini-projekt

**Projekt [02] Mini CRUD „Books"**

* Router: `/books`, `/books/:id`
* CRUD na localStorage.
* Użycie `provide/inject` np. do globalnego alert/notification systemu.
* Użycie dynamicznych routów i przedrostków typu `/books/edit/:id?`.

---

# ⭐ ETAP 3 — Praca z HTTP i architektura usług (3–5 dni)

### Wiedza

* Axios vs Fetch API → wybierasz fetch.
* Pattern „Service Layer”:
  osobne katalogi:

  ```
  /src/services/userService.ts
  /src/services/bookService.ts
  ```
* Error handling globalny: interceptor lub warstwa fetch-wrapper.
* AbortController.
* Globalny loading indicator: composable `useFetchState`.

### Mini-projekt

**Projekt [03] API Client Layer**

* Tworzysz `apiClient.ts`.
* Dodajesz retry + exponential backoff dla wybranych endpointów.
* Tworzysz prosty UI + tabela z paginacją (np. użytkownicy z placeholder JSONPlaceholder API).
* Obsługa błędów → snackbar lub toast.

---

# ⭐ ETAP 4 — Formy i walidacje (2–3 dni)

### Wiedza

* `v-model` (różne warianty, modyfikatory).
* Dynamiczne formularze.
* Walidacja (`yup` / `vee-validate` / własne composables).
* Controlled vs uncontrolled components.
* Reużywalne inputy: `BaseInput.vue`.

### Mini-projekt

**Projekt [04] Massive Form**

* Form z:

  * kilkoma typami inputów
  * walidacjami
  * dynamicznym renderowaniem pól
* Dodajesz debounce i watch do autouzupełniania.

---

# ⭐ ETAP 5 — Stan globalny: Pinia (3–4 dni)

### Wiedza

* Dlaczego Pinia → różnice z Vuex.
* Store patterns:

  * global store
  * slices/modules
  * object-based vs function-based stores
* Actions vs getters.
* Persisted state.
* Architektura folderów:

  ```
  /stores/
  /services/
  /composables/
  /components/
  ```

### Mini-projekt

**Projekt [05] Global Auth Store**

* Logowanie + odświeżanie tokenów.
* User profile store.
* „HasRole” / „IsAdmin” z meta w routerze.
* Persist token w localStorage + auto-logout.

---

# ⭐ ETAP 6 — Reużywalność i architektura dużych aplikacji (5–6 dni)

### Wiedza

* Budowanie composables:

  * `usePagination()`
  * `useTable()`
  * `useToggle()`
  * `useFetch()`
* Smart vs dumb components.
* Patterns:

  * Presenter–Container
  * Component-driven development
  * Headless components
* Slottable components (advanced): `v-slot`, scoped slots.
* Przemyślany podział na moduły.

### Mini-projekt

**Projekt [06] Headless Table Component**

* Tworzysz komponent:

  ```
  <BaseTable :items="..." :columns="...">
    <template #row="{ item }"> ...custom... </template>
  </BaseTable>
  ```
* Sortowanie, filtrowanie, paginacja.
* Reużywalna logika w composables.

---

# ⭐ ETAP 7 — Optymalizacja i Performance (5–7 dni)

### Wiedza (twarde inżynieryjne rzeczy)

* Render pipeline Vue:
  dependency tracking → effect → virtual DOM render → DOM patch.
* Memoization:

  * `computed` jako cache.
  * `shallowRef`, `shallowReactive`.
* Optymalizacje zaawansowane:

  * `v-memo`
  * `defineOptions({ name: ... })`
  * `keep-alive` w routerze.
  * Lazy loading heavy components.
  * Chunk splitting w Vite.
* Perf debugging:

  * `app.config.performance = true`
  * Vue DevTools → Flamegraph rendering.

### Mini-projekt

**Projekt [07] 1k rows table benchmark**

* Stworzyć tabelę na 1000–5000 wierszy.
* Przetestować:

  * naive version
  * reactive optimizations (shallowRef, pagination)
  * virtual scrolling (Vue Virtual Scroll List)

---

# ⭐ ETAP 8 — Integracja z TypeScript (3–5 dni)

### Wiedza

* Typowanie propsów (`defineProps<{...}>()`).
* Typowanie emitsów.
* Typowanie store’ów Pinia.
* Typowanie API responses.
* `interface` vs `type` w praktyce.
* TS generics w composables (ważna część!):

  ```
  function usePaginated<T>(fetcher: () => Promise<T[]>)
  ```

### Mini-projekt

**Projekt [08] Typed API SDK**

* Tworzysz warstwę API z silnym typowaniem
  (np. `/api/users`, `/api/orders`).
* TS generics w `useQuery<T>()`.

---

# ⭐ ETAP 9 — Testy (2–4 dni)

### Wiedza

* Vue Test Utils.
* Testy komponentów: snapshoty, mount, shallowMount.
* Testy composables osobno.
* Testy Pinia stores.
* Testy routera (navigation guard logic).

### Mini-projekt

**Projekt [09] Test Lab**

* Jednostkowe testy:

  * komponentu input
  * store Pinia
  * composable `useCounter`
  * router guard

---

# ⭐ ETAP 10 — Architektura profesjonalnej aplikacji CRUD (pełny projekt)

### Finalny cel

Budujesz projekt, który ma strukturę podobną do **Twojego `diet-app`**:

* 20–30 komponentów.
* Reużywalne layouty.
* Warstwa usług `/services`.
* Warstwa composables `/composables`.
* Globalny Pinia.
* Routing + guards.
* Responsywność + dark mode.
* Kod w całości w TypeScript.
* Modularyzacja widoków i pod-widoków.
* Integracja z realnym backendem.

### Projekt [10] Full CRUD System (Twojego wyboru):

* Inventory manager
* Notes/Tasks
* Meal planner (rewrite mini wersji dietonez)
* Budget tracker
* Library system

Z naciskiem na:

* struktury danych
* side effects
* architekturę folderów
* kompozycję logicznych warstw

---

# 🧠 (Opcjonalnie) ETAP 11 — Zaawansowane koncepcje

Dla ciekawych:

* Server-side rendering: **Nuxt 3**
* Micro-frontends + Vue
* Web Workers z Vue
* Storybook + komponenty headless
* Custom directives (`v-focus`, `v-on-click-outside`)
* Render functions i JSX w Vue
