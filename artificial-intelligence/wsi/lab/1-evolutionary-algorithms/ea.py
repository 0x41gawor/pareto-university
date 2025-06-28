from util import *
import numpy as np

np.set_printoptions(precision=4, suppress=True)

N = 100 #population_size
func = 'happycat' # objective function
mu = 30 # number of parents
lmbda = 30 # number of children
mp = 0.05 # mutation probability in single child 0.05 means 5%
sigma = 4.0 # mutation strengh, mean value of step 
print(f"EA Params: N={N}, func='{func}', μ={mu}, λ={lmbda}, mp={mp}, σ={sigma}")

p0 = initialization(population_size=N)

evals_num=0     # number of evaluations done
pop_counter=0   # population counter

P = p0
evals = []
while evals_num < 10000:
    evals = evaluate(P, func)
    evals_num += N
    print("🔎 Best Individual in Population", pop_counter, end=" --> ")
    print_best_solution(P, evals)
    T = selection(P, evals, mu, 3)
    O = crossover(T, lmbda)
    O = mutate(O, mp=0.5, sigma=2.0)
    evals = evaluate(np.vstack((P,O)), func)
    evals_num += (N+lmbda)
    P = succession(np.vstack((P, O)), evals, N)
    pop_counter += 1

print("🔎 Best Individual in Population", pop_counter, end=" --> ")
print_best_solution(P, evals)