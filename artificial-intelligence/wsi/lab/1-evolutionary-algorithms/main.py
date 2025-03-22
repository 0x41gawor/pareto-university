from util import *
import numpy as np

np.set_printoptions(precision=4, suppress=True)

population_size = 100
func = 'happycat'
k = 30 # number of parents
c = 30 # number of children
mp = 0.05 # mutation probability in single child 0.05 means 5%
sigma = 4.0 # mutation strengh, mean value of step 

p0 = initialization(population_size=population_size)

evals_num=0     # number of evaluations done
pop_counter=0   # population counter

P = p0
evals = []
while evals_num < 10000:
    evals = evaluate(P, func)
    evals_num += population_size
    print("🔎 Best Individual in Population ", pop_counter)
    print_best_solution(P, evals)
    T = selection(P, evals, k, 3)
    O = crossover(T, c)
    O = mutate(O, mp=0.5, sigma=2.0)
    evals = evaluate(np.vstack((P,O)), func)
    evals_num += (population_size+c)
    P = succession(np.vstack((P, O)), evals, population_size)
    pop_counter += 1

print("🔎 Best Individual in Population ", pop_counter)
print_best_solution(P, evals)