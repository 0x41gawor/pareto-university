from util import *

population_size = 20
func = 'happycat'
k = 5

# Example usage
p0 = initialization(population_size=population_size)
evals = evaluate(p0, func)

T = tournament_selection(p0, evals, k, 3)

print(p0)
print(evals)
print(T)
