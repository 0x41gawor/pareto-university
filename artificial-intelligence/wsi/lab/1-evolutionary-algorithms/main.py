from util import *

population_size = 10
func = 'happycat'
k = 5 # number of parents
c = 5 # number of children
mp = 0.05 # mutation probability in single child 0.05 means 5%
sigma = 2.0 # mutation strengh, mean value of step 

# Example usage
p0 = initialization(population_size=population_size)
evals = evaluate(p0, func)

T = selection(p0, evals, k, 3)
O = crossover(T, c)

print("Population")
print(p0)
print("Selected parents")
print(T)
print("Children")
print(O)
O = mutate(O, mp=0.5, sigma=2.0)
print("Mutated children")
print(O)
P = p0
evals = evaluate(np.vstack((P,O)), func)
P = succession(np.vstack((P, O)), evals, population_size)
print("Successed population")
print(P)
