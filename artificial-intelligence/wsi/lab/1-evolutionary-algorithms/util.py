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


def tournament_selection(population, fitnesses, k=10, tournament_size=3):
    """
    Perform tournament selection.
    
    Args:
        population (np.ndarray): Population array of shape (N, D)
        fitnesses (np.ndarray): Fitness values of shape (N,)
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
        best_idx = indices[np.argmin(fitnesses[indices])]  # assuming minimization
        selected.append(population[best_idx])

    return np.array(selected)
