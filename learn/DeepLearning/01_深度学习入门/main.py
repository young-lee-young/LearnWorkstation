import numpy as np
import matplotlib.pyplot as plt

# 创建数据
x = np.linspace(-20, 20, 300)
y = np.linspace(-5, 5, 300)
X, Y = np.meshgrid(x, y)
Z = (1/20) * X**2 + Y**2

# ===== 3D 曲面图 =====
fig = plt.figure(figsize=(10, 7))
ax = fig.add_subplot(111, projection='3d')
ax.plot_surface(X, Y, Z, cmap='coolwarm', edgecolor='none', alpha=0.9)
ax.set_xlabel('X')
ax.set_ylabel('Y')
ax.set_zlabel('Z')
ax.set_title(r'$f(x,y) = \frac{1}{20}x^2 + y^2$', fontsize=16)
plt.savefig('3d_surface.png', dpi=150)
plt.show()

# ===== 等高线图 =====
plt.figure(figsize=(8, 6))
cs = plt.contour(X, Y, Z, levels=20, cmap='coolwarm')
plt.clabel(cs, inline=True, fontsize=8)
plt.xlabel('X')
plt.ylabel('Y')
plt.title(r'$f(x,y) = \frac{1}{20}x^2 + y^2$', fontsize=16)
plt.axis('equal')
plt.savefig('contour.png', dpi=150)
plt.show()
