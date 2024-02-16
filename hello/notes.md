# Notes for Hello World chapter

- I merged this chapter with the DI chapter.

- Table-driven tests seem cool, but I guess it may lead to some level of "auto-piloting" when creating tests.

- From what I have seen in the Go standard library, the idea of injecting functions and dependencies is quite popular. It leads to cleaner and more reusable code. One must be careful when using this strategy to not create too much indirection, leading to over-engineerd code.

- This approach makes code much easier to test, reuse, and compose. Also leads to much DRYer code.
