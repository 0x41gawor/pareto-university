from benchmarkfcns import griewank, happycat
import numpy as np


def initialization(population_size=100, dimension=4, lower_bound=-100, upper_bound=100):
    """
    Initialize a population of individuals with random values.

    Each individual is represented as a real-valued vector within the given bounds.

    Parameters
    ----------
    population_size : int, optional
        Number of individuals in the population (default is 100).
    dimension : int, optional
        Number of dimensions (genes) for each individual (default is 4).
    lower_bound : float, optional
        Minimum value for each gene (default is -100).
    upper_bound : float, optional
        Maximum value for each gene (default is 100).

    Returns
    -------
    np.ndarray
        A NumPy array of shape (population_size, dimension) with uniformly sampled values
        in the range [lower_bound, upper_bound].
    """
    return np.random.uniform(low=lower_bound, high=upper_bound, size=(population_size, dimension))


def evaluate(population, func_name='happycat'):
    """
    Evaluate the fitness of each individual in the population using the specified benchmark function.

    Parameters
    ----------
    population : np.ndarray
        Array of shape (N, D) representing the population of N individuals with D-dimensional vectors.
    func_name : str, optional
        Name of the benchmark function to use. Must be one of {"griewank", "happycat"} (default is "happycat").

    Returns
    -------
    np.ndarray
        Array of shape (N,) containing the fitness values for each individual.

    Raises
    ------
    ValueError
        If the specified func_name is not recognized.
    """
    FUNCTION_MAP = {
        "griewank": griewank,
        "happycat": happycat
    }
    if func_name not in FUNCTION_MAP:
        raise ValueError(f"Unknown function name: {func_name}")
    
    func = FUNCTION_MAP[func_name]
    return func(population.astype(float)) 


def selection(population, evals, k=10, tournament_size=3):
    """
    Perform tournament selection.
    
    Args:
        population (np.ndarray): Population array of shape (N, D)
        evals (np.ndarray): Fitness values of shape (N,)
        k (int): Number of individuals to select
        tournament_size (int): Number of candidates per tournament

    Returns:
        np.ndarray: Selected individuals of shape (k, D)
    """
    N = population.shape[0]
    selected = []

    for _ in range(k):
        # Randomly select tournament_size individuals (by index)
        indices = np.random.choice(N, tournament_size, replace=False)
        best_idx = indices[np.argmin(evals[indices])]  # assuming minimization
        selected.append(population[best_idx])

    return np.array(selected)

import numpy as np

def crossover(population, c):
    """
    Generate a new population of children using one-point crossover.

    Parameters
    ----------
    population : np.ndarray
        The selected parent population of shape (k, d), where k is the number of parents and d is the dimensionality.
    c : int
        Number of children to produce.

    Returns
    -------
    np.ndarray
        A NumPy array of shape (c, d) containing the child population.
    """

    def reproduce(p1, p2):
        """
        Perform one-point crossover between two individuals.

        Parameters
        ----------
        p1, p2 : np.ndarray
            Parent individuals of shape (d,)

        Returns
        -------
        np.ndarray
            Child individual of shape (d,)
        """
        d = p1.shape[0]
        point = np.random.randint(1, d)  # crossover point in [1, d-1]
        child = np.concatenate((p1[:point], p2[point:]))
        return child

    k, d = population.shape
    children = []

    for _ in range(c):
        # Randomly select two parents from the selected pool
        p1, p2 = population[np.random.choice(k, 2, replace=True)]
        child = reproduce(p1, p2)
        children.append(child)

    return np.array(children)

import numpy as np

def mutate(population, mp=0.05, sigma=1.0):
    """
    Apply mutation to a population of individuals.

    Parameters
    ----------
    population : np.ndarray
        Population array of shape (n, d), where n is the number of individuals and d is the number of genes.
    mp : float
        Mutation probability for each individual (default is 0.05).
    sigma : float
        Standard deviation of the normal distribution used in mutation (default is 1.0).

    Returns
    -------
    np.ndarray
        Mutated population (a copy of the original if no mutations occur).
    """
    n, d = population.shape
    mutated = population.copy()

    for i in range(n):
        if np.random.rand() < mp:
            # Random direction: +1 or -1 for each gene
            noise = np.random.randn(d)  # from N(0, 1)
            delta =  sigma * noise
            mutated[i] += delta

    return mutated

def succession(population, evals, k):
    """
    Perform elitist succession: select the top-k individuals with the best (lowest) fitness.

    Parameters
    ----------
    population : np.ndarray
        Array of individuals of shape (n, d), where n is the population size and d is the dimensionality.
    evals : np.ndarray
        Array of fitness values of shape (n,). Lower values are considered better.
    k : int
        Number of individuals to keep for the next generation.

    Returns
    -------
    np.ndarray
        Array of the top-k individuals of shape (k, d).
    """
    if k > len(population):
        raise ValueError("k cannot be greater than the population size.")

    # Get indices of the top-k lowest fitness values
    best_indices = np.argsort(evals)[:k]

    # Return the best individuals
    return population[best_indices]
