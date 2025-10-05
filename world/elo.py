import csv
from datetime import datetime
from collections import defaultdict

# lista sklepów (plików CSV w katalogu bieżącym)
shops = [
    "biedronka", "carrefour", "kfc", "lidl", "mcdonalds",
    "putka", "spozywczy", "stokrotka", "w_bulce", "zabka", "zdrowe"
]

gas_stats = ["shell", "orlen", "bp", "circle_k", "amic", "moya", "mol"]

def parse_file(filename):
    """Parsuje plik CSV i zwraca sumy miesięczne {YYYY-MM: suma}"""
    monthly_sums = defaultdict(float)

    with open(filename, encoding="utf-8") as f:
        reader = csv.reader(f, delimiter=";")
        for row in reader:
            if not row or len(row) < 6:
                continue

            date_str = row[1]   # druga kolumna
            amount_str = row[5] # kwota

            try:
                amount = float(amount_str.replace(",", "."))
            except ValueError:
                continue

            dt = datetime.strptime(date_str, "%d-%m-%Y")
            month_key = dt.strftime("%Y-%m")
            monthly_sums[month_key] += amount

    return monthly_sums


# globalna suma dla wszystkich sklepów
total_sums = defaultdict(float)

# główna pętla po plikach
for shop in gas_stats:
    filename = f"{shop}.csv"
    sums = parse_file(filename)

    print(f"\n=== {shop.upper()} ===")
    for month in sorted(sums.keys()):
        print(f"{month}: {sums[month]:.2f} PLN")
        total_sums[month] += sums[month]

# drukujemy całość na końcu
print("\n=== ŁĄCZNIE ===")
for month in sorted(total_sums.keys()):
    print(f"{month}: {total_sums[month]:.2f} PLN")