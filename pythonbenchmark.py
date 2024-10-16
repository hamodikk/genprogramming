import pandas as pd
import statsmodels.api as sm
import timeit

# Define the anscombe data frame
anscombe = pd.DataFrame({
    'x1': [10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5],
    'x2': [10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5],
    'x3': [10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5],
    'x4': [8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8],
    'y1': [8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68],
    'y2': [9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74],
    'y3': [7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73],
    'y4': [6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89]
})

# Function to fit all models
def fit_all_models():
    set_I_design_matrix = sm.add_constant(anscombe['x1'])
    sm.OLS(anscombe['y1'], set_I_design_matrix).fit()

    set_II_design_matrix = sm.add_constant(anscombe['x2'])
    sm.OLS(anscombe['y2'], set_II_design_matrix).fit()

    set_III_design_matrix = sm.add_constant(anscombe['x3'])
    sm.OLS(anscombe['y3'], set_III_design_matrix).fit()

    set_IV_design_matrix = sm.add_constant(anscombe['x4'])
    sm.OLS(anscombe['y4'], set_IV_design_matrix).fit()

# Timing the entire fitting process
execution_time = timeit.timeit(fit_all_models, number=10000)
print(f"Average execution time for fitting all models over 1000 runs: {execution_time / 10000:.4f} seconds")