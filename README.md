# _Collatz Conjecture_ visualiser

A simple application to visualize solving-loop of ***Collatz Conjectur***, takes two arguments `N`, `C`. <br>
Where `N` represents base-positive-number, and `C` represents looping-count(if it's not provided, application will loop infinitely)

## Problem explanation
The 3x + 1 problem is most simply stated in terms of the Collatz function C(x) 
defined on integers as “multiply by three and add one” for odd integers and 
“divide by two” for even integers. That is,
```
 C(x) = {
            (if x ? ODD) 3x + 1
            (if x ? EVEN) x / 2
 }
```
The 3x + 1 problem (or Collatz problem) is to prove that starting from any positive integer,
some iterate of this function takes the value 1. The problem other names: it has also been called Kakutani’s problem, the Syracuse problem, and Ulam’s problem. <br>
***For more information refer to **[References](https://github.com/theiskaa/3Nplus1#references)** header***

## Visualiser Application (REPL) Overview
<img weight=200px src="https://user-images.githubusercontent.com/59066341/155001073-476e85db-bb98-4a40-960c-e77b4d02f56a.gif">

## References
> Inspired from [Veritasium - The Simplest Math Problem No One Can Solve - Collatz Conjecture](https://www.youtube.com/watch?v=094y1Z2wpJg&t=812s)

- [Lagarias, J. C. (2003). The 3x+ 1 problem: An annotated bibliography (1963–1999). The ultimate challenge: the 3x, 1, 267-341](https://ve42.co/Lagarias2003)
- [Lagarias, J. C. (2006). The 3x+ 1 problem: An annotated bibliography, II (2000-2009). arXiv preprint math/0608208](https://ve42.co/Lagarias2006)
- [Tao, T (2020). The Notorious Collatz Conjecture](https://ve42.co/Tao2020)
