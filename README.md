# Mycelium

This is a hobby project to build a programming language that solves a lot of the things that annoy me about modern microservice development.

The end goal is a full, high-level programming language (called Spore) that, upon deployment to the Mycelium Virtual Machine Network will self-organize into optimal microservices.

Each function is essentially a nanoservice and can run on any machine within the Mycelium network that fulfils its requirements.


### Tentative Plan:
1. ~~Build out (most of) the interpreter for Mycelium Bytecode for _most_ operations ("CALL" will be local only)~~
2. Implement Mycelium networking (connecting different machines, sending messages back and forth)
3. Implement the "CALL" bytecode op to take advantage of multiple machines
4. Build "Object"/ Plugin system for machine-dependent and non-pure functions
5. Begin function profiling tie-ins
6. Begin Semantic Hashing investigation
7. Begin Spore language design (Design language, implement ANTLR spec, build compiler, etc.)
8. Begin self-organization work

### Additional Goals/ Plans/ Misc. Sub-projects
- Graphic visualizer & interaction for network & node behavior & simulation
  - Would be nice to see pulse messages through the system (and how it diminishes further it goes)
  - Would be nice to see connection strengths, etc.
  - Would be nice to be able to sever entire network segments and/or core regional connections to see how it heals itself
  - 