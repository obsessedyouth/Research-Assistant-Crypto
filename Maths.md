# Maths

I'm using `Latex` so the reader can parse the equations with the proper mathematical notation and
generally because I'm more comfortable writing equations using this.
Although github's flavor of markdown does not render it, these are embedded snippets from
[code cogs](https://www.codecogs.com/latex/eqneditor.php). You can find a pdf version with
detail in the `root` of this repository as [Math.pdf](https://github.com/obsessedyouth/Research-Assistant-Crypto/blob/main/Math.pdf) or the actual `tex` code in [equations](https://github.com/obsessedyouth/Research-Assistant-Crypto/blob/main/equations/Maths.tex).
You can scan through this markdown as a summary, I didn't solve step by step here, but simply talk
about my experience solving this problem.

## Question

Over all real numbers, find the minimum value of a positive real number,
y such that:


![equation one](https://latex.codecogs.com/svg.latex?y%20%3D%20sqrt%28%28x+6%29%5E2%20+%2025%29%20+%20sqrt%28%28x-6%29%5E2%20+%20121%29)

## Answer

*Another trick question...sorta*

To find the minimum/maximum/inflexion point of a differential equation roughly three steps are needed:

- Find the first differential ![rate of change](https://latex.codecogs.com/svg.latex?dy/dx)

- Find the value of x when ![limit of x](https://latex.codecogs.com/svg.latex?dy/dx=0)

- Make an educated guess about the the "shape" of the curve and the nature of the stationary point(s). 
By finding the second differential or otherwise and substituting the value of x.

Before all that gotta reduce those annoying squares before a math sin is committed(see pdf)

![reduced](https://latex.codecogs.com/svg.latex?y%20%3D%20sqrt%28x%5E2%20&plus;%2012x%20&plus;%2061%29%20&plus;%20sqrt%28x%5E2%20-12x%20&plus;%20157%29)

ahh that's so much better. Now we can sort of do step one, and differentiate(see pdf)

![differential](https://latex.codecogs.com/svg.latex?dy/dx%20%3D%202x%20&plus;%2012/2%20*%20sqrt%28x%5E2%20&plus;%2012x%20&plus;%2061%29%20&plus;%202x%20-%2012%20/%202%20*%20sqrt%28x%5E2%20-%2012x%20&plus;%20157%29)

Look at that a properly scary equation unlike the first :smiling_imp: The funny thing about this question is that you 
cannot "solve for x" which I figured out after about like 6 hours of insane differentiation and running into 
weird complex numbers that did not converge to any real number, then [I plotted the graph here](https://www.symbolab.com/solver/minimum-calculator/minimum%20y%20%3D%20sqrt%5Cleft(%5Cleft(x%2B6%5Cright)%5E%7B2%7D%20%2B%2025%5Cright)%20%2B%20sqrt%5Cleft(%5Cleft(x-6%5Cright)%5E%7B2%7D%20%2B%20121%5Cright)
 at once I understood what I was solving for. It's one of the creepy `V` shaped ones. I felt tricked lolol, like someone took me around the
  world of concave and convex graphs and decided to say "hahaha" did you really understand the question?
 
 Let's examine that.
 
 Over all **real numbers**, find the **minimum value** of a **positive real number**
 
 So obviously the answer will not be an imaginary complex number, ugh and MORE importantly it will never be negative.
 To find the minimum value of `y` since we know for a fact that there exist no roots to the eqn for `x` and therefore a 
 single point of inflexion which can only exist at `x >= 0` where `0` is the minimum of x, implied by the question.
 
 Now you have a boring solution, simple substitution and you're on your way, proth and black-scholes were much 
 more fascinating. :sigh:
 
