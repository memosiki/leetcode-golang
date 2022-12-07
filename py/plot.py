import matplotlib.pyplot as plt
import numpy as np

data = np.array(
    [[0, 0], [0, 2], [2, 0], [1, 1]]
)

x, y = data.T
plt.scatter(x, y)

plt.grid(True)
plt.show()
