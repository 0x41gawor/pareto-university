# Pierwsze przybliżenie

## Podstawy BPMN

Cel BPMN:

> Zasadniczym celem BPMN jest zapewnienie notacji, która może być łatwo zrozumiana
> przez wszystkich użytkowników biznesowych, począwszy od analityków biznesowych
> tworzących pierwsze szkice procesów, a skończywszy na developerach technicznych
> odpowiedzialnych za wdrażanie technologii, które będą wykonywać te procesy
> i finalnie osobach zarządzających, które będą kierować procesami oraz je
> monitorować. […] Celem BPMN jest ustandaryzowanie modelu i notacji procesu
> biznesowego w odniesieniu do wielu różnych notacji modelowania i punktów
> widzenia. Dzięki temu BPMN zapewni prostą metodę komunikowania informacji
> o procesie innym użytkownikom biznesowym, osobom wdrażającym procesy, 
> klientom i dostawcom. […] BPMN zapewni organizacjom możliwość zrozumienia
> ich wewnętrznych procedur biznesowych z wykorzystaniem notacji graficznej oraz
> da możliwość komunikowania tych procedur w ustandaryzowany sposób. […]
> Dzięki temu organizacje będą mogły zrozumieć, jak działają one same oraz inne
> podmioty, z którymi współpracują. Organizacje będą też mogły szybciej dopasowywać
> się do nowych warunków — zarówno wewnętrznych, jak i związanych ze współpracą
> biznesową.

Trudne zadanie przed BPMN połączyć takie 3 grupy stąd BPMN proponuje koncepcję **poziomów zgodności**, która dzieli elementy BPMN na 3 grupy.

**Poziom opisowy** - najbardziej podstawowe elementy. Diagramy niewiele się różnią od schematów blokowych, ale są ustandaryzowane i czytelne.

**Poziom analityczny** - te elementy pozwalają na tworzenie bardziej wyrafinowanych modeli, skierowanych do odbiorcy zainteresowanego analizą procesów na potrzeby ich automatyzacji.

**Poziom ostatni** - pozwalają w zamyśle na automatyzację (przez silnik?), choć nic nie stoi na przeszkodzie by korzystać z nich w celach dokumentacyjnych.

![image-20250315232947127](img/image-20250315232947127.png)

1a i 1b to **pools** - nazwa ma sugerować, że proces zawsze płynie.  W basenie możemy mieć **swimlanes** (1c), żeby pokazać kolejny level hierarchi wykonawcy czynności.

**Events** (2a i 2b) - to pewien stan rzeczywistości, który może reprezentować rozpoczęcie naszego procesu lub jego rezultat.

**Activities** pokazują jaka praca jest wykonywana, mogą to być **tasks** (3) oraz **subprocesses** (4).

Procesy rzadko są sekwencjami, aby móc je rozdzielać na ścieżki mamy **gates** (5).

![image-20250315233853206](img/image-20250315233853206.png)

Najważniejszy rodzaj strzałek to te czarne (A) czyli **sequence flow**.

PRZEPŁYW PROCESU (TOKEN) MOŻE PŁYNĄĆ TYLKO W JEDNYM BASENIE.

Jak w takim razie pokazać komunikację różnych podmiotów?

Do tego służy kolejny rodzaj strzałek **message flow** (B).  On może być tylko między basenami.

Aby przekazywać dodatkowe info ale bez wpływania na elementy opisane wyżej to wykorzystujemy specjalne do tego **artefakty**.

Do tej kategorii zalicza się **text annotation** (C) oraz **group** (6b).

Gdy text annotation odnosi się do konkretnego elementu to możemu użyć konktora **association**. 

By pokazać w jaki sposób płyną dane w procese mamy **data object** (7a) oraz **data association** (D). 7b pokazuje **data storage**.

## Bramki

### Bramka XOR.

Żeton idzie tylko jedną ścieżką, pierwszą która spełni warunek. Nie ma rozdzielania żetona. Można zaznaczyć default scieżkę.

![image-20250315222944983](img/image-20250315222944983.png)

### Bramka AND (pararell)

![image-20250315223227301](img/image-20250315223227301.png)

Tu bramka bierze jeden żeton i wypuszcza 3. Potem procesy B, C i D trzymają przez różne czasy żetony i do E żeton trafia 3-krotnie. 

Więc log pokaże np. ACEFDEFBEF

![image-20250315223424026](img/image-20250315223424026.png)

Tu ta druga bramka jest telepatycznie powiązana z pierwszą i skoro tamta wypuściła 3 żetony, to ta czeka na 3 żetony i dopiero wypuści jeden na wyjściu, jak na inpucie dostanie 3.

### Bramka OR

Ta dostaje na input token i wypuszcza po jednym dla każdego prawdziwego warunku. Druga czeka na dokładnie tyle tokenów ile wypuściła ta pierwsza.

![image-20250315223813262](img/image-20250315223813262.png)

## Zadania

Mamy dwa typy obiektów do opisania aktywności - zadania i podprocesy.

![image-20250315224431247](img/image-20250315224431247.png)

### Zadanie użytkownika (user task)

System przypisze to zadanie jakiemuś użytkownikowi. Może to być w systemie CRM, ERP, "help desk" itp.

### Zadanie typu usługa (service task)

Silnik automatyzujący proces woła inne oprogramowanie udostępniające pewną usługę (możemy odwołać się do koncepcji SOA, architektury zorientowanej na usługi, albo też bardziej współczesnej koncepcji mikrousług).

### Zadanie manualne

Czyli takie, które nie jest robione z pomocą silnika procesowego. Oprócz zadań bez komputera, może to być też coś wykonywane w innych programach po prostu.

### Zadania: wysłane i odbiór

Te 3 już wystarczą do podstawowego modelowania. Kolejnym krokiem są zadania stosowane do opisywania zautomatyzowanego przesyłania komunikatów.

![image-20250315225820915](img/image-20250315225820915.png)

W BPMN jest ogólna zasada, że elementy ciemne są aktywne (rzucające), a jasne (otwarte) są pasywne (chwytające).

Zadania typów pasywnych czekają z tokenem i nie puszczą go, aż nie dostaną komunikatu od zewnątrz.

Ogólnie to nie musisz pokazywać drugiego basenu jeśli to nie jest istotne.

