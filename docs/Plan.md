# Plan

The tentative plan for the work that will go into Mycelium.

These are high level buckets for now mostly as I'm not entirely sure how they're going to be broken out/ subdivided yet
until I start working on them. As I get more into the weeds on things, I'll update this to show what's complete, what's 
left, etc.

The bytecode section is a good example of one that's been broken out further (since that one is mostly done at this point).


## Spore


### Bytecode

#### Arithmetic Operations
- [x] Addition
- [x] Subtraction
- [ ] Division
- [x] Modulo
- [x] Multiplication

#### Logical Operations
- [ ] And
- [ ] Or
- [ ] Not

#### Comparison Operations
- [ ] Compare
- [x] Equals
- [x] Less Than
- [x] Greater Than

#### Branch Operations
- [x] Jump
- [x] Jump False
- [ ] Jump Success

#### Function Operations
- [x] Call
- [x] Return

#### Data Structure Operations
- [ ] Make Struct
- [ ] Make Array

#### Stack Operations 
- [x] Duplicate
- [x] Pop

#### Load Operations
- [x] Load Local
- [ ] Load Field
- [x] Load Constant

#### Store Operations
- [x] Store Local
- [ ] Store Field
- [x] Store Constant


### Language
- [ ] ANTLR Spec
- [ ] Compiler


## Virtual Machine Network (VMN)


- [ ] (Spike) Build out initial network organization algorithm
- [ ] (Spike) Define node -> node communication method
- [ ] (Spike) Define node -> console communication method
- [ ] Build out network layer so Mycelium nodes can talk to each other
- [ ] Build out fake network layer to simulate a network without too much (internet) network overhead
- [ ] (Spike) Build out health & healing algorithm

- [ ] (Spike) Identify semantic hashing algorithm

- [ ] Adjust "CALL" bytecode op to take advantage of multiple machines
- [ ] Build out deployment framework to distribute functions through the VMN
- [ ] Build out initial function profiling
  - [ ] CPU Cost?
  - [ ] Memory Cost?
  - [ ] Time Cost?

- [ ] Build "Object"/ Plugin system for machine-dependent & non-pure functions



## Console
- [ ] Visualize graph from communication reported by node(s)
  - [ ] Connection strengths
- [ ] Visualize message propagation
- [ ] Visualize node health
- [ ] See function profile results

